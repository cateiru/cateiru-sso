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
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestLogin(t *testing.T) {
	createSession := func(ctx context.Context, user *models.User) string {
		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)

		s := models.Session{
			ID:     sessionToken,
			UserID: user.ID,

			Period: time.Now().Add(C.SessionDBPeriod),
		}
		err = s.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		return sessionToken
	}
	createRefresh := func(ctx context.Context, user *models.User, sessionToken string) string {
		refreshToken, err := lib.RandomStr(63)
		require.NoError(t, err)
		id := ulid.Make()
		idBin, err := id.MarshalBinary()
		require.NoError(t, err)
		r := models.Refresh{
			ID:        refreshToken,
			UserID:    user.ID,
			HistoryID: idBin,

			Period: time.Now().Add(C.RefreshDBPeriod),
		}
		if sessionToken != "" {
			r.SessionID = null.NewString(sessionToken, true)
		}
		err = r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		return refreshToken
	}
	s := src.NewSession(C, DB)

	// セッショントークンを使用してログインします
	// 一般的なログイン方法
	t.Run("成功: セッショントークン", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		sessionToken := createSession(ctx, &u)

		cookie := &http.Cookie{
			Name:  C.SessionCookie.Name,
			Value: sessionToken,
		}

		loginUser, setCookies, err := s.Login(ctx, []*http.Cookie{cookie})
		require.NoError(t, err)

		require.Len(t, setCookies, 0, "Cookieは更新しない")
		require.Equal(t, loginUser.ID, u.ID)
	})

	// リフレッシュトークンを使用してログインをすると、セッショントークン、リフレッシュトークン
	// の値を更新してログインします。
	// 通常、リフレッシュトークンでログインするケースは、
	// 1. アカウントの変更
	// 2. セッショントークンの有効期限切れ
	// です。
	t.Run("成功: リフレッシュトークン", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		sessionToken := createSession(ctx, &u)
		refreshToken := createRefresh(ctx, &u, sessionToken)

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID)
		refreshCookie := &http.Cookie{
			Name:  refreshCookieName,
			Value: refreshToken,
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}

		loginUser, setCookies, err := s.Login(ctx, []*http.Cookie{refreshCookie, loginUserCookie})
		require.NoError(t, err)

		require.Len(t, setCookies, 4, "Cookieが更新される")
		require.Equal(t, loginUser.ID, u.ID)

		// 前のリフレッシュトークン、セッショントークンはDBから削除されている
		refresh, err := models.Refreshes(
			models.RefreshWhere.ID.EQ(refreshToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, refresh)
		session, err := models.Sessions(
			models.SessionWhere.ID.EQ(sessionToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, session)

		// 新しいCookieを検証する
		// 面倒くさいのでsession-tokenとrefresh-tokenのみ
		var newSessionCookie *http.Cookie = nil
		var newRefreshCookie *http.Cookie = nil
		for _, c := range setCookies {
			switch c.Name {
			case C.SessionCookie.Name:
				newSessionCookie = c
			case refreshCookieName:
				newRefreshCookie = c
			}
		}
		require.NotNil(t, newSessionCookie)
		require.NotNil(t, newRefreshCookie)
		newSession, err := models.Sessions(
			models.SessionWhere.ID.EQ(newSessionCookie.Value),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, newSession)
		newRefresh, err := models.Refreshes(
			models.RefreshWhere.ID.EQ(newRefreshCookie.Value),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, newRefresh)
	})

	// Cookieが何も存在しない場合、エラーのみを返します
	t.Run("エラー: Cookieが空", func(t *testing.T) {
		ctx := context.Background()
		_, setCookies, err := s.Login(ctx, []*http.Cookie{})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
		require.Len(t, setCookies, 0)
	})

	// LoginStateCookieはJSでログイン状態を見るためのCookie
	// もし、このCookieが存在しているのに他のセッションCookieが存在しない場合、
	// エラーになって、LoginStateCookieは削除されます
	t.Run("エラー: LoginStateCookieのみがある", func(t *testing.T) {
		ctx := context.Background()
		cookie := &http.Cookie{
			Name:  C.LoginStateCookie.Name,
			Value: "1",
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{cookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// セッションCookieは削除される
		require.Len(t, setCookies, 1)
		require.Equal(t, setCookies[0].Name, C.LoginStateCookie.Name)
		require.Equal(t, setCookies[0].MaxAge, -1)
	})

	t.Run("エラー: セッショントークンはあるが値が空", func(t *testing.T) {
		ctx := context.Background()

		cookie := &http.Cookie{
			Name:  C.SessionCookie.Name,
			Value: "",
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{cookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// セッショントークンが削除される
		require.Len(t, setCookies, 1)
	})
	t.Run("エラー: リフレッシュトークンはあるが値が空", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		refreshCookie := &http.Cookie{
			Name:  fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID),
			Value: "",
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{refreshCookie, loginUserCookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// 失敗したリフレッシュトークンと、LoginUser cookieを削除
		require.Len(t, setCookies, 2)
	})

	// セッショントークンが不正な値（=DBに存在しない値）の場合、
	// セッショントークンのCookieを削除します。
	t.Run("エラー: セッショントークンが不正", func(t *testing.T) {
		ctx := context.Background()
		cookie := &http.Cookie{
			Name:  C.SessionCookie.Name,
			Value: "123abc",
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{cookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// セッションCookieは削除される
		require.Len(t, setCookies, 1)
		require.Equal(t, setCookies[0].Name, C.SessionCookie.Name)
		require.Equal(t, setCookies[0].MaxAge, -1)
	})

	// リフレッシュトークンは存在しているが、値が不正（=DBに無い）場合
	// 該当のリフレッシュトークンと、LoginUser cookieを削除する
	t.Run("エラー: リフレッシュトークンが不正", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		refreshCookie := &http.Cookie{
			Name:  fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID),
			Value: "123abc",
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}

		// 複数ログインして言う場合のリフレッシュトークン
		// 失敗してもこれは削除されない
		otherUserRefreshCookie := &http.Cookie{
			Name:  fmt.Sprintf("%s-%s", C.RefreshCookie.Name, "dummyuser123"),
			Value: "123abc",
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{
			refreshCookie,
			loginUserCookie,
			otherUserRefreshCookie,
		})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// 失敗したリフレッシュトークンと、LoginUser cookieを削除
		require.Len(t, setCookies, 2)
	})
	t.Run("エラー: セッショントークンの有効期限が切れている", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		sessionToken := createSession(ctx, &u)

		// セッショントークンの有効期限をきらす
		ss, err := models.Sessions(
			models.SessionWhere.ID.EQ(sessionToken),
		).One(ctx, DB)
		require.NoError(t, err)
		ss.Period = time.Now().Add(-10 * time.Hour)
		_, err = ss.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		sessionCookie := &http.Cookie{
			Name:  C.SessionCookie.Name,
			Value: sessionToken,
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{sessionCookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// セッショントークンが削除される
		require.Len(t, setCookies, 1)
	})
	t.Run("エラー: リフレッシュトークンの有効期限が切れている", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		sessionToken := createSession(ctx, &u)
		refreshToken := createRefresh(ctx, &u, sessionToken)

		// リフレッシュトークンの有効期限をきらす
		r, err := models.Refreshes(
			models.RefreshWhere.ID.EQ(refreshToken),
		).One(ctx, DB)
		require.NoError(t, err)
		r.Period = time.Now().Add(-10 * time.Hour)
		_, err = r.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID)
		refreshCookie := &http.Cookie{
			Name:  refreshCookieName,
			Value: refreshToken,
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{refreshCookie, loginUserCookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// リフレッシュトークンと、LoginUser cookieが削除される
		require.Len(t, setCookies, 2)
	})

	// リフレッシュトークンはCookieに存在するが、どのユーザでログインできるかを入れるLoginUser cookie
	// が存在していないため、ログインできるユーザがわからずログインできない
	t.Run("エラー: リフレッシュトークンはあるが、LoginUserが無い", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		refreshToken := createRefresh(ctx, &u, "")

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID)
		refreshCookie := &http.Cookie{
			Name:  refreshCookieName,
			Value: refreshToken,
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{refreshCookie})
		require.EqualError(t, err, "code=403, message=you may be able to log in with another account, unique=9")

		require.Len(t, setCookies, 0)
	})

	// ログインしている
	t.Run("エラー: 該当のリフレッシュトークンが無い", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		refreshToken := createRefresh(ctx, &u2, "")

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u2.ID)
		refreshCookie := &http.Cookie{
			Name:  refreshCookieName,
			Value: refreshToken,
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{refreshCookie, loginUserCookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// LoginUser cookieのみが削除される
		require.Len(t, setCookies, 1)
	})
	t.Run("エラー: LoginUserとリフレッシュトークンのユーザが違う", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		refreshToken := createRefresh(ctx, &u2, "")

		refreshCookieName := fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID)
		refreshCookie := &http.Cookie{
			Name:  refreshCookieName,
			Value: refreshToken, // u2のリフレッシュトークン
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{refreshCookie, loginUserCookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		// リフレッシュトークンとLoginUser cookieのみが削除される
		require.Len(t, setCookies, 2)
	})
	t.Run("エラー: ユーザが存在しない", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		sessionToken := createSession(ctx, &u)

		// ユーザを削除する
		_, err := u.Delete(ctx, DB)
		require.NoError(t, err)

		cookie := &http.Cookie{
			Name:  C.SessionCookie.Name,
			Value: sessionToken,
		}

		_, setCookies, err := s.Login(ctx, []*http.Cookie{cookie})
		require.EqualError(t, err, "code=403, message=login failed, unique=8")

		require.Len(t, setCookies, 1)
	})
}

func TestLogout(t *testing.T) {
	createSession := func(ctx context.Context, user *models.User) string {
		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)

		s := models.Session{
			ID:     sessionToken,
			UserID: user.ID,

			Period: time.Now().Add(C.SessionDBPeriod),
		}
		err = s.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		return sessionToken
	}
	createRefresh := func(ctx context.Context, user *models.User, sessionToken string) string {
		refreshToken, err := lib.RandomStr(63)
		require.NoError(t, err)
		id := ulid.Make()
		idBin, err := id.MarshalBinary()
		require.NoError(t, err)
		r := models.Refresh{
			ID:        refreshToken,
			UserID:    user.ID,
			HistoryID: idBin,

			Period: time.Now().Add(C.RefreshDBPeriod),
		}
		if sessionToken != "" {
			r.SessionID = null.NewString(sessionToken, true)
		}
		err = r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		return refreshToken
	}

	s := src.NewSession(C, DB)

	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		sessionToken := createSession(ctx, &u)
		refreshToken := createRefresh(ctx, &u, sessionToken)

		sessionCookie := &http.Cookie{
			Name:  C.SessionCookie.Name,
			Value: sessionToken,
		}
		refreshCookie := &http.Cookie{
			Name:  fmt.Sprintf("%s-%s", C.RefreshCookie.Name, u.ID),
			Value: refreshToken,
		}
		loginUserCookie := &http.Cookie{
			Name:  C.LoginUserCookie.Name,
			Value: string(u.ID),
		}
		loginStateCookie := &http.Cookie{
			Name:  C.LoginStateCookie.Name,
			Value: "1",
		}
		cookies, err := s.Logout(ctx, []*http.Cookie{
			sessionCookie,
			refreshCookie,
			loginUserCookie,
			loginStateCookie,
		}, &u)
		require.NoError(t, err)

		// すべてのCookieが更新される
		require.Len(t, cookies, 4)

		// 消えている
		for _, cookie := range cookies {
			require.Equal(t, cookie.MaxAge, -1, cookie.Name)
		}

		// セッショントークンが削除されている
		existsSession, err := models.Sessions(
			models.SessionWhere.ID.EQ(sessionToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existsSession)

		// リフレッシュトークンが削除されている
		existsRefresh, err := models.Refreshes(
			models.RefreshWhere.ID.EQ(refreshToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existsRefresh)
	})
}

