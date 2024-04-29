ARG ARCH=

FROM golang:1.22-alpine as builder

RUN apk update && apk add --no-cache git bash

WORKDIR /starlink_exporter
COPY go.mod go.sum /starlink_exporter/
RUN go get ./...
ADD . /starlink_exporter

RUN go build -o application ./cmd/starlink_exporter/main.go

FROM alpine:latest

COPY --from=builder /starlink_exporter/application /starlink_exporter


EXPOSE 9817
ENTRYPOINT ["/starlink_exporter"]
