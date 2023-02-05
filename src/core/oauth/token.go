package oauth

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils"
	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
}

func TokenEndpoint(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	query := r.Form

	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	switch query.Get("grant_type") {
	case "authorization_code":
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			return status.NewBadRequestError(errors.New("authorization heder required")).Caller()
		}
		authSplitted := strings.Split(auth, " ")
		if authSplitted[0] != "Basic" && len(authSplitted[1]) == 0 {
			return status.NewBadRequestError(errors.New("authorization heder must be basic")).Caller()
		}

		resp, err := AuthorizationCode(ctx, db, query, authSplitted[1])
		if err != nil {
			return err
		}
		net.ResponseOK(w, resp)
		return nil
	case "refresh_token":
		resp, err := Refresh(ctx, db, query)
		if err != nil {
			return err
		}
		net.ResponseOK(w, resp)
		return nil
	default:
		return status.NewBadRequestError(errors.New("grant_type required"))
	}
}

func AuthorizationCode(ctx context.Context, db *database.Database, query url.Values, tokenSecret string) (*TokenResponse, error) {
	request := TokenRequest{
		GrantType:   "authorization_code",
		Code:        query.Get("code"),
		RedirectUri: query.Get("redirect_uri"),
	}

	accessToken, err := request.Required(ctx, db)
	if err != nil {
		return nil, err
	}

	service, err := models.GetSSOServiceByClientId(ctx, db, accessToken.ClientID)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// token secretを検証する
	if service.TokenSecret != tokenSecret {
		return nil, status.NewForbiddenError(errors.New("secret")).Caller()
	}

	user, err := models.GetUserDataByUserID(ctx, db, accessToken.UserId.UserId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return nil, status.NewInternalServerErrorError(errors.New("user is empty")).Caller()
	}

	jwt := NewJWT(user, accessToken.ClientID, accessToken.Create)

	idToken, err := jwt.ConvertJWT()
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// refresh tokenが設定されてない場合は新しくつくる
	var refreshToken string
	if accessToken.SSORefreshToken == "" {
		refreshToken = utils.CreateID(0)

		refresh := models.SSORefreshToken{
			SSOAccessToken:  accessToken.SSOAccessToken,
			SSORefreshToken: refreshToken,
			ClientID:        accessToken.ClientID,
			RedirectURI:     accessToken.RedirectURI,
			Period: models.Period{
				CreateDate: time.Now(),
				PeriodDay:  150, // Googleが6ヶ月なのでこのくらいでもいいでしょう！
			},
			UserId: accessToken.UserId,
		}

		if err := refresh.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller()
		}
	} else {
		refreshToken = accessToken.SSORefreshToken
	}

	return &TokenResponse{
		AccessToken:  accessToken.SSOAccessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    accessToken.PeriodMinute*60 + accessToken.PeriodHour*3600 + accessToken.PeriodDay*86400,
		IDToken:      idToken,
	}, nil
}

func Refresh(ctx context.Context, db *database.Database, query url.Values) (*TokenResponse, error) {
	request := RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     query.Get("client_id"),
		ClientSecret: query.Get("client_secret"),
		RefreshToken: query.Get("refresh_token"),
		Scope:        strings.Split(query.Get("scope"), " "),
	}

	refresh, err := request.Required(ctx, db)
	if err != nil {
		return nil, err
	}

	service, err := models.GetSSOServiceByClientId(ctx, db, refresh.ClientID)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if service == nil {
		return nil, status.NewBadRequestError(err).Caller()
	}

	// token secretを検証する
	if service.TokenSecret != request.ClientSecret {
		return nil, status.NewForbiddenError(errors.New("secret")).Caller()
	}

	// access tokenは削除する
	if err := models.DeleteAccessTokenByAccessToken(ctx, db, refresh.SSOAccessToken); err != nil {
		return nil, status.NewInsufficientStorageError(err).Caller()
	}

	// refresh tokenも削除する
	if err := models.DeleteSSORefreshTokenByRefreshToken(ctx, db, request.RefreshToken); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	accessToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

	access := models.SSOAccessToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: refreshToken,

		ClientID:    refresh.ClientID,
		RedirectURI: refresh.RedirectURI,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 5,
		},

		UserId: refresh.UserId,
	}

	if err := access.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// 新しくrefresh tokenを作る
	newRefresh := models.SSORefreshToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: refreshToken,

		ClientID: refresh.ClientID,

		RedirectURI: refresh.RedirectURI,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  150, // Googleが6ヶ月なのでこのくらいでもいいでしょう！
		},
		UserId: refresh.UserId,
	}

	if err := newRefresh.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// id tokenも返す
	user, err := models.GetUserDataByUserID(ctx, db, refresh.UserId.UserId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return nil, status.NewInternalServerErrorError(errors.New("user is empty")).Caller()
	}

	jwt := NewJWT(user, refresh.ClientID, access.Create)

	idToken, err := jwt.ConvertJWT()
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    access.PeriodMinute*60 + access.PeriodHour*3600 + access.PeriodDay*86400,
		IDToken:      idToken,
	}, nil
}
