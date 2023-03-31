#!/bin/bash

# APIをHTTPS(TLS)でデバッグするために公開鍵と秘密鍵を生成します。
#
# `openssl`コマンドが必要です。`

if [ -e certificates/*.key ]; then
    echo "SSL用の鍵は作成済みです"
    echo "新しく作成する場合は rm -rf certificates/tls.* を実行してください"
else
    # create key
    openssl req -x509 -out certificates/tls.crt -keyout certificates/tls.key \
    -newkey rsa:2048 -nodes -sha256 -days 3650 \
    -subj '/CN=localhost' -extensions EXT -config ./scripts/openssl.conf
fi
