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
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

	SessionTest(t, h.AccountLogoutHandler, func(ctx context.Context, u *models.User) *mock.MockHandler {
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

	SessionTest(t, h.AccountDeleteHandler, func(ctx context.Context, u *models.User) *mock.MockHandler {
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
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountOTPPublicKeyHandler, func(ctx context.Context, u *models.User) *mock.MockHandler {
		RegisterPassword(t, ctx, u)

		m, err := mock.NewMock("", http.MethodPost, "/")
		require.NoError(t, err)
		return m
	})

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)
		cookies := RegisterSession(t, ctx, &u)

		m, err := mock.NewMock("", http.MethodPost, "/")
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPPublicKeyHandler(c)
		require.NoError(t, err)

		var response = src.AccountOTPPublic{}
		err = m.Json(&response)
		require.NoError(t, err)

		session, err := models.RegisterOtpSessions(
			models.RegisterOtpSessionWhere.ID.EQ(response.OTPSession),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, session.PublicKey, response.PublicKey)
	})

	t.Run("失敗: パスワードを設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := mock.NewMock("", http.MethodPost, "/")
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPPublicKeyHandler(c)
		require.EqualError(t, err, "code=400, message=no registered password, unique=11")
	})
}

func TestAccountOTPHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerOtpSession := func(ctx context.Context, u *models.User) (string, string) {
		otp, err := lib.NewOTP(C.OTPIssuer, u.UserName)
		require.NoError(t, err)

		session, err := lib.RandomStr(31)
		require.NoError(t, err)
		otpRegisterSession := models.RegisterOtpSession{
			ID:        session,
			UserID:    u.ID,
			PublicKey: otp.GetPublic(),
			Secret:    otp.GetSecret(),
			Period:    time.Now().Add(C.OTPRegisterSessionPeriod),
		}
		err = otpRegisterSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return session, otp.GetSecret()
	}

	SessionTest(t, h.AccountOTPHandler, func(ctx context.Context, u *models.User) *mock.MockHandler {
		RegisterPassword(t, ctx, u)
		session, _ := registerOtpSession(ctx, u)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 新規に追加", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		session, secretKey := registerOtpSession(ctx, &u)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.NoError(t, err)

		backups := []string{}
		err = m.Json(&backups)
		require.NoError(t, err)
		require.Len(t, backups, int(C.OTPBackupCount))

		// セッションは削除
		existsRegisterOTPSession, err := models.RegisterOtpSessions(
			models.RegisterOtpSessionWhere.ID.EQ(session),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existsRegisterOTPSession)

		// OTPが設定されている
		otp, err := models.Otps(
			models.OtpWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, otp.Secret, secretKey)
	})

	t.Run("成功: OTPを更新する", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		secret, backups := RegisterOTP(t, ctx, &u)
		session, secretKey := registerOtpSession(ctx, &u)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.NoError(t, err)

		// backupsがすべて更新されている
		newBackups := []string{}
		err = m.Json(&newBackups)
		require.NoError(t, err)
		require.Len(t, newBackups, int(C.OTPBackupCount))
		for i, b := range backups {
			for bb := range newBackups[i:] {
				require.NotEqual(t, b, bb)
			}
		}

		// セッションは削除
		existsRegisterOTPSession, err := models.RegisterOtpSessions(
			models.RegisterOtpSessionWhere.ID.EQ(session),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existsRegisterOTPSession)

		// OTPが更新されている
		otp, err := models.Otps(
			models.OtpWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.NotEqual(t, otp.Secret, secret)
	})

	t.Run("失敗: セッションが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		_, secretKey := registerOtpSession(ctx, &u)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.EqualError(t, err, "code=400, message=otp_session is empty")
	})

	t.Run("失敗: セッションが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		_, secretKey := registerOtpSession(ctx, &u)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("otp_session", "hogehoge")
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.EqualError(t, err, "code=400, message=invalid otp_session")
	})

	t.Run("失敗: セッションの有効期限が切れている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		session, secretKey := registerOtpSession(ctx, &u)

		// 有効期限 - 10日
		s, err := models.RegisterOtpSessions(
			models.RegisterOtpSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		s.Period = time.Now().Add(-24 * 10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("失敗: セッションのリトライ回数上限を超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		session, secretKey := registerOtpSession(ctx, &u)

		// リトライ上限にする
		s, err := models.RegisterOtpSessions(
			models.RegisterOtpSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		s.RetryCount = C.OTPRegisterLimit
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.EqualError(t, err, "code=403, message=exceeded retry, unique=4")
	})

	t.Run("失敗: OTPが認証できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		RegisterPassword(t, ctx, &u)
		session, _ := registerOtpSession(ctx, &u)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", "123456")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.EqualError(t, err, "code=403, message=failed otp validate")

		// リトライカウントが++される
		s, err := models.RegisterOtpSessions(
			models.RegisterOtpSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, s.RetryCount, uint8(1))
	})

	t.Run("失敗: パスワードを設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)
		session, secretKey := registerOtpSession(ctx, &u)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPHandler(c)
		require.EqualError(t, err, "code=400, message=no registered password, unique=11")
	})
}

func TestAccountDeleteOTPHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountDeleteOTPHandler, func(ctx context.Context, u *models.User) *mock.MockHandler {
		password := "password"
		RegisterPassword(t, ctx, u, password)
		RegisterOTP(t, ctx, u)

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		return m
	})

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountDeleteOTPHandler(c)
		require.NoError(t, err)

		// OTPが削除されている
		existOTP, err := models.Otps(
			models.OtpWhere.UserID.EQ(u.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existOTP)

		// OTPのバックアップが削除されている
		backups, err := models.OtpBackups(
			models.OtpBackupWhere.UserID.EQ(u.ID),
		).All(ctx, DB)
		require.NoError(t, err)
		require.Len(t, backups, 0)
	})

	t.Run("失敗: パスワードが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountDeleteOTPHandler(c)
		require.EqualError(t, err, "code=400, message=password is empty")
	})

	t.Run("失敗: パスワードが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", "hogehoge")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountDeleteOTPHandler(c)
		require.EqualError(t, err, "code=400, message=failed password, unique=8")
	})

	t.Run("失敗: OTPをそもそも設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountDeleteOTPHandler(c)
		require.EqualError(t, err, "code=400, message=otp is not registered")
	})

	t.Run("失敗: そもそもパスワードを設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountDeleteOTPHandler(c)
		require.EqualError(t, err, "code=400, message=no registered password, unique=11")
	})
}

func TestAccountOTPBackupHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: バックアップコードが返される", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPBackupHandler(c)
		require.NoError(t, err)

		backups := []string{}
		err = m.Json(&backups)
		require.NoError(t, err)

		b, err := models.OtpBackups(
			models.OtpBackupWhere.UserID.EQ(u.ID),
		).All(ctx, DB)
		require.NoError(t, err)

		dbBackupCodes := make([]string, len(b))
		for i, bb := range b {
			dbBackupCodes[i] = bb.Code
		}

		require.EqualValues(t, backups, dbBackupCodes)
	})

	t.Run("OTPが設定されていない=バックアップコードが無いと空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", password)
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPBackupHandler(c)
		require.NoError(t, err)

		backups := []string{}
		err = m.Json(&backups)
		require.NoError(t, err)

		require.Len(t, backups, 0)
	})

	t.Run("失敗: パスワードが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPBackupHandler(c)
		require.EqualError(t, err, "code=400, message=password is empty")
	})

	t.Run("失敗: パスワードが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := contents.NewMultipart()
		form.Insert("password", "password124")
		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPBackupHandler(c)
		require.EqualError(t, err, "code=400, message=failed password, unique=8")
	})
}

func TestAccountPasswordHandler(t *testing.T) {
	t.Run("成功: 新規に作成できる", func(t *testing.T) {})

	t.Run("失敗: すでに作成している", func(t *testing.T) {})

	t.Run("失敗: 新しいパスワードが要件不足", func(t *testing.T) {})
}

func TestAccountUpdatePasswordHandler(t *testing.T) {
	t.Run("成功: 更新できる", func(t *testing.T) {})

	t.Run("失敗: 前のパスワードが空", func(t *testing.T) {})

	t.Run("失敗: 前のパスワードが不正", func(t *testing.T) {})

	t.Run("失敗: 新しいパスワードが要件不足", func(t *testing.T) {})

}

func TestAccountBeginWebauthnHandler(t *testing.T) {
	t.Run("成功: webauthnのチャレンジを取得できる", func(t *testing.T) {})
}

func TestAccountWebauthnHandler(t *testing.T) {
	t.Run("成功", func(t *testing.T) {})

	t.Run("失敗: application/jsonじゃない", func(t *testing.T) {})

	t.Run("失敗: セッションが空", func(t *testing.T) {})

	t.Run("失敗: セッションが不正", func(t *testing.T) {})

	t.Run("失敗: セッションの有効期限切れ", func(t *testing.T) {})

	t.Run("失敗: セッションのidentifierが不正", func(t *testing.T) {})
}

func TestAccountCertificatesHandler(t *testing.T) {

	t.Run("成功: パスワード、OTP、Passkeyすべて設定している", func(t *testing.T) {})

	t.Run("成功: パスワードのみ", func(t *testing.T) {})

	t.Run("成功: パスワード、OTP", func(t *testing.T) {})

	t.Run("成功: Passkeyのみ", func(t *testing.T) {})
}

func TestAccountForgetPasswordHandler(t *testing.T) {
	t.Run("成功: メールを送信できる", func(t *testing.T) {})

	t.Run("失敗: メールアドレスが空", func(t *testing.T) {})

	t.Run("失敗: メールアドレスが存在しない", func(t *testing.T) {})

	t.Run("失敗: reCAPTCHA失敗", func(t *testing.T) {})

	t.Run("失敗: パスワード設定していない", func(t *testing.T) {})

	t.Run("失敗: すでにセッションが存在している", func(t *testing.T) {})
}

func TestAccountReRegisterAvailableTokenHandler(t *testing.T) {
	t.Run("セッションが存在する", func(t *testing.T) {})

	t.Run("セッションが存在しない", func(t *testing.T) {})

	t.Run("セッションが有効期限切れ", func(t *testing.T) {})

	t.Run("セッションは使用済み", func(t *testing.T) {})
}

func TestAccountReRegisterPasswordHandler(t *testing.T) {
	t.Run("成功: パスワードを新しくできる", func(t *testing.T) {})

	t.Run("失敗: reCAPTCHA失敗", func(t *testing.T) {})

	t.Run("失敗: tokenが空", func(t *testing.T) {})

	t.Run("失敗: tokenが不正", func(t *testing.T) {})

	t.Run("失敗: セッションが有効期限切れ", func(t *testing.T) {})

	t.Run("セッションは使用済み", func(t *testing.T) {})

	t.Run("emailが不正", func(t *testing.T) {})
}
