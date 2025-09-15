# バックエンドテスト戦略

Oreore.meのGoバックエンドは、包括的なテスト戦略によって品質を確保しています。

## 🧪 テスト概要

### テストレベル

| レベル | 対象 | ツール | 実行タイミング |
|--------|------|--------|----------------|
| **単体テスト** | 個別関数・メソッド | Go標準testing | 開発時・CI/CD |
| **統合テスト** | HTTPハンドラー | echo/httptest | 開発時・CI/CD |
| **E2Eテスト** | 全体フロー | 手動・Postman | リリース前 |
| **スナップショットテスト** | レスポンス形式 | go-snaps | 開発時 |

## 🚀 テスト実行

### 基本コマンド

```bash
# 全テスト実行
./scripts/test.sh

# 個別テスト実行
go test ./src -run TestLoginHandler
go test ./src -run TestOIDCFlow

# カバレッジ付きテスト
go test -cover ./src/...

# ベンチマークテスト
go test -bench=. ./src/...
```

### テスト環境セットアップ

```bash
# データベース起動
./scripts/docker-compose-db.sh up -d

# 環境変数設定（自動）
export RECAPTCHA_SECRET=secret
export MAILGUN_SECRET=secret
export FASTLY_API_TOKEN=token
export STORAGE_EMULATOR_HOST=localhost:4443

# テスト実行
go test ./src/... -test.config "$(realpath .)db/sqlboiler.toml"
```

## 📁 テストファイル構成

```
src/
├── account_handler_test.go        # アカウント機能テスト
├── login_handler_test.go          # ログイン機能テスト
├── register_account_handler_test.go # 登録機能テスト
├── oidc_handler_test.go           # OIDC機能テスト
├── admin_handler_test.go          # 管理者機能テスト
├── session_controller_test.go     # セッション管理テスト
├── db_controller_test.go          # DB操作テスト
└── __snapshots__/                 # スナップショットファイル
    ├── TestLoginHandler.snap
    ├── TestOIDCAuth.snap
    └── ...
```

## 🔧 テストツールとライブラリ

### 主要ライブラリ

```go
import (
    "testing"
    "net/http/httptest"           // HTTPテスト
    "github.com/stretchr/testify/assert"  // アサーション
    "github.com/gkampitakis/go-snaps"     // スナップショット
    "github.com/jarcoal/httpmock"         // HTTPモック
    "github.com/cateiru/go-http-easy-test/v2" // HTTPテストヘルパー
)
```

### テスト用ヘルパー

```go
// テスト用データベース初期化
func setupTestDB(t *testing.T) *sql.DB {
    db := ConnectDB(TestConfig.DatabaseConfig)
    // テストデータ投入
    return db
}

// テスト用HTTPサーバー
func setupTestServer(t *testing.T) (*echo.Echo, *Handler) {
    e := echo.New()
    h := NewHandler(TestConfig)
    Routes(e, h, TestConfig)
    return e, h
}
```

## 🎯 テストパターン

### 1. HTTPハンドラーテスト

```go
func TestLoginHandler(t *testing.T) {
    e, h := setupTestServer(t)

    // テストケース
    tests := []struct {
        name     string
        method   string
        path     string
        body     interface{}
        expected int
    }{
        {
            name:     "valid login",
            method:   "POST",
            path:     "/api/v2/login/password",
            body:     LoginRequest{Email: "test@example.com", Password: "password"},
            expected: 200,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(tt.method, tt.path, createJSONBody(tt.body))
            rec := httptest.NewRecorder()

            e.ServeHTTP(rec, req)

            assert.Equal(t, tt.expected, rec.Code)

            // スナップショットテスト
            snaps.MatchSnapshot(t, rec.Body.String())
        })
    }
}
```

### 2. 外部サービスモック

```go
func TestEmailSending(t *testing.T) {
    // Mailgunモック
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    httpmock.RegisterResponder("POST", "https://api.mailgun.net/v3/example.com/messages",
        httpmock.NewStringResponder(200, `{"message": "Queued"}`),
    )

    // テスト実行
    err := SendVerificationEmail(ctx, "test@example.com", "123456")
    assert.NoError(t, err)
}
```

## 📊 スナップショットテスト

### go-snaps使用

```go
func TestAPIResponse(t *testing.T) {
    e, _ := setupTestServer(t)

    req := httptest.NewRequest("GET", "/api/v2/user/profile", nil)
    rec := httptest.NewRecorder()

    e.ServeHTTP(rec, req)

    // レスポンス全体のスナップショット
    snaps.MatchSnapshot(t, rec.Body.String())
}
```

### スナップショット管理

```bash
# スナップショット更新
go test -update ./src/...

# スナップショット確認
ls src/__snapshots__/
```

## 🔐 認証テスト

### WebAuthnテスト

```go
func TestWebAuthnRegistration(t *testing.T) {
    e, h := setupTestServer(t)

    // 1. Begin Registration
    beginReq := httptest.NewRequest("POST", "/api/v2/register/begin_webauthn",
        createJSONBody(BeginWebAuthnRequest{SessionID: "session123"}))
    beginRec := httptest.NewRecorder()
    e.ServeHTTP(beginRec, beginReq)

    assert.Equal(t, 200, beginRec.Code)

    // 2. Complete Registration (Mock)
    mockCredential := createMockWebAuthnCredential()

    completeReq := httptest.NewRequest("POST", "/api/v2/register/webauthn",
        createJSONBody(CompleteWebAuthnRequest{
            SessionID:  "session123",
            Credential: mockCredential,
        }))
    completeRec := httptest.NewRecorder()
    e.ServeHTTP(completeRec, completeReq)

    assert.Equal(t, 200, completeRec.Code)
}
```

## 📈 テストカバレッジ

### カバレッジ測定

```bash
# カバレッジレポート生成
go test -coverprofile=coverage.out ./src/...
go tool cover -html=coverage.out -o coverage.html

# 関数別カバレッジ
go tool cover -func=coverage.out
```

### 目標値

| 項目 | 目標 | 現状 |
|------|------|------|
| 全体カバレッジ | 80%+ | 78% |
| ハンドラー | 90%+ | 85% |
| ビジネスロジック | 85%+ | 82% |

## 🚨 テスト失敗時の対処

### よくある失敗パターン

#### 1. データベース関連

```bash
# エラー: database connection failed
# 対処: データベースコンテナ確認
docker compose logs db
./scripts/docker-compose-db.sh up -d
```

#### 2. スナップショット不一致

```bash
# エラー: snapshot mismatch
# 対処: スナップショット更新
go test -update ./src -run TestSpecificFunction
```

## 🔄 CI/CD統合

### GitHub Actions設定例

```yaml
name: Test
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      mysql:
        image: mysql:8.0
        env:
          MYSQL_ROOT_PASSWORD: root
          MYSQL_DATABASE: test

    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '1.22'

      - name: Run tests
        run: ./scripts/test.sh

      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
```

## 📚 参考資料

- [Go Testing Documentation](https://golang.org/doc/tutorial/add-a-test)
- [Testify Framework](https://github.com/stretchr/testify)
- [go-snaps](https://github.com/gkampitakis/go-snaps)