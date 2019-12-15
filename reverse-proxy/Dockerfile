FROM golang:alpine

WORKDIR /usr/share/reverse_proxy

COPY src/reverse-proxy.go .

RUN go build

CMD ["./reverse_proxy"]
