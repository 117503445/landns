#!/usr/bin/env sh

set -e

docker build -t 117503445/landns-builder -f builder.Dockerfile .
docker run --rm -v "$(pwd):/workspace" -v "GOCACHE:/root/.cache/go-build" -v "GOMODCACHE:/go/pkg/mod" 117503445/landns-builder
# docker run --rm -it --entrypoint sh -v "$(pwd):/workspace" -v "GOCACHE:/root/.cache/go-build" -v "GOMODCACHE:/go/pkg/mod" 117503445/landns-builder 