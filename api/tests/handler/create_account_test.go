package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/cateiru-sso/api/handler"
)

func createAccountServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", handler.CreateHandler)
	mux.HandleFunc("/create/verify", handler.CreateVerifyHandler)
	mux.HandleFunc("/create/info", handler.CreateInfoHandler)

	return mux
}

// アカウント作成のテスト
//
//	メールアドレス、PWをpostしてメール認証開始
//	↓
//	メールトークンをクエリパラメータに含めたリンク踏む（GET）
//	↓
//	buffer tokenを使用してユーザ情報を入力し、アカウント作成
func TestCreateAccount(t *testing.T) {
	app := createAccountServer()
	server := httptest.NewServer(app)
	defer server.Close()

	// jar, err := cookiejar.New(nil)
	// require.NoError(t, err, "cookiejarでエラー")
	// client := &http.Client{Jar: jar}
}
