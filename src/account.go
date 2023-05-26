package src

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
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

	OTP         bool      `json:"otp"`
	OtpModified null.Time `json:"otp_modified"`
}

type AccountReRegisterPasswordIsSession struct {
	Active bool `json:"active"`
}

type AccountWebauthnDevice struct {
	ID uint64 `json:"id"`

	Device   null.String `json:"device,omitempty"`
	Os       null.String `json:"os,omitempty"`
	Browser  null.String `json:"browser,omitempty"`
	IsMobile null.Bool   `json:"is_mobile,omitempty"`
	IP       string      `json:"ip"`

	Created time.Time `json:"created"`
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
// `login_history_id`を指定するとそのHistory IDと繋がっているセッションを遠隔で削除する
func (h *Handler) AccountLogoutHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c, true)
	if err != nil {
		return err
	}

	// `login_history_id`がある場合は、そのログイン履歴を削除する
	loginHistoryId := c.FormValue("login_history_id")
	if loginHistoryId != "" {
		loginHistoryIdInt, err := strconv.ParseUint(loginHistoryId, 10, 64)
		if err != nil {
			return NewHTTPError(http.StatusBadRequest, "invalid type login_history_id")
		}

		loginHistory, err := models.LoginHistories(
			models.LoginHistoryWhere.ID.EQ(uint(loginHistoryIdInt)),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return NewHTTPError(http.StatusBadRequest, "invalid login_history_id")
		}
		if err != nil {
			return err
		}

		// 自分のリフレッシュトークンを取得する
		myRefreshCookieName := fmt.Sprintf("%s-%s", h.C.RefreshCookie.Name, user.ID)
		var myRefreshCookie http.Cookie
		for _, cookie := range c.Cookies() {
			if cookie.Name == myRefreshCookieName {
				myRefreshCookie = *cookie
			}
		}

		refresh, err := models.Refreshes(
			models.RefreshWhere.HistoryID.EQ(loginHistory.RefreshID),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return NewHTTPError(http.StatusBadRequest, "invalid login_history_id")
		}
		if err != nil {
			return err
		}

		// 自分のリフレッシュトークンは削除できない
		if refresh.ID == myRefreshCookie.Value {
			return NewHTTPError(http.StatusBadRequest, "cannot logout myself")
		}

		if refresh.SessionID.Valid {
			// リフレッシュトークンに紐づくセッショントークンを削除（ある場合）
			_, err = models.Sessions(
				models.SessionWhere.ID.EQ(refresh.SessionID.String),
			).DeleteAll(ctx, h.DB)
			if err != nil {
				return err
			}
		}

		// リフレッシュトークンを削除
		_, err = models.Refreshes(
			models.RefreshWhere.HistoryID.EQ(loginHistory.RefreshID),
		).DeleteAll(ctx, h.DB)
		if err != nil {
			return err
		}

		return nil
	}

	setCookies, err := h.Session.Logout(ctx, c.Cookies(), user)
	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}
	if err != nil {
		return err
	}

	return nil
}

// アカウントを削除する
func (h *Handler) AccountDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()

	cookies := c.Cookies()

	user, err := h.Session.SimpleLogin(ctx, c, true)
	if err != nil {
		return err
	}

	// TODO: 色々削除する

	setCookies, err := h.Session.Logout(ctx, cookies, user)
	for _, cookie := range setCookies {
		c.SetCookie(cookie)
	}
	if err != nil {
		return err
	}

	return nil
}

// OTPのpublic keyを返す
func (h *Handler) AccountOTPPublicKeyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	existPassword, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existPassword {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNoRegisteredPassword, "no registered password")
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

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	existPassword, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existPassword {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNoRegisteredPassword, "no registered password")
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

	return c.JSON(http.StatusOK, backups)
}

// OTPを削除する
func (h *Handler) AccountDeleteOTPHandler(c echo.Context) error {
	ctx := c.Request().Context()

	password := c.FormValue("password")
	if password == "" {
		return NewHTTPError(http.StatusBadRequest, "password is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// パスワードの検証
	p, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNoRegisteredPassword, "no registered password")
	}
	if !h.Password.VerifyPassword(password, p.Hash, p.Salt) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrLoginFailed, "failed password")
	}

	existOtp, err := models.Otps(
		models.OtpWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existOtp {
		return NewHTTPError(http.StatusBadRequest, "otp is not registered")
	}

	// OTP削除
	if _, err := models.Otps(
		models.OtpWhere.UserID.EQ(user.ID),
	).DeleteAll(ctx, h.DB); err != nil {
		return err
	}
	// OTPのバックアップ削除
	if _, err := models.OtpBackups(
		models.OtpBackupWhere.UserID.EQ(user.ID),
	).DeleteAll(ctx, h.DB); err != nil {
		return err
	}

	return nil
}

