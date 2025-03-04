# Usage

## Version 2

```sh
docker run --name vpn --rm -d \
    --env ANYCONNECT_SERVER=vpn.example.com \
    --env ANYCONNECT_USER=username \
    --env ANYCONNECT_PASSWORD=password \
    --env PROXY_TARGET_IP=127.0.0.1:8888 \
    --cap-add NET_ADMIN --privileged \
    -p 8000:8000 \
    ghcr.io/jusjira/anyconnect-vpn-proxy:v2
```

## Version 3

> Multiple Endpoint Support

```sh
docker run --name vpn --rm -d \
    --env ANYCONNECT_SERVER=vpn.example.com \
    --env ANYCONNECT_USER=username \
    --env ANYCONNECT_PASSWORD=password \
    --env PROXY_TARGET1=127.0.0.1:3000 \
    --env PROXY_TARGET2=127.0.0.1:3001 \
    --env PROXY_TARGET2_PATH=/api \
    --cap-add NET_ADMIN --privileged \
    -p 8000:8000 \
    ghcr.io/jusjira/anyconnect-vpn-proxy:v3
```

## Docker Compose

```yaml
services:
 vpn-client:
    image: ghcr.io/jusjira/anyconnect-vpn-proxy:v2
    cap_add:
     - NET_ADMIN
    secrets:
      - vpn-username
      - vpn-password
    environment:
      - ANYCONNECT_SERVER=vpn.example.com
      - ANYCONNECT_USER_FILE=/run/secrets/vpn-username
      - ANYCONNECT_PASSWORD_FILE=/run/secrets/vpn-password
      - PROXY_TARGET_IP=127.0.0.1:3001
    
secrets:
  vpn-username:
    external: true
  vpn-password:
    external: true
```
