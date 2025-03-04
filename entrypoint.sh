#!/bin/sh

# Start OpenConnect VPN client in the background
echo "Starting VPN connection..."
# Check if password file is provided and exists
if [ ! -z "$ANYCONNECT_PASSWORD_FILE" ] && [ -f "$ANYCONNECT_PASSWORD_FILE" ]; then
  echo "Using password from file..."
  ANYCONNECT_PASSWORD=$(cat "$ANYCONNECT_PASSWORD_FILE")
fi

# Check if password file is provided and exists
if [ ! -z "$ANYCONNECT_USER_FILE" ] && [ -f "$ANYCONNECT_USER_FILE" ]; then
  echo "Using password from file..."
  ANYCONNECT_USER=$(cat "$ANYCONNECT_USER_FILE")
fi

( echo yes; echo $ANYCONNECT_PASSWORD ) | openconnect $ANYCONNECT_SERVER --user=$ANYCONNECT_USER --timestamp --background

# Give the VPN a moment to establish
sleep 5

# Start the Go reverse proxy
echo "Starting reverse proxy..."
./vpn-proxy -targetIP=$PROXY_TARGET_IP -port=$PROXY_PORT

# If the Go binary exits, also terminate the VPN connection
# This ensures clean shutdown
pid=$(pgrep openconnect)
if [ ! -z "$pid" ]; then
  kill $pid
fi