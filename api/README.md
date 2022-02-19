# CateiruSSO API

[![codecov](https://codecov.io/gh/cateiru/cateiru-sso/branch/main/graph/badge.svg?token=YNVP7LX4WK)](https://codecov.io/gh/cateiru/cateiru-sso)
[![Lint](https://github.com/cateiru/cateiru-sso/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/cateiru/cateiru-sso/actions/workflows/test.yml)

## 環境変数

```env
# デプロイモード
# `production` or other
# `production`を設定するとdebugログが表示されません
DEPLOY_MODE=

# datastoreのホスト
# 通常、GCPサービス上にデプロイされると自動で追加されます
DATASTORE_EMULATOR_HOST=

# datastoreのプロジェクトID
DATASTORE_PROJECT_ID=

# ワンタイムパスワードなどに表示するISSUER
# サービス名
ISSUER=

# reCAPTCHAのsecret
RECAPTCHA_SECRET=

# mail gunのAPI KEY
# メール送信に使用します
MAILGUN_API_KEY=

# メール送信者のドメイン
MAIL_FROM_DOMAIN=

# メール送信者のメールアドレス
SENDER_MAIL_ADDRESS=

# Datastoreの親レベルのkey名
# デフォルトは`cateiru-sso`です
DATASTORE_PARENT_KEY=

# サイトのドメイン（パス）
SITE_DOMAIN=

# APIのドメイン（パス）
API_DOMAIN=

# cookieに適用するドメイン
# サイト、APIのドメインのルートドメインである必要があります
COOKIE_DOMAIN=

# adminのメールアドレスとパスワード
# 初回ログイン時にこの値を使用します
# adminユーザは、ログイン後、ワンタイムパスワードとパスワードの変更をする必要があります
ADMIN_MAIL=
ADMIN_PASSWORD=

# cloud storageのURL
STORAGE_URL=

# workerのパスワード
WORKER_PASSWORD=
```

## テスト

```bash
make test
```

## Dev

```bash
make dev
```

### [カスタム] DBの実行

```bash
# start db
docker-compose up -d

# stop db
docker-compose down --rmi all
```
