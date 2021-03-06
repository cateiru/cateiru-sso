package oauth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func ServiceLogin(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	var request Service

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	c := common.NewCert(w, r)
	if err := c.Login(ctx, db); err != nil {
		return err
	}

	service, err, errCode := request.Required(ctx, db)
	if err != nil {
		return status.NewBadRequestError(err).Caller().AddCode(errCode)
	}

	// roleが設定している場合、そのユーザは対象のroleがあるかチェックする
	if len(service.AllowRoles) != 0 {
		roles, err := models.GetRoleByUserID(ctx, db, c.UserId)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		if roles == nil {
			return status.NewInternalServerErrorError(errors.New("role empty")).Caller()
		}

		ok := false

		for _, role := range roles.Role {
			for _, requiredRole := range service.AllowRoles {
				if role == requiredRole {
					ok = true
					break
				}
			}
		}

		if !ok {
			return status.NewBadRequestError(errors.New("role")).Caller().AddCode(net.NoRole)
		}
	}

	accessToken, err := LoginOAuth(ctx, db, service.ClientID, request.RedirectURL, c.UserId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	log := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service.ClientID,

		UserId: models.UserId{
			UserId: c.UserId,
		},
	}

	if err := log.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	net.ResponseOK(w, LoginResponse{
		AccessToken: accessToken,
	})

	return nil
}

func LoginOAuth(ctx context.Context, db *database.Database, clientId string, redirectUri string, userId string) (string, error) {
	accessToken := utils.CreateID(0)

	access := models.SSOAccessToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: "", // まだ空（初回アクセス時に適用させる）

		ClientID: clientId,

		RedirectURI: redirectUri,

		Create: time.Now(),

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 5,
		},

		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err := access.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	return accessToken, nil
}
