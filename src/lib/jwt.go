package lib

import (
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt"
)

// jwk を返す
func JsonWebKeys(publicKeyFilePath string, algorithm string, use string, keyId string) (*jose.JSONWebKey, error) {
	bytes, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}

	// ref. https://stackoverflow.com/questions/70718821/go-rsa-load-public-key
	spkiBlock, _ := pem.Decode(bytes)
	publicKey, err := x509.ParsePKIXPublicKey(spkiBlock.Bytes)
	if err != nil {
		return nil, err
	}

	pub := jose.JSONWebKey{
		Key:       publicKey,
		KeyID:     keyId,
		Algorithm: algorithm,
		Use:       use,
	}

	return &pub, nil
}

// JWTを署名する
// 秘密鍵は都度生成
// TODO: テスト
func SignJwt(claims *jwt.Claims, secretKeyFilePath string) (string, error) {
	secret, err := os.ReadFile(secretKeyFilePath)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, *claims)

	return token.SignedString(secret)
}
