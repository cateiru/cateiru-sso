package routes

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/handler"
)

func handleFunc(mux *http.ServeMux, path string, hand func(w http.ResponseWriter, r *http.Request)) {
	formattedPath := config.Defs.ApiVersion + path

	mux.HandleFunc(formattedPath, hand)
}

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("/", handler.RootHandler)

	handleFunc(mux, "/create", handler.CreateHandler)
	handleFunc(mux, "/create/verify", handler.CreateVerifyHandler)
	handleFunc(mux, "/create/info", handler.CreateInfoHandler)

	handleFunc(mux, "/check/username", handler.CheckHandler)

	handleFunc(mux, "/login", handler.LoginHandler)
	handleFunc(mux, "/login/onetime", handler.LoginOnetimeHandler)

	handleFunc(mux, "/me", handler.MeHandler)

	handleFunc(mux, "/admin/role", handler.AdminRoleHandler)
	handleFunc(mux, "/admin/user", handler.AdminUserHandler)
	handleFunc(mux, "/admin/ban", handler.AdminBanHandler)
	handleFunc(mux, "/admin/certlog", handler.AdminMailCertLog)
	handleFunc(mux, "/admin/worker", handler.AdminWorker)

	handleFunc(mux, "/pro/sso", handler.ProSSOHandler)
	handleFunc(mux, "/pro/sso/image", handler.ProSSOImage)

	handleFunc(mux, "/password/forget", handler.PasswordForgetHandler)
	handleFunc(mux, "/password/forget/accept", handler.PasswordForgetAcceptHandler)

	handleFunc(mux, "/user/mail", handler.UserMailHandler)
	handleFunc(mux, "/user/password", handler.UserPasswordHandler)
	handleFunc(mux, "/user/otp", handler.UserOnetimePWHandler)
	handleFunc(mux, "/user/otp/backup", handler.UserOnetimePWBackupHandler)
	handleFunc(mux, "/user/otp/me", handler.UserOTPMeHandler)
	handleFunc(mux, "/user/info", handler.UserInfoChangeHandler)
	handleFunc(mux, "/user/history/access", handler.UserAccessHandler)
	handleFunc(mux, "/user/history/login", handler.UserHistoryHandler)
	handleFunc(mux, "/user/avatar", handler.UserAvatarHandler)
	handleFunc(mux, "/user/oauth", handler.UserOAuthHandler)

	handleFunc(mux, "/logout", handler.LogoutHandler)

	handleFunc(mux, "/oauth/preview", handler.OAuthPreview)
	handleFunc(mux, "/oauth/login", handler.OAuthLogin)
	handleFunc(mux, "/oauth/token", handler.OAuthToken)
	handleFunc(mux, "/oauth/jwt/key", handler.OAuthJWTKey)
}
