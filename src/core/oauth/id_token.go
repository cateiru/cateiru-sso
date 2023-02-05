package oauth

import (
	"context"
	"errors"

	"github.com/cateiru/cateiru-sso/src/core/common"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

// ref. http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#IDToken
type IDToken struct {
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	AuthTime string `json:"auth_time"`
	Nonce    string `json:"nonce"`
}

// ref. RFC6749: OAuth2.0 ---- 4.1.3.  Access Token Request
type TokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

func (c *TokenRequest) Required(ctx context.Context, db *database.Database) (*models.SSOAccessToken, error) {
	if c.GrantType != "authorization_code" {
		return nil, status.NewBadRequestError(errors.New("grant_type must be `authorization_code`")).Caller()
	}
	if len(c.Code) == 0 {
		return nil, status.NewBadRequestError(errors.New("code is null")).Caller()
	}
	if len(c.RedirectUri) == 0 {
		return nil, status.NewBadRequestError(errors.New("redirect_uri is null")).Caller()
	}

	accessToken, err := models.GetAccessTokenByAccessToken(ctx, db, c.Code)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if accessToken == nil {
		return nil, status.NewBadRequestError(errors.New("code is failed")).Caller()
	}

	if common.CheckExpired(&accessToken.Period) {
		return nil, status.NewBadRequestError(errors.New("expired")).Caller().AddCode(net.TimeOutError)
	}

	if c.RedirectUri != accessToken.RedirectURI {
		return nil, status.NewBadRequestError(errors.New("redirect uri")).Caller()
	}

	return accessToken, nil
}

// ref. OpenIDConnect 1.0: 12.1. Refresh Request
type RefreshRequest struct {
	GrantType    string   `json:"grant_type"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
}

func (c *RefreshRequest) Required(ctx context.Context, db *database.Database) (*models.SSORefreshToken, error) {
	if c.GrantType != "refresh_token" {
		return nil, status.NewBadRequestError(errors.New("grant_type must be `authorization_code`")).Caller()
	} else if len(c.ClientID) == 0 {
		return nil, status.NewBadRequestError(errors.New("client id is null")).Caller()
	} else if len(c.ClientSecret) == 0 {
		return nil, status.NewBadRequestError(errors.New("client secret is null")).Caller()
	} else if len(c.RefreshToken) == 0 {
		return nil, status.NewBadRequestError(errors.New("refresh token is null")).Caller()
	}

	isOpenIDScope := false
	for _, v := range c.Scope {
		if v == "openid" {
			isOpenIDScope = true
			break
		}
	}
	if !isOpenIDScope {
		return nil, status.NewBadRequestError(errors.New("no exist openid value in scope filed")).Caller()
	}

	refresh, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, c.RefreshToken)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if refresh == nil {
		return nil, status.NewBadRequestError(errors.New("refresh")).Caller()
	}

	if common.CheckExpired(&refresh.Period) {
		return nil, status.NewBadRequestError(errors.New("expired")).Caller().AddCode(net.TimeOutError)
	}

	return refresh, nil
}