// OTPのバックアップコードを返す
func (h *Handler) AccountOTPBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	password := c.FormValue("password")
	if password == "" {
		return NewHTTPError(http.StatusBadRequest, "password is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// パスワードの検証
	p, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNoRegisteredPassword, "no registered password")
	}
	if !h.Password.VerifyPassword(password, p.Hash, p.Salt) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrLoginFailed, "failed password")
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

	return c.JSON(http.StatusOK, backupCodes)
}

func (h *Handler) AccountPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	newPassword := c.FormValue("new_password")
	if !lib.ValidatePassword(newPassword) {
		return NewHTTPError(http.StatusBadRequest, "invalid password")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	hash, salt, err := h.Password.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// パスワード更新は、現在のパスワードの検証も行うのでこのハンドラでは新規
	// 作成のみを受付る
	existPassword, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if existPassword {
		return NewHTTPError(http.StatusBadRequest, "password is already exists")
	}

	password := models.Password{
		UserID: user.ID,
		Salt:   salt,
		Hash:   hash,
	}
	if err := password.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// パスワード更新
func (h *Handler) AccountUpdatePasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	newPassword := c.FormValue("new_password")
	if !lib.ValidatePassword(newPassword) {
		return NewHTTPError(http.StatusBadRequest, "invalid password")
	}
	oldPassword := c.FormValue("old_password")
	if oldPassword == "" {
		return NewHTTPError(http.StatusBadRequest, "password is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// パスワードの検証
	p, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNoRegisteredPassword, "no registered password")
	}
	if !h.Password.VerifyPassword(oldPassword, p.Hash, p.Salt) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrLoginFailed, "failed password")
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

	return nil
}

func (h *Handler) AccountBeginWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	webauthnUser, err := NewWebAuthnUserFromDB(ctx, h.DB, user)
	if err != nil {
		return err
	}
	creation, s, err := h.WebAuthn.BeginRegistration(webauthnUser)
	if err != nil {
		return err
	}

	row := types.JSON{}
	if err = row.Marshal(s); err != nil {
		return err
	}
	webauthnSessionId, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	webauthnSession := models.WebauthnSession{
		ID:     webauthnSessionId,
		UserID: null.NewString(user.ID, true),

		Row: row,

		Period:     time.Now().Add(h.C.WebAuthnSessionPeriod),
		Identifier: 3,
	}
	if err := webauthnSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     h.C.WebAuthnSessionCookie.Name,
		Value:    webauthnSessionId,
		Path:     h.C.WebAuthnSessionCookie.Path,
		Domain:   h.C.SiteHost.Host,
		Secure:   h.C.WebAuthnSessionCookie.Secure,
		HttpOnly: h.C.WebAuthnSessionCookie.HttpOnly,
		MaxAge:   h.C.WebAuthnSessionCookie.MaxAge,
		SameSite: h.C.WebAuthnSessionCookie.SameSite,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, creation)
}

// Webuathnを登録しているデバイスの一覧を返す
func (h *Handler) AccountWebauthnRegisteredDevicesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	devices, err := models.Webauthns(
		models.WebauthnWhere.UserID.EQ(user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	d := make([]AccountWebauthnDevice, len(devices))
	for i, device := range devices {
		d[i] = AccountWebauthnDevice{
			ID: device.ID,

			Device:   device.Device,
			Os:       device.Os,
			Browser:  device.Browser,
			IsMobile: device.IsMobile,
			IP:       net.IP.To16(device.IP).String(),

			Created: device.Created,
		}
	}

	return c.JSON(http.StatusOK, d)
}

// Passkeyの新規追加
// 本当は、更新時にはすでに登録しているPasskeyを求めたいけどフローが複雑になるので後々
func (h *Handler) AccountWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if c.Request().Header.Get("Content-Type") != "application/json" {
		return NewHTTPError(http.StatusBadRequest, "invalid content-type")
	}

	webauthnToken, err := c.Cookie(h.C.WebAuthnSessionCookie.Name)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "session is empty")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	credential, _, err := h.RegisterWebauthn(ctx, c.Request().Body, webauthnToken.Value, 3)
	if err != nil {
		return err
	}

	rowCredential := types.JSON{}
	if err := rowCredential.Marshal(credential); err != nil {
		return err
	}

	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}
	ip := c.RealIP()

	auth := models.Webauthn{
		UserID:     user.ID,
		Credential: rowCredential,

		Device:   null.NewString(ua.Device, true),
		Os:       null.NewString(ua.OS, true),
		Browser:  null.NewString(ua.Browser, true),
		IsMobile: null.NewBool(ua.IsMobile, true),

		IP: net.ParseIP(ip),
	}
	if err := auth.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (h *Handler) AccountDeleteWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	webauthnId := c.QueryParam("webauthn_id")
	if webauthnId == "" {
		return NewHTTPError(http.StatusBadRequest, "webauthn_id is empty")
	}
	parsedWebauthnId, err := strconv.Atoi(webauthnId)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "webauthn_id is invalid")
	}

	user, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// パスワードが設定されていない場合、WebAuthnはすべて削除することができない
	hasSetPassword, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !hasSetPassword {
		webauthnCount, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(user.ID),
		).Count(ctx, h.DB)
		if err != nil {
			return err
		}

		if webauthnCount <= 1 {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrNoMoreAuthentication, "webauthn must be set at least one")
		}
	}

	_, err = models.Webauthns(
		models.WebauthnWhere.UserID.EQ(user.ID),
		models.WebauthnWhere.ID.EQ(uint64(parsedWebauthnId)),
	).DeleteAll(ctx, h.DB)
	if err != nil {
		return err
	}

	return nil
}

