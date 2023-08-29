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

# -f オプションでも使いたいので、ここで判定する
if [ -n "$1" ]; then
    FILE="-f $1"
else
    FILE=""
fi

echo $FILE

for OP_DATABASE in "${OP_DATABASES[@]}" ; do
    echo "Migrate ${OP_DATABASE}..."

    MYSQL_DSN="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${HOST}:3306)/${OP_DATABASE}"
    docker compose $FILE exec db bash -c "migrate -path ${MIGRATIONS_PATH} -database \"${MYSQL_DSN}\" up"
done

