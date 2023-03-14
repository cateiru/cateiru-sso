package src

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo, h *Handler) {
	e.GET("/", h.Root)

	// アカウント登録
	// フロー:
	// 1. `/v2/register/email/send`にEmailを渡して確認コードをEmailに送信
	// 2. `/v2/register/email/verify`に確認コードを入力してEmailを認証
	// 3. `/v2/register/webauthn`か`/v2/register/password`で認証を追加
	e.POST("/v2/register/email/send", h.SendEmailVerifyHandler)
	e.POST("/v2/register/email/resend", h.ReSendVerifyEmailHandler) // メールの再送信
	e.POST("/v2/register/email/verify", h.RegisterVerifyEmailHandler)
	e.POST("/v2/register/begin_webauthn", h.RegisterBeginWebAuthnHandler) // Passkeyの前処理
	e.POST("/v2/register/webauthn", h.RegisterWebAuthnHandler)
	e.POST("/v2/register/password", h.RegisterPasswordHandler)

	// ログイン
	e.POST("/v2/login/user", h.LoginUserHandler)                    // emailでユーザのアバターとuser nameを返す
	e.POST("/v2/login/begin_webauthn", h.LoginBeginWebauthnHandler) // Passkeyの前処理
	e.POST("/v2/login/webathn", h.LoginWebauthnHandler)
	e.POST("/v2/login/password", h.LoginPasswordHandler)
	e.POST("/v2/login/otp", h.LoginOTPHandler)

	// アカウントの認証周り操作
	e.GET("/v2/account/list", h.AccountListHandler)
	e.POST("/v2/account/switch", h.AccountSwitchHandler) // ログインアカウントの変更
	e.HEAD("/v2/account/logout", h.AccountLogoutHandler)
	e.HEAD("/v2/account/delete", h.AccountDeleteHandler)
	e.GET("/v2/account/otp", h.AccountOTPPublicKeyHandler)               // OTPのpublic keyを返す
	e.POST("/v2/account/otp", h.AccountOTPHandler)                       // OTP設定
	e.POST("/v2/account/otp/delete", h.AccountDeleteOTPHandler)          // OTP削除
	e.POST("/v2/account/otp/backups", h.AccountOTPBackupHandler)         // OTPのバックアップコード
	e.POST("/v2/account/password", h.AccountPasswordHandler)             // パスワードの新規作成
	e.PUT("/v2/account/password/update", h.AccountUpdatePasswordHandler) // パスワードの更新
	e.POST("/v2/account/begin_webauthn", h.AccountBeginWebauthnHandler)
	e.POST("/v2/account/webauthn", h.AccountWebauthnHandler)        // passkey更新or新規作成
	e.GET("/v2/account/certificates", h.AccountCertificatesHandler) // 認証の設定情報

	e.POST("/v2/account/forget/password", h.AccountForgetPasswordHandler)                          // パスワード再登録リクエスト
	e.POST("/v2/account/reregistration/available_token", h.AccountReRegisterAvailableTokenHandler) // そのセッションが有効かどうか判定する
	e.POST("/v2/account/reregistration/password", h.AccountReRegisterPasswordHandler)

	// ユーザ情報
	e.GET("/v2/user/me", h.Root)
	e.PUT("/v2/user", h.Root) // ユーザ情報の更新
	e.PUT("/v2/user/setting", h.Root)
	e.GET("/v2/user/brand", h.Root)
	e.POST("/v2/user/email", h.Root)          // Email変更
	e.POST("/v2/user/email/register", h.Root) // Email変更確認コード打つ
	e.POST("/v2/user/avatar", h.Root)         // アバター画像の設定
	e.DELETE("/v2/user/avatar", h.Root)
	e.GET("/v2/user/client/login", h.Root)   // ログインしているSSOクライアント
	e.GET("/v2/user/client/history", h.Root) // クライアントのログイン履歴
	e.POST("/v2/user/client/logout", h.Root) // クライアントからログアウト

	// 履歴
	e.GET("/v2/history/login_devices", h.Root) // 現在ログインしているデバイス
	e.GET("/v2/history/login", h.Root)         // ログイン履歴
	e.GET("/v2/history/try_login", h.Root)     // ログイントライ履歴

	// 通知
	e.GET("/v2/notice", h.Root)
	e.POST("/v2/notice/read", h.Root) // 既読にする

	// SSOクライアント
	e.GET("/v2/client/list", h.Root) // クライアント一覧
	e.GET("/v2/client", h.Root)
	e.POST("/v2/client", h.Root) // クライアント新規作成
	e.PUT("/v2/client", h.Root)  // クライアントの編集
	e.DELETE("/v2/client", h.Root)
	e.GET("/v2/client/login_users", h.Root) // ログインしているユーザ一覧

	// OIDC
	e.POST("/v2/oidc/require", h.Root)
	e.GET("/v2/oidc/cert/available_passkey", h.Root)
	e.POST("/v2/oidc/cert/begin_webauthn", h.Root)
	e.POST("/v2/oidc/cert/webathn", h.Root)
	e.POST("/v2/oidc/cert/password", h.Root)
	e.GET("/v2/oidc/quiz", h.Root)
	e.POST("/v2/oidc/quiz", h.Root)
	e.POST("/v2/oid/cert/otp", h.Root)
	e.POST("/v2/oidc/login", h.Root)

	// API
	e.POST("/api/v2/jwt_key", h.Root)
	e.POST("/api/v2/key", h.Root)
	e.POST("/api/v2/login", h.Root)

	// CDN通したり、バッチ処理したり
	e.GET("/internal/v2/avatar", h.Root)
	e.GET("/internal/v2/worker", h.Root)

	// 管理者用
	e.GET("/v2/admin/users", h.Root) // ユーザ一覧
	e.GET("/v2/admin/user_detail", h.Root)
	e.POST("/v2/admin/brand", h.Root) // ブランドの付与、削除
	e.POST("/v2/admin/staff", h.Root) // スタッフフラグの付与、削除
	e.POST("/v2/admin/broadcast", h.Root)
	e.GET("/v2/admin/user/role", h.Root)
	e.POST("/v2/admin/user/role", h.Root)
}
