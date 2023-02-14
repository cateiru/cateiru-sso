package src

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// 最初にメールアドレス宛に確認コードを送信する
// アカウント作成フローの一番はじめ
func (h *Handler) SendEmailVerifyHandler(c echo.Context) error {
	email := c.FormValue("email")
	recaptcha := c.FormValue("recaptcha")
	ip := c.RealIP()

	// Emailの形式が正しいか検証する
	if !lib.ValidateEmail(email) {
		return lib.NewHTTPError(http.StatusBadRequest, "invalid email")
	}

	// reCAPTCHA
	if h.C.UseReCaptcha {
		order, err := h.ReCaptcha.ValidateOrder(recaptcha, ip)
		if err != nil {
			return err
		}
		// 検証に失敗した or スコアが閾値以下の場合はエラーにする
		if !order.Success || order.Score < h.C.ReCaptchaAllowScore {
			return lib.NewHTTPUniqueError(http.StatusBadRequest, lib.ErrReCaptcha, "reCAPTCHA validation failed")
		}
	}

	userData, err := ParseUA(c.Request())
	if err != nil {
		return err
	}

	// メールを送信するのでログを出す
	L.Info("mail",
		zap.String("Email", email),
		zap.String("IP", ip),
		zap.String("Device", userData.Device),
		zap.String("Browser", userData.Browser),
		zap.String("OS", userData.OS),
		zap.Bool("IsMobile", userData.IsMobile),
	)

	return nil
}
