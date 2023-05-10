#!/bin/bash

# MySQLに接続するためのスクリプト
# 使用法:
#  ./scripts/sql.sh

docker compose -f ./docker/docker-compose.db.yaml exec db bash -c "mysql -u docker -pdocker cateiru-sso"
