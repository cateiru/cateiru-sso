# Oreore.me ドキュメンテーション

Oreore.meは、WebAuthnベースの認証を提供するSSO（Single Sign-On）Identity Providerです。

## アーキテクチャ概要

- **バックエンド**: Go + Echo + SQLBoiler + MySQL
- **フロントエンド**: Next.js 13 (App Router) + Chakra UI + Jotai + SWR
- **認証**: WebAuthn + JWT/OIDC
- **インフラ**: Docker Compose (ローカル) + Google Cloud Platform (本番)

## ドキュメント構成

### 🚀 セットアップ

- [開発環境構築](./setup/development-setup.md) - 開発環境の立ち上げ手順
- [環境変数](./setup/environment-variables.md) - 必要な環境変数の詳細
- [Docker構成](./setup/docker-setup.md) - Docker Composeの詳細設定

### 🔧 バックエンド

- [アーキテクチャ](./backend/architecture.md) - Goバックエンドの構成と設計
- [テスト](./backend/testing.md) - テスト実行方法と戦略
- [データベース](./backend/database.md) - データベース設計とマイグレーション
- [認証システム](./backend/authentication.md) - WebAuthn、JWT、OIDC実装

### 🎨 フロントエンド

- [アーキテクチャ](./frontend/architecture.md) - Next.jsフロントエンドの構成
- [コンポーネント](./frontend/components.md) - コンポーネント設計と構造
- [テスト](./frontend/testing.md) - Storybook使用方法
- [状態管理](./frontend/state-management.md) - Jotaiによる状態管理

### 🚢 デプロイメント

- [本番環境](./deployment/production.md) - 本番環境へのデプロイ
- [ステージング](./deployment/staging.md) - ステージング環境
- [監視](./deployment/monitoring.md) - 監視とログ

## クイックスタート

```bash
# 開発環境立ち上げ
docker compose up

# http://localhost:3000 でアクセス可能
```

## 主な機能

- **WebAuthn認証**: パスワードレス認証
- **OIDC Provider**: OAuth2.0/OpenID Connect対応
- **多要素認証**: TOTP (Google Authenticator等)
- **組織管理**: 組織単位でのユーザー・クライアント管理
- **セッション管理**: セキュアなセッション管理
- **メール認証**: アカウント作成・パスワード再設定

## 技術スタック

### バックエンド
- **Go 1.22+**: メイン開発言語
- **Echo v4**: Webフレームワーク
- **SQLBoiler**: ORMライブラリ
- **MySQL**: データベース
- **WebAuthn**: パスワードレス認証
- **JWT**: トークンベース認証

### フロントエンド
- **Next.js 13**: Reactフレームワーク（App Router使用）
- **Chakra UI**: UIコンポーネントライブラリ
- **Jotai**: 状態管理ライブラリ
- **SWR**: データフェッチライブラリ
- **TypeScript**: 型安全性

### インフラ・ツール
- **Docker & Docker Compose**: 開発環境
- **Google Cloud Platform**: 本番環境
- **Fastly**: CDN
- **Cloud Storage**: オブジェクトストレージ
- **Storybook**: コンポーネント開発

## 貢献方法

1. このリポジトリをフォーク
2. 機能ブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add some amazing feature'`)
4. ブランチをプッシュ (`git push origin feature/amazing-feature`)
5. Pull Requestを作成

## ライセンス

このプロジェクトのライセンス情報については、LICENSEファイルを参照してください。

## サポート

問題や質問がある場合は、GitHubのIssuesで報告してください。