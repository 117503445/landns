#!/usr/bin/env sh

set -e

./scripts/build.sh

docker build -t 117503445/landns -f landns.Dockerfile .
docker build -t 117503445/landns-agent -f landns-agent.Dockerfile .