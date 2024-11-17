#!/bin/bash

go build -o ./docker/echo1/echo ./cmd && \
docker build -t echo:0.0.1 -f ./docker/Dockerfile ./