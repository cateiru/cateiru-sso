# 環境変数

Oreore.meでは、環境に応じて設定を切り替えるために環境変数を使用しています。

## 設定ファイル

- **バックエンド**: [`src/config.go`](../../src/config.go)
- **フロントエンド**: [`utils/config.ts`](../../utils/config.ts)

## 環境モード

システムは4つの環境モードをサポートしています：

| 環境 | 説明 | バックエンド設定 |
|------|------|------------------|
| `local` | ローカル開発環境 | `LocalConfig` |
| `test` | テスト環境 | `TestConfig` |
| `cloudrun` | 本番環境 | `CloudRunConfig` |
| `cloudrun-staging` | ステージング環境 | `CloudRunStagingConfig` |

## フロントエンド環境変数

### 必須設定

```bash
# 開発時はデフォルト値が使用されるため設定不要
NODE_ENV=development
```

### オプション設定

| 変数名 | 説明 | 例 | デフォルト |
|--------|------|-----|------------|
| `NEXT_PUBLIC_API_HOST` | APIホスト | `https://api.example.com` | 相対パス |
| `NEXT_PUBLIC_RE_CAPTCHA` | reCAPTCHAサイトキー | `6LeIxAcTAAAAAJcZVRqyHh...` | - |
| `NEXT_PUBLIC_GOOGLE_ANALYTICS_ID` | Google Analytics ID | `G-XXXXXXXXXX` | - |
| `NEXT_PUBLIC_PUBLICATION_TYPE` | 公開環境タイプ | `staging` | - |
| `NEXT_PUBLIC_REVISION` | コミットハッシュ | `a1b2c3d` | `-` |
| `NEXT_PUBLIC_BRANCH_NAME` | ブランチ名 | `feature/auth` | `-` |

## バックエンド環境変数

### データベース関連

| 変数名 | 説明 | 例 |
|--------|------|-----|
| `DB_USER` | データベースユーザー | `app_user` |
| `DB_PASSWORD` | データベースパスワード | `secure_password` |
| `INSTANCE_CONNECTION_NAME` | Cloud SQLインスタンス名 | `project:region:instance` |

### 外部サービス

| 変数名 | 説明 | 例 |
|--------|------|-----|
| `RECAPTCHA_SECRET` | reCAPTCHAシークレット | `6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe` |
| `MAILGUN_SECRET` | Mailgunシークレット | `key-1234567890abcdef1234567890abcdef` |
| `FASTLY_API_TOKEN` | Fastly APIトークン | `abcdef1234567890abcdef1234567890abcdef12` |

### ストレージ

| 変数名 | 説明 | 例 |
|--------|------|-----|
| `STORAGE_EMULATOR_HOST` | Cloud Storageエミュレータ | `localhost:4443` |
| `STORAGE_URL` | ストレージURL | `localhost:4443` |

## 環境別設定例

### ローカル開発環境

```bash
# .env.local (開発時は通常不要)
NODE_ENV=development
```

### テスト環境

```bash
# テスト実行時に自動設定
export RECAPTCHA_SECRET=secret
export MAILGUN_SECRET=secret
export FASTLY_API_TOKEN=token
export STORAGE_EMULATOR_HOST=localhost:4443
export STORAGE_URL=localhost:4443
```

### ステージング環境

```bash
NODE_ENV=production
NEXT_PUBLIC_PUBLICATION_TYPE=staging
NEXT_PUBLIC_API_HOST=https://staging.oreore.me
NEXT_PUBLIC_RE_CAPTCHA=your-staging-recaptcha-key

# バックエンド
DB_USER=staging_user
DB_PASSWORD=staging_password
INSTANCE_CONNECTION_NAME=project:region:staging-instance
RECAPTCHA_SECRET=your-staging-recaptcha-secret
MAILGUN_SECRET=your-mailgun-secret
FASTLY_API_TOKEN=your-fastly-token
```

### 本番環境

```bash
NODE_ENV=production
NEXT_PUBLIC_API_HOST=https://oreore.me
NEXT_PUBLIC_RE_CAPTCHA=your-production-recaptcha-key
NEXT_PUBLIC_GOOGLE_ANALYTICS_ID=your-ga-id

# バックエンド
DB_USER=production_user
DB_PASSWORD=production_password
INSTANCE_CONNECTION_NAME=project:region:production-instance
RECAPTCHA_SECRET=your-production-recaptcha-secret
MAILGUN_SECRET=your-mailgun-secret
FASTLY_API_TOKEN=your-fastly-token
```

## 設定詳細

### reCAPTCHA設定

- **ローカル/テスト**: 無効（`UseReCaptcha: false`）
- **ステージング**: 無効（テスト用）
- **本番**: 有効（スコア閾値: 50）

### メール送信設定

- **ローカル/テスト**: 無効（`SendMail: false`）
- **ステージング/本番**: Mailgun経由で有効

### CORS設定

- **ローカル**: 無制限（開発用）
- **本番**: 厳格な設定

### セキュリティ設定

| 設定項目 | ローカル | 本番 |
|----------|----------|------|
| CSRF対策 | 無効 | 有効 |
| Cookie Secure | false | true |
| HTTPS必須 | false | true |

## 設定の優先順位

1. 環境変数
2. 設定ファイル内のデフォルト値

## セキュリティ注意事項

### 🔒 機密情報管理

- **絶対にコミットしない**
  - データベースパスワード
  - API秘密鍵
  - reCAPTCHAシークレット
  - 外部サービストークン

- **環境変数で管理する**
  ```bash
  # ❌ 悪い例 - ハードコード
  const apiKey = "sk-1234567890abcdef"

  # ✅ 良い例 - 環境変数
  const apiKey = process.env.API_SECRET_KEY
  ```

### 🔐 本番環境設定

- すべての機密情報を環境変数として設定
- Cloud Runの場合はSecret Managerを使用
- 定期的なトークンローテーション

## トラブルシューティング

### よくある問題

#### 1. 環境変数が読み込まれない

```bash
# 環境変数確認
echo $NEXT_PUBLIC_API_HOST
env | grep NEXT_PUBLIC

# Next.jsの場合は再起動が必要
pnpm dev
```

#### 2. reCAPTCHAエラー

```bash
# ローカル環境ではreCAPTCHAは無効
# ステージングでテストしたい場合は設定を変更
```

#### 3. データベース接続エラー

```bash
# 接続情報確認
./scripts/sql.sh

# 環境変数確認
echo $DB_USER
echo $DB_PASSWORD
```

### デバッグ方法

#### バックエンド設定確認

```go
// src/config.go内でログ出力
L.Debug("Config loaded",
    zap.String("mode", config.Mode),
    zap.String("host", config.Host.String()))
```

#### フロントエンド設定確認

```typescript
// utils/config.ts内でconsole.log
console.log('Config:', {
  mode: config.mode,
  apiHost: config.apiHost,
  title: config.title
});
```

## 関連ドキュメント

- [開発環境構築](./development-setup.md)
- [Docker設定](./docker-setup.md)
- [バックエンドアーキテクチャ](../backend/architecture.md)