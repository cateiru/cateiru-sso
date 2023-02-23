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
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestSendEmailVerifyHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功する", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		form := contents.NewMultipart()
		form.Insert("email", email)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.SendEmailVerifyHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA token is empty")
	})

	t.Run("reCAPTCHAのチャレンジ失敗", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		s := createSession(email, 1)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		s := createSession(email, 1)

		// すでに5回ミスった
		s.RetryCount = 5
		_, err = s.Update(ctx, DB, boil.Infer())
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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		s := createSession(email, 0)

		// 有効期限 - 10日
		s.Period = s.Period.Add(-24 * 10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
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
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

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
