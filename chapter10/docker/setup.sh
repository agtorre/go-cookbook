#!/usr/bin/env bash

pushd example
env GOOS=linux go build -ldflags "-X main.version=1.0 -X main.builddate=$(date +%s)"
popd
docker build . -t example
docker run -d -p 8000:8000 example 
