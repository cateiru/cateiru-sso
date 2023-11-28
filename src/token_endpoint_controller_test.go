package src_test

import "testing"

func TestClientAuthentication(t *testing.T) {
	t.Run("成功: Basic認証", func(t *testing.T) {})

	t.Run("成功: POST", func(t *testing.T) {})

	t.Run("失敗: どの認証も無い", func(t *testing.T) {})

	t.Run("失敗: Basic認証でAuthorizationの形式が不正", func(t *testing.T) {})

	t.Run("失敗: Basic認証でBase64デコードに失敗", func(t *testing.T) {})

	t.Run("失敗: Basic認証でクライアントが存在しない", func(t *testing.T) {})

	t.Run("失敗: POSTでクライアントが存在しない", func(t *testing.T) {})

	t.Run("失敗: Basic認証でクライアントシークレットが不正", func(t *testing.T) {})

	t.Run("失敗: POSTでクライアントシークレットが不正", func(t *testing.T) {})
}
