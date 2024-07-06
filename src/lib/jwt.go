package lib

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// jwk を返す
func JsonWebKeys(publicKeyFilePath string, algorithm string, use string, keyId string) (jwk.Key, error) {
	bytes, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}

	keyset, err := jwk.ParseKey(bytes, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	keyset.Set(jwk.KeyIDKey, keyId)
	keyset.Set(jwk.AlgorithmKey, algorithm)
	keyset.Set(jwk.KeyUsageKey, use)

	return keyset, nil
}

// JWTを署名する
// 秘密鍵は都度読み込む
func SignJwt(claims jwt.Claims, secretKeyFilePath string) (string, error) {
	secret, err := os.ReadFile(secretKeyFilePath)
	if err != nil {
		return "", err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(secret)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}
