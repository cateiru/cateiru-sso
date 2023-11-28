package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/go-webauthn/webauthn/protocol"
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
	webauthnCredentials, err := models.Webauthns(
		models.WebauthnWhere.UserID.EQ(user.ID),
	).All(ctx, db)
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

	credentials := []webauthn.Credential{}
	for _, webauthnCredential := range webauthnCredentials {
		credential := &webauthn.Credential{}
		if err := webauthnCredential.Credential.Unmarshal(credential); err != nil {
			return nil, err
		}
		credentials = append(credentials, *credential)
	}

	return &WebAuthnUser{
		ID:         []byte(user.ID),
		Credential: credentials,

		Name:        user.UserName,
		DisplayName: displayName,
		Icon:        icon,
	}, nil
}

func NewWebAuthnUserNoCredential(user *models.User) (*WebAuthnUser, error) {
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

	return &WebAuthnUser{
		ID:         []byte(user.ID),
		Credential: []webauthn.Credential{},

		Name:        user.UserName,
		DisplayName: displayName,
		Icon:        icon,
	}, nil
}

// ユーザが存在しない状態でWebAuthnを登録する場合に使用します
func NewWebAuthnUserRegister(email string, id []byte) (*WebAuthnUser, error) {
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
func (h *Handler) RegisterWebauthn(ctx context.Context, body io.Reader, webauthnSessionToken string, identifier int8) (*webauthn.Credential, webauthn.User, error) {
	response, err := h.WebAuthn.ParseCreate(body)
	if err != nil {
		return nil, nil, NewHTTPError(http.StatusBadRequest, err)
	}

	webauthnSession, err := models.WebauthnSessions(
		models.WebauthnSessionWhere.ID.EQ(webauthnSessionToken),
		models.WebauthnSessionWhere.Identifier.EQ(identifier),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, NewHTTPError(http.StatusForbidden, "invalid webauthn token")
	}
	if err != nil {
		return nil, nil, err
	}

	if time.Now().After(webauthnSession.Period) {
		// webauthnセッションは削除
		_, err := webauthnSession.Delete(ctx, h.DB)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, NewHTTPUniqueError(http.StatusForbidden, ErrExpired, "expired token")
	}

	// Rowから取得する
	session := new(webauthn.SessionData)
	err = webauthnSession.Row.Unmarshal(session)
	if err != nil {
		return nil, nil, err
	}

	createUser := func() (*WebAuthnUser, error) {
		if !webauthnSession.UserID.Valid {
			return nil, NewHTTPError(http.StatusInternalServerError, "user not found")
		}

		user, err := models.Users(
			models.UserWhere.ID.EQ(webauthnSession.UserID.String),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return &WebAuthnUser{
				ID:         []byte(webauthnSession.UserID.String),
				Credential: []webauthn.Credential{},

				// 認証には使わないので空
				Name:        "",
				DisplayName: "",
				Icon:        "",
			}, nil
		}
		if err != nil {
			return nil, err
		}

		// user存在するならちゃんと埋めてあげる
		return NewWebAuthnUserNoCredential(user)
	}
	webauthnUser, err := createUser()
	if err != nil {
		return nil, nil, err
	}

	credential, err := h.WebAuthn.FinishRegistration(webauthnUser, *session, response)
	if err != nil {
		return nil, nil, NewHTTPError(http.StatusForbidden, err)
	}

	// WebauthnSessionは削除
	if _, err := webauthnSession.Delete(ctx, h.DB); err != nil {
		return nil, nil, err
	}

	return credential, webauthnUser, nil
}

// Webauthnでログインする
func (h *Handler) LoginWebauthn(ctx context.Context, body io.Reader, webauthnSessionToken string, identifier int8) (*models.User, error) {
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

	// idからユーザーとクレデンシャルを引く
	var loginUser models.User
	handler := func(rawID, userHandle []byte) (user webauthn.User, err error) {
		u, err := models.Users(
			models.UserWhere.ID.EQ(string(userHandle)),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewHTTPError(http.StatusForbidden, "invalid user")
		}
		if err != nil {
			return nil, err
		}
		loginUser = *u
		return NewWebAuthnUserFromDB(ctx, h.DB, u)
	}

	_, err = h.WebAuthn.FinishLogin(handler, *session, response)
	if protocolError, ok := err.(*protocol.Error); ok {
		return nil, NewHTTPUniqueError(http.StatusBadRequest, ErrLoginFailed, protocolError.Details)
	}
	if err != nil {
		return nil, NewHTTPError(http.StatusForbidden, err)
	}

	// WebauthnSessionは削除
	if _, err := webauthnSession.Delete(ctx, h.DB); err != nil {
		return nil, err
	}

	return &loginUser, nil
}
