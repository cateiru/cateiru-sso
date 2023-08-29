#!/bin/bash

# マイグレーションを実行する

# Docker db の MySQL
MYSQL_USER="docker"
MYSQL_PASSWORD="docker"

OP_DATABASES=(
    "cateiru-sso"
    "cateiru-sso-test"
)

HOST="db"

MIGRATIONS_PATH="/migrations"

for OP_DATABASE in "${OP_DATABASES[@]}" ; do
    MYSQL_DSN="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${HOST}:3306)/${OP_DATABASE}"
    docker compose exec backend_app sh -c "migrate -path ${MIGRATIONS_PATH} -database \"${MYSQL_DSN}\" up"
done

