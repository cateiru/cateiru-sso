# Cateiru認証 API

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

# パスワードのハッシュをするときのseed
# このseedは一度デプロイしたら変更しません。変更した場合すべてのハッシュ化されたPWが検証できなくなります。
# このSEEDとハッシュ化されたPWが外部に漏れた場合、パスワードが脆弱になる可能性があります。
# （ハッシュ元のPWが脆弱な場合、全探索で検証される可能性）
PW_HASH_SEED=

# send gridのAPI KEY
# メール送信に使用します
SENDGRID_API_KEY=

# メール送信者の名前
MAIL_FROM_NAME=

# メール受信者の名前
MAIL_FROM_ADDRESS=

# Datastoreの親レベルのkey名
# デフォルトは`cateiru-sso`です
DATASTORE_PARENT_KEY=

# サイトのドメイン
SITE_DOMAIN=

# APIのドメイン
API_DOMAIN=

# adminのメールアドレスとパスワード
# 初回ログイン時にこの値を使用します
# adminユーザは、ログイン後、ワンタイムパスワードとパスワードの変更をする必要があります
ADMIN_MAIL=
ADMIN_PASSWORD=
```

## テスト

```bash
make test
```

### [カスタム] DBの実行

```bash
# start db
docker-compose up -d

# stop db
docker-compose down --rmi all
```
