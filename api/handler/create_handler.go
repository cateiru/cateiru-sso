package handler

import (
	"net/http"

	"golang.org/x/net/websocket"
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
		// プロトコルをWSにアップデーーーーート！！！
		s := websocket.Server{
			Handler: websocket.Handler(createVerifyWSHandler),
		}
		s.ServeHTTP(w, r)
	case http.MethodPost:
		createVerifyPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// メール認証を適用(?) みたいなの
func CreateAcceptHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createAcceptPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ワンタイムパスワードの処理とか
func CreateOnetimeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		createOnetimeGetHandler(w, r)
	case http.MethodPost:
		createOnetimePostHandler(w, r)
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

}

// アカウント作成（メール認証前）
// パスワード、メールアドレスを取得してメール認証を開始します
func createPostHandler(w http.ResponseWriter, r *http.Request) {

}

// メールアドレス確認待機Websocket
func createVerifyWSHandler(ws *websocket.Conn) {
}

// メールアドレスから開いたときににトークンを送信してcookie作成
func createVerifyPostHandler(w http.ResponseWriter, r *http.Request) {
}

// WS接続中メールアドレスで確認できた場合にWSを閉じてここにトークンを送信しcookie作成
func createAcceptPostHandler(w http.ResponseWriter, r *http.Request) {
}

// ワンタイムパスワードのトークンを取得する
func createOnetimeGetHandler(w http.ResponseWriter, r *http.Request) {
}

// ワンタイムパスワードを設定する
// cookieで認証します
func createOnetimePostHandler(w http.ResponseWriter, r *http.Request) {
}
