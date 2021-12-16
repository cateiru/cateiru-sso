package secure

import (
	"bytes"

	"github.com/cateiru/cateiru-sso/api/utils"
	"golang.org/x/crypto/argon2"
)

type HashedPassword struct {
	Key  []byte
	Salt []byte
}

// パスワードとSEEDを使用してパスワードをハッシュ化します
// DBにパスワードを保存する場合は必ずハッシュ化して保存すること
//
//	- ハッシュアルゴリズム: Argon2
//	- Time: 1s
//	- Memory cost: 64*1024 (64MB)
//	- Thured: 1
//	- key length: 32
func PWHash(pw string) *HashedPassword {
	salt := utils.NewHash(utils.UUID()).SHA256Byte()
	key := argon2.IDKey([]byte(pw), salt, 1, 64*1024, 4, 32)

	return &HashedPassword{
		Key:  key,
		Salt: salt,
	}
}

// パスワードを検証します
func ValidatePW(pw string, hashedPw []byte, salt []byte) bool {
	key := argon2.IDKey([]byte(pw), []byte(salt), 1, 64*1024, 4, 32)

	return bytes.Equal(key, hashedPw)
}
