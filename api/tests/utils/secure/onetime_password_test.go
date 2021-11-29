package secure_test

import (
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

// OTPが生成でき、検証し成功するかチェック
func TestValidateOnetimePassword(t *testing.T) {
	t.Setenv("ONETIME_PASSWORD_ISSUER", "test_issuer")

	accountName := "TestUser"

	_otp, err := secure.NewOnetimePassword(accountName)
	require.NoError(t, err, "OTPが作成できない")

	public := _otp.GetPublic()
	n := time.Now().UTC()

	w, err := otp.NewKeyFromURL(public)
	require.NoError(t, err, "publicを解析できない")

	code, err := totp.GenerateCode(w.Secret(), n)
	require.NoError(t, err, "URLからパスコードを生成できない")

	isValidate := secure.ValidateOnetimePassword(code, _otp.GetSecret())

	require.True(t, isValidate, "検証が失敗した")

}

// 違うパスコードで検証が失敗するか
func TestValidateFailed(t *testing.T) {
	t.Setenv("ONETIME_PASSWORD_ISSUER", "test_issuer")

	accountName := "TestUser"

	_otp, err := secure.NewOnetimePassword(accountName)
	require.NoError(t, err, "OTPが作成できない")

	n := time.Now().UTC()

	code, err := totp.GenerateCode(_otp.GetSecret(), n)
	require.NoError(t, err, "URLからパスコードを生成できない")

	// 同じアカウント名で別の鍵生成
	secondOtp, err := secure.NewOnetimePassword(accountName)
	require.NoError(t, err, "OTPが作成できない")

	isValidate := secure.ValidateOnetimePassword(code, secondOtp.GetSecret())

	require.False(t, isValidate, "検証が違うsecretで成功してしまった")
}

// 複数のOTPは全部違うsecretか
func TestOTPMulti(t *testing.T) {
	t.Setenv("ONETIME_PASSWORD_ISSUER", "test_issuer")
	secrets := []string{}
	accountName := "TestUser"

	length := 10

	for i := 0; length > i; i++ {
		_otp, err := secure.NewOnetimePassword(accountName)
		require.NoError(t, err, "OTPを作成できない")
		secrets = append(secrets, _otp.GetSecret())
	}

	for i, a := range secrets {
		for j := i + 1; length > j; j++ {
			require.NotEqual(t, a, secrets[j])
		}
	}
}
