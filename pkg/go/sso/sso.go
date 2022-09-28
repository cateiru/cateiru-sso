package sso

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt"
)

const TOKEN_ENDPOINT = "https://api.sso.cateiru.com/v1/oauth/token"
const PUBLIC_KEY_ENDPOINT = "https://api.sso.cateiru.com/v1/oauth/jwt/key"

type PublicKey struct {
	Pkcs8 string `json:"pkcs8"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
}

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

// Get JWT Public key.
func GetPublicKey() (string, error) {
	res, err := http.DefaultClient.Get(PUBLIC_KEY_ENDPOINT)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New("connection failed")
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	var r PublicKey

	if err := json.Unmarshal(buf.Bytes(), &r); err != nil {
		return "", err
	}

	return r.Pkcs8, nil
}

func GetToken(code string, redirect string, tokenSecret string) (*TokenResponse, error) {
	encodedCode := url.QueryEscape(code)
	encodedRedirect := url.QueryEscape(redirect)

	uri := fmt.Sprintf("%s?grant_type=authorization_code&code=%s&redirect_uri=%s",
		TOKEN_ENDPOINT, encodedCode, encodedRedirect)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", tokenSecret))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("connection failed")
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	var r TokenResponse

	if err := json.Unmarshal(buf.Bytes(), &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func Refresh(refreshToken string, clientId string, clientSecret string, scope []string) (*TokenResponse, error) {
	encodedRefreshToken := url.QueryEscape(refreshToken)
	encodedClientId := url.QueryEscape(clientId)
	encodedClientSecret := url.QueryEscape(clientSecret)
	encodedScope := url.QueryEscape(strings.Join(scope, " "))

	uri := fmt.Sprintf("%s?grant_ype=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s&scope=%s",
		TOKEN_ENDPOINT, encodedClientId, encodedClientSecret, encodedRefreshToken, encodedScope)

	res, err := http.DefaultClient.Get(uri)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("connection failed")
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	var r TokenResponse

	if err := json.Unmarshal(buf.Bytes(), &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func ValidateIDToken(idToken string) (*Claims, error) {
	pks8, err := GetPublicKey()
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pks8))
	if err != nil {
		return nil, err
	}

	var claims Claims

	token, err := jwt.ParseWithClaims(idToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("IDToken is not valid")
	}

	return &claims, nil
}
