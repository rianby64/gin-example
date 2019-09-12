#!/bin/bash

./run-go.sh &&

docker build -t gin-example-test -f test.Dockerfile . &&

docker run \
--rm \
--name gin-example-test \
--net gin-example-net \
gin-example-test &&

docker container stop gin-example-go &&
docker network rm gin-example-net
