# Cateiru SSO

Cateiru's Single Sign On

## Quick Start

- require
  - docker
    - docker compose >= [v2.20.0](https://github.com/docker/compose/releases/tag/v2.20.0)

```bash
./docker-compose up

# http://localhost:3000
# （APIは http://localhost:8080 ）
```

## Storybook, Test and Lint

```bash
# DBは起動しておく
./script/docker-compose-db.sh up -d

# Go
go mod download
./script/test

# Next.js
yarn
yarn lint

# Storybook
yarn build-storybook
yarn storybook
# http://localhost:6006
```

## Environments

- Goは[./src/config.go](./src/config.go)に`os.Getenv`があります。
- Next.jsは[./utils/config.ts](./utils/config.ts)に`process.env`があります。

```env
# APIのホスト
# Next.js側でAPIに接続する際に使用
NEXT_PUBLIC_API_HOST=[API host]

# reCAPTCHAのトークン
NEXT_PUBLIC_RE_CAPTCHA=[token]

# GAのトークン
NEXT_PUBLIC_GOOGLE_ANALYTICS_ID=[token]

# ステージング環境などの場合に設定します
NEXT_PUBLIC_PUBLICATION_TYPE=[publication type]

# コミットハッシュ
# Cloud Buildで自動的に埋めています
NEXT_PUBLIC_REVISION=[hash]

# ブランチ名
# Cloud Buildで自動的に埋めています
NEXT_PUBLIC_BRANCH_NAME=[branch name]

# reCAPTCHAのシークレット
# ローカル、テストではreCAPTCHAは使用しないので空でOK
RECAPTCHA_SECRET=[secret]

# mailgunのシークレット
# ローカル、テストではメールを送信しないので空でOK
MAILGUN_SECRET=[secert]

# fastryのトークン
# ローカル、テストではfastlyは使っていないので空でOK
FASTLY_API_TOKEN=[token]
```

## MySQLに入る

```bash
./scripts/sql.sh
```
