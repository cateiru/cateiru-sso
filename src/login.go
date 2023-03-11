package src

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type LoginUser struct {
	Avatar            null.String `json:"avatar"`
	UserName          string      `json:"user_name"`
	AvailablePasskey  bool        `json:"available_passkey"`
	AvailablePassword bool        `json:"available_password"`
	AutoUsePasskey    bool        `json:"auto_use_passkey"`
}

type LoginResponse struct {
	User *models.User `json:"user,omitempty"`
	OTP  string       `json:"otp,omitempty"`
}

// ユーザの情報を返す
// BOT使われると困るのでreCAPTCHA使いながら
func (h *Handler) LoginUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userNameOrEmail := c.FormValue("username_or_email")
	recaptcha := c.FormValue("recaptcha")
	ip := c.RealIP()

	if userNameOrEmail == "" {
		return NewHTTPError(http.StatusBadRequest, "username_or_email is empty")
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

	user, err := FindUserByUserNameOrEmail(ctx, h.DB, userNameOrEmail)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNotFoundUser, "user not found")
	}
	if err != nil {
		return err
	}

	availablePasskey := false
	autoUsePasskey := false
	availablePassword := false

	// Passkeyの判定
	ua, err := ParseUA(c.Request())
	if err != nil {
		return err
	}
	passkeyLoginDevices, err := models.PasskeyLoginDevices(
		models.PasskeyLoginDeviceWhere.UserID.EQ(user.ID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	if len(passkeyLoginDevices) >= 1 {
		availablePasskey = true

		passkey, err := models.Passkeys(
			models.PasskeyWhere.UserID.EQ(user.ID),
		).One(ctx, h.DB)
		if err != nil {
			return err
		}

		if ua.Browser != "" && ua.OS != "" {
			for _, devices := range passkeyLoginDevices {
				// Passkeyを登録したOSと同じOSであれば自動ログイン
				// iCloudなどOSで共有可能な場合があるので
				if devices.IsRegisterDevice {
					if devices.Os.String == ua.OS {
						autoUsePasskey = true
						break
					}
				}
				// 過去にログインしたことあるブラウザ
				if devices.Os.String == ua.OS &&
					devices.Browser.String == ua.Browser &&
					devices.Device.String == ua.Device &&
					(!devices.IsMobile.Valid || devices.IsMobile.Bool == ua.IsMobile) { // IsMobileがある場合はそれも使用して判定する
					autoUsePasskey = true
					break
				}

				// BackupStateがtrueの場合（iCloudなどで共有されている場合）は
				// OSで判定する
				if passkey.FlagBackupState && lib.ValidateOS(devices.Os.String, ua.OS) {
					autoUsePasskey = true
					break
				}
			}
		}
	}

	// パスワードの設定
	passwordExists, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	availablePassword = passwordExists

	// PasskeyかPasswordかならずどちらかは存在するはず
	if !availablePasskey && !availablePassword {
		return NewHTTPError(http.StatusInternalServerError, "no certificate")
	}

	loginUser := &LoginUser{
		Avatar:            user.Avatar,
		UserName:          user.UserName,
		AvailablePasskey:  availablePasskey,
		AutoUsePasskey:    autoUsePasskey,
		AvailablePassword: availablePassword,
	}
	return c.JSON(http.StatusOK, loginUser)
}

// Passkeyでログインするために、Challengeなどを返す
func (h *Handler) LoginBeginWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userNameOrEmail := c.FormValue("username_or_email")
	recaptcha := c.FormValue("recaptcha")
	ip := c.RealIP()

	if userNameOrEmail == "" {
		return NewHTTPError(http.StatusBadRequest, "username_or_email is empty")
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

	user, err := FindUserByUserNameOrEmail(ctx, h.DB, userNameOrEmail)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNotFoundUser, "user not found")
	}
	if err != nil {
		return err
	}

	webauthnUser, err := NewWebAuthnUserFromDB(ctx, h.DB, user)
	if err != nil {
		return err
	}
	webauthnSessionId, err := lib.RandomStr(31)
	if err != nil {
		return err
	}

	creation, s, err := h.WebAuthn.BeginLogin(webauthnUser)
	if err != nil {
		return err
	}

	row := types.JSON{}
	if err = row.Marshal(s); err != nil {
		return err
	}

	webauthnSession := models.WebauthnSession{
		ID:               webauthnSessionId,
		UserID:           null.NewString(user.ID, true),
		WebauthnUserID:   s.UserID,
		UserDisplayName:  s.UserDisplayName,
		Challenge:        s.Challenge,
		UserVerification: string(s.UserVerification),
		Row:              row,

		Period: time.Now().Add(h.C.WebAuthnSessionPeriod),
	}
	if err := webauthnSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     h.C.WebAuthnSessionCookie.Name,
		Value:    webauthnSessionId,
		Path:     h.C.WebAuthnSessionCookie.Path,
		Domain:   h.C.Host.Host,
		Secure:   h.C.WebAuthnSessionCookie.Secure,
		HttpOnly: h.C.WebAuthnSessionCookie.HttpOnly,
		MaxAge:   h.C.WebAuthnSessionCookie.MaxAge,
		SameSite: h.C.WebAuthnSessionCookie.SameSite,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, creation)
}

