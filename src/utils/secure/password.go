package secure

import (
	"bytes"
	"errors"

	"github.com/cateiru/cateiru-sso/src/utils"
	"golang.org/x/crypto/argon2"
)

type HashedPassword struct {
	Key  []byte
	Salt []byte
}

// パスワードとSEEDを使用してパスワードをハッシュ化します
// DBにパスワードを保存する場合は必ずハッシュ化して保存すること
//
// 最大長: 128文字
//
//   - ハッシュアルゴリズム: Argon2
//   - Time: 1s
//   - Memory cost: 64*1024 (64MB)
//   - Thured: 1
//   - key length: 32
func PWHash(pw string) (*HashedPassword, error) {
	if len(pw) > 128 {
		return nil, errors.New("password length is up to 128 characters")
	}

	salt := utils.NewHash(utils.UUID()).SHA256Byte()
	key := argon2.IDKey([]byte(pw), salt, 1, 64*1024, 4, 32)

	return &HashedPassword{
		Key:  key,
		Salt: salt,
	}, nil
}

// パスワードを検証します
func ValidatePW(pw string, hashedPw []byte, salt []byte) bool {
	key := argon2.IDKey([]byte(pw), []byte(salt), 1, 64*1024, 4, 32)

	return bytes.Equal(key, hashedPw)
}
