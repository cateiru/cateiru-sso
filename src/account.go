package src

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type AccountUser struct {
	UserName string `json:"user_name"`
	ID       string `json:"id"`
	Avatar   string `json:"avatar,omitempty"`
}

type AccountOTPPublic struct {
	OTPSession string `json:"otp_session"`
	PublicKey  string `json:"public_key"`
}

type AccountReRegisterPasswordTemplate struct {
	SessionToken string
	Email        string
	PeriodTime   time.Time
	Now          time.Time
}

type AccountCertificates struct {
	Password bool `json:"password"`
	OTP      bool `json:"otp"`
	Passkey  bool `json:"passkey"`
}

// ログイン可能なアカウントのリストを返すハンドラ
// 複数アカウントにログインしている場合のやつ
func (h *Handler) AccountListHandler(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.Session.LoggedInAccounts(ctx, c.Cookies())
	if err != nil {
		return err
	}

	accountUsers := make([]AccountUser, len(users))
	for i, u := range users {
		accountUsers[i] = AccountUser{
			UserName: u.UserName,
			ID:       u.ID,
			Avatar:   u.Avatar.String,
		}
	}

	return c.JSON(http.StatusOK, accountUsers)
}

// 別のアカウントにログインするためのやつ
// cookieを変えるだけなので、ブラウザをリロードする必要がある
func (h *Handler) AccountSwitchHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.FormValue("user_id")
	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "user_id is empty")
	}
	cookies := c.Cookies()

	_, _, err := h.Session.Login(ctx, cookies, true)
	if err != nil {
		return err
	}

	setCookies, err := h.Session.SwitchAccount(ctx, c.Cookies(), userId)
	if err != nil {
		return err
	}
	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return nil
}

// アカウントからログアウトする
func (h *Handler) AccountLogoutHandler(c echo.Context) error {
	ctx := c.Request().Context()

	cookies := c.Cookies()

	user, _, err := h.Session.Login(ctx, cookies, true)
	if err != nil {
		return err
	}

	setCookies, err := h.Session.Logout(ctx, cookies, user)
	if err != nil {
		return err
	}
	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return nil
}

// アカウントを削除する
func (h *Handler) AccountDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()

	cookies := c.Cookies()

	user, _, err := h.Session.Login(ctx, cookies, true)
	if err != nil {
		return err
	}

	// 色々削除する
	// TODO

	setCookies, err := h.Session.Logout(ctx, cookies, user)
	if err != nil {
		return err
	}
	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return nil
}

// OTPのpublic keyを返す
func (h *Handler) AccountOTPPublicKeyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, setCookies, err := h.Session.Login(ctx, c.Cookies())
	if err != nil {
		return err
	}

	otp, err := lib.NewOTP(h.C.OTPIssuer, user.UserName)
	if err != nil {
		return err
	}

	session, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	otpRegisterSession := models.RegisterOtpSession{
		ID:        session,
		UserID:    user.ID,
		PublicKey: otp.GetPublic(),
		Secret:    otp.GetSecret(),
		Period:    time.Now().Add(h.C.OTPRegisterSessionPeriod),
	}
	if err := otpRegisterSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, AccountOTPPublic{
		OTPSession: session,
		PublicKey:  otp.GetPublic(),
	})
}

// otpの設定
func (h *Handler) AccountOTPHandler(c echo.Context) error {
	ctx := c.Request().Context()

	otpSessionToken := c.FormValue("otp_session")
	if otpSessionToken == "" {
		return NewHTTPError(http.StatusBadRequest, "otp_session is empty")
	}
	code := c.FormValue("code")
	if code == "" {
		return NewHTTPError(http.StatusBadRequest, "code is empty")
	}

	user, setCookies, err := h.Session.Login(ctx, c.Cookies())
	if err != nil {
		return err
	}

	otpSession, err := models.RegisterOtpSessions(
		models.RegisterOtpSessionWhere.ID.EQ(otpSessionToken),
		qm.And("user_id = ?", user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "invalid otp_session")
	}
	if err != nil {
		return err
	}
	// セッションが有効期限切れの場合
	if time.Now().After(otpSession.Period) {
		// セッションは削除
		if _, err := otpSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}
	if otpSession.RetryCount >= h.C.OTPRegisterLimit {
		// セッションは削除
		if _, err := otpSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExceededRetry, "exceeded retry")
	}
	otpSession.RetryCount++

	if !lib.ValidateOTP(code, otpSession.Secret) {
		// リトライ回数を++するのでUPDATE
		if _, err := otpSession.Update(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		return NewHTTPError(http.StatusForbidden, "failed otp validate")
	}

	otp := models.Otp{
		UserID: user.ID,
		Secret: otpSession.Secret,
	}
	if err := otp.Upsert(ctx, h.DB, boil.Infer(), boil.Infer()); err != nil {
		return err
	}

	// バックアップコードを新規作成
	// すでにバックアップコードがある場合は一旦削除
	if _, err := models.OtpBackups(
		models.OtpBackupWhere.UserID.EQ(user.ID),
	).DeleteAll(ctx, h.DB); err != nil {
		return err
	}
	backups := make([]string, h.C.OTPBackupCount)
	for i := uint8(0); h.C.OTPBackupCount > i; i++ {
		code, err := lib.RandomStr(15)
		if err != nil {
			return err
		}
		backups[i] = code

		// FIXME: Bulk Insertを使いたいがSQLBoilerに実装が無いので後々検討したい
		backupDB := models.OtpBackup{
			UserID: user.ID,
			Code:   code,
		}
		if err := backupDB.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
	}

	// セッションは削除
	if _, err := otpSession.Delete(ctx, h.DB); err != nil {
		return err
	}

	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, backups)
}

