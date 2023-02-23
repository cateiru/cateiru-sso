package src

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo, h *Handler) {
	e.GET("/", h.Root)

	// アカウント登録
	e.POST("/v2/register/email/send", h.SendEmailVerifyHandler)
	e.POST("/v2/register/email/resend", h.ReSendVerifyEmailHandler)
	e.POST("/v2/register/email/verify", h.RegisterVerifyEmailHandler)
	e.POST("/v2/register/begin_webauthn", h.RegisterBeginWebAuthn)
	e.POST("/v2/register/webauthn", h.RegisterBeginWebAuthn)
	e.POST("/v2/register/password", h.RegisterPassword)
}
