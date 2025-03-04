FROM alpine:latest

# Install required packages
RUN apk update && \
    apk add --no-cache \
    openconnect \
    ca-certificates \
    openssl \
    bash \
    curl \
    ip6tables \
    iptables \
    && rm -rf /var/cache/apk/*

# Default working directory
WORKDIR /vpn

# Make entrypoint script executable
COPY . ./
RUN chmod +x /vpn/entrypoint.sh

# Install Golang
RUN apk add --no-cache go git

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Copy Go source code

# Build the Go application
RUN go mod tidy
RUN go build -o /vpn/vpn-proxy /vpn/main.go

EXPOSE 8000
# Default command
ENTRYPOINT ["/vpn/entrypoint.sh"]