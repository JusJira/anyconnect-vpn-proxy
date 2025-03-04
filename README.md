# Usage

```sh
docker run --name vpn --rm -d \
    --env ANYCONNECT_SERVER=vpn.example.com \
    --env ANYCONNECT_USER=username \
    --env ANYCONNECT_PASSWORD=password \
    --env PROXY_TARGET_IP=127.0.0.1:8888 \
    --cap-add NET_ADMIN --privileged \
    -p 8000:8000 \
    ghcr.io/jusjira/anyconnect-vpn-proxy:latest
```
