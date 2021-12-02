// ハッシュ生成
//
// Example:
//	seed := "hoge"
//	seed2 := "fuga"
//	hash := NewHash(seed)
//	hash = hash.AddSeed(seed2)
//
//	fmt.Println(hash.SHA256())
//	fmt.Println(hash.SHA256Byte())
//
package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
)

type Hash struct {
	Seeds []string
}

func NewHash(seed string) *Hash {
	return &Hash{
		Seeds: []string{seed},
	}
}

// seedを追加する
func (c *Hash) AddSeed(seed string) *Hash {
	return &Hash{
		Seeds: append(c.Seeds, seed),
	}
}

// SHA256 hashを生成する
func (c *Hash) SHA256() string {
	return hex.EncodeToString(c.SHA256Byte())
}

// SHA256 hashをbyteで生成する
func (c *Hash) SHA256Byte() []byte {
	hash := sha256.Sum256([]byte(strings.Join(c.Seeds, "")))
	return hash[:]
}

// パスワードとSEEDを使用してパスワードをハッシュ化します
// DBにパスワードを保存する場合は必ずハッシュ化して保存すること
func PWHash(pw string) string {
	hash := NewHash(pw).AddSeed(os.Getenv("PW_HASH_SEED"))
	return hash.SHA256()
}
