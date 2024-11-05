#!/bin/bash

export PATH=$PATH:/usr/local/go/bin
go build -o ./docker/echo1/echo ./cmd/echo/ && \
docker build -t echo1:0.0.2 -f ./docker/echo1/dockerfile.echo ./docker/echo1