package lib_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/require"
)

type WebAuthnUser struct{}

func (w *WebAuthnUser) WebAuthnID() []byte {
	return []byte("test1234")
}

func (w *WebAuthnUser) WebAuthnName() string {
	return "test1234"
}

func (w *WebAuthnUser) WebAuthnDisplayName() string {
	return "test hoge"
}

func (w *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{}
}

func (w *WebAuthnUser) WebAuthnIcon() string {
	return ""
}

func TestWebAuthn(t *testing.T) {
	c := &webauthn.Config{
		RPDisplayName: "Cateiru SSO Test",
		RPID:          "localhost:3000",
		RPOrigins:     []string{"localhost:3000", "localhost:8080"},
	}

	w, err := lib.NewWebAuthn(c)
	require.NoError(t, err)

	u := &WebAuthnUser{}

	t.Run("Register", func(t *testing.T) {
		creation, session, err := w.BeginRegistration(u)
		require.NoError(t, err)

		// FIXME: どうにかしたい
		require.NotNil(t, creation)
		require.NotNil(t, session)
	})

	t.Run("Login", func(t *testing.T) {
		creation, session, err := w.BeginLogin()
		// WebAuthnCredentialsが空なのでエラーになる
		require.Error(t, err)

		// FIXME: どうにかしたい
		require.Nil(t, creation)
		require.Nil(t, session)
	})
}
