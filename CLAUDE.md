# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

Oreore.meは、Identity Provider（IdP）として機能するSSOシステムです。GoバックエンドとNext.jsフロントエンドで構成されています。

## 開発環境

### 必須コマンド

**ローカル開発環境立ち上げ:**
```bash
docker compose up
# http://localhost:3000 (ユーザー向け)
# http://localhost:3002 (管理者向け)
```

**Goテスト実行:**
```bash
./scripts/docker-compose-db.sh up -d  # DBを先に起動
go mod download
./scripts/test.sh
```

**フロントエンドの開発:**
```bash
pnpm i
pnpm lint          # TypeScriptとESLintチェック
pnpm fix           # 自動修正
pnpm storybook     # http://localhost:6006
```

**データベース操作:**
```bash
./scripts/sql.sh   # MySQLコンソールに接続
```

### システム構成

**ローカル開発環境:**
- `backend_app` (Go): ポート8080
- `frontend_app` (Next.js): ポート3001
- `nginx` (プロキシ): ポート3000, 3002
- `db` (MySQL): ポート3306
- `gcs` (オブジェクトストレージ): ポート4443

**バックエンドアーキテクチャ:**
- エントリーポイント: `main.go`
- ソースコード: `src/` ディレクトリ
- フレームワーク: Echo (ルーティング)
- ORM: SQLBoiler
- 認証: WebAuthn + JWT
- テスト: Go標準のtestingパッケージ + go-snaps (スナップショットテスト)
- ホットリロード: Air（`.air.toml`設定）

**フロントエンドアーキテクチャ:**
- フレームワーク: Next.js 13 (App Router)
- UIライブラリ: Chakra UI
- 状態管理: Jotai
- データフェッチング: SWR
- フォーム: React Hook Form + Zod
- 認証: WebAuthn
- テスト: Storybook
- リンター: GTS (Google TypeScript Style)

**ディレクトリ構造:**
- `src/`: Goバックエンドソースコード
- `app/`: Next.js App Routerのページ
- `components/`: 再利用可能なReactコンポーネント
- `stories/`: Storybookストーリー（機能別に分類）
- `utils/`: 共通ユーティリティ関数
- `db/`: データベーススキーマとマイグレーション
- `scripts/`: 開発用スクリプト
- `templates/`: Go HTMLテンプレート
- `docker/`: Docker設定ファイル

### データベースマイグレーション

1. `db/schema.sql`を編集
2. `./scripts/setup_migrate.sh [マイグレーション名]`でマイグレーションファイル生成
3. `./scripts/migrate.sh up`でマイグレーション実行
4. `./scripts/sqlboiler.sh`でSQLBoilerモデル再生成

### 環境変数

**バックエンド設定:** `src/config.go`
**フロントエンド設定:** `utils/config.ts`

開発時は多くの環境変数は空でも動作します（reCAPTCHA、Mailgun、Fastlyなど）。

### テスト戦略

**Goテスト:**
- `./scripts/test.sh`でバックエンドテスト実行
- SQLBoilerを使用したDBテスト
- HTTPテスト用の`go-http-easy-test`
- スナップショットテストで出力検証

**フロントエンドテスト:**
- Storybookでコンポーネントの動作確認
- `pnpm lint`でコード品質チェック

### 特殊な注意点

- 管理画面アクセスはローカル環境のみ`admin@local.test`を使用
- パスワード再設定URLはDEBUGログに出力される
- WebAuthnを使用しているため、HTTPS環境で認証テストが必要
- SQLBoilerのモデルは自動生成のため手動編集禁止
- Air設定でGoファイル変更時の自動ビルド・再起動