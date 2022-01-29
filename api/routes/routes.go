package routes

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/handler"
)

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("/", handler.RootHandler)

	mux.HandleFunc("/create", handler.CreateHandler)
	mux.HandleFunc("/create/verify", handler.CreateVerifyHandler)
	mux.HandleFunc("/create/info", handler.CreateInfoHandler)

	mux.HandleFunc("/check/username", handler.CheckHandler)

	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/login/onetime", handler.LoginOnetimeHandler)
	mux.HandleFunc("/login/sso", handler.LoginSSOHandler)

	mux.HandleFunc("/me", handler.MeHandler)

	mux.HandleFunc("/admin/pro", handler.AdminProHandler)
	mux.HandleFunc("/admin/user", handler.AdminUserHandler)
	mux.HandleFunc("/admin/ban", handler.AdminBanHandler)
	mux.HandleFunc("/admin/status", handler.AdminStatusHandler)

	mux.HandleFunc("/pro/sso", handler.ProSSOHandler)

	mux.HandleFunc("/password/forget", handler.PasswordForgetHandler)
	mux.HandleFunc("/password/forget/accept", handler.PasswordForgetAcceptHandler)

	mux.HandleFunc("/user/mail", handler.UserMailHandler)
	mux.HandleFunc("/user/password", handler.UserPasswordHandler)
	mux.HandleFunc("/user/otp", handler.UserOnetimePWHandler)
	mux.HandleFunc("/user/otp/backup", handler.UserOnetimePWBackupHandler)
	mux.HandleFunc("/user/info", handler.UserInfoChangeHandler)
	mux.HandleFunc("/user/history/access", handler.UserAccessHandler)
	mux.HandleFunc("/user/history/login", handler.UserHistoryHandler)
	mux.HandleFunc("/user/avatar", handler.UserAvatarHandler)

	mux.HandleFunc("/logout", handler.LogoutHandler)

	mux.HandleFunc("/oauth/cert", handler.OAuthCertHandler)
	mux.HandleFunc("/oauth/update", handler.OAuthUpdateHandler)
}
