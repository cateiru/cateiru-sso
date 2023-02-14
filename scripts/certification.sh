#!/bin/bash

# APIをHTTPS(TLS)でデバッグするために公開鍵と秘密鍵を生成します。
#
# `openssl`コマンドが必要です。`

if [ -e certificates/*.key ]; then
    echo "SSL用の鍵は作成済みです"
else
    # create key
    openssl req -x509 -out certificates/tls.crt -keyout certificates/tls.key \
    -newkey rsa:2048 -nodes -sha256 \
    -subj '/CN=100.125.206.35' -extensions EXT -config ./scripts/openssl.conf
fi
