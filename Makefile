.PHONY: test format build
SHELL := /bin/bash

test:
	go test -v -coverprofile=cover.out.tmp -coverpkg=./... ./...

format:	# Formats Go files
	goimports -local "github.com/idoqo/starlink_exporter/" -l -w ./
	for f in $(git status -s | cut -f3 -d ' '); do go fmt "$f"; done
	gci write --section Standard --section Default --section "Prefix(github.com/idoqo/starlink_exporter)" $(shell find . -type f -name "*.go")

build:
	go build -o ./dist/starlink_exporter ./cmd/starlink_exporter/main.go