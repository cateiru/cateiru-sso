package oauth

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/go-http-error/httperror/status"
)

func TokenEndpoint(w http.ResponseWriter, r *http.Request, query url.Values) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	switch query.Get("grant_type") {
	case "authorization_code":
		return AuthorizationCode(ctx, db, query)
	case "refresh_token":
		return Refresh(ctx, db, query)
	default:
		return status.NewBadRequestError(errors.New("grant_type required"))
	}
}

func AuthorizationCode(ctx context.Context, db *database.Database, query url.Values) error {
	// request := TokenRequest{
	// 	GrantType:   "authorization_code",
	// 	Code:        query.Get("code"),
	// 	RedirectUri: query.Get("redirect_uri"),
	// }

	// accessToken, err := request.Required(ctx, db)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func Refresh(ctx context.Context, db *database.Database, query url.Values) error {
	// request := RefreshRequest{
	// 	GrantType:    "refresh_token",
	// 	ClientID:     query.Get("client_id"),
	// 	ClientSecret: query.Get("client_secret"),
	// 	RefreshToken: query.Get("refresh_token"),
	// 	Scope:        strings.Split(query.Get("scope"), " "),
	// }

	return nil
}
