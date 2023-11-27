package lib

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/go-jose/go-jose/v3"
)

// jwk を返す
func JsonWebKeys(publicKeyFileName string, algorithm string, use string, keyId string) (*jose.JSONWebKey, error) {
	bytes, err := os.ReadFile(publicKeyFileName)
	if err != nil {
		return nil, err
	}

	// ref. https://stackoverflow.com/questions/70718821/go-rsa-load-public-key
	spkiBlock, _ := pem.Decode(bytes)
	var publicKey *rsa.PublicKey
	pubInterface, _ := x509.ParsePKIXPublicKey(spkiBlock.Bytes)
	publicKey = pubInterface.(*rsa.PublicKey)

	pub := jose.JSONWebKey{
		Key:       publicKey,
		KeyID:     keyId,
		Algorithm: algorithm,
		Use:       use,
	}

	return &pub, nil
}
