package src_test

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// TODO

func TestHistoryClientLoginHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerClientRefresh := func(clientId string, u *models.User) string {
		id, err := lib.RandomStr(63)
		require.NoError(t, err)

		scope := []string{
			"openid",
			"email",
		}
		scopesJson := types.JSON{}
		err = scopesJson.Marshal(&scope)
		require.NoError(t, err)

		refresh := models.ClientRefresh{
			ID:     id,
			UserID: u.ID,

			ClientID: clientId,

			Scopes: scopesJson,
			Period: time.Now().Add(C.ClientRefreshPeriod),
		}
		err = refresh.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return id
	}

	t.Run("成功: ログインしているクライアントが返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)
		registerClientRefresh(clientId, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryClientLoginHandler(c)
		require.NoError(t, err)

		response := []src.ClientLoginResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 1)
		require.Equal(t, response[0].Client.ClientID, clientId)
	})

	t.Run("成功: クライアントが存在していなくても返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, err := lib.RandomStr(31)
		require.NoError(t, err)
		registerClientRefresh(clientId, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryClientLoginHandler(c)
		require.NoError(t, err)

		response := []src.ClientLoginResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 1)
		require.Nil(t, response[0].Client)
	})

	t.Run("成功: 何もログインしていないときは空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryClientLoginHandler(c)
		require.NoError(t, err)

		response := []src.ClientLoginResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 0)
	})
}

func TestHistoryClientHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	registerClientLoginHistory := func(clientId string, u *models.User) {
		history := models.LoginClientHistory{
			ClientID: clientId,
			UserID:   u.ID,

			Device:   null.NewString("", false),
			Os:       null.NewString("Windows", true),
			Browser:  null.NewString("Chrome", true),
			IsMobile: null.NewBool(false, true),

			IP: net.ParseIP("10.0.1.1"),
		}
		err := history.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
	}

	adminEmail := RandomEmail(t)
	adminUser := RegisterUser(t, ctx, adminEmail)

	clientId, _ := RegisterClient(t, ctx, &adminUser)

	t.Run("成功: ログイン履歴が返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		registerClientLoginHistory(clientId, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryClientHandler(c)
		require.NoError(t, err)

		response := []src.ClientLoginHistoryResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 1)
	})

	t.Run("成功: クライアントが存在しなくても返る", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, err := lib.RandomStr(31)
		require.NoError(t, err)
		registerClientLoginHistory(clientId, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryClientHandler(c)
		require.NoError(t, err)

		response := []src.ClientLoginHistoryResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 1)
		require.Nil(t, response[0].Client)
	})

	t.Run("成功: 最新順に並んでいる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		registerClientLoginHistory(clientId, &u)
		time.Sleep(1 * time.Second)
		registerClientLoginHistory(clientId, &u)

		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryClientHandler(c)
		require.NoError(t, err)

		response := []src.ClientLoginHistoryResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 2)

		require.True(t, response[0].Created.After(response[1].Created))
	})
}

func TestHistoryLoginDeviceHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: ログインしているデバイスを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterSession(t, ctx, &u)
		cookies := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryLoginDeviceHandler(c)
		require.NoError(t, err)

		response := []src.LoginDeviceResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 2)
	})

	t.Run("成功: 履歴はあるが、リフレッシュトークンが存在しない場合は返さない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterSession(t, ctx, &u)
		cookies := RegisterSession(t, ctx, &u)

		refreshes, err := models.Refreshes(
			models.RefreshWhere.UserID.EQ(u.ID),
		).All(ctx, DB)
		require.NoError(t, err)
		_, err = refreshes[0].Delete(ctx, DB)
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryLoginDeviceHandler(c)
		require.NoError(t, err)

		response := []src.LoginDeviceResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 1)
	})
}

func TestHistoryLoginHistoryHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: ログイン履歴を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		refreshId, err := lib.RandomBytes(16)
		require.NoError(t, err)
		history := models.LoginHistory{
			UserID:    u.ID,
			RefreshID: refreshId,

			IP: net.ParseIP("10.0.0.1"),
		}
		err = history.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryLoginHistoryHandler(c)
		require.NoError(t, err)

		response := []src.LoginDeviceResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 2)
	})
}

func TestHistoryLoginTryHistoryHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: ログイントライ履歴を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		loginTryHistory := models.LoginTryHistory{
			UserID: u.ID,

			IP: net.ParseIP("10.0.0.1"),
		}
		err := loginTryHistory.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)

		c := m.Echo()

		err = h.HistoryLoginTryHistoryHandler(c)
		require.NoError(t, err)

		response := []src.LoginTryHistoryResponse{}
		require.NoError(t, m.Json(&response))
		require.Len(t, response, 1)
	})
}
