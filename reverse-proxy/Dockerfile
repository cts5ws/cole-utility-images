FROM golang:alpine

WORKDIR /usr/share/reverse-proxy

COPY src/ ./src

CMD ["go", "run", "github.com/cts5ws/server"]
