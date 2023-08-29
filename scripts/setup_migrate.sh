#!/bin/bash

# マイグレーション用のup.sql, down.sql を作成します。
# 使い方:
# ./scripts/setup_migrate.sh test

ARGS=$@

docker compose exec db bash -c "migrate create -ext sql -dir /migrations $ARGS && chown -R 1000:1000 /migrations"
