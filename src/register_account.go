package src

import (
	"database/sql"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
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

type RegisterVerifyEmailResponse struct {
	RemainingCount uint8 `json:"remaining_count"`
	Verified       bool  `json:"verified"`
}

// 最初にメールアドレス宛に確認コードを送信する
// アカウント作成フローの一番はじめ
// Emailを送信するのでreCAPTCHA使う
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
	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.Email.EQ(email),
	).One(ctx, h.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if registerSession != nil {
		if registerSession.Period.After(time.Now()) {
			return NewHTTPUniqueError(http.StatusBadRequest, ErrSessionExists, "session exists")
		}
		// 有効期限が切れているので削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}
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

	userData, err := h.ParseUA(c.Request())
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
		Data:              GenerateEmailData(r, h.C),
		PlainTextFileName: "register.gtpl",
		HTMLTextFileName:  "register.html",
	}
	msg, id, err := h.Sender.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", email),
			zap.String("Subject", m.Subject),
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
		zap.String("Subject", m.Subject),
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

// 確認コードを再送する
// 再送すると、確認コードは別のものに変更される
// Emailを送信するのでreCAPTCHA使う
func (h *Handler) ReSendVerifyEmailHandler(c echo.Context) error {
	ctx := c.Request().Context()

	token := c.Request().Header.Get("X-Register-Token") // SendEmailVerifyHandlerのレスポンスToken

	recaptcha := c.FormValue("recaptcha")
	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "token is empty")
	}
	if recaptcha == "" {
		return NewHTTPError(http.StatusBadRequest, "reCAPTCHA token is empty")
	}

	userData, err := h.ParseUA(c.Request())
	if err != nil {
		return err
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

	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "token is invalid")
	}
	if err != nil {
		return err
	}

	// 有効期限が切れた場合
	if time.Now().After(registerSession.Period) {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	// メール送信上限を超えた場合
	// 認証はできるのでセッションのレコードは削除しない
	if registerSession.SendCount >= h.C.RegisterEmailSendLimit {
		return NewHTTPUniqueError(http.StatusTooManyRequests, ErrEmailSendingLimit, "email sending limit")
	}

	// リトライ回数が指定回数を超えた場合、失敗させる
	// ブルートフォースアタック対策
	// 普通は失敗しないよね
	if registerSession.RetryCount >= h.C.RegisterSessionRetryLimit {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}

		// スパムだった場合を考えてログを出す
		L.Info("exceeded retry",
			zap.String("Email", registerSession.Email),
			zap.String("IP", ip),
			zap.String("Device", userData.Device),
			zap.String("Browser", userData.Browser),
			zap.String("OS", userData.OS),
			zap.Bool("IsMobile", userData.IsMobile),
		)

		return NewHTTPUniqueError(http.StatusTooManyRequests, ErrExceededRetry, "exceeded retry")
	}

	// codeを更新させる
	code, err := lib.RandomNumber(6)
	if err != nil {
		return err
	}
	registerSession.VerifyCode = code

	registerSession.SendCount++

	// 対象のメールアドレスにメールを送信する
	r := RegisterEmailVerify{
		Code:     code,
		Email:    registerSession.Email,
		Time:     time.Now(),
		UserData: userData,
	}
	m := &lib.MailBody{
		EmailAddress:      registerSession.Email,
		Subject:           "【再送】メールアドレスの登録確認",
		Data:              r,
		PlainTextFileName: "register.gtpl",
		HTMLTextFileName:  "register.html",
	}
	msg, id, err := h.Sender.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", registerSession.Email),
			zap.String("Subject", m.Subject),
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
		zap.String("Email", registerSession.Email),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", ip),
		zap.String("Device", userData.Device),
		zap.String("Browser", userData.Browser),
		zap.String("OS", userData.OS),
		zap.Bool("IsMobile", userData.IsMobile),
	)

	if _, err := registerSession.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// メールアドレスに送られた確認コードを入力してEmailを認証させるハンドラー
