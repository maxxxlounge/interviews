#!/usr/bin/env bash

docker pull golang
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest go test -v ./...
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest go build -v
docker run -p 8888:80  --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest ./SouthAfricanNumber
