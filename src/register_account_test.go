package src_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func TestSendEmailVerifyHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功する", func(t *testing.T) {
		email := RandomEmail(t)

		form := contents.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.NoError(t, err)

		m.Ok(t)

		resp := &src.RegisterEmailResponse{}
		require.NoError(t, m.Json(resp))
		require.NotNil(t, resp.Token)

		s, err := models.RegisterSessions(
			models.RegisterSessionWhere.Email.EQ(email),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, s.ID, resp.Token)
		require.Equal(t, s.RetryCount, uint8(0))
		require.False(t, s.EmailVerified)
		require.NotNil(t, s.VerifyCode)
	})

	t.Run("Emailが不正な形式の場合エラー", func(t *testing.T) {
		email := "hogehoge124"

		form := contents.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.EqualError(t, err, "code=400, message=invalid email")
	})

	t.Run("reCAPTCHA tokenが無いとエラー", func(t *testing.T) {
		email := RandomEmail(t)

		form := contents.NewMultipart()
		form.Insert("email", email)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA token is empty")
	})

	t.Run("reCAPTCHAのチャレンジ失敗", func(t *testing.T) {
		email := RandomEmail(t)

		form := contents.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "fail")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})

	t.Run("すでにセッションテーブルにEmailが存在している場合はエラー", func(t *testing.T) {
		email := RandomEmail(t)

		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		// Emailのセッションを格納する
		sessionDB := models.RegisterSession{
			ID:         session,
			Email:      email,
			VerifyCode: "123456",

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		// アクセスする
		form := contents.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.EqualError(t, err, "code=400, message=session exists, unique=2")
	})

	t.Run("すでにメールアドレスが登録されている", func(t *testing.T) {
		email := RandomEmail(t)

		RegisterUser(t, ctx, email)

		// アクセスする
		form := contents.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.EqualError(t, err, "code=400, message=impossible register, unique=3")
	})
}

func TestReSendVerifyEmailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// セッションを作成する
	createSession := func(email string, sendCount uint8) *models.RegisterSession {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:         session,
			Email:      email,
			VerifyCode: "123456",
			SendCount:  sendCount,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		s, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		return s
	}

	t.Run("成功する", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 1)

		form := contents.NewMultipart()
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.NoError(t, err)

		m.Ok(t)

		resendSession, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(s.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.NotEqual(t, s.VerifyCode, resendSession.VerifyCode, "Codeが変わっている")
		require.Equal(t, resendSession.SendCount, uint8(2))

		require.False(t, resendSession.EmailVerified, "まだ認証は完了されてない")
	})

	t.Run("tokenが空だとエラー", func(t *testing.T) {
		form := contents.NewMultipart()
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=400, message=token is empty")
	})

	t.Run("recaptchaが空だとエラー", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 1)

		form := contents.NewMultipart()
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA token is empty")
	})

	t.Run("recaptchaのチャレンジが失敗", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 1)

		form := contents.NewMultipart()
		form.Insert("recaptcha", "fail")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})

	t.Run("tokenが不正", func(t *testing.T) {
		form := contents.NewMultipart()
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", "123")
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=400, message=token is invalid")
	})

	t.Run("tokenの有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 1)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err := s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("メールの送信上限を超えた", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 3)

		form := contents.NewMultipart()
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=429, message=email sending limit, unique=6")
	})

	t.Run("リトライ回数が上限を超えていた", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 1)

		// すでに5回ミスった
		s.RetryCount = 5
		_, err := s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("recaptcha", "123abc")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.ReSendVerifyEmailHandler(c)
		require.EqualError(t, err, "code=429, message=exceeded retry, unique=4")
	})
}

func TestRegisterVerifyEmailHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// セッションを作成する
	createSession := func(email string, retryCount uint8) *models.RegisterSession {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:         session,
			Email:      email,
			VerifyCode: "123456",
			RetryCount: retryCount,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		s, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		return s
	}

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 0)

		form := contents.NewMultipart()
		form.Insert("code", s.VerifyCode)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterVerifyEmailHandler(c)
		require.NoError(t, err)

		m.Ok(t)

		resp := &src.RegisterVerifyEmailResponse{}
		err = m.Json(resp)
		require.NoError(t, err)

		require.Equal(t, resp.Verified, true)
		require.Equal(t, resp.RemainingCount, uint8(4))

		resendSession, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(s.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.True(t, resendSession.EmailVerified)
		require.Equal(t, resendSession.RetryCount, uint8(1))
	})

	t.Run("tokenがない場合エラー", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 0)

		form := contents.NewMultipart()
		form.Insert("code", s.VerifyCode)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterVerifyEmailHandler(c)
		require.EqualError(t, err, "code=400, message=token is empty")
	})

	t.Run("tokenが不正", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 0)

		form := contents.NewMultipart()
		form.Insert("code", s.VerifyCode)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", "123")
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterVerifyEmailHandler(c)
		require.EqualError(t, err, "code=400, message=token is invalid")
	})

	t.Run("tokenの有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, 0)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err := s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("code", s.VerifyCode)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterVerifyEmailHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("リトライ回数上限", func(t *testing.T) {
		email := RandomEmail(t)

		// すでに5回リトライ済み
		s := createSession(email, 5)

		form := contents.NewMultipart()
		form.Insert("code", s.VerifyCode)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterVerifyEmailHandler(c)
		require.EqualError(t, err, "code=429, message=exceeded retry, unique=4")
	})
}

func TestRegisterBeginWebAuthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// セッションを作成する
	createSession := func(email string, verified bool) *models.RegisterSession {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:            session,
			Email:         email,
			EmailVerified: verified,
			VerifyCode:    "123456",
			RetryCount:    1,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		s, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		return s
	}

	t.Run("成功する", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)

		m, err := mock.NewMock("", http.MethodPost, "/")
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterBeginWebAuthnHandler(c)
		require.NoError(t, err)

		// response
		resp := new(protocol.CredentialCreation)
		err = m.Json(resp)
		require.NoError(t, err)

		require.Equal(t, resp.Response.User.Name, email)
		require.Equal(t, resp.Response.User.DisplayName, email)

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

		require.Equal(t, webauthnSession.Challenge, resp.Response.Challenge.String())
		require.Equal(t, protocol.URLEncodedBase64(webauthnSession.WebauthnUserID).String(), resp.Response.User.ID)
		require.Equal(t, webauthnSession.UserDisplayName, "") // 定義はされているのに何故か代入していない
		require.Equal(t, webauthnSession.Identifier, int8(1))

		// rowにjsonが入っている
		sessionFromRow := new(webauthn.SessionData)
		err = webauthnSession.Row.Unmarshal(sessionFromRow)
		require.NoError(t, err)

		require.Equal(t, sessionFromRow.Challenge, resp.Response.Challenge.String())
	})

	t.Run("tokenが無い", func(t *testing.T) {
		m, err := mock.NewMock("", http.MethodPost, "/")
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterBeginWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=token is empty")
	})

	t.Run("tokenが不正", func(t *testing.T) {
		m, err := mock.NewMock("", http.MethodPost, "/")
		m.R.Header.Add("X-Register-Token", "123")
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterBeginWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=token is invalid")

	})

	t.Run("tokenが有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err := s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodPost, "/")
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterBeginWebAuthnHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")

	})

	t.Run("まだEmailの認証が終わっていない", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		s := createSession(email, false)

		m, err := mock.NewMock("", http.MethodPost, "/")
		m.R.Header.Add("X-Register-Token", s.ID)
		require.NoError(t, err)
		c := m.Echo()

		err = h.RegisterBeginWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=Email is not verified, unique=7")
	})
}

func TestRegisterWebAuthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// セッションを作成する
	createSession := func(email string, verified bool) *models.RegisterSession {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:            session,
			Email:         email,
			EmailVerified: verified,
			VerifyCode:    "123456",
			RetryCount:    1,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		s, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		return s
	}

	// Webauthnのセッションを作成する
	registerWebauthnSession := func(email string, identifier int8) string {
		user, err := src.NewWebAuthnUserRegister(email)
		require.NoError(t, err)
		webauthnSessionId, err := lib.RandomStr(31)
		require.NoError(t, err)

		_, s, err := h.WebAuthn.BeginRegistration(user)
		require.NoError(t, err)

		row := types.JSON{}
		err = row.Marshal(s)
		require.NoError(t, err)

		webauthnSession := models.WebauthnSession{
			ID:               webauthnSessionId,
			WebauthnUserID:   s.UserID,
			UserDisplayName:  s.UserDisplayName,
			Challenge:        s.Challenge,
			UserVerification: string(s.UserVerification),
			Row:              row,

			Period:     time.Now().Add(h.C.WebAuthnSessionPeriod),
			Identifier: identifier,
		}
		err = webauthnSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		return webauthnSessionId
	}

	t.Run("成功する", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)
		webauthnSession := registerWebauthnSession(email, 1)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.NoError(t, err)

		// userが返ってきているか
		responseUser := new(models.User)
		require.NoError(t, m.Json(responseUser))
		require.NotNil(t, responseUser)

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

		require.Equal(t, sessionUser.ID, responseUser.ID)

		// passkeyが保存されているか
		existsPasskey, err := models.Passkeys(
			models.PasskeyWhere.UserID.EQ(responseUser.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existsPasskey)

		passkeyLoginDeviceCount, err := models.PasskeyLoginDevices(
			models.PasskeyLoginDeviceWhere.UserID.EQ(responseUser.ID),
		).Count(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, passkeyLoginDeviceCount, int64(1))

		// WebauthnSessionは削除されている
		existsWebauthnSession, err := models.WebauthnSessionExists(ctx, DB, webauthnSession)
		require.NoError(t, err)
		require.False(t, existsWebauthnSession)

		// registerSessionは削除されている
		existsRegisterSession, err := models.RegisterSessionExists(ctx, DB, s.ID)
		require.NoError(t, err)
		require.False(t, existsRegisterSession)
	})

	t.Run("失敗: X-Register-Tokenがない", func(t *testing.T) {
		email := RandomEmail(t)

		webauthnSession := registerWebauthnSession(email, 1)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=token is empty")
	})

	t.Run("失敗: webauthnTokenのCookieが無い", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=http: named cookie not present")
	})

	t.Run("失敗:  X-Register-Tokenが不正", func(t *testing.T) {
		email := RandomEmail(t)

		webauthnSession := registerWebauthnSession(email, 1)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", "hogehoge")
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=token is invalid")
	})

	t.Run("失敗: X-Register-Tokenがまだ認証完了していない", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, false) // 認証終わってない
		webauthnSession := registerWebauthnSession(email, 1)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=403, message=Email is not verified, unique=7")
	})

	t.Run("失敗: X-Register-Tokenの有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)
		webauthnSession := registerWebauthnSession(email, 1)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err := s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("失敗: webauthnTokenの値が不正", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: "hogehoge",
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=403, message=invalid webauthn token")
	})

	t.Run("失敗: webauthnTokenのIdentifierが違う", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)
		webauthnSession := registerWebauthnSession(email, 5)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=403, message=invalid webauthn token")
	})

	t.Run("失敗: webauthnTokenが有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)
		webauthnSession := registerWebauthnSession(email, 1)

		// 有効期限切れにする
		session, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(webauthnSession),
		).One(ctx, DB)
		require.NoError(t, err)
		session.Period = time.Now().Add(-10 * time.Hour)
		_, err = session.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := mock.NewJson("/", "", http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("application/json以外", func(t *testing.T) {
		email := RandomEmail(t)

		s := createSession(email, true)
		webauthnSession := registerWebauthnSession(email, 1)

		m, err := mock.NewMock("", http.MethodPost, "/")
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)
		cookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: webauthnSession,
		}
		m.Cookie([]*http.Cookie{cookie})

		c := m.Echo()

		err = h.RegisterWebAuthnHandler(c)
		require.EqualError(t, err, "code=400, message=invalid content-type")
	})
}

func TestRegisterPasswordHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	// セッションを作成する
	createSession := func(email string, verified bool) *models.RegisterSession {
		session, err := lib.RandomStr(31)
		require.NoError(t, err)

		sessionDB := models.RegisterSession{
			ID:            session,
			Email:         email,
			EmailVerified: verified,
			VerifyCode:    "123456",
			RetryCount:    1,

			Period: time.Now().Add(h.C.RegisterSessionPeriod),
		}
		err = sessionDB.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		s, err := models.RegisterSessions(
			models.RegisterSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		return s
	}

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		s := createSession(email, true)

		password := "password123456789"

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.NoError(t, err)

		// userが返ってきているか
		responseUser := &models.User{}
		require.NoError(t, m.Json(responseUser))
		require.NotNil(t, responseUser)

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

		require.Equal(t, sessionUser.ID, responseUser.ID)

		// パスワードは保存されている
		passwordModel, err := models.Passwords(
			models.PasswordWhere.UserID.EQ(responseUser.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		validated := C.Password.VerifyPassword(password, passwordModel.Hash, passwordModel.Salt)
		require.True(t, validated)

		// registerSessionは削除されている
		existsRegisterSession, err := models.RegisterSessionExists(ctx, DB, s.ID)
		require.NoError(t, err)
		require.False(t, existsRegisterSession)
	})

	t.Run("失敗: X-Register-Tokenが無い", func(t *testing.T) {
		password := "password123456789"

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=token is empty")
	})
	t.Run("失敗: X-Register-Tokenの値が不正", func(t *testing.T) {
		password := "password123456789"

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", "hogehoge123")

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=token is invalid")
	})
	t.Run("失敗: X-Register-Tokenがまだ認証完了していない", func(t *testing.T) {
		email := RandomEmail(t)
		s := createSession(email, false)

		password := "password123456789"

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=Email is not verified, unique=7")
	})
	t.Run("失敗: X-Register-Tokenの有効期限が切れている", func(t *testing.T) {
		email := RandomEmail(t)
		s := createSession(email, true)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err := s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		password := "password123456789"

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("失敗: パスワードが無い", func(t *testing.T) {
		email := RandomEmail(t)
		s := createSession(email, true)

		password := ""

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=password is empty")
	})
	t.Run("失敗: パスワードが範囲外", func(t *testing.T) {
		email := RandomEmail(t)
		s := createSession(email, true)

		password := "あああああああああああああああああああああああああああああああああああああああああああああああああああああ"

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.R.Header.Add("X-Register-Token", s.ID)

		c := m.Echo()

		err = h.RegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=bad password")
	})
}