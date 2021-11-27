package handler

import "net/http"

// Proユーザの操作をする
func AdminProHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		adminProGetHandler(w, r)
	case http.MethodPost:
		adminProPostHandler(w, r)
	case http.MethodDelete:
		adminProDeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// 全ユーザ情報取得やユーザの削除
func AdminUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		adminUserGetHandler(w, r)
	case http.MethodDelete:
		adminUserDeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ユーザをban
func AdminBanHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		adminBanPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// workerログ取得やworkerを動かす
func AdminStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		adminStatusGetHandler(w, r)
	case http.MethodPost:
		adminStatusPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// Proユーザ一覧を取得
// adminユーザのみ
func adminProGetHandler(w http.ResponseWriter, r *http.Request) {
}

// Proユーザ追加
// adminユーザのみ
func adminProPostHandler(w http.ResponseWriter, r *http.Request) {
}

// Proユーザを削除
// adminユーザのみ
func adminProDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// 全ユーザ情報取得
// adminユーザのみ
func adminUserGetHandler(w http.ResponseWriter, r *http.Request) {
}

// `?id=[id]`を指定して該当ユーザを削除
func adminUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// ユーザをメールアドレスでBanする
func adminBanPostHandler(w http.ResponseWriter, r *http.Request) {
}

// Workerの動作ログを取得
func adminStatusGetHandler(w http.ResponseWriter, r *http.Request) {
}

// mail_tokenなどの有効期限切れのエンティティを削除など操作をする
// workerはここをcronで叩く（TODO: 専用のPWを使用する）
func adminStatusPostHandler(w http.ResponseWriter, r *http.Request) {
}
