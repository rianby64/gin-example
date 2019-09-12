#!/bin/bash

docker container stop gin-example-go
docker network create gin-example-net

docker build -t gin-example-go -f go.Dockerfile . &&

docker run -d \
--rm \
--name gin-example-go \
--net gin-example-net \
gin-example-go
