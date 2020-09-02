### How to Build & Run Service
----
#### via terminal
**1. Build app**: ```$> go build -v ./cmd/position-service```

**2. Run server**: ```$> ./position-service -config-path configs/config.toml```
#### via docker
**1. Build image**: ```$> docker build -t position-service .```

**2. Run container**: ```$> docker run --rm -d --name position-service -p 8089:8089 position-service:latest```
