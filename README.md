# Oreore.me

IdP

## Quick Start

- 必要なもの
  - docker
  - docker compose

```bash
docker compose up

# http://localhost:3000
# （APIは http://localhost:8080 ）
```

### 管理画面に入る方法

> [!WARNING]
> この機能はローカル環境のみ有効です。他の環境では`staff`テーブルにINSERTしてください。

ローカル環境では、`admin@local.test`というメールアドレスをもつユーザーが作成されています。[パスワード再設定](http://localhost:3000/forget_password)からパスワードを再設定してください。

再設定用のURLはDEBUGログに以下のように出力されます。

```log
2023-09-03T13:34:11.747+0900       DEBUG   src/handler.go:160      send mail       {"email_address": "admin@local.test", "subject": "パスワードを再設定してください", "data": {"URL":"http://localhost:3000/forget_password/reregister?email=admin%40local.test&token=8K7R0stblqJLp8AyIOh3yzFYYSQl3RA","UserName":"admin","Expiration":"2023-09-03T13:39:11.747804793+09:00","BrandName":"oreore.me local","BrandUrl":"http://localhost:3000","BrandImageUrl":"https://todo","BrandDomain":"localhost:3000","Email":"admin@local.test"}}
```

## Storybook, Test and Lint

```bash
# DBは起動しておく
./script/docker-compose-db.sh up -d

# Go test
go mod download
./script/test

# Next.js lint
pnpm i
pnpm lint

# Storybook
pnpm storybook
# http://localhost:6006
```

## データベースのマイグレーション

```bash
docker compose up

# マイグレーション用の`up`, `down`ファイルを作成します。
# `YYYYMMDDhhmmss_[name].[up|down].sql` のファイルが `db/migrations/` に追加されます。
./script/setup_migrate.sh [name]

# go-migrateを使用してマイグレーションを実行します
# ローカル環境で使用するDBとテスト環境で使用するDBの2つでマイグレーションが実行されます。
./script/migrate.sh
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

## oreore.me 変更状況

- [x] ドメイン
- [x] CORS
- [x] Email
- [x] WebAuthn
- [x] タイトル
- [x] ストレージ
- [x] OTPのIssuer
- [ ] ブランチ名
- [x] ローカルデータベース
- [ ] 本番データベース
- [ ] Goのパッケージ名
- [ ] Storybook
- [ ] README
