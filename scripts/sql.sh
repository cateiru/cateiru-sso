#!/bin/bash

# MySQLに接続するためのスクリプト
# 使用法:
#  ./scripts/sql.sh

docker compose exec db bash -c "mysql -u docker -pdocker local"
