package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define command-line flags
	targetIP := flag.String("target", "127.0.0.1:8888", "Target IP and port to proxy to (e.g. 192.168.1.100:8080)")
	port := flag.String("port", "8000", "Port to run the proxy server on")
	flag.Parse()

	// Parse the target URL
	targetURL, err := url.Parse("http://" + *targetIP)
	if err != nil {
		log.Fatalf("Error parsing target URL: %v", err)
	}

	// Create a new reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Configure the director function (optional: for additional request modifications)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		fmt.Printf("Proxying request to: %s %s\n", req.Method, req.URL.Path)
	}

	// Create a new HTTP server
	server := http.Server{
		Addr:    ":" + *port,
		Handler: proxy,
	}

	// Start the server
	fmt.Printf("Starting reverse proxy server on port %s -> %s\n", *port, *targetIP)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