// OTPのバックアップコードを返す
func (h *Handler) AccountOTPBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, setCookies, err := h.Session.Login(ctx, c.Cookies())
	if err != nil {
		return err
	}

	backups, err := models.OtpBackups(
		models.OtpBackupWhere.UserID.EQ(user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	backupCodes := make([]string, len(backups))
	for i, b := range backups {
		backupCodes[i] = b.Code
	}

	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, backupCodes)
}

func (h *Handler) AccountPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	newPassword := c.FormValue("new_password")
	if !lib.ValidatePassword(newPassword) {
		return NewHTTPError(http.StatusBadRequest, "invalid password")
	}

	user, setCookies, err := h.Session.Login(ctx, c.Cookies())
	if err != nil {
		return err
	}

	hash, salt, err := h.Password.HashPassword(newPassword)
	if err != nil {
		return err
	}

	password := models.Password{
		UserID: user.ID,
		Salt:   salt,
		Hash:   hash,
	}
	if err := password.Upsert(ctx, h.DB, boil.Infer(), boil.Infer()); err != nil {
		return err
	}

	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return nil
}

func (h *Handler) AccountBeginWebauthnHandler(c echo.Context) error {
	return nil
}

func (h *Handler) AccountWebauthnHandler(c echo.Context) error {
	return nil
}

// アカウントの認証情報を返す
// パスワードの設定可否、Passkeyの設定可否、OTPの設定可否
func (h *Handler) AccountCertificatesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, setCookies, err := h.Session.Login(ctx, c.Cookies())
	if err != nil {
		return err
	}

	password, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	otp, err := models.Otps(
		models.OtpWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	passkey, err := models.Passkeys(
		models.PasskeyWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, AccountCertificates{
		Password: password,
		OTP:      otp,
		Passkey:  passkey,
	})
}

// メールアドレスを指定して、パスワード再設定メールを送信する
// - Passkeyのみ設定している場合は使えない（パスワードを設定している場合のみ）
// - メールアドレスはアカウントが存在している必要がある
// - メールを送信するのでreCAPTCHA使う
// - 疲労攻撃の対策で、一度送信したらそのメールアドレスの再送に時間を持たせる
func (h *Handler) AccountForgetPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.FormValue("email")
	if !lib.ValidateEmail(email) {
		return NewHTTPError(http.StatusBadRequest, "invalid email")
	}
	recaptcha := c.FormValue("recaptcha")
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}

	ip := c.RealIP()

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

	user, err := models.Users(
		models.UserWhere.Email.EQ(email),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNotFoundUser, "user not found")
	}
	if err != nil {
		return err
	}

	// ユーザーがPasswordを登録しているか
	existPassword, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existPassword {
		return NewHTTPError(http.StatusBadRequest, "no registered password")
	}

	// すでにセッションが存在している
	existSession, err := models.ReregistrationPasswordSessions(
		qm.Where("email = ?", email),
		qm.And("period_clear > NOW()"),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if existSession {
		return NewHTTPError(http.StatusBadRequest, "already sessions")
	}

	token, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	session := models.ReregistrationPasswordSession{
		ID:          token,
		Email:       email,
		Period:      time.Now().Add(h.C.ReregistrationPasswordSessionPeriod),
		PeriodClear: time.Now().Add(h.C.ReregistrationPasswordSessionClearPeriod),
	}
	if err := session.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}

	m := &lib.MailBody{
		EmailAddress: email,
		Subject:      "パスワードを再設定してください",
		Data: AccountReRegisterPasswordTemplate{
			SessionToken: token,
			Now:          time.Now(),
			PeriodTime:   time.Now().Add(h.C.ReregistrationPasswordSessionPeriod),
			Email:        email,
		},
		PlainTextFileName: "forget_reregistration_password.gtpl",
		HTMLTextFileName:  "forget_reregistration_password.html",
	}

	msg, id, err := h.Sender.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", email),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", ip),
			zap.String("Device", ua.Device),
			zap.String("Browser", ua.Browser),
			zap.String("OS", ua.OS),
			zap.Bool("IsMobile", ua.IsMobile),
		)
		return err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("Email", email),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", ip),
		zap.String("Device", ua.Device),
		zap.String("Browser", ua.Browser),
		zap.String("OS", ua.OS),
		zap.Bool("IsMobile", ua.IsMobile),
	)

	return nil
}

// パスワード再設定
// BOT来ると困るのでreCAPTCHA使いながら
func (h *Handler) AccountReRegisterPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.FormValue("email")
	if !lib.ValidateEmail(email) {
		return NewHTTPError(http.StatusBadRequest, "invalid email")
	}
	recaptcha := c.FormValue("recaptcha")
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}
	token := c.FormValue("reregister_token")
	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "reregister_token is empty")
	}
	newPassword := c.FormValue("new_password")
	if !lib.ValidatePassword(newPassword) {
		return NewHTTPError(http.StatusBadRequest, "bad password")
	}

	ip := c.RealIP()

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

	session, err := models.ReregistrationPasswordSessions(
		qm.Where("id = ?", token),
		qm.And("completed = FALSE"),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "invalid token")
	}
	if err != nil {
		return err
	}
	// セッションが有効期限切れの場合
	if time.Now().After(session.Period) {
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}
	if session.Email != email {
		return NewHTTPError(http.StatusForbidden, "email is different")
	}

	hash, salt, err := h.Password.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// パスワードを更新する
	passwordDB, err := models.Passwords(
		qm.InnerJoin("user on user.id = password.id"),
		qm.Where("user.email = ?", session.Email),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}
	passwordDB.Hash = hash
	passwordDB.Salt = salt
	if _, err := passwordDB.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	// 使用済みフラグを立てる
	session.Completed = true
	if _, err := session.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}
