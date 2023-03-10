package src_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestLoginUserHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: Email", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.Equal(t, response.UserName, user.UserName)
		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.False(t, response.AvailablePasskey)
		require.False(t, response.AutoUsePasskey)
		require.True(t, response.AvailablePassword)
	})
	t.Run("成功: ユーザ名", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", user.UserName)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.Equal(t, response.UserName, user.UserName)
		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.False(t, response.AvailablePasskey)
		require.False(t, response.AutoUsePasskey)
		require.True(t, response.AvailablePassword)
	})
	t.Run("成功: アバターあり", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		user.Avatar = null.NewString("https://example.com/avatar", true)
		_, err := user.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		RegisterPassword(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.Equal(t, response.Avatar.String, "https://example.com/avatar")
		require.Equal(t, response.UserName, user.UserName)
		require.False(t, response.AvailablePasskey)
		require.False(t, response.AutoUsePasskey)
		require.True(t, response.AvailablePassword)
	})
	t.Run("成功: passkey登録していて、登録したデバイスでを使用", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
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
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.True(t, response.AvailablePasskey)
		require.True(t, response.AutoUsePasskey)
		require.False(t, response.AvailablePassword)
	})
	t.Run("成功: passkey登録していて、過去にpasskeyでログインしたデバイスを使用", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &user)

		// Passkeyログイン履歴追加
		passkeyHistory := models.PasskeyLoginDevice{
			UserID:   user.ID,
			Device:   null.NewString("iPhone", true),
			Os:       null.NewString("iOS", true),
			Browser:  null.NewString("Safari", true),
			IsMobile: null.NewBool(true, true),
		}
		err := passkeyHistory.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		userData := &src.UserData{
			Browser:  "Safari",
			OS:       "iOS",
			Device:   "iPhone",
			IsMobile: true,
		}
		SetUserData(t, m, userData)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.True(t, response.AvailablePasskey)
		require.True(t, response.AutoUsePasskey)
		require.False(t, response.AvailablePassword)
	})
	t.Run("成功: passkey登録しているけどログインしたことないデバイス", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		userData := &src.UserData{
			Browser:  "Safari",
			OS:       "iOS",
			Device:   "iPhone",
			IsMobile: true,
		}
		SetUserData(t, m, userData)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.True(t, response.AvailablePasskey)
		require.False(t, response.AutoUsePasskey)
		require.False(t, response.AvailablePassword)
	})
	t.Run("成功: passkeyとパスワードどっちも登録している", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &user)
		RegisterPassword(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
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
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.True(t, response.AvailablePasskey)
		require.True(t, response.AutoUsePasskey)
		require.True(t, response.AvailablePassword)
	})
	t.Run("成功: BackupStateがtrueでOSが共有可能", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &user)

		// UAを変更する
		passkeyHistory, err := models.PasskeyLoginDevices(
			models.PasskeyLoginDeviceWhere.UserID.EQ(user.ID),
		).One(ctx, DB)
		require.NoError(t, err)
		passkeyHistory.Os = null.NewString("macOS", true)
		_, err = passkeyHistory.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("username_or_email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		userData := &src.UserData{
			Device:   "",
			OS:       "iOS",
			Browser:  "Chrome",
			IsMobile: true,
		}
		SetUserData(t, m, userData)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.NoError(t, err)

		response := src.LoginUser{}
		err = m.Json(&response)
		require.NoError(t, err)

		require.False(t, response.Avatar.Valid)
		require.Equal(t, response.UserName, user.UserName)
		require.True(t, response.AvailablePasskey)
		require.True(t, response.AutoUsePasskey)
		require.False(t, response.AvailablePassword)
	})
	t.Run("失敗: username_or_emailが空", func(t *testing.T) {
		form := contents.NewMultipart()
		form.Insert("username_or_email", "")
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
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
		form := contents.NewMultipart()
		form.Insert("username_or_email", "aaaa")
		form.Insert("recaptcha", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
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
		require.EqualError(t, err, "code=400, message=user not found, unique=10")
	})
	t.Run("失敗: recaptchaが空", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", user.UserName)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA token is empty")
	})
	t.Run("失敗: reCAPTCHAチャレンジ失敗", func(t *testing.T) {
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &user)

		form := contents.NewMultipart()
		form.Insert("username_or_email", user.UserName)
		form.Insert("recaptcha", "fail")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		c := m.Echo()

		err = h.LoginUserHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})
}

func TestLoginBeginWebauthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "aaaa")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginBeginWebauthnHandler(c)
		require.NoError(t, err)

		// response
		t.Log(m.W.Body.String())
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

		require.Equal(t, webauthnSession.Challenge, resp.Response.Challenge.String())
		require.Equal(t, webauthnSession.UserID.String, u.ID)

		// rowにjsonが入っている
		sessionFromRow := new(webauthn.SessionData)
		err = webauthnSession.Row.Unmarshal(sessionFromRow)
		require.NoError(t, err)

		require.Equal(t, sessionFromRow.Challenge, resp.Response.Challenge.String())
	})

	t.Run("失敗: reCAPTCHA失敗", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPasskey(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "fail")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginBeginWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})

	t.Run("失敗: ユーザが存在しない", func(t *testing.T) {
		form := contents.NewMultipart()
		form.Insert("username_or_email", "hogehoge")
		form.Insert("recaptcha", "aaa")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginBeginWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=user not found, unique=10")
	})

	t.Run("失敗: ユーザがpasskeyを登録していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		form := contents.NewMultipart()
		form.Insert("username_or_email", u.Email)
		form.Insert("recaptcha", "aaaa")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		c := m.Echo()

		err = h.LoginBeginWebauthnHandler(c)
		require.EqualError(t, err, "code=403, message=passkey was not registered")
	})
}
