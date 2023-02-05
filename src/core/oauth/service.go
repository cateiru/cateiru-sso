package oauth

import (
	"context"
	"errors"

	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils/net"
)

type Service struct {
	Scope        []string `json:"scope"`
	ResponseType string   `json:"response_type"`
	ClientID     string   `json:"client_id"`
	RedirectURL  string   `json:"redirect_uri"`
	State        string   `json:"state"`
	Prompt       string   `json:"prompt"`

	FromURL string `json:"from_url"`
}

func (c *Service) Required(ctx context.Context, db *database.Database) (*models.SSOService, error, int) {
	if len(c.Scope) == 0 {
		return nil, errors.New("scope is null"), net.IncorrectOIDC
	} else if len(c.ResponseType) == 0 {
		return nil, errors.New("response type is null"), net.IncorrectOIDC
	} else if len(c.ClientID) == 0 {
		return nil, errors.New("client id is null"), net.IncorrectOIDC
	} else if len(c.RedirectURL) == 0 {
		return nil, errors.New("redirect url is null"), net.IncorrectOIDC
	} else if len(c.FromURL) == 0 {
		return nil, errors.New("from url is null"), net.IncorrectOIDC
	}

	isOpenIDScope := false
	for _, v := range c.Scope {
		if v == "openid" {
			isOpenIDScope = true
			break
		}
	}
	if !isOpenIDScope {
		return nil, errors.New("no exist openid value in scope filed"), net.IncorrectOIDC
	}

	ssoService, err := models.GetSSOServiceByClientId(ctx, db, c.ClientID)
	if err != nil {
		return nil, err, 1
	}

	if ssoService == nil {
		return nil, errors.New("service not found"), net.NotExistService
	}

	isRedirectURL := false
	for _, url := range ssoService.ToUrl {
		if url == c.RedirectURL {
			isRedirectURL = true
			break
		}
	}
	if !isRedirectURL {
		return nil, errors.New("redirect url is not exist"), net.NoRedirectURI
	}

	isFromURL := false
	for _, url := range ssoService.FromUrl {
		if url == c.FromURL {
			isFromURL = true
			break
		}
	}
	if !isFromURL {
		return nil, errors.New("from url is not"), net.NoRefererURI
	}

	return ssoService, nil, 0
}
