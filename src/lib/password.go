package lib

import (
	"bytes"

	"golang.org/x/crypto/argon2"
)

type PasswordInterface interface {
	HashPassword(password string) ([]byte, []byte, error)
	VerifyPassword(password string, hashedPassword []byte, salt []byte) bool
}

type Password struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

func (p *Password) HashPassword(password string) ([]byte, []byte, error) {
	salt, err := RandomBytes(int(p.KeyLen))
	if err != nil {
		return nil, nil, err
	}

	return argon2.IDKey([]byte(password), salt, p.Time, p.Memory, p.Threads, p.KeyLen), salt, nil
}

func (p *Password) VerifyPassword(password string, hashedPassword []byte, salt []byte) bool {
	hashed := argon2.IDKey([]byte(password), salt, p.Time, p.Memory, p.Threads, p.KeyLen)

	return bytes.Equal(hashedPassword, hashed)
}
