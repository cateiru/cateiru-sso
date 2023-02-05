package lib

import (
	"crypto/rand"

	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OTP struct {
	GenerateOpt *totp.GenerateOpts
	Key         *otp.Key
}

// OTP認証を使用します
//
// OTP仕様
//   - 有効時間: 30s
//   - Secretバイト長: 20byte
//   - TOTPハッシュ桁: 6桁
//   - ハッシュアルゴリズム: SHA1
//   - 乱数生成: rand.Reader
func NewOTP(issuer string, accountName string) (*OTP, error) {

	ops := totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		Period:      30,
		SecretSize:  20,
		Secret:      []byte(uuid.NewString()),
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        rand.Reader,
	}

	key, err := totp.Generate(ops)
	if err != nil {
		return nil, err
	}

	return &OTP{
		GenerateOpt: &ops,
		Key:         key,
	}, nil
}

// ワンタイムパスワードのpublic
// ユーザはこのkeyをAuthenticator appに入力することで、ワンタイムパスワードを登録できます。
func (o *OTP) GetPublic() string {
	return o.Key.String()
}

// ワンタイムパスワードのSecret keyを取得する。
// このSecretはサーバーサイドで検証するためユーザには提供しないようにしてください。
//
// Example:
//
//	secret = o.GetSecret()
//	passcode = ""
//	if totp.Validate(passcode, secret) {
//		...
//	}
func (o *OTP) GetSecret() string {
	return o.Key.Secret()
}

// ワンタイムパスワードを検証する。
// passcodeは、ユーザから取得したパスコード。
// secretは、サーバー内で保存するkey
func ValidateOTP(passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}