// アカウントの認証情報を返す
// パスワードの設定可否、Passkeyの設定可否、OTPの設定可否
func (h *Handler) AccountCertificatesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.Session.SimpleLogin(ctx, c)
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
	).One(ctx, h.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	otpModified := null.NewTime(time.Time{}, false)
	if otp != nil {
		otpModified = null.TimeFrom(otp.Modified)
	}

	return c.JSON(http.StatusOK, AccountCertificates{
		Password:    password,
		OTP:         otp != nil,
		OtpModified: otpModified,
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

	// すでにセッションが存在している
	reRegisterSession, err := models.ReregistrationPasswordSessions(
		models.ReregistrationPasswordSessionWhere.Email.EQ(email),
	).One(ctx, h.DB)
	// レコードが無い以外のエラーは500にする
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	} else if err == nil {
		// period_clearが有効期限過ぎていたら削除してそのまま続ける
		if time.Now().After(reRegisterSession.PeriodClear) {
			if _, err := reRegisterSession.Delete(ctx, h.DB); err != nil {
				return err
			}
		} else {
			return NewHTTPError(http.StatusBadRequest, "already sessions")
		}
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

	// ログイントライ履歴を保存する
	loginTryHistory := models.LoginTryHistory{
		UserID: user.ID,

		Device:   null.NewString(ua.Device, true),
		Os:       null.NewString(ua.OS, true),
		Browser:  null.NewString(ua.Browser, true),
		IsMobile: null.NewBool(ua.IsMobile, true),

		IP: net.ParseIP(ip),

		Identifier: 1, // パスワード再登録なので1
	}
	if err := loginTryHistory.Insert(ctx, h.DB, boil.Infer()); err != nil {
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

// パスワード再設定のトークンが有効化どうか
func (h *Handler) AccountReRegisterAvailableTokenHandler(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.FormValue("email")
	if !lib.ValidateEmail(email) {
		return NewHTTPError(http.StatusBadRequest, "invalid email")
	}
	token := c.FormValue("reregister_token")
	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "reregister_token is empty")
	}

	session, err := models.ReregistrationPasswordSessions(
		qm.Where("id = ?", token),
		qm.And("completed = FALSE"),
		qm.And("email = ?", email),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return c.JSON(http.StatusOK, &AccountReRegisterPasswordIsSession{
			Active: false,
		})
	}
	if err != nil {
		return err
	}
	// PeriodClear時間超えていたらDBのほうは消す
	if time.Now().After(session.PeriodClear) {
		if _, err := session.Delete(ctx, h.DB); err != nil {
			return err
		}
	}

	// セッションが有効期限切れの場合
	if time.Now().After(session.Period) {
		return c.JSON(http.StatusOK, &AccountReRegisterPasswordIsSession{
			Active: false,
		})
	}
	if session.Email != email {
		return c.JSON(http.StatusOK, &AccountReRegisterPasswordIsSession{
			Active: false,
		})
	}

	return c.JSON(http.StatusOK, &AccountReRegisterPasswordIsSession{
		Active: true,
	})
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

	// ユーザーを引く
	user, err := models.Users(
		models.UserWhere.Email.EQ(email),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "invalid email")
	}
	if err != nil {
		return err
	}

	passwordDB, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if passwordDB == nil {
		// 新たにパスワードを作成する
		passwordDB = &models.Password{
			UserID: user.ID,
			Hash:   hash,
			Salt:   salt,
		}
		if err := passwordDB.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
	} else {
		// パスワードを更新する
		passwordDB.Hash = hash
		passwordDB.Salt = salt
		if _, err := passwordDB.Update(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
	}

	// 使用済みフラグを立てる
	session.Completed = true
	if _, err := session.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}
