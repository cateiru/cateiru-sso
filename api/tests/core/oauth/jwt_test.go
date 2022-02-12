package oauth_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/oauth"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestCreateJWT(t *testing.T) {
	config.TestInit(t)

	dummy := tools.NewDummyUser()

	user := models.User{
		FirstName:         "太郎",
		LastName:          "テスト",
		UserName:          "Taro",
		UserNameFormatted: "taro",

		Mail:      "taro@example.com",
		Theme:     "dark",
		AvatarUrl: "https://example.com/hogehoge",

		Role: []string{},

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	claims := oauth.NewJWT(&user, utils.CreateID(30), time.Now())

	result, err := claims.ConvertJWT()
	require.NoError(t, err)

	// --- 検証

	key, err := oauth.GetPublicKey()
	require.NoError(t, err)

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(key.PKCS8)
	require.NoError(t, err)

	token, err := jwt.Parse(result, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return verifyKey, nil
	})
	require.NoError(t, err)
	require.True(t, token.Valid)

	require.NotEmpty(t, token.Raw)
}
