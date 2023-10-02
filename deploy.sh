#!/bin/bash

set -e

TAG="go_zero:$(date +%F_%H-%M-%S)"

docker build -t "$TAG" ./

docker rm -f go_zero

docker run -d \
    --network host \
    --name guofeng-api \
    --restart=on-failure:10 \
    --log-driver json-file \
    --log-opt max-size=500m \
    --log-opt max-file=2 \
    --env-file /home/ubuntu/devops/guofeng/config.env \
    "$TAG"

echo "deploy go_zero success"