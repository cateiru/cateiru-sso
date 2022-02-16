package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/admin"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

// Proユーザの操作をする
func AdminRoleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		adminRolePostHandler(w, r)
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

func AdminMailCertLog(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		adminGetMailCertLog(w, r)
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

// Roleの追加、削除
// adminユーザのみ
func adminRolePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := admin.AdminRoleHand(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// 全ユーザ情報取得
// adminユーザのみ
func adminUserGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := admin.AllUsersHand(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// `?id=[id]`を指定して該当ユーザを削除
func adminUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := admin.DeleteUserHand(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ユーザをメールアドレスでBanする
func adminBanPostHandler(w http.ResponseWriter, r *http.Request) {
}

func adminGetMailCertLog(w http.ResponseWriter, r *http.Request) {

}

// Workerの動作ログを取得
func adminStatusGetHandler(w http.ResponseWriter, r *http.Request) {
}

// mail_tokenなどの有効期限切れのエンティティを削除など操作をする
// workerはここをcronで叩く（TODO: 専用のPWを使用する）
func adminStatusPostHandler(w http.ResponseWriter, r *http.Request) {
}
