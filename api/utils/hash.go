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
	c.Seeds = append(c.Seeds, seed)
	return c
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
