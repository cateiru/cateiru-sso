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
	e.POST("/v2/register/begin_webauthn", h.RegisterBeginWebAuthn) // Passkeyの前処理
	e.POST("/v2/register/webauthn", h.RegisterBeginWebAuthn)
	e.POST("/v2/register/password", h.RegisterPassword)
}
