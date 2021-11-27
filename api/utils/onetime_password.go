// ワンタイムパスワードを作成、検証します。
//
// Example:
//	accountName := "example@example.com"
//	otp, _ := OnetimePasswordNew(accountName)
//
//	secret := GetSecret() // データベースに保存
//	public := GetPublic() // ユーザに提供
//
//	passcode := Send(public) // Public()は例: ユーザにpublicを提供しそれを使用してパスコードを生成してもらいそれをもらう
//	if ValidateOnetimePassword(passcode, secret) {
//		fmt.Println("OK")
//		SaveSecret(secret) // SaveSecret()は例: secretを保存する
//	} else {
//		fmt.Println("NO")
//	}
//
// Note: secretをデータベースに保存する前に一度、ユーザにパスコードを生成してもらい結果を検証してください
package utils

import (
	"crypto/rand"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OnetimePassword struct {
	GenerateOpt *totp.GenerateOpts
	Key         *otp.Key
}

// OTP認証を使用します
//
// OTP仕様
//	- 有効時間: 30s
//	- Secretバイト長: 20byte
//	- TOTPハッシュ桁: 6桁
//	- ハッシュアルゴリズム: SHA1
//	- 乱数生成: rand.Reader
func OnetimePasswordNew(accountName string) (*OnetimePassword, error) {
	ops := totp.GenerateOpts{
		Issuer:      ONETIME_PASSWORD_ISSUER,
		AccountName: accountName,
		Period:      30,
		SecretSize:  20,
		Secret:      ONETIME_PASSWORD_SECRET,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        rand.Reader,
	}

	key, err := totp.Generate(ops)
	if err != nil {
		return nil, err
	}

	return &OnetimePassword{
		GenerateOpt: &ops,
		Key:         key,
	}, nil
}

// ワンタイムパスワードのpublic
// ユーザはこのkeyをAuthenticator appに入力することで、ワンタイムパスワードを登録できます。
func (o *OnetimePassword) GetPublic() string {
	return o.Key.String()
}

// ワンタイムパスワードのSecret keyを取得する。
// このSecretはサーバーサイドで検証するためユーザには提供しないようにしてください。
//
// Example:
//	secret = o.GetSecret()
//	passcode = ""
//	if totp.Validate(passcode, secret) {
//		...
//	}
func (o *OnetimePassword) GetSecret() string {
	return o.Key.Secret()
}

// ワンタイムパスワードを検証する。
// passcodeは、ユーザから取得したパスコード。
// secretは、サーバー内で保存するkey
func ValidateOnetimePassword(passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}
