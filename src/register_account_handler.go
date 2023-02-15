package src

import (
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

// テンプレートにわたすやつ
type RegisterEmailVerify struct {
	Code     string
	Email    string
	Time     time.Time
	UserData *UserData
}

type RegisterEmailResponse struct {
	Token string `json:"register_token"`
}

// 最初にメールアドレス宛に確認コードを送信する
// アカウント作成フローの一番はじめ
func (h *Handler) SendEmailVerifyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.FormValue("email")
	recaptcha := c.FormValue("recaptcha")
	ip := c.RealIP()

	// Emailの形式が正しいか検証する
	if !lib.ValidateEmail(email) {
		return NewHTTPError(http.StatusBadRequest, "invalid email")
	}
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}

	// reCAPTCHA
	if h.C.UseReCaptcha {
		order, err := h.ReCaptcha.ValidateOrder(recaptcha, ip)
		if err != nil {
			return err
		}
		// 検証に失敗した or スコアが閾値以下の場合はエラーにする
		if !order.Success || order.Score < h.C.ReCaptchaAllowScore {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrReCaptcha, "reCAPTCHA validation failed")
		}
	}

	// セッションテーブルにEmailが存在しているか
	// 有効期限が切れるまで同じメールアドレスでセッションを作れないようにする
	// スパム防止のため
	existsEmailInRegisterSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.Email.EQ(email),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if existsEmailInRegisterSession {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrSessionExists, "session exists")
	}

	// すでに登録されているメールアドレスの場合は登録できない（Emailがユニークなため）
	exitsEmailInUser, err := models.Users(
		models.UserWhere.Email.EQ(email),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if exitsEmailInUser {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrImpossibleRegisterAccount, "impossible register")
	}

	userData, err := ParseUA(c.Request())
	if err != nil {
		return err
	}
	session, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	code, err := lib.RandomNumber(6)
	if err != nil {
		return err
	}

	// セッションを作成する
	// email_verifiedはデフォルトfalse
	// retry_countはデフォルト0
	sessionDB := models.RegisterSession{
		ID:         session,
		Email:      email,
		VerifyCode: code,

		Period: time.Now().Add(h.C.RegisterSessionPeriod),
	}
	if err := sessionDB.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	// 対象のメールアドレスにメールを送信する
	r := RegisterEmailVerify{
		Code:     code,
		Email:    email,
		Time:     time.Now(),
		UserData: userData,
	}
	m := &lib.MailBody{
		EmailAddress:      email,
		Subject:           "メールアドレスの登録確認",
		Data:              r,
		PlainTextFileName: "register.gtpl",
		HTMLTextFileName:  "register.html",
	}
	msg, id, err := h.Sender.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", email),
			zap.Error(err),
			zap.String("IP", ip),
			zap.String("Device", userData.Device),
			zap.String("Browser", userData.Browser),
			zap.String("OS", userData.OS),
			zap.Bool("IsMobile", userData.IsMobile),
		)
		return err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("Email", email),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", ip),
		zap.String("Device", userData.Device),
		zap.String("Browser", userData.Browser),
		zap.String("OS", userData.OS),
		zap.Bool("IsMobile", userData.IsMobile),
	)

	resp := &RegisterEmailResponse{
		Token: session,
	}
	return c.JSON(http.StatusOK, resp)
}
