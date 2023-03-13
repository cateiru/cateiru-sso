package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

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
		if user.MiddleName.Valid {
			displayName = fmt.Sprintf("%s %s %s", user.FamilyName.String, user.MiddleName.String, user.GivenName.String)
		} else {
			displayName = fmt.Sprintf("%s %s", user.FamilyName.String, user.GivenName.String)
		}
	} else {
		// 設定していない場合はEmail
		displayName = user.Email
	}

	icon := ""
	if user.Avatar.Valid {
		icon = user.Avatar.String
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

func NewWebauthnUserFromUser(user *models.User) (*WebAuthnUser, error) {
	displayName := ""
	if user.FamilyName.Valid && user.GivenName.Valid {
		// 名前が設定されている場合はフォーマットする
		// 順序は姓名と、日本式
		if user.MiddleName.Valid {
			displayName = fmt.Sprintf("%s %s %s", user.FamilyName.String, user.MiddleName.String, user.GivenName.String)
		} else {
			displayName = fmt.Sprintf("%s %s", user.FamilyName.String, user.GivenName.String)
		}
	} else {
		// 設定していない場合はEmail
		displayName = user.Email
	}

	icon := ""
	if user.Avatar.Valid {
		icon = user.Avatar.String
	}

	id, err := lib.RandomBytes(64)
	if err != nil {
		return nil, err
	}

	return &WebAuthnUser{
		ID:         id,
		Credential: []webauthn.Credential{},

		Name:        user.UserName,
		DisplayName: displayName,
		Icon:        icon,
	}, nil
}

// 新規作成で使用するwebauthnのユーザを作成する
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

// Webauthnを登録する
func (h *Handler) RegisterWebauthn(ctx context.Context, body io.Reader, webauthnSessionToken string, identifier int8) (*webauthn.Credential, error) {
	response, err := h.WebAuthn.ParseCreate(body)
	if err != nil {
		return nil, NewHTTPError(http.StatusBadRequest, err)
	}

	webauthnSession, err := models.WebauthnSessions(
		models.WebauthnSessionWhere.ID.EQ(webauthnSessionToken),
		models.WebauthnSessionWhere.Identifier.EQ(identifier),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewHTTPError(http.StatusForbidden, "invalid webauthn token")
	}
	if err != nil {
		return nil, err
	}

	if time.Now().After(webauthnSession.Period) {
		// webauthnセッションは削除
		_, err := webauthnSession.Delete(ctx, h.DB)
		if err != nil {
			return nil, err
		}
		return nil, NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	// Rowから取得する
	session := new(webauthn.SessionData)
	err = webauthnSession.Row.Unmarshal(session)
	if err != nil {
		return nil, err
	}

	webauthnUser := &WebAuthnUser{
		ID:         webauthnSession.WebauthnUserID,
		Credential: []webauthn.Credential{},

		Name:        "",
		DisplayName: webauthnSession.UserDisplayName,
		Icon:        "",
	}

	credential, err := h.WebAuthn.FinishRegistration(webauthnUser, *session, response)
	if err != nil {
		return nil, NewHTTPError(http.StatusForbidden, err)
	}

	// WebauthnSessionは削除
	if _, err := webauthnSession.Delete(ctx, h.DB); err != nil {
		return nil, err
	}

	return credential, nil
}

// Webauthnでログインする
func (h *Handler) LoginWebauthn(ctx context.Context, body io.Reader, webauthnSessionToken string, identifier int8, before func(u *models.User) error) (*models.User, error) {
	response, err := h.WebAuthn.ParseLogin(body)
	if err != nil {
		return nil, NewHTTPError(http.StatusBadRequest, err)
	}

	webauthnSession, err := models.WebauthnSessions(
		models.WebauthnSessionWhere.ID.EQ(webauthnSessionToken),
		models.WebauthnSessionWhere.Identifier.EQ(identifier),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewHTTPError(http.StatusForbidden, "invalid webauthn token")
	}
	if err != nil {
		return nil, err
	}
	if !webauthnSession.UserID.Valid || webauthnSession.UserID.String == "" {
		return nil, NewHTTPError(http.StatusInternalServerError, "user is empty")
	}

	if time.Now().After(webauthnSession.Period) {
		// webauthnセッションは削除
		_, err := webauthnSession.Delete(ctx, h.DB)
		if err != nil {
			return nil, err
		}
		return nil, NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	// Rowから取得する
	session := new(webauthn.SessionData)
	err = webauthnSession.Row.Unmarshal(session)
	if err != nil {
		return nil, err
	}

	webauthnUser := &WebAuthnUser{
		ID:         webauthnSession.WebauthnUserID,
		Credential: []webauthn.Credential{},

		Name:        "",
		DisplayName: webauthnSession.UserDisplayName,
		Icon:        "",
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(webauthnSession.UserID.String),
	).One(ctx, h.DB)
	if err != nil {
		return nil, err
	}

	if err := before(user); err != nil {
		return nil, err
	}

	_, err = h.WebAuthn.FinishLogin(webauthnUser, *session, response)
	if err != nil {
		return nil, NewHTTPError(http.StatusForbidden, err)
	}

	// WebauthnSessionは削除
	if _, err := webauthnSession.Delete(ctx, h.DB); err != nil {
		return nil, err
	}

	return user, nil
}
