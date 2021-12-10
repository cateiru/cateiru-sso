package routes

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/handler"
)

func Routes(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("/", handler.RootHandler)

	mux.HandleFunc("/create", handler.CreateHandler)
	mux.HandleFunc("/create/verify", handler.CreateVerifyHandler)
	mux.HandleFunc("/create/info", handler.CreateInfoHandler)

	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/login/onetime", handler.LoginOnetimeHandler)
	mux.HandleFunc("/login/sso", handler.LoginSSOHandler)

	mux.HandleFunc("/me", handler.MeHandler)

	mux.HandleFunc("/admin/pro", handler.AdminProHandler)
	mux.HandleFunc("/admin/user", handler.AdminUserHandler)
	mux.HandleFunc("/admin/ban", handler.AdminBanHandler)
	mux.HandleFunc("/admin/status", handler.AdminStatusHandler)

	mux.HandleFunc("/pro/sso", handler.ProSSOHandler)

	mux.HandleFunc("/user/mail", handler.UserMailHandler)
	mux.HandleFunc("/user/password", handler.UserPasswordHandler)
	mux.HandleFunc("/user/password/forget", handler.UserPasswordForgetHandler)
	mux.HandleFunc("/user/onetime", handler.UserOnetimeHandler)
	mux.HandleFunc("/user/onetime/backup", handler.UserOnetimeBackupGetHandler)
	mux.HandleFunc("/user/access", handler.UserAccessHandler)
	mux.HandleFunc("/user/history", handler.UserHistoryHandler)

	mux.HandleFunc("/logout", handler.LogoutHandler)

	mux.HandleFunc("/oauth/cert", handler.OAuthCertHandler)
	mux.HandleFunc("/oauth/update", handler.OAuthUpdateHandler)

	return mux
}
