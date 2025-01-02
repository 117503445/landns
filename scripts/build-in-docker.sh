#!/usr/bin/env sh

set -e
sh
protoc --go_out=. --go_opt=paths=source_relative --twirp_out=. ./pkg/rpcgen/landns.proto
go build -o landns ./cmd/landns
go build -o landns-agent ./cmd/landns-agent