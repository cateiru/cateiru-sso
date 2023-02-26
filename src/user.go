package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
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

// webauthn用のuserとsessionをDBから取得して組み立てる
func NewWebAuthnUserSession(ctx context.Context, db *sql.DB, webauthnSession string) (*WebAuthnUser, *webauthn.SessionData, error) {
	auth, err := models.WebauthnSessions(
		models.WebauthnSessionWhere.ID.EQ(webauthnSession),
	).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, NewHTTPError(http.StatusForbidden, "invalid webauthn token")
	}
	if err != nil {
		return nil, nil, err
	}

	if time.Now().After(auth.Period) {
		// webauthnセッションは削除
		_, err := auth.Delete(ctx, db)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, NewHTTPError(http.StatusForbidden, "invalid webauthn token")
	}

	session := &webauthn.SessionData{
		Challenge:        auth.Challenge,
		UserID:           auth.WebauthnUserID,
		UserDisplayName:  auth.UserDisplayName,
		UserVerification: protocol.UserVerificationRequirement(auth.UserVerification),
		// FIXME: 他のプロパティはどうなんだろう？
	}

	return &WebAuthnUser{
		ID:         auth.WebauthnUserID,
		Credential: []webauthn.Credential{},

		Name:        "",
		DisplayName: auth.UserDisplayName,
		Icon:        "",
	}, session, nil
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

// ユーザを新規に作成する
// 最初は、ユーザ名などの情報はデフォルト値に設定する（ユーザ登録フローの簡略化のため）
func RegisterUser(ctx context.Context, db *sql.DB, email string) (*models.User, error) {
	// もう一度Emailが登録されていないか確認する
	exist, err := models.Users(models.UserWhere.Email.EQ(email)).Exists(ctx, db)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, NewHTTPUniqueError(http.StatusBadRequest, ErrImpossibleRegisterAccount, "impossible register account")
	}

	id := ulid.Make()
	idBin, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}

	u := models.User{
		ID:    idBin,
		Email: email,
	}
	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	L.Info("register user",
		zap.String("email", email),
	)

	return models.Users(
		models.UserWhere.ID.EQ(idBin),
	).One(ctx, db)
}
