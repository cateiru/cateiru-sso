package src_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h, err := src.NewHandler(DB, C)
	require.NoError(t, err)

	require.NotNil(t, h.DB)
	require.NotNil(t, h.C)
	require.NotNil(t, h.ReCaptcha)
}

func TestParseUA(t *testing.T) {
	t.Run("UA-CH", func(t *testing.T) {
		r := http.Request{
			Header: http.Header{
				"User-Agent":         {""}, // UAはない
				"Sec-Ch-Ua":          {`"Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110"`},
				"Sec-Ch-Ua-Platform": {`"Windows"`},
				"Sec-Ch-Ua-Mobile":   {"?0"},
			},
		}

		d, err := src.ParseUA(&r)
		require.NoError(t, err)

		require.Equal(t, d.Browser, "Google Chrome")
		require.Equal(t, d.Device, "")
		require.Equal(t, d.OS, "Windows")
		require.False(t, d.IsMobile)
	})

	t.Run("UA", func(t *testing.T) {
		r := http.Request{
			Header: http.Header{
				"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"},
			},
		}

		d, err := src.ParseUA(&r)
		require.NoError(t, err)

		require.Equal(t, d.Browser, "Chrome")
		require.Equal(t, d.Device, "")
		require.Equal(t, d.OS, "Windows")
		require.False(t, d.IsMobile)
	})
}

type ReCaptchaMock struct{}

func (c *ReCaptchaMock) ValidateOrder(token string, remoteIp string) (*lib.RecaptchaResponse, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	// failedにするとreCAPTCHAを失敗させる
	if token == "fail" {
		return &lib.RecaptchaResponse{
			Success:     true,
			Score:       0,
			Action:      "",
			ChallengeTS: time.Now(),
			Hostname:    "",
			ErrorCodes:  []string{},
		}, nil
	}

	return &lib.RecaptchaResponse{
		Success:     true,
		Score:       100,
		Action:      "",
		ChallengeTS: time.Now(),
		Hostname:    "",
		ErrorCodes:  []string{},
	}, nil
}

type SenderMock struct{}

func (c *SenderMock) Send(m *lib.MailBody) (string, string, error) {
	return "ok", "200", nil
}

type WebAuthnMock struct{}

func (a *WebAuthnMock) BeginRegistration(user webauthn.User) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	// TODO
	return nil, nil, nil
}
func (a *WebAuthnMock) FinishRegistration(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialCreationData) (*webauthn.Credential, error) {
	// TODO
	return nil, nil
}
func (a *WebAuthnMock) BeginLogin(user webauthn.User) (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	// TODO
	return nil, nil, nil
}
func (a *WebAuthnMock) FinishLogin(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error) {
	// TODO
	return nil, nil
}

// テスト用のダーミハンドラーを作成する
//
// モックしているやつ
// - ReCaptcha
// - Sender
func NewTestHandler(t *testing.T) *src.Handler {
	return &src.Handler{
		DB:        DB,
		C:         C,
		ReCaptcha: &ReCaptchaMock{},
		Sender:    &SenderMock{},
		WebAuthn:  &WebAuthnMock{},
	}
}
