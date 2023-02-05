package user

import (
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type SSOLoginLog struct {
	ClientID string `json:"client_id"`

	Name        string `json:"name"`
	ServiceIcon string `json:"service_icon"`

	Logs []models.SSOServiceLog `json:"logs"`
}

func OAuthShow(w http.ResponseWriter, r *http.Request) error {
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

	serviceLogs, err := models.GetSSOServiceLogsByUserId(ctx, db, c.UserId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	log := []SSOLoginLog{}

	clientIds := make(map[string][]models.SSOServiceLog, len(serviceLogs))
	for _, serviceLog := range serviceLogs {
		already := false
		for id := range clientIds {
			if id == serviceLog.ClientID {
				already = true
			}
		}
		if !already {
			clientIds[serviceLog.ClientID] = []models.SSOServiceLog{serviceLog}
		} else {
			clientIds[serviceLog.ClientID] = append(clientIds[serviceLog.ClientID], serviceLog)
		}
	}

	for clientId, serviceLogs := range clientIds {
		service, err := models.GetSSOServiceByClientId(ctx, db, clientId)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		log = append(log, SSOLoginLog{
			ClientID: clientId,

			Name:        service.Name,
			ServiceIcon: service.ServiceIcon,

			Logs: serviceLogs,
		})
	}

	net.ResponseOK(w, log)

	return nil
}

func DeleteOAth(w http.ResponseWriter, r *http.Request) error {
	clientId, err := net.GetQuery(r, "id")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}
	if clientId == "" {
		return status.NewBadRequestError(errors.New("client id required")).Caller()
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

	// ClientIDが存在するか確認する
	service, err := models.GetSSOServiceByClientId(ctx, db, clientId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if service == nil {
		return status.NewBadRequestError(errors.New("service empty")).Caller()
	}

	if err := models.DeleteAccessTokenByUserIdAndClientId(ctx, db, c.UserId, clientId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteSSORefreshTokenByUserIdAndClientID(ctx, db, c.UserId, clientId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if err := models.DeleteSSOServiceLogByUserIDAndClientId(ctx, db, c.UserId, clientId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
