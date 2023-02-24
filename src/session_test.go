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
