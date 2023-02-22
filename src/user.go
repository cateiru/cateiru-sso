package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/go-webauthn/webauthn/webauthn"
)

type WebAuthnUser struct {
	// size of 64 bytes
	ID         []byte
	Credential []webauthn.Credential
	// 表示用
	Name        string
	DisplayName string
	Icon        string
}

// DBからWebAuthnにわたす用のユーザを作成します
// ユーザはログインしている必要があり、かつpasskeyが登録されている必要があります
func NewWebAuthnUserFromDB(ctx context.Context, db *sql.DB, user *models.User) (*WebAuthnUser, error) {
	passkey, err := models.Passkeys(models.PasskeyWhere.UserID.EQ(user.ID)).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewHTTPError(403, "passkey was not registered")
	}
	if err != nil {
		return nil, err
	}

	displayName := ""
	if user.FamilyName.Valid && user.GivenName.Valid {
		// 名前が設定されている場合はフォーマットする
		// 順序は姓名と、日本式
		if !user.MiddleName.Valid {
			displayName = fmt.Sprintf("%s %s %s", user.FamilyName.String, user.MiddleName.String, user.GivenName.String)
		} else {
			displayName = fmt.Sprintf("%s %s", user.FamilyName.String, user.GivenName.String)
		}
	} else {
		// 設定していない場合はEmail
		displayName = user.Email
	}

	icon := ""
	if user.Avater.Valid {
		icon = user.Avater.String
	}

	credential := new(webauthn.Credential)
	if err := passkey.Credential.Unmarshal(credential); err != nil {
		return nil, err
	}

	return &WebAuthnUser{
		ID: passkey.WebauthnUserID,
		Credential: []webauthn.Credential{
			*credential,
		},

		Name:        user.UserName,
		DisplayName: displayName,
		Icon:        icon,
	}, nil
}

// webauthnのユーザを作成する
// WebAuthnで使用するユーザIDを生成します
func NewWebAuthnUserRegister(email string) (*WebAuthnUser, error) {
	id, err := lib.RandomBytes(64)
	if err != nil {
		return nil, err
	}

	return &WebAuthnUser{
		ID:         id,
		Credential: []webauthn.Credential{},

		Name:        email,
		DisplayName: email,
		Icon:        "",
	}, nil
}

func (w *WebAuthnUser) WebAuthnID() []byte {
	return w.ID
}

func (w *WebAuthnUser) WebAuthnName() string {
	return w.Name
}

func (w *WebAuthnUser) WebAuthnDisplayName() string {
	return w.DisplayName
}

func (w *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return w.Credential
}

func (w *WebAuthnUser) WebAuthnIcon() string {
	return w.Icon
}
