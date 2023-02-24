package src_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func TestRegisterUser(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		u, err := src.RegisterUser(ctx, DB, email)
		require.NoError(t, err)

		require.Equal(t, u.Email, email)
		require.Len(t, u.UserName, 8)
	})

	t.Run("すでにEmailが存在している場合はエラー", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		RegisterUser(t, ctx, email)

		_, err := src.RegisterUser(ctx, DB, email)
		require.EqualError(t, err, "code=400, message=impossible register account, unique=3")
	})
}
