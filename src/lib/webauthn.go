package lib

import (
	"io"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type WebAuthnInterface interface {
	BeginRegistration(user webauthn.User) (*protocol.CredentialCreation, *webauthn.SessionData, error)
	ParseCreate(body io.Reader) (pcc *protocol.ParsedCredentialCreationData, err error)
	ParseLogin(body io.Reader) (pcc *protocol.ParsedCredentialAssertionData, err error)
	FinishRegistration(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialCreationData) (*webauthn.Credential, error)
	BeginLogin() (*protocol.CredentialAssertion, *webauthn.SessionData, error)
	FinishLogin(handler webauthn.DiscoverableUserHandler, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error)
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

func (a *WebAuthn) ParseCreate(body io.Reader) (pcc *protocol.ParsedCredentialCreationData, err error) {
	return protocol.ParseCredentialCreationResponseBody(body)
}

// 登録: WebAuthnの検証
//
// response create is...
// `response, err := protocol.ParseCredentialCreationResponseBody(r.Body)`
func (a *WebAuthn) FinishRegistration(user webauthn.User, session webauthn.SessionData, response *protocol.ParsedCredentialCreationData) (*webauthn.Credential, error) {
	return a.W.CreateCredential(user, session, response)
}

// ログイン: config作成
func (a *WebAuthn) BeginLogin() (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	return a.W.BeginDiscoverableLogin()
}

func (a *WebAuthn) ParseLogin(body io.Reader) (pcc *protocol.ParsedCredentialAssertionData, err error) {
	return protocol.ParseCredentialRequestResponseBody(body)
}

// ログイン: 検証
func (a *WebAuthn) FinishLogin(handler webauthn.DiscoverableUserHandler, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error) {
	return a.W.ValidateDiscoverableLogin(handler, session, response)
}
