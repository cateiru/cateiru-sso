package src_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func TestNewWebAuthUserRegister(t *testing.T) {
	r, err := lib.RandomStr(10)
	require.NoError(t, err)
	email := fmt.Sprintf("%s@exmaple.com", r)

	user, err := src.NewWebAuthnUserRegister(email)
	require.NoError(t, err)

	t.Run("それぞれのメソッドが正しく返せる", func(t *testing.T) {
		require.Len(t, user.WebAuthnID(), 64)
		require.Equal(t, user.WebAuthnName(), email)
		require.Equal(t, user.WebAuthnDisplayName(), email)
		require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{})
		require.Equal(t, user.WebAuthnIcon(), "")
	})
}

func TestNewWebauthnUserFromUser(t *testing.T) {
	t.Run("それぞれのメソッドが正しく返せる", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		user, err := src.NewWebauthnUserFromUser(&u)
		require.NoError(t, err)

		require.Len(t, user.WebAuthnID(), 64)
		require.Equal(t, user.WebAuthnName(), u.UserName)
		require.Equal(t, user.WebAuthnDisplayName(), u.Email)
		require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{})
		require.Equal(t, user.WebAuthnIcon(), "")
	})

	t.Run("FamilyNameとGivenNameが設定されている", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		u.FamilyName = null.NewString("Family", true)
		u.GivenName = null.NewString("Given", true)

		user, err := src.NewWebauthnUserFromUser(&u)
		require.NoError(t, err)

		require.Len(t, user.WebAuthnID(), 64)
		require.Equal(t, user.WebAuthnName(), u.UserName)
		require.Equal(t, user.WebAuthnDisplayName(), "Family Given")
		require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{})
		require.Equal(t, user.WebAuthnIcon(), "")
	})

	t.Run("MiddleNameがある", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		u.FamilyName = null.NewString("Family", true)
		u.GivenName = null.NewString("Given", true)
		u.MiddleName = null.NewString("Middle", true)

		user, err := src.NewWebauthnUserFromUser(&u)
		require.NoError(t, err)

		require.Len(t, user.WebAuthnID(), 64)
		require.Equal(t, user.WebAuthnName(), u.UserName)
		require.Equal(t, user.WebAuthnDisplayName(), "Family Middle Given")
		require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{})
		require.Equal(t, user.WebAuthnIcon(), "")
	})

	t.Run("icon設定されている", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		u.Avatar = null.NewString("https://example.com/avater", true)

		user, err := src.NewWebauthnUserFromUser(&u)
		require.NoError(t, err)

		require.Len(t, user.WebAuthnID(), 64)
		require.Equal(t, user.WebAuthnName(), u.UserName)
		require.Equal(t, user.WebAuthnDisplayName(), u.Email)
		require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{})
		require.Equal(t, user.WebAuthnIcon(), "https://example.com/avater")
	})
}

func TestNewWebAuthnUserFromDB(t *testing.T) {
	ctx := context.Background()

	t.Run("成功", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		u := RegisterUser(t, ctx, email)

		uid, err := lib.RandomBytes(64)
		require.NoError(t, err)

		dummyCredential := webauthn.Credential{}
		dummyCredential.ID = uid
		dummyCredentialBytes, err := json.Marshal(dummyCredential)
		require.NoError(t, err)

		passkey := models.Passkey{
			UserID:          u.ID,
			WebauthnUserID:  uid,
			Credential:      dummyCredentialBytes,
			FlagBackupState: false,
		}
		err = passkey.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		user, err := src.NewWebAuthnUserFromDB(ctx, DB, &u)
		require.NoError(t, err)

		t.Run("それぞれのメソッドが正しく返せる", func(t *testing.T) {
			require.Len(t, user.WebAuthnID(), 64)
			require.Equal(t, user.WebAuthnName(), u.UserName)
			require.Equal(t, user.WebAuthnDisplayName(), email)
			require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{dummyCredential})
			require.Equal(t, user.WebAuthnIcon(), "")
		})
	})

	t.Run("名前が設定されている場合、displayNameは名前になる", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@exmaple.com", r)

		u := RegisterUser(t, ctx, email)

		u.FamilyName = null.NewString("Test", true)
		u.GivenName = null.NewString("Taro", true)
		_, err = u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		uid, err := lib.RandomBytes(64)
		require.NoError(t, err)

		dummyCredential := webauthn.Credential{}
		dummyCredential.ID = uid
		dummyCredentialBytes, err := json.Marshal(dummyCredential)
		require.NoError(t, err)

		passkey := models.Passkey{
			UserID:          u.ID,
			WebauthnUserID:  uid,
			Credential:      dummyCredentialBytes,
			FlagBackupState: false,
		}
		err = passkey.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		user, err := src.NewWebAuthnUserFromDB(ctx, DB, &u)
		require.NoError(t, err)

		t.Run("それぞれのメソッドが正しく返せる", func(t *testing.T) {
			require.Len(t, user.WebAuthnID(), 64)
			require.Equal(t, user.WebAuthnName(), u.UserName)
			require.Equal(t, user.WebAuthnDisplayName(), "Test Taro")
			require.Equal(t, user.WebAuthnCredentials(), []webauthn.Credential{dummyCredential})
			require.Equal(t, user.WebAuthnIcon(), "")
		})
	})
}

