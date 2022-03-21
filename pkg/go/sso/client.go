package sso

import (
	"fmt"
	"net/url"
)

const SSO_ENDPOINT = "https://sso.cateiru.com/sso/login"

// Create Login URI
func CreateURI(clientId string, redirect string) string {
	encodedClientId := url.QueryEscape(clientId)
	encodedRedirect := url.QueryEscape(redirect)

	return fmt.Sprintf("%s?scope=openid&response_type=code&client_id=%s&redirect_uri=%s&prompt=consent", SSO_ENDPOINT, encodedClientId, encodedRedirect)
}
