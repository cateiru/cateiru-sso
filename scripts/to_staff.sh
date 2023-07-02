#!/bin/bash
# 最近作成されたユーザーをスタッフにするスクリプト

docker compose -f ./docker/docker-compose.db.yaml exec db bash -c "mysql -u docker -pdocker cateiru-sso -e'INSERT INTO staff (user_id) VALUES ((SELECT id FROM user ORDER BY created_at DESC LIMIT 1));'"
