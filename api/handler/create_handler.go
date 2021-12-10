package handler

import (
	"net/http"

	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

// アカウント作成: 初期
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// メール認証
func CreateVerifyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		createVerifyWSHandler(w, r)
	case http.MethodPost:
		createVerifyPostHandler(w, r)
	case http.MethodHead:
		createAcceptPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ユーザ情報決定
func CreateInfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createInfoPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ユーザの情報を追加する
// 名前、テーマ、プロフィール画像
// アカウント登録時のみ有効: 変更時は/userでやる
func createInfoPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := createaccount.CreateInfoHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// アカウント作成（メール認証前）
// パスワード、メールアドレスを取得してメール認証を開始します
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := createaccount.CreateTemporaryHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// メールアドレス確認待機Websocket
func createVerifyWSHandler(w http.ResponseWriter, r *http.Request) {
	if err := createaccount.MailVerifyObserve(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// メールアドレスから開いたときににトークンを送信してcookie作成
func createVerifyPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := createaccount.CreateVerifyHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// WS接続中メールアドレスで確認できた場合にWSを閉じてここにトークンを送信しcookie作成
func createAcceptPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := createaccount.CreateAcceptHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
