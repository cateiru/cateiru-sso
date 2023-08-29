#!/bin/bash

# 使い方:
# ./scripts/setup_migrate.sh test

ARGS=$@

docker compose exec backend_app sh -c "migrate create -ext sql -dir /migrations $ARGS && chown -R 1000:1000 /migrations"
