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
