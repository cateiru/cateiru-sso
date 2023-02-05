package oauth

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils"
	"github.com/golang-jwt/jwt"
)

const JWT_PRIVATE_KEY_PATH = "jwt/jwt_key.rsa"
const JWT_PUBLIC_PKCS8_PATH = "jwt/jwt_key.rsa.pkcs8"

type JWTPublic struct {
	PKCS8 string `json:"pkcs8"`
}

// ref. http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#StandardClaims
type Claims struct {
	Name                string `json:"name"`
	GivenName           string `json:"given_name"`
	FamilyName          string `json:"family_name"`
	MiddleName          string `json:"middle_name"`
	NickName            string `json:"nick_name"`
	PreferredUserName   string `json:"preferred_username"`
	Profile             string `json:"profile"`
	Picture             string `json:"picture"`
	Website             string `json:"website"`
	Email               string `json:"email"`
	EmailVerified       bool   `json:"email_verified"`
	Gender              string `json:"gender"`
	Birthdate           string `json:"birthdate"`
	Zoneinfo            string `json:"zoneinfo"`
	Locale              string `json:"locale"`
	PhoneNumber         string `json:"phone_number"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	UpdatedAt           int64  `json:"updated_at"`

	ID    string `json:"id"`
	Role  string `json:"role"`
	Theme string `json:"theme"`

	Iat      int64 `json:"iat"`
	AuthTime int64 `json:"auth_time"`

	jwt.StandardClaims
}

func NewJWT(user *models.User, clientId string, authTime time.Time) *Claims {
	standartClaims := jwt.StandardClaims{
		Issuer:    fmt.Sprintf("https://%s", config.Defs.SiteDomain),
		Subject:   utils.UUID(),
		Audience:  clientId,
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	}

	claims := &Claims{
		StandardClaims: standartClaims,

		Iat:      time.Now().Unix(),
		AuthTime: authTime.Unix(),

		Name:                fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		GivenName:           user.FirstName,
		FamilyName:          user.LastName,
		MiddleName:          "", // TODO: あとで追加したい気がする
		NickName:            user.UserNameFormatted,
		PreferredUserName:   user.UserName,
		Profile:             "",
		Picture:             user.AvatarUrl,
		Website:             "",
		Email:               user.Mail,
		EmailVerified:       true,         // ユーザはすべて認証済みなので必ずtrueになる
		Gender:              "",           // TODO: あとで追加したい気がするけど世間的に空にしていたほうが良いかもしれない
		Birthdate:           "",           // ここも追加したい
		Zoneinfo:            "Asia/Tokyo", // 日本語しか対応していないので日本標準時にしてしまう
		Locale:              "ja-JP",      // 上に同じ
		PhoneNumber:         "",
		PhoneNumberVerified: false,
		UpdatedAt:           time.Now().Unix(), // めんどくさいので今の時間

		ID:    user.UserId.UserId,
		Role:  strings.Join(user.Role, " "),
		Theme: user.Theme,
	}

	return claims
}

func (c *Claims) ConvertJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)

	signBytes, err := ioutil.ReadFile(JWT_PRIVATE_KEY_PATH)
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}

	return token.SignedString(signKey)
}

func GetPublicKey() (*JWTPublic, error) {
	pkcs8, err := ioutil.ReadFile(JWT_PUBLIC_PKCS8_PATH)
	if err != nil {
		return nil, err
	}

	return &JWTPublic{
		PKCS8: string(pkcs8),
	}, nil
}
