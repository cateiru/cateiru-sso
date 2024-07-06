#!/bin/bash

# マイグレーション用のup.sql, down.sql を作成します。
# 使い方:
# ./scripts/setup_migrate.sh test

docker compose exec db sh /scripts/migrate.sh $1
