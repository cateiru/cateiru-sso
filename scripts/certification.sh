#!/bin/bash

# APIをHTTPS(TLS)でデバッグするために公開鍵と秘密鍵を生成します。
#
# `openssl`コマンドが必要です。`

if [ -e certificates/*.key ]; then
    echo "SSL用の鍵は作成済みです"
else
    # create key
    openssl req -x509 -out certificates/localhost.crt -keyout certificates/localhost.key \
    -newkey rsa:2048 -nodes -sha256 \
    -subj '/CN=localhost' -extensions EXT -config ./scripts/openssl.conf
fi
