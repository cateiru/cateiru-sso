package lib

import (
	"os"

	"github.com/go-jose/go-jose/v3"
	"golang.org/x/crypto/ssh"
)

// jwk を返す
func JsonWebKeys(publicKeyFileName string, algorithm string, use string, keyId string) (*jose.JSONWebKey, error) {
	bytes, err := os.ReadFile(publicKeyFileName)
	if err != nil {
		return nil, err
	}
	key, err := ssh.ParsePublicKey(bytes)
	if err != nil {
		return nil, err
	}

	pub := jose.JSONWebKey{
		Key:       key,
		KeyID:     keyId,
		Algorithm: algorithm,
		Use:       use,
	}

	return &pub, nil
}
