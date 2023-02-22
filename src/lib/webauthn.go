package lib

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type WebAuthnInterface interface {
	BeginRegistration(user webauthn.User) (*protocol.CredentialCreation, *webauthn.SessionData, error)
	FinishRegistration(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialCreationData) (*webauthn.Credential, error)
	BeginLogin(user webauthn.User) (*protocol.CredentialAssertion, *webauthn.SessionData, error)
	FinishLogin(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error)
}

type WebAuthn struct {
	W *webauthn.WebAuthn
}

func NewWebAuthn(c *webauthn.Config) (*WebAuthn, error) {

	w, err := webauthn.New(c)
	if err != nil {
		return nil, err
	}

	return &WebAuthn{
		W: w,
	}, nil
}

// 登録: WebAuthnのconfig作成
func (a *WebAuthn) BeginRegistration(user webauthn.User) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	return a.W.BeginRegistration(user)
}

// 登録: WebAuthnの検証
//
// response create is...
// `response, err := protocol.ParseCredentialCreationResponseBody(r.Body)`
func (a *WebAuthn) FinishRegistration(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialCreationData) (*webauthn.Credential, error) {
	return a.W.CreateCredential(user, session, response)
}

// ログイン: config作成
func (a *WebAuthn) BeginLogin(user webauthn.User) (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	return a.W.BeginLogin(user)
}

// ログイン: 検証
//
// response create is...
// `response, err := protocol.ParseCredentialRequestResponseBody(r.Body)`
func (a *WebAuthn) FinishLogin(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error) {
	return a.W.ValidateLogin(user, session, response)
}
