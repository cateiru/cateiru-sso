package oauth

import "time"

// ref. http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#IDToken
type IDToken struct {
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	AuthTime string `json:"auth_time"`
	Nonce    string `json:"nonce"`

	Claim
}

// ref. http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#StandardClaims
type Claim struct {
	Sub                 string    `json:"sub"`
	Name                string    `json:"name"`
	GivenName           string    `json:"given_name"`
	FamilyName          string    `json:"family_name"`
	MiddleName          string    `json:"middle_name"`
	PreferredUserName   string    `json:"preferred_username"`
	Profile             string    `json:"profile"`
	Picture             string    `json:"picture"`
	Website             string    `json:"website"`
	Email               string    `json:"email"`
	EmailVerified       bool      `json:"email_verified"`
	Gender              string    `json:"gender"`
	Birthdate           time.Time `json:"birthdate"`
	Zoneinfo            string    `json:"zoneinfo"`
	PhoneNumber         string    `json:"phone_number"`
	PhoneNumberVerified bool      `json:"phone_number_verified"`
	UpdatedAt           string    `json:"updated_at"`
}
