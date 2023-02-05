#!/bin/bash

# docker-composeでデータベースを起動するための便利スクリプト
# 使い方: `./scripts/docker-compose-db.sh up -d`

docker-compose -f docker-compose.db.yaml $@
