package src_test

import (
	"errors"
	"io"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

// テスト用のmock オブジェクト
type ReCaptchaMock struct{}
type SenderMock struct{}
type WebAuthnMock struct {
	M *lib.WebAuthn
}
type CDNMock struct {
}

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

func (c *SenderMock) Send(m *lib.MailBody) (string, string, error) {
	return "ok", "200", nil
}

func (a *WebAuthnMock) BeginRegistration(user webauthn.User) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	// 影響は無いのでmockしない
	return a.M.BeginRegistration(user)
}
func (a *WebAuthnMock) ParseCreate(body io.Reader) (pcc *protocol.ParsedCredentialCreationData, err error) {
	return &protocol.ParsedCredentialCreationData{}, nil
}
func (a *WebAuthnMock) FinishRegistration(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialCreationData) (*webauthn.Credential, error) {
	id, err := lib.RandomBytes(64)
	if err != nil {
		return nil, err
	}
	return &webauthn.Credential{
		ID: id,
		Flags: webauthn.CredentialFlags{
			BackupState: true,
		},
	}, nil
}
func (a *WebAuthnMock) BeginLogin(user webauthn.User) (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	// 影響は無いのでmockしない
	return a.M.BeginLogin(user)
}
func (a *WebAuthnMock) ParseLogin(body io.Reader) (pcc *protocol.ParsedCredentialAssertionData, err error) {
	return &protocol.ParsedCredentialAssertionData{}, nil
}
func (a *WebAuthnMock) FinishLogin(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error) {
	return &webauthn.Credential{}, nil
}

func (c *CDNMock) Purge(url string) error {
	return nil
}
func (c *CDNMock) SoftPurge(url string) error {
	return nil
}
