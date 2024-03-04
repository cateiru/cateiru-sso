package lib_test

import (
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

// OTPが生成でき、検証し成功するかチェック
func TestOTP(t *testing.T) {
	accountName := "TestUser"
	issuer := "Issuer"

	o, err := lib.NewOTP(issuer, accountName)
	require.NoError(t, err)

	o2, err := lib.NewOTP(issuer, accountName)
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		public := o.GetPublic()
		n := time.Now().UTC()

		w, err := otp.NewKeyFromURL(public)
		require.NoError(t, err)

		code, err := totp.GenerateCode(w.Secret(), n)
		require.NoError(t, err)

		require.True(t, lib.ValidateOTP(code, o.GetSecret()))
	})

	t.Run("failed", func(t *testing.T) {
		public := o.GetPublic()
		n := time.Now().UTC()

		w, err := otp.NewKeyFromURL(public)
		require.NoError(t, err)

		code, err := totp.GenerateCode(w.Secret(), n)
		require.NoError(t, err)

		require.False(t, lib.ValidateOTP(code, o2.GetSecret()))
	})
}