func TestNewRegisterSession(t *testing.T) {
	s := src.NewSession(C, DB)

	t.Run("セッションを登録できる", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		ua := &src.UserData{
			Browser:  "Chrome",
			OS:       "Windows",
			Device:   "",
			IsMobile: false,
		}
		ip := "198.51.100.5"

		registerSession, err := s.NewRegisterSession(ctx, &u, ua, ip)
		require.NoError(t, err)

		t.Run("ログが保存されている", func(t *testing.T) {
			historyCount, err := models.LoginHistories(
				models.LoginHistoryWhere.UserID.EQ(u.ID),
			).Count(ctx, DB)
			require.NoError(t, err)

			require.Equal(t, historyCount, int64(1), "ログイン履歴が保存されている")
		})
		t.Run("セッションがDBにある", func(t *testing.T) {
			existsSession, err := models.Sessions(
				models.SessionWhere.ID.EQ(registerSession.SessionToken),
			).Exists(ctx, DB)
			require.NoError(t, err)
			require.True(t, existsSession)
		})
		t.Run("リフレッシュがDBにある", func(t *testing.T) {
			existsRefresh, err := models.Refreshes(
				models.RefreshWhere.ID.EQ(registerSession.RefreshToken),
			).Exists(ctx, DB)
			require.NoError(t, err)
			require.True(t, existsRefresh)
		})
		t.Run("ログのIDとリフレッシュのIDが同じ", func(t *testing.T) {
			existsLogFromRefresh, err := models.LoginHistories(
				qm.InnerJoin("refresh ON refresh.history_id = login_history.refresh_id"),
			).Exists(ctx, DB)
			require.NoError(t, err)
			require.True(t, existsLogFromRefresh)
		})
		t.Run("cookieを作成できる", func(t *testing.T) {
			cookies := registerSession.InsertCookie(C)

			require.Len(t, cookies, 4)
		})
	})
}
