package src_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestApiOpenidConfigurationHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.ApiOpenidConfigurationHandler(c)
		require.NoError(t, err)

		response := src.OpenidConfiguration{}
		require.NoError(t, m.Json(&response))

		snaps.MatchSnapshot(t, response)
	})
}

func TestJwksJsonHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.JwksJsonHandler(c)
		require.NoError(t, err)

		data := m.W.Body.Bytes()

		snaps.MatchSnapshot(t, string(data))
	})
}

func TestWebIdentityHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")

		c := m.Echo()

		err = h.WebIdentityHandler(c)
		require.NoError(t, err)

		response := src.WebIdentityResponse{}
		require.NoError(t, m.Json(&response))

		snaps.MatchSnapshot(t, response.ProvidersUrl)
	})
}

func TestFedCMConfigHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")

		c := m.Echo()

		err = h.FedCMConfigHandler(c)
		require.NoError(t, err)

		response := src.FedCMConfigResponse{}
		require.NoError(t, m.Json(&response))

		require.NotEmpty(t, response.AccountsEndpoint)
		require.NotEmpty(t, response.ClientMetadataEndpoint)
		require.NotEmpty(t, response.IdAssertionEndpoint)

		require.Equal(t, response.Branding.BackgroundColor, C.BrandBackgroundColor)
		require.Equal(t, response.Branding.Color, C.BrandColor)
		require.Equal(t, response.Branding.Name, C.BrandName)
	})

	t.Run("失敗: Sec-Fetch-Destがwebidentityでない", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.FedCMConfigHandler(c)
		require.EqualError(t, err, "code=400, message=invalid request")
	})
}

func TestFedCMAccountsHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)
	t.Run("ユーザー数0", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")

		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.NoError(t, err)

		response := src.FedCMAccountsResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Accounts, 0)
	})

	t.Run("ユーザー数1", func(t *testing.T) {
		email1 := RandomEmail(t)

		u1 := RegisterUser(t, ctx, email1)

		cookies := RegisterSession(t, ctx, &u1)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)
		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.NoError(t, err)

		response := src.FedCMAccountsResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Accounts, 1)

		require.Equal(t, response.Accounts[0].ID, u1.ID)
		require.Equal(t, response.Accounts[0].Name, u1.UserName)
		require.Equal(t, response.Accounts[0].Email, u1.Email)
		require.Equal(t, response.Accounts[0].GivenName, u1.GivenName.String)
		require.Equal(t, response.Accounts[0].Picture, u1.Avatar.String)
	})

	t.Run("ユーザー数2", func(t *testing.T) {
		email1 := RandomEmail(t)
		email2 := RandomEmail(t)

		u1 := RegisterUser(t, ctx, email1)
		u2 := RegisterUser(t, ctx, email2)

		cookies := RegisterSession(t, ctx, &u1, &u2)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)
		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.NoError(t, err)

		response := src.FedCMAccountsResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Accounts, 2)
	})

	t.Run("失敗: Sec-Fetch-Destがwebidentityでない", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.EqualError(t, err, "code=400, message=invalid request")
	})
}

func TestFedCMClientMetadataHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")

		c := m.Echo()

		err = h.FedCMClientMetadataHandler(c)
		require.NoError(t, err)

		response := src.FedCMClientMetadataResponse{}
		require.NoError(t, m.Json(&response))

		snaps.MatchSnapshot(t, response)
	})

	t.Run("失敗: Sec-Fetch-Destがwebidentityでない", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.FedCMClientMetadataHandler(c)
		require.EqualError(t, err, "code=400, message=invalid request")
	})
}

func TestFedCMIdAssertionHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: token取得可能", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		cookies := RegisterSession(t, ctx, &u)

		param := url.Values{}
		param.Set("account_id", u.ID)
		param.Set("client_id", client.ClientID)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.NoError(t, err)

		response := src.FedCMIdAssertionResponse{}
		require.NoError(t, m.Json(&response))

		code := response.Token

		oathSession, err := models.OauthSessions(
			models.OauthSessionWhere.Code.EQ(code),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, oathSession.UserID, u.ID)
		require.Equal(t, oathSession.ClientID, client.ClientID)
		require.False(t, oathSession.Nonce.Valid)
	})

	t.Run("成功: token取得可能(nonceあり)", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		cookies := RegisterSession(t, ctx, &u)

		nonce, err := lib.RandomStr(31)
		require.NoError(t, err)

		param := url.Values{}
		param.Set("account_id", u.ID)
		param.Set("client_id", client.ClientID)
		param.Set("nonce", nonce)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.NoError(t, err)

		response := src.FedCMIdAssertionResponse{}
		require.NoError(t, m.Json(&response))

		code := response.Token

		oathSession, err := models.OauthSessions(
			models.OauthSessionWhere.Code.EQ(code),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, oathSession.UserID, u.ID)
		require.Equal(t, oathSession.ClientID, client.ClientID)
		require.True(t, oathSession.Nonce.Valid)
		require.Equal(t, oathSession.Nonce.String, nonce)
	})

	t.Run("失敗: account_idがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		cookies := RegisterSession(t, ctx, &u)

		nonce, err := lib.RandomStr(31)
		require.NoError(t, err)

		param := url.Values{}
		param.Set("account_id", "")
		param.Set("client_id", client.ClientID)
		param.Set("nonce", nonce)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.EqualError(t, err, "code=400, message=account_id is required")
	})

	t.Run("失敗: account_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		cookies := RegisterSession(t, ctx, &u)

		nonce, err := lib.RandomStr(31)
		require.NoError(t, err)

		param := url.Values{}
		param.Set("account_id", "invalid")
		param.Set("client_id", client.ClientID)
		param.Set("nonce", nonce)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.EqualError(t, err, "code=400, message=user not found")
	})

	t.Run("失敗: account_idのユーザーは存在するがリフレッシュトークンが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		nonce, err := lib.RandomStr(31)
		require.NoError(t, err)

		param := url.Values{}
		param.Set("account_id", u.ID)
		param.Set("client_id", client.ClientID)
		param.Set("nonce", nonce)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.EqualError(t, err, "code=403, message=login failed, unique=8")
	})

	t.Run("失敗: client_idがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		nonce, err := lib.RandomStr(31)
		require.NoError(t, err)

		param := url.Values{}
		param.Set("account_id", u.ID)
		param.Set("client_id", "")
		param.Set("nonce", nonce)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: client_idの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookies := RegisterSession(t, ctx, &u)

		nonce, err := lib.RandomStr(31)
		require.NoError(t, err)

		param := url.Values{}
		param.Set("account_id", u.ID)
		param.Set("client_id", "invalid")
		param.Set("nonce", nonce)

		m, err := easy.NewURLEncoded("/", http.MethodPost, param)
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")
		m.Cookie(cookies)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is invalid")
	})

	t.Run("失敗: Sec-Fetch-Destがwebidentityでない", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.FedCMIdAssertionHandler(c)
		require.EqualError(t, err, "code=400, message=invalid request")
	})
}
