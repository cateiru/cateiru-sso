package sso_test

import (
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/pkg/go/sso"
	"github.com/stretchr/testify/require"
)

func TestCreateURI(t *testing.T) {
	clientID := "dummy"
	redirect := "https://example.com"

	uri := sso.CreateURI(clientID, redirect)

	u, err := url.Parse(uri)
	require.NoError(t, err)

	require.Equal(t, u.Host, "sso.cateiru.com")
	require.Equal(t, u.Scheme, "https")

	require.Equal(t, u.Query().Get("scope"), "openid")
	require.Equal(t, u.Query().Get("response_type"), "code")
	require.Equal(t, u.Query().Get("prompt"), "consent")

	c, err := url.QueryUnescape(u.Query().Get("client_id"))
	require.NoError(t, err)
	require.Equal(t, c, clientID)

	r, err := url.QueryUnescape(u.Query().Get("redirect_uri"))
	require.NoError(t, err)
	require.Equal(t, r, redirect)
}