func (h *Handler) RegisterVerifyEmailHandler(c echo.Context) error {
	ctx := c.Request().Context()

	token := c.Request().Header.Get("X-Register-Token") // SendEmailVerifyHandlerのレスポンスToken
	code := c.FormValue("code")

	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "token is empty")
	}

	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "token is invalid")
	}
	if err != nil {
		return err
	}

	// 有効期限が切れた場合
	if time.Now().After(registerSession.Period) {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	// リトライ回数が指定回数を超えた場合、失敗させる
	// ブルートフォースアタック対策
	// 普通は失敗しないよね
	if registerSession.RetryCount >= h.C.RegisterSessionRetryLimit {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}

		ua, err := h.ParseUA(c.Request())
		if err != nil {
			return err
		}
		ip := c.RealIP()

		// スパムだった場合を考えてログを出す
		L.Info("exceeded retry",
			zap.String("Email", registerSession.Email),
			zap.String("IP", ip),
			zap.String("Device", ua.Device),
			zap.String("Browser", ua.Browser),
			zap.String("OS", ua.OS),
			zap.Bool("IsMobile", ua.IsMobile),
		)

		return NewHTTPUniqueError(http.StatusTooManyRequests, ErrExceededRetry, "exceeded retry")
	}

	registerSession.RetryCount++
	verify := registerSession.VerifyCode == code

	// 確認コードが正しい場合はOK
	if verify {
		registerSession.EmailVerified = true
	}

	if _, err := registerSession.Update(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	resp := &RegisterVerifyEmailResponse{
		RemainingCount: h.C.RegisterSessionRetryLimit - registerSession.RetryCount,
		Verified:       verify,
	}

	return c.JSON(http.StatusOK, resp)
}

