package src_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func TestLoginUserHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: Email", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &user)

		form := easy.NewMultipart()
		form.Insert("username_or_email", email)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
	})
	t.Run("成功: ユーザ名", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &user)

		form := easy.NewMultipart()
		form.Insert("username_or_email", user.UserName)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
	})
	t.Run("成功: アバターあり", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		user.Avatar = null.NewString("https://example.com/avatar", true)
		_, err := user.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		RegisterPassword(t, ctx, &user)

		form := easy.NewMultipart()
		form.Insert("username_or_email", email)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.Equal(t, response.Avatar.String, "https://example.com/avatar")
	})

	t.Run("失敗: username_or_emailが空", func(t *testing.T) {
		form := easy.NewMultipart()
		form.Insert("username_or_email", "")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		userData := &src.UserData{
			Device:   "",
			OS:       "Windows",
			Browser:  "Brave", // 登録時はChrome
			IsMobile: false,
		}
		SetUserData(t, m, userData)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.EqualError(t, err, "code=400, message=username_or_email is empty")
	})
	t.Run("失敗: username_or_emailの値が不正", func(t *testing.T) {
		form := easy.NewMultipart()
		form.Insert("username_or_email", "aaaa")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		userData := &src.UserData{
			Device:   "",
			OS:       "Windows",
			Browser:  "Brave", // 登録時はChrome
			IsMobile: false,
		}
		SetUserData(t, m, userData)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.EqualError(t, err, "code=404, message=user not found, unique=10")
	})
}

func TestLoginBeginWebauthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginBeginWebauthnHandler(c)
		require.NoError(t, err)

		// response
		resp := new(protocol.CredentialCreation)
		require.NoError(t, m.Json(resp))

		// cookie
		cookies := m.Response().Cookies()
		sessionCookie := new(http.Cookie)
		for _, cookie := range cookies {
			if cookie.Name == C.WebAuthnSessionCookie.Name {
				sessionCookie = cookie
			}
		}
		require.NotNil(t, sessionCookie)

		webauthnSession, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(sessionCookie.Value),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, webauthnSession.UserID.Valid)
		require.Equal(t, webauthnSession.Identifier, int8(2))

		// rowにjsonが入っている
		sessionFromRow := new(webauthn.SessionData)
		err = webauthnSession.Row.Unmarshal(sessionFromRow)
		require.NoError(t, err)

		require.Equal(t, sessionFromRow.Challenge, resp.Response.Challenge.String())
	})
}

func TestLoginWebauthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// Webauthnのセッションを作成する
	registerWebauthnSession := func(identifier int8) string {
		webauthnSessionId, err := lib.RandomStr(31)
		require.NoError(t, err)

		_, s, err := h.WebAuthn.BeginLogin()
		require.NoError(t, err)

		row := types.JSON{}
		err = row.Marshal(s)
		require.NoError(t, err)

		webauthnSession := models.WebauthnSession{
			ID:  webauthnSessionId,
			Row: row,

			Period:     time.Now().Add(h.C.WebAuthnSessionPeriod),
			Identifier: identifier,
		}
		err = webauthnSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return webauthnSessionId
	}

	t.Run("成功", func(t *testing.T) {
		webauthnSession := registerWebauthnSession(2)

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})
		m.R.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36`)
		c := m.Echo()

		err = h.LoginWebauthnHandler(c)
		require.NoError(t, err)

		// userが返ってきているか
		response := new(src.LoginResponse)
		require.NoError(t, m.Json(response))
		require.NotNil(t, response)

		// cookieは設定されているか（セッショントークンのみ見る）
		var sessionCookie *http.Cookie = nil
		for _, cookie := range m.Response().Cookies() {
			if cookie.Name == C.SessionCookie.Name {
				sessionCookie = cookie
				break
			}
		}
		require.NotNil(t, sessionCookie)

		// セッションはBodyのユーザと同じか
		sessionUser, err := models.Users(
			qm.InnerJoin("session on session.user_id = user.id"),
			qm.Where("session.id = ?", sessionCookie.Value),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, sessionUser.ID, response.User.ID)

		// ログイントライ履歴が保存されている
		existsLoginTryHistory, err := models.LoginTryHistories(
			models.LoginTryHistoryWhere.UserID.EQ(sessionUser.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existsLoginTryHistory)

		// WebauthnSessionは削除されている
		existsWebauthnSession, err := models.WebauthnSessionExists(ctx, DB, webauthnSession)
		require.NoError(t, err)
		require.False(t, existsWebauthnSession)
	})

	t.Run("失敗: application/jsonじゃない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &u)
		webauthnSession := registerWebauthnSession(2)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})
		c := m.Echo()

		err = h.LoginWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=invalid content-type")
	})

	t.Run("失敗: セッションが空", func(t *testing.T) {
		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=http: named cookie not present")
	})

	t.Run("失敗: セッショントークンが不正", func(t *testing.T) {
		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: "hogehoge",
		}
		m.Cookie([]*http.Cookie{cookie})
		c := m.Echo()

		err = h.LoginWebauthnHandler(c)
		require.EqualError(t, err, "code=403, message=invalid webauthn token")
	})

	t.Run("失敗: identifierが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &u)
		webauthnSession := registerWebauthnSession(4)

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})
		c := m.Echo()

		err = h.LoginWebauthnHandler(c)
		require.EqualError(t, err, "code=403, message=invalid webauthn token")
	})

	t.Run("すでにログイン済み", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		TestWebAuthnUser = &u

		cookies := RegisterSession(t, ctx, &u)

		webauthnSession := registerWebauthnSession(2)

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})
		m.Cookie(cookies)
		m.R.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36`)
		c := m.Echo()

		err = h.LoginWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=already logged in")
	})
}

func TestLoginPasswordHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: パスワードのみ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u, "password123ABC123123")

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "password123ABC123123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.NoError(t, err)

		// userが返ってきているか
		response := new(src.LoginResponse)
		require.NoError(t, m.Json(response))
		require.NotNil(t, response)

		// cookieは設定されているか（セッショントークンのみ見る）
		var sessionCookie *http.Cookie = nil
		for _, cookie := range m.Response().Cookies() {
			if cookie.Name == C.SessionCookie.Name {
				sessionCookie = cookie
				break
			}
		}
		require.NotNil(t, sessionCookie)

		// セッションはBodyのユーザと同じか
		sessionUser, err := models.Users(
			qm.InnerJoin("session on session.user_id = user.id"),
			qm.Where("session.id = ?", sessionCookie.Value),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, sessionUser.ID, response.User.ID)

		// ログイントライ履歴が保存されている
		existsLoginTryHistory, err := models.LoginTryHistories(
			models.LoginTryHistoryWhere.UserID.EQ(u.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existsLoginTryHistory)
	})

	t.Run("成功: OTPが設定されている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u, "password123ABC123123")
		RegisterOTP(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "password123ABC123123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.NoError(t, err)

		response := new(src.LoginResponse)
		require.NoError(t, m.Json(response))
		require.NotNil(t, response)

		require.Nil(t, response.User)

		existOtpSession, err := models.OtpSessions(
			models.OtpSessionWhere.ID.EQ(response.OTP.Token),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existOtpSession)

		require.Equal(t, response.OTP.LoginUser.Avatar, u.Avatar)
	})

	t.Run("失敗: パスワードが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=password is empty")
	})

	t.Run("失敗: パスワードが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u, "password123ABC123123")

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "aaaaa")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.EqualError(t, err, "code=403, message=invalid password, unique=8")
	})

	t.Run("失敗: reCAPTCHA失敗", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u, "password123ABC123123")

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "fail")
		form.Insert("password", "password123ABC123123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})

	t.Run("失敗: パスワードを設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "password123ABC123123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=password not registered, unique=8")
	})

	t.Run("失敗: ユーザーが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u, "password123ABC123123")

		form := easy.NewMultipart()
		form.Insert("username_or_email", "hogehoge")
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "password123ABC123123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.EqualError(t, err, "code=404, message=user not found, unique=10")
	})

	t.Run("失敗: すでにログインしている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u, "password123ABC123123")

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("password", "password123ABC123123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.LoginPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=already logged in")
	})
}

func TestLoginOTPHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerOTPSession := func(u *models.User) string {
		otpSessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)
		otpSession := models.OtpSession{
			ID:     otpSessionToken,
			UserID: u.ID,

			Period: time.Now().Add(C.OTPSessionPeriod),
		}
		err = otpSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		return otpSessionToken
	}

	t.Run("成功: OTPコード", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		secret, _ := RegisterOTP(t, ctx, &u)
		code, err := totp.GenerateCode(secret, time.Now())
		require.NoError(t, err)

		otpSession := registerOTPSession(&u)

		form := easy.NewMultipart()
		form.Insert("otp_session", otpSession)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.NoError(t, err)

		// userが返ってきているか
		response := new(src.LoginResponse)
		require.NoError(t, m.Json(response))
		require.NotNil(t, response)

		// cookieは設定されているか（セッショントークンのみ見る）
		var sessionCookie *http.Cookie = nil
		for _, cookie := range m.Response().Cookies() {
			if cookie.Name == C.SessionCookie.Name {
				sessionCookie = cookie
				break
			}
		}
		require.NotNil(t, sessionCookie)

		// セッションはBodyのユーザと同じか
		sessionUser, err := models.Users(
			qm.InnerJoin("session on session.user_id = user.id"),
			qm.Where("session.id = ?", sessionCookie.Value),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, sessionUser.ID, response.User.ID)

		// OTPSessionは削除されている
		existsOtpSession, err := models.OtpSessionExists(ctx, DB, otpSession)
		require.NoError(t, err)
		require.False(t, existsOtpSession)
	})

	t.Run("成功: バックアップ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		_, backups := RegisterOTP(t, ctx, &u)
		backup := backups[0]

		otpSession := registerOTPSession(&u)

		form := easy.NewMultipart()
		form.Insert("otp_session", otpSession)
		form.Insert("code", backup)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.NoError(t, err)

		// userが返ってきているか
		response := new(src.LoginResponse)
		require.NoError(t, m.Json(response))
		require.NotNil(t, response)

		// cookieは設定されているか（セッショントークンのみ見る）
		var sessionCookie *http.Cookie = nil
		for _, cookie := range m.Response().Cookies() {
			if cookie.Name == C.SessionCookie.Name {
				sessionCookie = cookie
				break
			}
		}
		require.NotNil(t, sessionCookie)

		// セッションはBodyのユーザと同じか
		sessionUser, err := models.Users(
			qm.InnerJoin("session on session.user_id = user.id"),
			qm.Where("session.id = ?", sessionCookie.Value),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, sessionUser.ID, response.User.ID)

		// OTPSessionは削除されている
		existsOtpSession, err := models.OtpSessionExists(ctx, DB, otpSession)
		require.NoError(t, err)
		require.False(t, existsOtpSession)

		// backupは1度使用したら削除される
		existOtpBackup, err := models.OtpBackups(
			models.OtpBackupWhere.Code.EQ(backup),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existOtpBackup)
	})

	t.Run("失敗: OTPが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOTP(t, ctx, &u)

		otpSession := registerOTPSession(&u)

		form := easy.NewMultipart()
		form.Insert("otp_session", otpSession)
		form.Insert("code", "123456")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// retry_count++
		s, err := models.OtpSessions(
			models.OtpSessionWhere.ID.EQ(otpSession),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, s.RetryCount, uint8(1))
	})

	t.Run("失敗: OTPが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterOTP(t, ctx, &u)

		otpSession := registerOTPSession(&u)

		form := easy.NewMultipart()
		form.Insert("otp_session", otpSession)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.EqualError(t, err, "code=400, message=code is empty")
	})

	t.Run("失敗: OTPのセッションが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		secret, _ := RegisterOTP(t, ctx, &u)
		code, err := totp.GenerateCode(secret, time.Now())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.EqualError(t, err, "code=400, message=otp_session is empty")
	})

	t.Run("失敗: OTPのセッションが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		secret, _ := RegisterOTP(t, ctx, &u)
		code, err := totp.GenerateCode(secret, time.Now())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("otp_session", "hogehoge")
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.EqualError(t, err, "code=400, message=invalid otp session")
	})

	t.Run("失敗: OTPセッションの有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		secret, _ := RegisterOTP(t, ctx, &u)
		code, err := totp.GenerateCode(secret, time.Now())
		require.NoError(t, err)

		otpSession := registerOTPSession(&u)

		// 有効期限 - 10日
		s, err := models.OtpSessions(
			models.OtpSessionWhere.ID.EQ(otpSession),
		).One(ctx, DB)
		require.NoError(t, err)
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("otp_session", otpSession)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("失敗: OTPセッションのリトライ上限超えた", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		secret, _ := RegisterOTP(t, ctx, &u)
		code, err := totp.GenerateCode(secret, time.Now())
		require.NoError(t, err)

		otpSession := registerOTPSession(&u)

		// リトライ回数を上限にする
		s, err := models.OtpSessions(
			models.OtpSessionWhere.ID.EQ(otpSession),
		).One(ctx, DB)
		require.NoError(t, err)
		s.RetryCount = h.C.OTPRetryLimit
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("otp_session", otpSession)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginOTPHandler(c)
		require.EqualError(t, err, "code=403, message=exceeded retry, unique=4")
	})
}
