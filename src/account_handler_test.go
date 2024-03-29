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
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
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

		m, err := easy.NewMock("/", http.MethodGet, "")
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
		m, err := easy.NewMock("/", http.MethodGet, "")
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

		m, err := easy.NewMock("/", http.MethodGet, "")
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

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("user_id", u.ID)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		c := m.Echo()

		err = h.AccountSwitchHandler(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
	})

	t.Run("ログインしているアカウントが1つだけの場合は変わらない", func(t *testing.T) {
		email1 := RandomEmail(t)
		u1 := RegisterUser(t, ctx, email1)

		cookies := RegisterSession(t, ctx, &u1)

		form := easy.NewMultipart()
		form.Insert("user_id", u1.ID)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("user_id", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("user_id", u2.ID)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

	SessionTest(t, h.AccountLogoutHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		return m
	})

	t.Run("ログアウトできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		session := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(session)
		c := m.Echo()

		err = h.AccountLogoutHandler(c)
		require.NoError(t, err)

		require.Equal(t, m.Response().Header.Get("Set-Login"), "logged-out")

		// すべてのCookieが削除されている
		cookies := m.Response().Cookies()
		for _, cookie := range cookies {
			require.Equal(t, cookie.MaxAge, -1)
		}
	})

	t.Run("login_history_idを指定して該当セッションを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		session := RegisterSession(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		RegisterSession(t, ctx, &u2)

		// user2のログイン履歴
		loginHistory, err := models.LoginHistories(
			models.LoginHistoryWhere.UserID.EQ(u2.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("login_history_id", fmt.Sprint(loginHistory.ID))
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(session)
		c := m.Echo()

		err = h.AccountLogoutHandler(c)
		require.NoError(t, err)

		require.Equal(t, m.Response().Header.Get("Set-Login"), "", "遠隔でログアウトであり、現在のブラウザはログイン状態なはずなので、Set-Loginヘッダーは空")

		// このCookieは削除しない
		cookies := m.Response().Cookies()
		require.Len(t, cookies, 0)

		// u2のセッションが全て削除されている
		existU2Session, err := models.Sessions(
			models.SessionWhere.UserID.EQ(u2.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existU2Session)

		// u2のリフレッシュトークンが全て削除されている
		existU2Refresh, err := models.Refreshes(
			models.RefreshWhere.UserID.EQ(u2.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existU2Refresh)
	})

	t.Run("失敗: 自分のセッションのlogin_history_idを指定しして削除はできない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		session := RegisterSession(t, ctx, &u)

		loginHistory, err := models.LoginHistories(
			models.LoginHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("login_history_id", fmt.Sprint(loginHistory.ID))
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(session)
		c := m.Echo()

		err = h.AccountLogoutHandler(c)
		require.EqualError(t, err, "code=400, message=cannot logout myself")
	})
}

func TestAccountDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodHead, "")
		require.NoError(t, err)
		return m
	})

	t.Run("削除されている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		session := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodHead, "")
		require.NoError(t, err)
		m.Cookie(session)
		c := m.Echo()

		err = h.AccountDeleteHandler(c)
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

	SessionTest(t, h.AccountOTPPublicKeyHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		RegisterPassword(t, ctx, u)

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)
		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodPost, "")
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

		m, err := easy.NewMock("/", http.MethodPost, "")
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

	SessionTest(t, h.AccountOTPHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		RegisterPassword(t, ctx, u)
		session, secretKey := registerOtpSession(ctx, u)

		code, err := totp.GenerateCode(secretKey, time.Now())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", "hogehoge")
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", "123456")
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("otp_session", session)
		form.Insert("code", code)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

	SessionTest(t, h.AccountDeleteOTPHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		password := "password"
		RegisterPassword(t, ctx, u, password)
		RegisterOTP(t, ctx, u)

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("password", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

	SessionTest(t, h.AccountOTPBackupHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		password := "password"
		RegisterPassword(t, ctx, u, password)
		RegisterOTP(t, ctx, u)

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: バックアップコードが返される", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		password := "password"
		RegisterPassword(t, ctx, &u, password)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		m, err := easy.NewFormData("/", http.MethodPost, form)
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

		form := easy.NewMultipart()
		form.Insert("password", "password124")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.AccountOTPBackupHandler(c)
		require.EqualError(t, err, "code=400, message=failed password, unique=8")
	})
}

func TestAccountPasswordHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountPasswordHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		password := "password_123456"

		form := easy.NewMultipart()
		form.Insert("new_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		return m
	})

	t.Run("成功: 新規に作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "password_123456"

		form := easy.NewMultipart()
		form.Insert("new_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountPasswordHandler(c)
		require.NoError(t, err)

		// パスワードが設定されている
		pw, err := models.Passwords(
			models.PasswordWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		ok := C.Password.VerifyPassword(password, pw.Hash, pw.Salt)
		require.True(t, ok)

		operationHistory, err := models.OperationHistories(
			models.OperationHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, operationHistory.Identifier, int8(7), "操作履歴が保存されている")
	})

	t.Run("失敗: すでに作成している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "password_123456"
		RegisterPassword(t, ctx, &u, password)

		form := easy.NewMultipart()
		form.Insert("new_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=password is already exists")
	})

	t.Run("失敗: 新しいパスワードが要件不足", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "p"

		form := easy.NewMultipart()
		form.Insert("new_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=invalid password")
	})
}

func TestAccountUpdatePasswordHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountUpdatePasswordHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		password := "password"
		newPassword := "password_123456"
		RegisterPassword(t, ctx, u, password)

		form := easy.NewMultipart()
		form.Insert("new_password", newPassword)
		form.Insert("old_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: 更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "password"
		newPassword := "password_123456"
		RegisterPassword(t, ctx, &u, password)

		form := easy.NewMultipart()
		form.Insert("new_password", newPassword)
		form.Insert("old_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountUpdatePasswordHandler(c)
		require.NoError(t, err)

		// パスワードが更新されている
		pw, err := models.Passwords(
			models.PasswordWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		ok := C.Password.VerifyPassword(newPassword, pw.Hash, pw.Salt)
		require.True(t, ok)

		operationHistory, err := models.OperationHistories(
			models.OperationHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, operationHistory.Identifier, int8(7), "操作履歴が保存されている")
	})

	t.Run("失敗: 前のパスワードが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "password"
		newPassword := "password_123456"
		RegisterPassword(t, ctx, &u, password)

		form := easy.NewMultipart()
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountUpdatePasswordHandler(c)
		require.EqualError(t, err, "code=400, message=password is empty")
	})

	t.Run("失敗: 前のパスワードが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "password"
		newPassword := "password_123456"
		RegisterPassword(t, ctx, &u, password)

		form := easy.NewMultipart()
		form.Insert("new_password", newPassword)
		form.Insert("old_password", "hogehoge123")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountUpdatePasswordHandler(c)
		require.EqualError(t, err, "code=400, message=failed password, unique=8")
	})

	t.Run("失敗: 新しいパスワードが要件不足", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		password := "password"
		newPassword := "p"
		RegisterPassword(t, ctx, &u, password)

		form := easy.NewMultipart()
		form.Insert("new_password", newPassword)
		form.Insert("old_password", password)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountUpdatePasswordHandler(c)
		require.EqualError(t, err, "code=400, message=invalid password")
	})

}

func TestAccountBeginWebauthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountBeginWebauthnHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: webauthnのチャレンジを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountBeginWebauthnHandler(c)
		require.NoError(t, err)

		// レスポンス
		response := protocol.CredentialCreation{}
		require.NoError(t, m.Json(&response))
		require.NotNil(t, response.Response)

		// Cookie
		var cookie *http.Cookie = nil
		for _, co := range m.Response().Cookies() {
			if co.Name == C.WebAuthnSessionCookie.Name {
				cookie = co
			}
		}
		require.NotNil(t, cookie)

		// セッションがある
		session, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(cookie.Value),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, session)
	})
}

func TestAccountWebauthnRegisteredDevicesHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountWebauthnRegisteredDevicesHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: 登録していないと空の配列が返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountWebauthnRegisteredDevicesHandler(c)
		require.NoError(t, err)

		// レスポンス
		response := []src.AccountWebauthnDevice{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 0)
	})

	t.Run("成功: 登録したデバイスが返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		RegisterPasskey(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountWebauthnRegisteredDevicesHandler(c)
		require.NoError(t, err)

		// レスポンス
		response := []src.AccountWebauthnDevice{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 1)
	})
}

func TestAccountWebauthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerWebauthnSession := func(u *models.User) string {
		webauthnUser, err := src.NewWebAuthnUserNoCredential(u)
		require.NoError(t, err)
		_, s, err := h.WebAuthn.BeginRegistration(webauthnUser)
		require.NoError(t, err)

		row := types.JSON{}
		err = row.Marshal(s)
		require.NoError(t, err)

		webauthnSessionId, err := lib.RandomStr(31)
		require.NoError(t, err)

		webauthnSession := models.WebauthnSession{
			ID:     webauthnSessionId,
			UserID: null.NewString(u.ID, true),
			Row:    row,

			Period:     time.Now().Add(h.C.WebAuthnSessionPeriod),
			Identifier: 3,
		}
		err = webauthnSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return webauthnSessionId
	}

	SessionTest(t, h.AccountWebauthnHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		session := registerWebauthnSession(u)
		sessionCookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: session,
		}

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie([]*http.Cookie{sessionCookie})

		return m
	})

	t.Run("成功: 新規追加", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		session := registerWebauthnSession(&u)
		sessionCookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: session,
		}

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		m.Cookie([]*http.Cookie{sessionCookie})

		c := m.Echo()

		err = h.AccountWebauthnHandler(c)
		require.NoError(t, err)

		// セッションは削除される
		existSession, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(session),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existSession)

		// passkeyが登録されている
		existsWebauthn, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existsWebauthn)
	})

	t.Run("失敗: application/jsonじゃない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		session := registerWebauthnSession(&u)
		sessionCookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: session,
		}

		m, err := easy.NewMock("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		m.Cookie([]*http.Cookie{sessionCookie})

		c := m.Echo()

		err = h.AccountWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=invalid content-type")
	})

	t.Run("失敗: セッションが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=session is empty")
	})

	t.Run("失敗: セッションが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		sessionCookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: "hogehoge123",
		}

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		m.Cookie([]*http.Cookie{sessionCookie})

		c := m.Echo()

		err = h.AccountWebauthnHandler(c)
		require.EqualError(t, err, "code=403, message=invalid webauthn token")
	})

	t.Run("失敗: セッションの有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		session := registerWebauthnSession(&u)
		sessionCookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: session,
		}

		// 有効期限+10h
		s, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		s.Period = s.Period.Add(-10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		m.Cookie([]*http.Cookie{sessionCookie})

		c := m.Echo()

		err = h.AccountWebauthnHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("失敗: セッションのidentifierが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		session := registerWebauthnSession(&u)
		sessionCookie := &http.Cookie{
			Name:  C.WebAuthnSessionCookie.Name,
			Value: session,
		}

		s, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(session),
		).One(ctx, DB)
		require.NoError(t, err)
		s.Identifier = 10
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewJson("/", http.MethodPost, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		m.Cookie([]*http.Cookie{sessionCookie})

		c := m.Echo()

		err = h.AccountWebauthnHandler(c)
		require.EqualError(t, err, "code=403, message=invalid webauthn token")
	})
}

func TestAccountDeleteWebauthnHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountDeleteWebauthnHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		RegisterPassword(t, ctx, u)
		RegisterPasskey(t, ctx, u)

		webauthn, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?webauthn_id=%d", webauthn.ID), http.MethodDelete, "")
		require.NoError(t, err)

		return m
	})

	t.Run("Webauthnを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)
		RegisterPasskey(t, ctx, &u)

		webauthn, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?webauthn_id=%d", webauthn.ID), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountDeleteWebauthnHandler(c)
		require.NoError(t, err)

		existsWebauthn, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existsWebauthn)
	})

	t.Run("失敗: IDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)
		RegisterPasskey(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?webauthn_id=hogehoge", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountDeleteWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=webauthn_id is invalid")
	})

	t.Run("失敗: IDを指定しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)
		RegisterPasskey(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountDeleteWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=webauthn_id is empty")
	})

	t.Run("失敗: パスワードを設定していない場合、すべて削除することはできない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		// passkeyを1つ作成する
		RegisterPasskey(t, ctx, &u)

		webauthn, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?webauthn_id=%d", webauthn.ID), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountDeleteWebauthnHandler(c)
		require.EqualError(t, err, "code=400, message=webauthn must be set at least one, unique=14")
	})

	t.Run("成功: WebAuthnが複数ある場合はパスワードを設定してくなくても削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		// passkeyを2つ作成する
		RegisterPasskey(t, ctx, &u)
		RegisterPasskey(t, ctx, &u)

		webauthn, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?webauthn_id=%d", webauthn.ID), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountDeleteWebauthnHandler(c)
		require.NoError(t, err)

		webauthnCount, err := models.Webauthns(
			models.WebauthnWhere.UserID.EQ(u.ID),
		).Count(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, webauthnCount, int64(1))
	})
}

func TestAccountCertificatesHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.AccountCertificatesHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		RegisterPassword(t, ctx, u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: パスワードのみ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountCertificatesHandler(c)
		require.NoError(t, err)

		response := src.AccountCertificates{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.Password)
		require.False(t, response.OTP)
		require.False(t, response.OtpUpdatedAt.Valid)
	})

	t.Run("成功: パスワード、OTP", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)
		RegisterOTP(t, ctx, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.AccountCertificatesHandler(c)
		require.NoError(t, err)

		response := src.AccountCertificates{}
		require.NoError(t, m.Json(&response))

		require.True(t, response.Password)
		require.True(t, response.OTP)
		require.True(t, response.OtpUpdatedAt.Valid)
	})
}

func TestAccountForgetPasswordHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: メールを送信できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.NoError(t, err)

		// ログイントライ履歴が保存されている
		existLoginTryHistory, err := models.LoginTryHistories(
			models.LoginTryHistoryWhere.UserID.EQ(u.ID),
			qm.And("identifier = 1"),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.True(t, existLoginTryHistory)

		// パスワード再設定セッションが保存されている
		existsReregistrationPasswordSession, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.Email.EQ(email),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existsReregistrationPasswordSession)
	})

	t.Run("成功: パスワード設定していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.NoError(t, err)

		// ログイントライ履歴が保存されている
		existLoginTryHistory, err := models.LoginTryHistories(
			models.LoginTryHistoryWhere.UserID.EQ(u.ID),
			qm.And("identifier = 1"),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.True(t, existLoginTryHistory)

		// パスワード再設定セッションが保存されている
		existsReregistrationPasswordSession, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.Email.EQ(email),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existsReregistrationPasswordSession)
	})

	t.Run("成功: period_clearの有効期限が切れてしまっている場合、DBから削除される", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		// 有効期限切れのセッションを作成する
		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.ReregistrationPasswordSession{
			ID:          token,
			Email:       email,
			Period:      time.Now().Add(-100 * time.Hour),
			PeriodClear: time.Now().Add(-10 * time.Hour),
		}
		err = session.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.NoError(t, err)

		// ログイントライ履歴が保存されている
		existLoginTryHistory, err := models.LoginTryHistories(
			models.LoginTryHistoryWhere.UserID.EQ(u.ID),
			qm.And("identifier = 1"),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.True(t, existLoginTryHistory)

		// パスワード再設定セッションが保存されている
		reRegistrationPasswordSession, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.Email.EQ(email),
		).Count(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, reRegistrationPasswordSession, int64(1), "前のセッションは削除されているので1つしかない")
	})

	t.Run("失敗: メールアドレスが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=invalid email")
	})

	t.Run("失敗: メールアドレスが存在しない", func(t *testing.T) {
		email := RandomEmail(t)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=user not found, unique=10")
	})

	t.Run("失敗: reCAPTCHA失敗", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "fail")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})

	t.Run("失敗: すでにセッションが存在している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		// セッションを作成する
		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.ReregistrationPasswordSession{
			ID:          token,
			Email:       email,
			Period:      time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			PeriodClear: time.Now().Add(C.ReregistrationPasswordSessionClearPeriod),
		}
		err = session.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountForgetPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=already sessions")
	})
}

func TestAccountReRegisterAvailableTokenHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerSession := func(u *models.User, period time.Time, completed bool) string {
		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.ReregistrationPasswordSession{
			ID:          token,
			Email:       u.Email,
			Period:      period,
			PeriodClear: time.Now().Add(C.ReregistrationPasswordSessionClearPeriod),

			Completed: completed,
		}

		err = session.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return token
	}

	t.Run("セッションが存在する", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			false,
		)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("reregister_token", token)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterAvailableTokenHandler(c)
		require.NoError(t, err)

		response := src.AccountReRegisterPasswordIsSession{}
		require.NoError(t, m.Json(&response))
		require.True(t, response.Active)
	})

	t.Run("セッションが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		RegisterUser(t, ctx, email)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("reregister_token", "hogehoge")
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterAvailableTokenHandler(c)
		require.NoError(t, err)

		response := src.AccountReRegisterPasswordIsSession{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.Active)
	})

	t.Run("セッションが有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		token := registerSession(
			&u,
			time.Now().Add(-10*time.Hour),
			false,
		)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("reregister_token", token)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterAvailableTokenHandler(c)
		require.NoError(t, err)

		response := src.AccountReRegisterPasswordIsSession{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.Active)

		// セッションは削除されない
		dbSession, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.ID.EQ(token),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, dbSession)
	})

	t.Run("period_clearの有効期限が切れてしまっている場合、DBから削除される", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		// 有効期限切れのセッションを作成する
		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.ReregistrationPasswordSession{
			ID:          token,
			Email:       email,
			Period:      time.Now().Add(-100 * time.Hour),
			PeriodClear: time.Now().Add(-10 * time.Hour),
		}
		err = session.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("reregister_token", token)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterAvailableTokenHandler(c)
		require.NoError(t, err)

		response := src.AccountReRegisterPasswordIsSession{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.Active)

		// セッションは削除されている
		dbSession, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.ID.EQ(token),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, dbSession)
	})

	t.Run("セッションは使用済み", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			true,
		)

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("reregister_token", token)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterAvailableTokenHandler(c)
		require.NoError(t, err)

		response := src.AccountReRegisterPasswordIsSession{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.Active)
	})

	t.Run("emailが不正", func(t *testing.T) {
		email := RandomEmail(t)
		email2 := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			false,
		)

		form := easy.NewMultipart()
		form.Insert("email", email2)
		form.Insert("reregister_token", token)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterAvailableTokenHandler(c)
		require.NoError(t, err)

		response := src.AccountReRegisterPasswordIsSession{}
		require.NoError(t, m.Json(&response))
		require.False(t, response.Active)
	})
}

func TestAccountReRegisterPasswordHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerSession := func(u *models.User, period time.Time, completed bool) string {
		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.ReregistrationPasswordSession{
			ID:          token,
			Email:       u.Email,
			Period:      period,
			PeriodClear: time.Now().Add(C.ReregistrationPasswordSessionClearPeriod),

			Completed: completed,
		}

		err = session.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return token
	}

	t.Run("成功: パスワードを新しくできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		// パスワードは存在している
		RegisterPassword(t, ctx, &u)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			false,
		)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("reregister_token", token)
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.NoError(t, err)

		// セッションは使用済みフラグが立ち、削除されない
		session, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.ID.EQ(token),
		).One(ctx, DB)
		require.NoError(t, err)
		require.True(t, session.Completed)

		// パスワードが更新されている
		password, err := models.Passwords(
			models.PasswordWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		isVerify := C.Password.VerifyPassword(newPassword, password.Hash, password.Salt)
		require.True(t, isVerify)

		operationHistory, err := models.OperationHistories(
			models.OperationHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, operationHistory.Identifier, int8(7))
	})

	t.Run("成功: パスワードを新規に作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			false,
		)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("reregister_token", token)
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.NoError(t, err)

		// セッションは使用済みフラグが立ち、削除されない
		session, err := models.ReregistrationPasswordSessions(
			models.ReregistrationPasswordSessionWhere.ID.EQ(token),
		).One(ctx, DB)
		require.NoError(t, err)
		require.True(t, session.Completed)

		// パスワードが新規に作成されている
		password, err := models.Passwords(
			models.PasswordWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		isVerify := C.Password.VerifyPassword(newPassword, password.Hash, password.Salt)
		require.True(t, isVerify)

		operationHistory, err := models.OperationHistories(
			models.OperationHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, operationHistory.Identifier, int8(7))
	})

	t.Run("失敗: reCAPTCHA失敗", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			false,
		)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "fail")
		form.Insert("reregister_token", token)
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=reCAPTCHA validation failed, unique=1")
	})

	t.Run("失敗: tokenが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=reregister_token is empty")
	})

	t.Run("失敗: tokenが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("reregister_token", "hogehoge")
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=invalid token")
	})

	t.Run("失敗: セッションが有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		token := registerSession(
			&u,
			time.Now().Add(-10*time.Hour),
			false,
		)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("reregister_token", token)
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")
	})

	t.Run("セッションは使用済み", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			true,
		)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("reregister_token", token)
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.EqualError(t, err, "code=400, message=invalid token")
	})

	t.Run("emailが不正", func(t *testing.T) {
		email := RandomEmail(t)
		email2 := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterPassword(t, ctx, &u)

		token := registerSession(
			&u,
			time.Now().Add(C.ReregistrationPasswordSessionPeriod),
			false,
		)

		newPassword := "password_1234567"

		form := easy.NewMultipart()
		form.Insert("email", email2)
		form.Insert("recaptcha", "hogehoge")
		form.Insert("reregister_token", token)
		form.Insert("new_password", newPassword)
		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = h.AccountReRegisterPasswordHandler(c)
		require.EqualError(t, err, "code=403, message=email is different")
	})
}
