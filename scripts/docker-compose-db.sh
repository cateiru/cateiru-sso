#!/bin/bash

docker compose -f ./docker/docker-compose.db.yaml -f ./docker/docker-compose.db-healthcheck.yaml $@
