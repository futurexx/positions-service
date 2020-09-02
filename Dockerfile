 
FROM golang:alpine
RUN apk add --no-cache git g++
ADD . /go/src/position-service
RUN  cd /go/src/position-service && go build -v ./cmd/position-service
ENTRYPOINT ["/go/src/position-service/position-service"]
CMD ["-config-path", "/go/src/position-service/configs/config_docker.toml"]