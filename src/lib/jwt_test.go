package lib_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

const JWT_PUBLIC_KEY_PATH = "jwt/test.pub.pkcs8"
const JWT_PRIVATE_KEY_PATH = "jwt/test"

func TestJsonWebKeys(t *testing.T) {
	resp, err := lib.JsonWebKeys(JWT_PUBLIC_KEY_PATH, "RS256", "sig", "test")
	require.NoError(t, err)

	snaps.MatchSnapshot(t, resp)
}

func TestSignJwt(t *testing.T) {
	claims := &jwt.StandardClaims{
		Audience:  "test",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        "test",
		Issuer:    "test",
		NotBefore: time.Now().Unix(),
		Subject:   "test",
	}

	signed, err := lib.SignJwt(claims, JWT_PRIVATE_KEY_PATH)
	require.NoError(t, err)

	// JWTを検証する
	public, err := os.ReadFile(JWT_PUBLIC_KEY_PATH)
	require.NoError(t, err)
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(public)
	require.NoError(t, err)

	parsedClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(signed, &parsedClaims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return verifyKey, nil
	})
	require.NoError(t, err)
	require.True(t, token.Valid)

	require.Equal(t, *claims, parsedClaims)
}
