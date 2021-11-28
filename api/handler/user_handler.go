package handler

import "net/http"

// 自分のメールアドレスの操作
func UserMailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userMailGetHandler(w, r)
	case http.MethodPost:
		userMailPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// パスワード変更
func UserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userPasswordPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// パスワード忘れの再登録用
func UserPasswordForgetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userPasswordForgetPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ワンタイムパスワードの無効化、有効化
func UserOnetimeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userOnetimePostHandler(w, r)
	case http.MethodGet:
		createOnetimeGetHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ワンタイムパスワードのバックアップコードを表示する
func UserOnetimeBackupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		UserOnetimeBackupGetHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ログインしているSSOの操作
func UserAccessHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userAccessGetHandler(w, r)
	case http.MethodPost:
		userAccessPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// アカウントのログイン履歴
func UserHistoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userHistoryGetHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// メールアドレス取得
// 全ユーザ
// `/me`でも取得できる
func userMailGetHandler(w http.ResponseWriter, r *http.Request) {
}

// メールアドレスの更新
func userMailPostHandler(w http.ResponseWriter, r *http.Request) {
}

// パスワード更新
// 現在のパスワードを送信するか、パスワード忘れ用の再登録トークンを送信
func userPasswordPostHandler(w http.ResponseWriter, r *http.Request) {
}

// パスワードを忘れた場合の再登録
//
// メールアドレスを送信して、そのメールアドレスの持ったアカウントが存在する場合に、
// トークンをパラメータに付与したURLをメール送信
// UserPasswordHandlerでPW変更する
func userPasswordForgetPostHandler(w http.ResponseWriter, r *http.Request) {
}

// ワンタイムパスワードのトークンURLを取得する
func createOnetimeGetHandler(w http.ResponseWriter, r *http.Request) {
}

// ワンタイムパスワードの無効化、有効化
func userOnetimePostHandler(w http.ResponseWriter, r *http.Request) {
}

// ワンタイムパスワードのバックアップコードを返す
func UserOnetimeBackupGetHandler(w http.ResponseWriter, r *http.Request) {
}

// ログインしているSSOを取得
func userAccessGetHandler(w http.ResponseWriter, r *http.Request) {
}

// ログインしているSSOからログアウト
func userAccessPostHandler(w http.ResponseWriter, r *http.Request) {
}

// アカウントのログイン履歴取得
func userHistoryGetHandler(w http.ResponseWriter, r *http.Request) {
}
