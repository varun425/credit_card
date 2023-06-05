#!/bin/bash

# Check if the CouchDB image is already present
if [[ "$(docker images -q couchdb:3.3.2 2> /dev/null)" == "" ]]; then
    # Pull the CouchDB image from Docker Hub
    docker pull couchdb:3.3.2
fi

# Check if the network exists, if not create it
NETWORK_NAME="poc"
if [[ "$(docker network ls -q -f name=${NETWORK_NAME})" == "" ]]; then
    docker network create ${NETWORK_NAME}
fi

# Start the CouchDB container using Docker Compose
docker-compose -f docker-compose.yaml up -d
