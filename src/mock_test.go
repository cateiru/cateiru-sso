package src_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// テスト用のmock オブジェクト
type ReCaptchaMock struct{}

type SenderMock struct{}

type WebAuthnMock struct {
	M *lib.WebAuthn
}
type CDNMock struct{}

type StorageMock struct {
	S lib.CloudStorageInterface
}

// --- ReCaptchaMock

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

// --- SenderMock

func (c *SenderMock) Send(m *lib.MailBody) (string, string, error) {
	return "ok", "200", nil
}

// --- WebAuthnMock

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

func (a *WebAuthnMock) BeginLogin() (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	// 影響は無いのでmockしない
	return a.M.BeginLogin()
}

func (a *WebAuthnMock) ParseLogin(body io.Reader) (pcc *protocol.ParsedCredentialAssertionData, err error) {
	return &protocol.ParsedCredentialAssertionData{}, nil
}

func (a *WebAuthnMock) FinishLogin(handler webauthn.DiscoverableUserHandler, session webauthn.SessionData, response *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error) {
	ctx := context.Background()

	r, err := lib.RandomStr(10)
	if err != nil {
		return nil, err
	}
	id := ulid.Make()

	u := models.User{
		ID:    id.String(),
		Email: fmt.Sprintf("%s@exmaple.com", r),
	}
	if err := u.Insert(ctx, DB, boil.Infer()); err != nil {
		return nil, err
	}

	if _, err := handler([]byte{}, []byte(u.ID)); err != nil {
		return nil, err
	}

	return &webauthn.Credential{}, nil
}

// --- CDNMock

func (c *CDNMock) Purge(url string) error {
	return nil
}
func (c *CDNMock) SoftPurge(url string) error {
	return nil
}

// --- StorageMock

func (c *StorageMock) Read(ctx context.Context, path string) ([]byte, string, error) {
	return c.S.Read(ctx, path)
}

func (c *StorageMock) Write(ctx context.Context, path string, data io.Reader, contentType string) error {
	return c.S.Write(ctx, path, data, contentType)
}

func (c *StorageMock) Delete(ctx context.Context, path string) error {
	// 削除部分だけモックする
	return nil
}
