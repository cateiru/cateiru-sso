package src_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

func TestClientAuthentication(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: Basic認証", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte(client.ClientID + ":" + client.ClientSecret))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Value))

		c := m.Echo()

		returnClient, err := h.ClientAuthentication(ctx, c)
		require.NoError(t, err)

		require.Equal(t, client.ClientID, returnClient.ClientID)
	})

	t.Run("成功: POST", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		pathParam := fmt.Sprintf("/?client_id=%s&client_secret=%s", client.ClientID, client.ClientSecret)

		m, err := easy.NewMock(pathParam, http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		returnClient, err := h.ClientAuthentication(ctx, c)
		require.NoError(t, err)

		require.Equal(t, client.ClientID, returnClient.ClientID)
	})

	t.Run("失敗: どの認証も無い", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=401, error=invalid_client, message=Invalid client authentication")

		wwwAuthenticate := m.Response().Header.Get("WWW-Authenticate")
		require.Equal(t, wwwAuthenticate, "Basic")
	})

	t.Run("失敗: Basic認証でAuthorizationの形式が不正", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte(client.ClientID + ":" + client.ClientSecret))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basicaaa %s", base64Value))

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid Authorization Header")
	})

	t.Run("失敗: Basic認証でBase64デコードに失敗", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "Basic hogehoge")

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid Authorization Header")

	})

	t.Run("失敗: Basic認証でクライアントが存在しない", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte("invalid_client" + ":" + client.ClientSecret))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Value))

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_id")
	})

	t.Run("失敗: POSTでクライアントが存在しない", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		pathParam := fmt.Sprintf("/?client_id=%s&client_secret=%s", "invalid_client", client.ClientSecret)

		m, err := easy.NewMock(pathParam, http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_id")
	})

	t.Run("失敗: Basic認証でクライアントシークレットが不正", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte(client.ClientID + ":" + "invalid_client_secret"))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Value))

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_secret")
	})

	t.Run("失敗: POSTでクライアントシークレットが不正", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		pathParam := fmt.Sprintf("/?client_id=%s&client_secret=%s", client.ClientID, "invalid_client_secret")

		m, err := easy.NewMock(pathParam, http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_secret")
	})
}
