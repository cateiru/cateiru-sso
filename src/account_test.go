package src_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/stretchr/testify/require"
)

func TestAccountListHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功", func(t *testing.T) {
		email1 := RandomEmail(t)
		email2 := RandomEmail(t)
		email3 := RandomEmail(t)

		u1 := RegisterUser(t, ctx, email1)
		u2 := RegisterUser(t, ctx, email2)
		u3 := RegisterUser(t, ctx, email3)

		cookies := RegisterSession(t, ctx, &u1, &u2, &u3)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountListHandler(c)
		require.NoError(t, err)

		accountUsers := []src.AccountUser{}
		err = m.Json(&accountUsers)
		require.NoError(t, err)

		require.Len(t, accountUsers, 3)

		u1Exist := false
		u2Exist := false
		u3Exist := false
		for _, u := range accountUsers {
			switch u.ID {
			case u1.ID:
				u1Exist = true
			case u2.ID:
				u2Exist = true
			case u3.ID:
				u3Exist = true
			}
		}
		require.True(t, u1Exist && u2Exist && u3Exist)
	})

	t.Run("Cookieに何もない場合は空", func(t *testing.T) {
		m, err := mock.NewGet("", "/")
		require.NoError(t, err)
		c := m.Echo()

		err = h.AccountListHandler(c)
		require.NoError(t, err)

		accountUsers := []src.AccountUser{}
		err = m.Json(&accountUsers)
		require.NoError(t, err)

		require.Len(t, accountUsers, 0)
	})

	t.Run("Cookieのトークンが不正な場合は空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID)
		cookie := &http.Cookie{
			Name:     refreshCookieName,
			Secure:   C.RefreshCookie.Secure,
			HttpOnly: C.RefreshCookie.HttpOnly,
			Path:     C.RefreshCookie.Path,
			MaxAge:   C.RefreshCookie.MaxAge,
			Expires:  time.Now().Add(time.Duration(C.RefreshCookie.MaxAge) * time.Second),
			SameSite: C.RefreshCookie.SameSite,

			Value: "aaaaa",
		}

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)
		m.Cookie([]*http.Cookie{
			cookie,
		})
		c := m.Echo()

		err = h.AccountListHandler(c)
		require.NoError(t, err)

		accountUsers := []src.AccountUser{}
		err = m.Json(&accountUsers)
		require.NoError(t, err)

		require.Len(t, accountUsers, 0)
	})
}

func TestAccountSwitchHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("アカウントを変えられる", func(t *testing.T) {
		email1 := RandomEmail(t)
		email2 := RandomEmail(t)
		u1 := RegisterUser(t, ctx, email1)
		u2 := RegisterUser(t, ctx, email2)

		cookies := RegisterSession(t, ctx, &u1, &u2)

		form := contents.NewMultipart()
		form.Insert("user_id", u2.ID)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountSwitchHandler(c)
		require.NoError(t, err)

		// LoginUserCookieが変わっている
		responseCookies := m.Response().Cookies()
		var loggedInUserCookie *http.Cookie = nil
		for _, c := range responseCookies {
			if c.Name == C.LoginUserCookie.Name {
				loggedInUserCookie = c
			}
		}
		require.NotNil(t, loggedInUserCookie)
		require.Equal(t, loggedInUserCookie.Value, u2.ID)
	})

	t.Run("ログインしているアカウントが存在しないと変えられない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		form := contents.NewMultipart()
		form.Insert("user_id", u.ID)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.AccountSwitchHandler(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
	})

	t.Run("ログインしているアカウントが1つだけの場合は変わらない", func(t *testing.T) {
		email1 := RandomEmail(t)
		u1 := RegisterUser(t, ctx, email1)

		cookies := RegisterSession(t, ctx, &u1)

		form := contents.NewMultipart()
		form.Insert("user_id", u1.ID)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountSwitchHandler(c)
		require.NoError(t, err)
	})

	t.Run("不正なIDを指定した", func(t *testing.T) {
		email1 := RandomEmail(t)
		u1 := RegisterUser(t, ctx, email1)

		cookies := RegisterSession(t, ctx, &u1)

		form := contents.NewMultipart()
		form.Insert("user_id", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountSwitchHandler(c)
		require.EqualError(t, err, "code=400, message=user not found")
	})

	t.Run("ログインしていいないIDを指定した", func(t *testing.T) {
		email1 := RandomEmail(t)
		email2 := RandomEmail(t)
		u1 := RegisterUser(t, ctx, email1)
		u2 := RegisterUser(t, ctx, email2)

		cookies := RegisterSession(t, ctx, &u1)

		form := contents.NewMultipart()
		form.Insert("user_id", u2.ID)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountSwitchHandler(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
	})
}

func TestAccountLogoutHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountLogoutHandler, func(u *models.User) *mock.MockHandler {
		m, err := mock.NewMock("", http.MethodHead, "/")
		require.NoError(t, err)
		return m
	})

	t.Run("ログアウトできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		session := RegisterSession(t, ctx, &u)

		m, err := mock.NewMock("", http.MethodHead, "/")
		require.NoError(t, err)
		m.Cookie(session)
		c := m.Echo()

		err = h.AccountLogoutHandler(c)
		require.NoError(t, err)

		// すべてのCookieが削除されている
		cookies := m.Response().Cookies()
		for _, cookie := range cookies {
			require.Equal(t, cookie.MaxAge, -1)
		}
	})
}

func TestAccountDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountDeleteHandler, func(u *models.User) *mock.MockHandler {
		m, err := mock.NewMock("", http.MethodHead, "/")
		require.NoError(t, err)
		return m
	})

	t.Run("削除されている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		session := RegisterSession(t, ctx, &u)

		m, err := mock.NewMock("", http.MethodHead, "/")
		require.NoError(t, err)
		m.Cookie(session)
		c := m.Echo()

		err = h.AccountLogoutHandler(c)
		require.NoError(t, err)

		// すべてのCookieが削除されている
		cookies := m.Response().Cookies()
		for _, cookie := range cookies {
			require.Equal(t, cookie.MaxAge, -1)
		}

		// TODO: 色々削除されているか確認する
	})
}

func TestAccountOTPPublicKeyHandler(t *testing.T) {

}

func TestAccountOTPHandler(t *testing.T) {

}

func TestAccountOTPBackupHandler(t *testing.T) {

}

func TestAccountPasswordHandler(t *testing.T) {

}

func TestAccountBeginWebauthnHandler(t *testing.T) {

}

func TestAccountWebauthnHandler(t *testing.T) {

}

func TestAccountCertificatesHandler(t *testing.T) {

}

func TestAccountForgetPasswordHandler(t *testing.T) {

}

func TestAccountReRegisterPasswordHandler(t *testing.T) {

}