// Passkeyでログインする
// WebauthnSessionCookieは有効期限が短いので削除しないが、DBのセッションは削除する
func (h *Handler) LoginWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if c.Request().Header.Get("Content-Type") != "application/json" {
		return NewHTTPError(http.StatusBadRequest, "invalid content-type")
	}

	webauthnToken, err := c.Cookie(h.C.WebAuthnSessionCookie.Name)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := h.LoginWebauthn(ctx, c.Request().Body, webauthnToken.Value)
	if err != nil {
		return err
	}

	ip := c.RealIP()
	ua, err := ParseUA(c.Request())
	if err != nil {
		return err
	}

	session, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	for _, cookie := range session.InsertCookie(h.C) {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, LoginResponse{
		User: user,
	})
}

// パスワードでログインする
// OTPを登録している場合はログインさせずに、OTPセッションを返します
// BOT使われると困るのでreCAPTCHA使いながら
func (h *Handler) LoginPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userNameOrEmail := c.FormValue("username_or_email")
	if userNameOrEmail == "" {
		return NewHTTPError(http.StatusBadRequest, "username_or_email is empty")
	}
	recaptcha := c.FormValue("recaptcha")
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}
	password := c.FormValue("password")
	if password == "" {
		return NewHTTPError(http.StatusBadRequest, "password is empty")
	}
	if !lib.ValidatePassword(password) {
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

	user, err := FindUserByUserNameOrEmail(ctx, h.DB, userNameOrEmail)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrNotFoundUser, "user not found")
	}
	if err != nil {
		return err
	}

	p, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "password not registered")
	}
	if err != nil {
		return err
	}

	if !h.Password.VerifyPassword(password, p.Hash, p.Salt) {
		return NewHTTPError(http.StatusForbidden, "invalid password")
	}

	otpRegistered, err := models.Otps(
		models.OtpWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	// OTPが設定されている場合
	// LoginOTPHandlerで再度OTPの認証を行うため、ここではログインさせないでOTPのセッションのみ作成して
	// そのまま返す
	if otpRegistered {
		otpSessionToken, err := lib.RandomStr(31)
		if err != nil {
			return err
		}
		otpSession := models.OtpSession{
			ID:     otpSessionToken,
			UserID: user.ID,

			Period: time.Now().Add(h.C.OTPSessionPeriod),
		}
		if err := otpSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return nil
		}

		return c.JSON(http.StatusOK, LoginResponse{
			OTP: otpSessionToken,
		})
	}

	// OTPが設定されていない場合はそのままログイン
	ua, err := ParseUA(c.Request())
	if err != nil {
		return err
	}

	session, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	for _, cookie := range session.InsertCookie(h.C) {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, LoginResponse{
		User: user,
	})
}

// OTPでログインする
func (h *Handler) LoginOTPHandler(c echo.Context) error {
	ctx := c.Request().Context()

	OTPSession := c.FormValue("otp_session")
	if OTPSession == "" {
		return NewHTTPError(http.StatusBadRequest, "otp_session is empty")
	}
	code := c.FormValue("code")
	if code == "" {
		return NewHTTPError(http.StatusBadRequest, "code is empty")
	}

	session, err := models.OtpSessions(
		models.OtpSessionWhere.ID.EQ(OTPSession),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "invalid otp session")
	}
	if err != nil {
		return err
	}

	if time.Now().After(session.Period) {
		// セッションは削除
		_, err := session.Delete(ctx, h.DB)
		if err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	if session.RetryCount >= h.C.OTPRetryLimit {
		// セッションは削除
		_, err := session.Delete(ctx, h.DB)
		if err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExceededRetry, "exceeded retry")
	}
	session.RetryCount++

	otp, err := models.Otps(
		models.OtpWhere.UserID.EQ(session.UserID),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	result := false
	if lib.ValidateOTPCode(code) {
		result = lib.ValidateOTP(code, otp.Secret.String)
	} else {
		// Backupからログインを試みる
		backups, err := models.OtpBackups(
			models.OtpBackupWhere.UserID.EQ(otp.UserID),
		).All(ctx, h.DB)
		if err != nil {
			return err
		}
		for _, backup := range backups {
			if backup.Code == code {
				// バックアップは1度使用したら削除する
				if _, err := backup.Delete(ctx, h.DB); err != nil {
					return err
				}
				result = true
				break
			}
		}
	}
	if !result {
		return NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "login failed")
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(session.UserID),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	ip := c.RealIP()
	ua, err := ParseUA(c.Request())
	if err != nil {
		return err
	}

	registerSession, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	for _, cookie := range registerSession.InsertCookie(h.C) {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, LoginResponse{
		User: user,
	})
}
