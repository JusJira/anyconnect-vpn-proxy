package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	// Define command-line flags
	targetIP1 := flag.String("target1", "127.0.0.1:8888", "First target IP and port to proxy to (e.g. 192.168.1.100:8080)")
	targetIP2 := flag.String("target2", "", "Second target IP and port to proxy to (optional)")
	target2Path := flag.String("target2path", "/api", "Path prefix to route to the second target")
	port := flag.String("port", "8000", "Port to run the proxy server on")
	flag.Parse()

	// Parse the first target URL
	targetURL1, err := url.Parse("http://" + *targetIP1)
	if err != nil {
		log.Fatalf("Error parsing first target URL: %v", err)
	}

	// Create the first reverse proxy
	proxy1 := httputil.NewSingleHostReverseProxy(targetURL1)

	// Configure director for first proxy
	originalDirector1 := proxy1.Director
	proxy1.Director = func(req *http.Request) {
		originalDirector1(req)
		fmt.Printf("Proxying request to target1: %s %s\n", req.Method, req.URL.Path)
	}

	var proxy2 *httputil.ReverseProxy

	// Setup second proxy if target2 is provided
	if *targetIP2 != "" {
		targetURL2, err := url.Parse("http://" + *targetIP2)
		if err != nil {
			log.Fatalf("Error parsing second target URL: %v", err)
		}

		proxy2 = httputil.NewSingleHostReverseProxy(targetURL2)

		// Configure director for second proxy
		originalDirector2 := proxy2.Director
		proxy2.Director = func(req *http.Request) {
			// Remove the path prefix before forwarding the request
			req.URL.Path = strings.TrimPrefix(req.URL.Path, *target2Path)
			if req.URL.Path == "" || req.URL.Path[0] != '/' {
				req.URL.Path = "/" + req.URL.Path
			}

			originalDirector2(req)
			fmt.Printf("Proxying request to target2: %s %s\n", req.Method, req.URL.Path)
		}
	}

	// Create router function
	router := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if proxy2 != nil && strings.HasPrefix(r.URL.Path, *target2Path) {
			// Route to second target if path matches and second target is configured
			proxy2.ServeHTTP(w, r)
		} else {
			// Default route to first target
			proxy1.ServeHTTP(w, r)
		}
	})

	// Create a new HTTP server
	server := http.Server{
		Addr:    ":" + *port,
		Handler: router,
	}

	// Start the server
	if *targetIP2 != "" {
		fmt.Printf("Starting multi-target proxy server on port %s\n", *port)
		fmt.Printf("- Default target: %s\n", *targetIP1)
		fmt.Printf("- Path %s -> %s\n", *target2Path, *targetIP2)
	} else {
		fmt.Printf("Starting single-target proxy server on port %s -> %s\n", *port, *targetIP1)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
