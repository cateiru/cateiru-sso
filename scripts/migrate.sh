#!/bin/bash

# マイグレーションを実行する

# Docker db の MySQL
MYSQL_USER="docker"
MYSQL_PASSWORD="docker"

OP_DATABASES=(
    "local"
    "test"
)

HOST="db"

MIGRATIONS_PATH="/migrations"
ARGS="$@"

for OP_DATABASE in "${OP_DATABASES[@]}" ; do
    echo "Migrate ${OP_DATABASE}..."

    MYSQL_DSN="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${HOST}:3306)/${OP_DATABASE}"
    docker compose exec db bash -c "migrate -path ${MIGRATIONS_PATH} -database \"${MYSQL_DSN}\" $ARGS"
done

