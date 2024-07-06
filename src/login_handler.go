package src

import (
	"context"
	"database/sql"
	"errors"
	"net"
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
	Avatar null.String `json:"avatar"`
}

type OTP struct {
	Token     string     `json:"token"`
	LoginUser *LoginUser `json:"login_user"`
}

type LoginResponse struct {
	User *UserMeResponse `json:"user,omitempty"`
	OTP  *OTP            `json:"otp,omitempty"`
}

// ユーザの情報を返す
// BOT使われると困るのでreCAPTCHA使いながら
func (h *Handler) LoginUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userNameOrEmail := c.FormValue("username_or_email")

	if userNameOrEmail == "" {
		return NewHTTPError(http.StatusBadRequest, "username_or_email is empty")
	}

	user, err := FindUserByUserNameOrEmail(ctx, h.DB, userNameOrEmail)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &LoginUser{
		Avatar: user.Avatar,
	})
}

// Passkeyでログインするために、Challengeなどを返す
func (h *Handler) LoginBeginWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	webauthnSessionId, err := lib.RandomStr(31)
	if err != nil {
		return err
	}

	// ユーザーは指定しない
	creation, s, err := h.WebAuthn.BeginLogin()
	if err != nil {
		return err
	}

	row := types.JSON{}
	if err = row.Marshal(s); err != nil {
		return err
	}

	webauthnSession := models.WebauthnSession{
		ID:  webauthnSessionId,
		Row: row,

		Period:     time.Now().Add(h.C.WebAuthnSessionPeriod),
		Identifier: 2,
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

// WebAuthnでログインする
// WebauthnSessionCookieは有効期限が短いので削除しないが、DBのセッションは削除する
func (h *Handler) LoginWebauthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if c.Request().Header.Get("Content-Type") != "application/json" {
		return NewHTTPError(http.StatusBadRequest, "invalid content-type")
	}
	oauthLoginSessionToken := c.Request().Header.Get("X-Oauth-Login-Session")

	webauthnToken, err := c.Cookie(h.C.WebAuthnSessionCookie.Name)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	ip := c.RealIP()
	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}

	user, err := h.LoginWebauthn(ctx, c.Request().Body, webauthnToken.Value, 2)
	if err != nil {
		return err
	}

	var loginTryHistoryIdentifier int8 = 0
	if oauthLoginSessionToken != "" {
		loginTryHistoryIdentifier = 2
	}

	// ログイントライ履歴を追加する
	loginTryHistory := models.LoginTryHistory{
		UserID:   user.ID,
		Device:   null.NewString(ua.Device, true),
		Os:       null.NewString(ua.OS, true),
		Browser:  null.NewString(ua.Browser, true),
		IsMobile: null.NewBool(ua.IsMobile, true),
		IP:       net.ParseIP(ip),

		Identifier: loginTryHistoryIdentifier,
	}
	if err := loginTryHistory.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	// X-Oauth-Login-Session が設定されている場合、そのセッションの `login_ok = true` にする
	if oauthLoginSessionToken != "" {
		if err := SetLoggedInOauthLoginSession(ctx, h.DB, oauthLoginSessionToken); err != nil {
			return err
		}

		// セッションを作り直したくないのでそのまま返してしまう
		userMe, err := h.LoginNoRegisterSession(ctx, c, user.ID)
		if err != nil {
			return err
		}
		if userMe != nil {
			// ref. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja#login-status-api
			c.Response().Header().Set("Set-Login", "logged-in")

			return c.JSON(http.StatusOK, LoginResponse{
				User: userMe,
				OTP:  nil,
			})
		}
	} else {
		if ok := h.Session.IsLoggedIn(ctx, c.Cookies(), user); ok {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrAlreadyLoggedIn, "already logged in")
		}
	}

	session, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	for _, cookie := range session.InsertCookie(h.C) {
		c.SetCookie(cookie)
	}

	isStaff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	JoinedOrganization, err := models.OrganizationUsers(
		models.OrganizationUserWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	userMe := UserMeResponse{
		UserInfo:           user,
		IsStaff:            isStaff,
		JoinedOrganization: JoinedOrganization,
	}

	// ref. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja#login-status-api
	c.Response().Header().Set("Set-Login", "logged-in")

	return c.JSON(http.StatusOK, LoginResponse{
		User: &userMe,
		OTP:  nil,
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
	oauthLoginSessionToken := c.Request().Header.Get("X-Oauth-Login-Session")

	ip := c.RealIP()
	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
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
	if err != nil {
		return err
	}

	var loginTryHistoryIdentifier int8 = 0
	if oauthLoginSessionToken != "" {
		loginTryHistoryIdentifier = 2
	}

	// ログイントライ履歴を保存する
	loginTryHistory := models.LoginTryHistory{
		UserID:   user.ID,
		Device:   null.NewString(ua.Device, true),
		Os:       null.NewString(ua.OS, true),
		Browser:  null.NewString(ua.Browser, true),
		IsMobile: null.NewBool(ua.IsMobile, true),
		IP:       net.ParseIP(ip),

		Identifier: loginTryHistoryIdentifier,
	}
	if err := loginTryHistory.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	p, err := models.Passwords(
		models.PasswordWhere.UserID.EQ(user.ID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrLoginFailed, "password not registered")
	}
	if err != nil {
		return err
	}

	if !h.Password.VerifyPassword(password, p.Hash, p.Salt) {
		return NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "invalid password")
	}

	// X-Oauth-Login-Session が設定されている場合、そのセッションの `login_ok = true` にする
	if oauthLoginSessionToken != "" {
		if err := SetLoggedInOauthLoginSession(ctx, h.DB, oauthLoginSessionToken); err != nil {
			return err
		}

		// セッションを作り直したくないのでそのまま返してしまう
		userMe, err := h.LoginNoRegisterSession(ctx, c, user.ID)
		if err != nil {
			return err
		}
		if userMe != nil {
			// ref. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja#login-status-api
			c.Response().Header().Set("Set-Login", "logged-in")

			return c.JSON(http.StatusOK, LoginResponse{
				User: userMe,
				OTP:  nil, // FIXME: OTPも使えるようにする
			})
		}
	} else {
		if ok := h.Session.IsLoggedIn(ctx, c.Cookies(), user); ok {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrAlreadyLoggedIn, "already logged in")
		}
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
			return err
		}

		return c.JSON(http.StatusOK, LoginResponse{
			OTP: &OTP{
				Token: otpSessionToken,
				LoginUser: &LoginUser{
					Avatar: user.Avatar,
				},
			},
			User: nil,
		})
	}

	// OTPが設定されていない場合はそのままログイン
	session, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	for _, cookie := range session.InsertCookie(h.C) {
		c.SetCookie(cookie)
	}

	isStaff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	JoinedOrganization, err := models.OrganizationUsers(
		models.OrganizationUserWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	userMe := UserMeResponse{
		UserInfo:           user,
		IsStaff:            isStaff,
		JoinedOrganization: JoinedOrganization,
	}

	// ref. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja#login-status-api
	c.Response().Header().Set("Set-Login", "logged-in")

	return c.JSON(http.StatusOK, LoginResponse{
		User: &userMe,
		OTP:  nil,
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
		result = lib.ValidateOTP(code, otp.Secret)
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
		// retry_countを更新するのでUPDATE
		if _, err := session.Update(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrLoginFailed, "login failed")
	}

	// セッションは削除する
	if _, err := session.Delete(ctx, h.DB); err != nil {
		return err
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(session.UserID),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	ip := c.RealIP()
	ua, err := h.ParseUA(c.Request())
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

	// X-Oauth-Login-Session が設定されている場合、そのセッションの `login_ok = true` にする
	oauthLoginSessionToken := c.Request().Header.Get("X-Oauth-Login-Session")
	if oauthLoginSessionToken != "" {
		if err := SetLoggedInOauthLoginSession(ctx, h.DB, oauthLoginSessionToken); err != nil {
			return err
		}
	}

	isStaff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	JoinedOrganization, err := models.OrganizationUsers(
		models.OrganizationUserWhere.UserID.EQ(user.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}

	userMe := UserMeResponse{
		UserInfo:           user,
		IsStaff:            isStaff,
		JoinedOrganization: JoinedOrganization,
	}

	// ref. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja#login-status-api
	c.Response().Header().Set("Set-Login", "logged-in")

	return c.JSON(http.StatusOK, LoginResponse{
		User: &userMe,
		OTP:  nil,
	})
}

func (h *Handler) LoginNoRegisterSession(ctx context.Context, c echo.Context, userId string) (*UserMeResponse, error) {
	sessionUser, err := h.Session.SimpleLogin(ctx, c, true)
	if errors.Is(err, ErrorLoginFailed) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// ユーザーが違う = 違うユーザーでログインしている
	if sessionUser.ID != userId {
		return nil, nil
	}

	isStaff, err := models.Staffs(
		models.StaffWhere.UserID.EQ(sessionUser.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return nil, err
	}

	JoinedOrganization, err := models.OrganizationUsers(
		models.OrganizationUserWhere.UserID.EQ(sessionUser.ID),
	).Exists(ctx, h.DB)
	if err != nil {
		return nil, err
	}

	userMe := &UserMeResponse{
		UserInfo:           sessionUser,
		IsStaff:            isStaff,
		JoinedOrganization: JoinedOrganization,
	}

	return userMe, nil
}
