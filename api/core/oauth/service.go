package oauth

import (
	"context"
	"errors"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
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

func (c *Service) Required(ctx context.Context, db *database.Database) (*models.SSOService, error) {
	if len(c.Scope) == 0 {
		return nil, errors.New("scope is null")
	} else if len(c.ResponseType) == 0 {
		return nil, errors.New("response type is null")
	} else if len(c.ClientID) == 0 {
		return nil, errors.New("client id is null")
	} else if len(c.RedirectURL) == 0 {
		return nil, errors.New("redirect url is null")
	} else if len(c.FromURL) == 0 {
		return nil, errors.New("from url is null")
	}

	isOpenIDScope := false
	for _, v := range c.Scope {
		if v == "openid" {
			isOpenIDScope = true
			break
		}
	}
	if !isOpenIDScope {
		return nil, errors.New("no exist openid value in scope filed")
	}

	ssoService, err := models.GetSSOServiceByClientId(ctx, db, c.ClientID)
	if err != nil {
		return nil, err
	}

	if ssoService == nil {
		return nil, errors.New("service not found")
	}

	isRedirectURL := false
	for _, url := range ssoService.ToUrl {
		if url == c.RedirectURL {
			isRedirectURL = true
			break
		}
	}
	if !isRedirectURL {
		return nil, errors.New("redirect url is not exist")
	}

	isFromURL := false
	for _, url := range ssoService.FromUrl {
		if url == c.FromURL {
			isFromURL = true
			break
		}
	}
	if !isFromURL {
		return nil, errors.New("from url is not")
	}

	return ssoService, nil
}
