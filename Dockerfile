FROM golang:1.17 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    CGO_ENABLED=0

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -ldflags="-w -s"

FROM alpine:3.14

COPY --from=builder /build/tinyurl /

ENTRYPOINT [ "/tinyurl" ]
