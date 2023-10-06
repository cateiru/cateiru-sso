package src

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo, h *Handler, c *Config) {
	e.GET("/", h.Root)

	version := e.Group("/v2")

	// CSRF設定
	// APIにはつけたくないのでここで定義している
	if c.EnableCSRFMeasures {
		version.Use(CSRFHandler)
	}

	// アカウント登録
	// フロー:
	// 1. `/v2/register/email/send`にEmailを渡して確認コードをEmailに送信
	// 2. `/v2/register/email/verify`に確認コードを入力してEmailを認証
	// 3. `/v2/register/webauthn`か`/v2/register/password`で認証を追加
	register := version.Group("/register")
	register.POST("/email/send", h.SendEmailVerifyHandler)                     // 確認コードを送信
	register.POST("/email/resend", h.ReSendVerifyEmailHandler)                 // メールの再送信
	register.POST("/email/verify", h.RegisterVerifyEmailHandler)               // Emailの認証
	register.POST("/begin_webauthn", h.RegisterBeginWebAuthnHandler)           // Passkeyの前処理
	register.POST("/webauthn", h.RegisterWebAuthnHandler)                      // Webauthnでアカウント作成
	register.POST("/password", h.RegisterPasswordHandler)                      // パスワードでアカウント作成
	register.POST("/invite_generate_session", h.RegisterInviteRegisterSession) // 招待メールからアカウント作成セッションを構築する

	// ログイン
	login := version.Group("/login")
	login.POST("/user", h.LoginUserHandler)                    // emailでユーザのアバターとuser nameを返す
	login.POST("/begin_webauthn", h.LoginBeginWebauthnHandler) // Passkeyの前処理
	login.POST("/webathn", h.LoginWebauthnHandler)             // WebAuthnでログイン
	login.POST("/password", h.LoginPasswordHandler)            // パスワードでログイン
	login.POST("/otp", h.LoginOTPHandler)                      // OTPで認証。設定している場合、/passwordでトークンが返るのでそれと一緒に送信する

	// アカウントの認証周り操作
	account := version.Group("/account")
	account.GET("/list", h.AccountListHandler)                      // ログインしているアカウント一覧
	account.POST("/switch", h.AccountSwitchHandler)                 // ログインアカウントの変更
	account.POST("/logout", h.AccountLogoutHandler)                 // ログアウト
	account.POST("/delete", h.AccountDeleteHandler)                 // アカウント削除
	account.GET("/otp", h.AccountOTPPublicKeyHandler)               // OTPのpublic keyを返す
	account.POST("/otp", h.AccountOTPHandler)                       // OTP設定
	account.POST("/otp/delete", h.AccountDeleteOTPHandler)          // OTP削除
	account.POST("/otp/backups", h.AccountOTPBackupHandler)         // OTPのバックアップコード
	account.POST("/password", h.AccountPasswordHandler)             // パスワードの新規作成
	account.PUT("/password/update", h.AccountUpdatePasswordHandler) // パスワードの更新
	account.POST("/begin_webauthn", h.AccountBeginWebauthnHandler)
	account.GET("/webauthn", h.AccountWebauthnRegisteredDevicesHandler)
	account.POST("/webauthn", h.AccountWebauthnHandler) // passkey新規作成
	account.DELETE("/webauthn", h.AccountDeleteWebauthnHandler)
	account.GET("/certificates", h.AccountCertificatesHandler) // 認証の設定情報

	account.POST("/forget/password", h.AccountForgetPasswordHandler)                          // パスワード再登録リクエスト
	account.POST("/reregistration/available_token", h.AccountReRegisterAvailableTokenHandler) // そのセッションが有効かどうか判定する
	account.POST("/reregistration/password", h.AccountReRegisterPasswordHandler)              // パスワード更新

	// ユーザ情報
	user := version.Group("/user")
	user.GET("/me", h.UserMeHandler)
	user.PUT("/", h.UserUpdateHandler)               // ユーザ情報の更新
	user.POST("/user_name", h.UserUserNameHandler)   // ユーザー名のチェック
	user.PUT("/setting", h.UserUpdateSettingHandler) // 設定の更新
	user.GET("/brand", h.UserBrandHandler)
	user.POST("/email", h.UserUpdateEmailHandler)                  // Email変更
	user.POST("/email/register", h.UserUpdateEmailRegisterHandler) // Email変更確認コード打つ
	user.POST("/avatar", h.UserAvatarHandler)                      // アバター画像の設定
	user.DELETE("/avatar", h.UserDeleteAvatarHandler)
	user.POST("/client/logout", h.UserLogoutClientHandler) // TODO: クライアントからログアウト

	// 履歴
	history := version.Group("/history")
	history.GET("/client/login", h.HistoryClientLoginHandler)  // ログインしているSSOクライアント
	history.GET("/client", h.HistoryClientHandler)             // クライアントのログイン履歴
	history.GET("/login_devices", h.HistoryLoginDeviceHandler) // 現在ログインしているデバイス
	history.GET("/login", h.HistoryLoginHistoryHandler)        // ログイン履歴
	history.GET("/try_login", h.HistoryLoginTryHistoryHandler) // ログイントライ履歴

	// 通知
	notice := version.Group("/notice")
	notice.GET("", h.Root)
	notice.POST("/read", h.Root) // 既読にする

	// SSOクライアント
	client := version.Group("/client")
	client.GET("", h.ClientHandler)
	client.POST("", h.ClientCreateHandler) // クライアント新規作成
	client.PUT("", h.ClientUpdateHandler)  // クライアントの編集
	client.DELETE("", h.ClientDeleteHandler)
	client.GET("/config", h.ClientConfigHandler)        // クライアントの設定
	client.DELETE("/image", h.ClientDeleteImageHandler) // クライアント画像の削除
	client.GET("/allow_user", h.ClientAllowUserHandler)
	client.POST("/allow_user", h.ClientAddAllowUserHandler)
	client.DELETE("/allow_user", h.ClientDeleteAllowUserHandler)
	client.GET("/login_users", h.ClientLoginUsersHandler) // ログインしているユーザ一覧

	org := version.Group("/org")
	org.GET("/list", h.OrgGetHandler)
	org.GET("/list/simple", h.OrgGetSimpleListHandler) // クライアント一覧で使用するorgのリスト
	org.GET("/detail", h.OrgGetDetailHandler)
	org.GET("/member", h.OrgGetMemberHandler)
	org.POST("/member", h.OrgPostMemberHandler) // 招待。アカウント登録しているユーザーに対して
	org.PUT("/member", h.OrgUpdateMemberHandler)
	org.DELETE("/member", h.OrgDeleteMemberHandler)

	org.GET("/member/invite", h.OrgInvitedMemberHandler)
	org.POST("/member/invite", h.OrgInviteNewMemberHandler)      // orgの招待。アカウント登録していないユーザーに対して
	org.DELETE("/member/invite", h.OrgInviteMemberDeleteHandler) // 招待のキャンセル

	// OIDC
	oidc := version.Group("/oidc")
	oidc.POST("/require", h.OIDCRequireHandler)
	oidc.POST("/cert/begin_webauthn", h.OIDCBeginWebAuthnHandler)
	oidc.POST("/cert/webathn", h.OIDCWebAuthnHandler)
	oidc.POST("/cert/password", h.OIDCPasswordHandler)
	oidc.POST("/cert/otp", h.OIDCOTPHandler)
	oidc.POST("/login", h.Root)

	// 管理者用
	admin := version.Group("/admin")
	admin.GET("/users", h.AdminUsersHandler) // ユーザ一覧
	admin.GET("/user_detail", h.AdminUserDetailHandler)
	admin.POST("/user/brand", h.AdminUserBrandHandler)         // ブランドの付与
	admin.DELETE("/user/brand", h.AdminUserBrandDeleteHandler) // ブランドの削除
	admin.POST("/staff", h.AdminStaffHandler)                  // スタッフフラグの付与、削除
	admin.POST("/broadcast", h.AdminBroadcastHandler)

	admin.GET("/brand", h.AdminBrandHandler)
	admin.POST("/brand", h.AdminBrandCreateHandler)
	admin.PUT("/brand", h.AdminBrandUpdateHandler)
	admin.DELETE("/brand", h.AdminBrandDeleteHandler)

	admin.GET("/orgs", h.AdminOrgHandler)                    // 全org参照
	admin.GET("/org", h.AdminOrgDetailHandler)               // org詳細を取得
	admin.POST("/org", h.AdminOrgCreateHandler)              // org作成
	admin.PUT("/org", h.AdminOrgUpdateHandler)               // org編集
	admin.DELETE("/org", h.AdminOrgDeleteHandler)            // org削除
	admin.DELETE("/org/image", h.AdminOrgDeleteImageHandler) // orgの画像削除

	admin.POST("/org/member", h.AdminOrgMemberJoinHandler)     // orgにメンバー追加
	admin.DELETE("/org/member", h.AdminOrgMemberRemoveHandler) // orgからメンバー削除

	admin.GET("/clients", h.AdminClientsHandler)            // クライアント一覧
	admin.GET("/client_detail", h.AdminClientDetailHandler) // クライアント詳細

	admin.GET("/template/:name", h.AdminPreviewTemplateHTMLHandler) // テンプレートのプレビュー

	admin.GET("/register_session", h.AdminRegisterSessionHandler)          // 登録のセッション一覧を表示する
	admin.DELETE("/register_session", h.AdminDeleteRegisterSessionHandler) // 登録のセッションを削除する

	// CDN通したり、バッチ処理したり
	// Basic Auth使う
	internal := e.Group("/internal")
	internal.Use(
		middleware.BasicAuth(func(userName, password string, ctx echo.Context) (bool, error) {
			if userName == c.InternalBasicAuthUserName && password == c.InternalBasicAuthPassword {
				return true, nil
			}
			return false, nil
		}),
	)
	internal.GET("/avatar/:key/:id", h.InternalAvatarHandler)
	internal.GET("/worker", h.InternalWorkerHandler)

	// API
	api := e.Group("/api/v2")
	api.POST("/v2/jwt_key", h.Root)
	api.POST("/v2/key", h.Root)
	api.POST("/v2/login", h.Root)
}
