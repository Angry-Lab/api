#!/bin/sh

DOCKER_IMAGE="${DOCKER_REGISTRY}/angry-lab/api:latest"

echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USER}" --password-stdin "${DOCKER_REGISTRY}"

docker build -t "${DOCKER_IMAGE}" .
docker push "${DOCKER_IMAGE}"
