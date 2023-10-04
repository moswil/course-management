#!/bin/bash

IMAGE_NAME=scr.shirwil.com/shirwil-dev/keycloak
IMAGE_TAG=v0.0.1

docker buildx build --load --platform linux/arm64 \
  --build-arg KEYCLOAK_VERSION=22.0.3 --build-arg K8S_MODE=false \
  -t $IMAGE_NAME:$IMAGE_TAG --progress=plain --no-cache .

# docker buildx build --push --platform linux/amd64,linux/arm64 \
#   --build-arg KEYCLOAK_VERSION=22.0.3 --build-arg K8S_MODE=false \
#   -t $IMAGE_NAME:$IMAGE_TAG --progress=plain --no-cache .
