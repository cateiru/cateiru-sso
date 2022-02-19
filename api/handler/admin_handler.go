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
	case http.MethodGet:
		adminBanGetHandler(w, r)
	case http.MethodPost:
		adminBanPostHandler(w, r)
	case http.MethodDelete:
		adminBanDeleteHandler(w, r)
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

func AdminWorker(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		adminWorkerGet(w, r)
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

func adminBanGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := admin.BanHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ユーザをメールアドレス、IPでBanする
func adminBanPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := admin.SetBanHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func adminBanDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := admin.DeleteBlocks(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func adminGetMailCertLog(w http.ResponseWriter, r *http.Request) {
	if err := admin.MailCertLogHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func adminWorkerGet(w http.ResponseWriter, r *http.Request) {

}
