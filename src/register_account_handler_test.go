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