// Passkeyを登録するために、Challengeなどを返す
func (h *Handler) RegisterBeginWebAuthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	token := c.Request().Header.Get("X-Register-Token") // SendEmailVerifyHandlerのレスポンスToken

	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "token is empty")
	}

	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "token is invalid")
	}
	if err != nil {
		return err
	}

	// まだ認証されていない場合
	if !registerSession.EmailVerified {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrEmailNotVerified, "Email is not verified")
	}

	// 有効期限が切れた場合
	if time.Now().After(registerSession.Period) {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	// リトライ回数や送信回数は認証されていたら用済みなので見ない
	id := ulid.Make().String()
	user, err := NewWebAuthnUserRegister(registerSession.Email, []byte(id))
	if err != nil {
		return err
	}
	webauthnSessionId, err := lib.RandomStr(31)
	if err != nil {
		return err
	}

	creation, s, err := h.WebAuthn.BeginRegistration(user)
	if err != nil {
		return err
	}

	row := types.JSON{}
	if err = row.Marshal(s); err != nil {
		return err
	}

	webauthnSession := models.WebauthnSession{
		ID:     webauthnSessionId,
		UserID: null.NewString(id, true),
		Row:    row,

		Period:     time.Now().Add(h.C.WebAuthnSessionPeriod),
		Identifier: 1,
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

// Passkeyによる認証の登録
// 事前にRegisterBeginWebAuthnを呼び出してtokenをcookieに付与させる必要がある
func (h *Handler) RegisterWebAuthnHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if c.Request().Header.Get("Content-Type") != "application/json" {
		return NewHTTPError(http.StatusBadRequest, "invalid content-type")
	}

	webauthnToken, err := c.Cookie(h.C.WebAuthnSessionCookie.Name)
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, err)
	}

	token := c.Request().Header.Get("X-Register-Token") // SendEmailVerifyHandlerのレスポンスToken
	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "token is empty")
	}

	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "token is invalid")
	}
	if err != nil {
		return err
	}

	// まだ認証されていない場合
	if !registerSession.EmailVerified {
		return NewHTTPUniqueError(http.StatusForbidden, ErrEmailNotVerified, "Email is not verified")
	}

	// 有効期限が切れた場合
	if time.Now().After(registerSession.Period) {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	credential, webauthnUser, err := h.RegisterWebauthn(ctx, c.Request().Body, webauthnToken.Value, 1)
	if err != nil {
		return err
	}

	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}
	ip := c.RealIP()

	var user *models.User
	joinedOrganization := false

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		// 登録フロー
		user, err = RegisterUser(ctx, h.DB, registerSession.Email, string(webauthnUser.WebAuthnID()))
		if err != nil {
			return err
		}

		// 認証を追加
		rowCredential := types.JSON{}
		if err := rowCredential.Marshal(credential); err != nil {
			return err
		}
		passkey := models.Webauthn{
			UserID:     user.ID,
			Credential: rowCredential,

			Device:   null.NewString(ua.Device, true),
			Os:       null.NewString(ua.OS, true),
			Browser:  null.NewString(ua.Browser, true),
			IsMobile: null.NewBool(ua.IsMobile, true),

			IP: net.ParseIP(ip),
		}
		if err := passkey.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}

		// registerSessionは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}

		// org_idが設定されている場合は、orgに所属させる
		if registerSession.OrgID.Valid {
			orgExist, err := models.Organizations(
				models.OrganizationWhere.ID.EQ(registerSession.OrgID.String),
			).Exists(ctx, tx)
			if err != nil {
				return err
			}
			// orgが存在している場合のみ招待する
			if orgExist {
				orgUser := models.OrganizationUser{
					OrganizationID: registerSession.OrgID.String,
					UserID:         user.ID,

					Role: "guest", // ゲストで登録する
				}
				if err := orgUser.Insert(ctx, tx, boil.Infer()); err != nil {
					return err
				}
				joinedOrganization = true
			} else {
				L.Error("org not found",
					zap.String("org_id", registerSession.OrgID.String),
					zap.String("email", registerSession.Email),
					zap.String("user_id", user.ID),
				)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	if user == nil {
		return NewHTTPError(http.StatusInternalServerError, "user is nil")
	}

	register, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	cookies := register.InsertCookie(h.C)
	for _, cookie := range cookies {
		c.SetCookie(cookie)
	}

	response := UserMeResponse{
		UserInfo:           user,
		Setting:            nil,
		IsStaff:            false, // 登録時はtrueになることはない
		JoinedOrganization: joinedOrganization,
	}

	return c.JSON(http.StatusCreated, &response)
}

// パスワードによる認証の登録
func (h *Handler) RegisterPasswordHandler(c echo.Context) error {
	ctx := c.Request().Context()

	password := c.FormValue("password")
	if password == "" {
		return NewHTTPError(http.StatusBadRequest, "password is empty")
	}
	if !lib.ValidatePassword(password) {
		return NewHTTPError(http.StatusBadRequest, "bad password")
	}

	token := c.Request().Header.Get("X-Register-Token") // SendEmailVerifyHandlerのレスポンスToken
	if token == "" {
		return NewHTTPError(http.StatusBadRequest, "token is empty")
	}

	registerSession, err := models.RegisterSessions(
		models.RegisterSessionWhere.ID.EQ(token),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "token is invalid")
	}
	if err != nil {
		return err
	}

	// まだ認証されていない場合
	if !registerSession.EmailVerified {
		return NewHTTPUniqueError(http.StatusBadRequest, ErrEmailNotVerified, "Email is not verified")
	}

	// 有効期限が切れた場合
	if time.Now().After(registerSession.Period) {
		// セッションは削除する
		if _, err := registerSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	hashedPassword, salt, err := h.Password.HashPassword(password)
	if err != nil {
		return err
	}

	var user *models.User
	joinedOrganization := false

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		user, err = RegisterUser(ctx, tx, registerSession.Email)
		if err != nil {
			return err
		}

		passwordModel := models.Password{
			UserID: user.ID,
			Salt:   salt,
			Hash:   hashedPassword,
		}
		if err := passwordModel.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		// registerSessionは削除する
		if _, err := registerSession.Delete(ctx, tx); err != nil {
			return err
		}

		// org_idが設定されている場合は、orgに所属させる
		if registerSession.OrgID.Valid {
			orgExist, err := models.Organizations(
				models.OrganizationWhere.ID.EQ(registerSession.OrgID.String),
			).Exists(ctx, tx)
			if err != nil {
				return err
			}
			// orgが存在している場合のみ招待する
			if orgExist {
				orgUser := models.OrganizationUser{
					OrganizationID: registerSession.OrgID.String,
					UserID:         user.ID,

					Role: "guest", // ゲストで登録する
				}
				if err := orgUser.Insert(ctx, tx, boil.Infer()); err != nil {
					return err
				}

				joinedOrganization = true
			} else {
				L.Error("org not found",
					zap.String("org_id", registerSession.OrgID.String),
					zap.String("email", registerSession.Email),
					zap.String("user_id", user.ID),
				)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	if user == nil {
		return NewHTTPError(http.StatusInternalServerError, "user is nil")
	}

	ua, err := h.ParseUA(c.Request())
	if err != nil {
		return err
	}
	ip := c.RealIP()
	register, err := h.Session.NewRegisterSession(ctx, user, ua, ip)
	if err != nil {
		return err
	}
	cookies := register.InsertCookie(h.C)
	for _, cookie := range cookies {
		c.SetCookie(cookie)
	}

	response := UserMeResponse{
		UserInfo:           user,
		Setting:            nil,
		IsStaff:            false, // 登録時はtrueになることはない
		JoinedOrganization: joinedOrganization,
	}

	return c.JSON(http.StatusCreated, &response)
}

// 招待メールからregister_sessionを作成する
func (h *Handler) RegisterInviteRegisterSession(c echo.Context) error {
	ctx := c.Request().Context()

	inviteToken := c.FormValue("invite_token")
	if inviteToken == "" {
		return NewHTTPError(http.StatusBadRequest, "invite_token is empty")
	}
	email := c.FormValue("email")
	if email == "" {
		return NewHTTPError(http.StatusBadRequest, "email is empty")
	}

	inviteOrgSession, err := models.InviteOrgSessions(
		models.InviteOrgSessionWhere.Token.EQ(inviteToken),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "invalid invite_token")
	}
	if err != nil {
		return err
	}

	if email != inviteOrgSession.Email {
		return NewHTTPError(http.StatusBadRequest, "invalid email")
	}

	// セッションが有効期限切れの場合
	if time.Now().After(inviteOrgSession.Period) {
		// セッションは削除
		if _, err := inviteOrgSession.Delete(ctx, h.DB); err != nil {
			return err
		}
		return NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	session, err := lib.RandomStr(31)
	if err != nil {
		return err
	}

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		// トークンを使用した招待ではRegisterSessionの有効期限が切れていなくても
		// 作成可能とする
		registerSession, err := models.RegisterSessions(
			models.RegisterSessionWhere.Email.EQ(email),
		).One(ctx, h.DB)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		if registerSession != nil {
			if _, err := registerSession.Delete(ctx, h.DB); err != nil {
				return err
			}
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

		// orgの存在確認をする
		orgExists, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(inviteOrgSession.OrgID),
		).Exists(ctx, h.DB)
		if err != nil {
			return err
		}
		if !orgExists {
			return NewHTTPError(http.StatusBadRequest, "invalid invite_token")
		}

		newRegisterSession := &models.RegisterSession{
			ID: session,

			Email:         inviteOrgSession.Email,
			EmailVerified: true,     // 対象のメールアドレスに送信するため、確認済みとする
			VerifyCode:    "000000", // 使用しないので

			OrgID: null.NewString(inviteOrgSession.OrgID, true),

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		if err := newRegisterSession.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		// inviteOrgSessionは削除する
		if _, err := inviteOrgSession.Delete(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	response := &RegisterEmailResponse{
		Token: session,
	}

	return c.JSON(http.StatusOK, response)
}
