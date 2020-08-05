# What is Explorer Gate?

Noah Gate is a service which provides to clients publish prepared transactions to noah Network

<p align="center" background="black"><img src="https://raw.githubusercontent.com/noah-blockchain/explorer-gate/master/noah-explorer.jpeg" width="400"></p>


## Related services:
- [explorer-extender](https://github.com/noah-blockchain/noah-explorer-extender)
- [explorer-api](https://github.com/noah-blockchain/noah-explorer-api)
- [explorer-validators](https://github.com/noah-blockchain/noah-explorer-validators) - API for validators meta
- [explorer-tools](https://github.com/noah-blockchain/noah-explorer-tools) - common packages
- [explorer-genesis-uploader](https://github.com/noah-blockchain/explorer-genesis-uploader)

## How to use this image

```bash
docker run -d --name gate  \
    -e GATE_DEBUG=true \
    -e GATE_PORT=9000 \
    -e BASE_COIN=MNT \
    -e NODE_API=https://texasnet.node-api.noah.network/ \
    -e NODE_API_TIMEOUT=30
```

## ... via docker-compose


Example ```docker-compose.yml``` for Minoahnter Explorer Genesis Uploader:


```yml
version: '3.6'

services:
  app:
    image: noahteam/explorer-gate:latest
  ports:
      - 9000:9000
  environment:
      GATE_DEBUG: true
      GATE_PORT: 9000
      BASE_COIN: MNT
      NODE_API: https://noah-node-1.testnet.noah.network:8841/
      NODE_API_TIMEOUT: 30
```
