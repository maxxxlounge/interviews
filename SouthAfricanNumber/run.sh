#!/usr/bin/env bash

docker pull golang
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest rm ./SouthAfricanNumber
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest go test -v ./... -cover
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest go build -v -o SouthAfricanNumber
docker run -p 8888:80  --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest ./SouthAfricanNumber