func TestRegisterWebauthn(t *testing.T) {
	h := NewTestHandler(t)
	createWebauthSession := func(ctx context.Context, userId []byte) string {
		webauthnSessionId, err := lib.RandomStr(31)
		require.NoError(t, err)

		challenge, err := lib.RandomStr(10)
		require.NoError(t, err)

		session := &webauthn.SessionData{
			Challenge:        challenge,
			UserID:           userId,
			UserDisplayName:  "test taro",
			UserVerification: protocol.VerificationRequired,
		}
		row := types.JSON{}
		err = row.Marshal(session)
		require.NoError(t, err)

		webauthnSession := models.WebauthnSession{
			ID:               webauthnSessionId,
			WebauthnUserID:   session.UserID,
			UserDisplayName:  session.UserDisplayName,
			Challenge:        session.Challenge,
			UserVerification: string(session.UserVerification),
			Row:              row,

			Period: time.Now().Add(C.WebAuthnSessionPeriod),
		}
		err = webauthnSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return webauthnSessionId
	}

	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		userId, err := lib.RandomBytes(10)
		require.NoError(t, err)

		webauthSession := createWebauthSession(ctx, userId)

		credential, err := h.RegisterWebauthn(ctx, nil, webauthSession)
		require.NoError(t, err)
		require.NotNil(t, credential)

		// webauthnSessionは削除されている
		existsWebauthnSession, err := models.WebauthnSessionExists(ctx, DB, webauthSession)
		require.NoError(t, err)
		require.False(t, existsWebauthnSession)
	})

	t.Run("有効期限切れ", func(t *testing.T) {
		ctx := context.Background()
		userId, err := lib.RandomBytes(10)
		require.NoError(t, err)

		webauthSession := createWebauthSession(ctx, userId)

		// 有効期限を切らす
		s, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(webauthSession),
		).One(ctx, DB)
		require.NoError(t, err)
		s.Period = time.Now().Add(-10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		_, err = h.RegisterWebauthn(ctx, nil, webauthSession)
		require.EqualError(t, err, "code=403, message=expired token, unique=5")

		// セッションは削除されている
		exists, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(webauthSession),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, exists)
	})
}

func TestLoginWebauthn(t *testing.T) {
	h := NewTestHandler(t)
	createWebauthSession := func(ctx context.Context, userId []byte, u *models.User) string {
		webauthnSessionId, err := lib.RandomStr(31)
		require.NoError(t, err)

		challenge, err := lib.RandomStr(10)
		require.NoError(t, err)

		session := &webauthn.SessionData{
			Challenge:        challenge,
			UserID:           userId,
			UserDisplayName:  "test taro",
			UserVerification: protocol.VerificationRequired,
		}
		row := types.JSON{}
		err = row.Marshal(session)
		require.NoError(t, err)

		webauthnSession := models.WebauthnSession{
			ID:               webauthnSessionId,
			UserID:           null.NewString(u.ID, true),
			WebauthnUserID:   session.UserID,
			UserDisplayName:  session.UserDisplayName,
			Challenge:        session.Challenge,
			UserVerification: string(session.UserVerification),
			Row:              row,

			Period: time.Now().Add(C.WebAuthnSessionPeriod),
		}
		err = webauthnSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		return webauthnSessionId
	}

	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)
		userId, err := lib.RandomBytes(10)
		require.NoError(t, err)

		webauthSession := createWebauthSession(ctx, userId, &user)

		loggedInUser, err := h.LoginWebauthn(ctx, nil, webauthSession, func(u *models.User) error { return nil })
		require.NoError(t, err)

		require.Equal(t, loggedInUser.ID, user.ID)

		// webauthnSessionは削除されている
		existsWebauthnSession, err := models.WebauthnSessionExists(ctx, DB, webauthSession)
		require.NoError(t, err)
		require.False(t, existsWebauthnSession)
	})

	t.Run("有効期限切れ", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		user := RegisterUser(t, ctx, email)
		userId, err := lib.RandomBytes(10)
		require.NoError(t, err)

		webauthSession := createWebauthSession(ctx, userId, &user)

		// 有効期限を切らす
		s, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(webauthSession),
		).One(ctx, DB)
		require.NoError(t, err)
		s.Period = time.Now().Add(-10 * time.Hour)
		_, err = s.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		_, err = h.LoginWebauthn(ctx, nil, webauthSession, func(u *models.User) error { return nil })
		require.EqualError(t, err, "code=403, message=expired token, unique=5")

		// セッションは削除されている
		exists, err := models.WebauthnSessions(
			models.WebauthnSessionWhere.ID.EQ(webauthSession),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, exists)
	})
}
