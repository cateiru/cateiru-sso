package pro

import (
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type SSOService struct {
	LoginCount int `json:"login_count"`

	models.SSOService
}

func GetSSOServices(w http.ResponseWriter, r *http.Request) error {
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

	// proユーザのみ使用可
	if err := common.ProOnly(ctx, db, c.UserId); err != nil {
		return err
	}

	services, err := models.GetSSOServiceByUserID(ctx, db, c.UserId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	servicesInCount := []SSOService{}

	// それぞれのサービスの利用者数をカウントする
	for _, service := range services {
		count, err := models.CountSSOServiceLogByClientId(ctx, db, service.ClientID)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		servicesInCount = append(servicesInCount, SSOService{
			LoginCount: count,

			SSOService: service,
		})
	}

	net.ResponseOK(w, servicesInCount)

	return nil
}

func DeleteService(w http.ResponseWriter, r *http.Request) error {
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

	// proユーザのみ使用可
	if err := common.ProOnly(ctx, db, c.UserId); err != nil {
		return err
	}

	clientId, err := net.GetQuery(r, "id")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	service, err := models.GetSSOServiceByClientId(ctx, db, clientId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if service == nil {
		return status.NewBadRequestError(errors.New("service not found")).Caller()
	}
	if service.UserId.UserId != c.UserId {
		return status.NewBadRequestError(errors.New("user failed")).Caller()
	}

	// serviceを削除する
	if err := models.DeleteSSOServiceByClientId(ctx, db, service.ClientID); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// そのserviceのログインログを削除する
	if err := models.DeleteSSOServiceLogByClientId(ctx, db, service.ClientID); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// AccessTokenを削除する（今ログインを試みている場合など）
	if err := models.DeleteAccessTokenByClientID(ctx, db, service.ClientID); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// SSO Refresh tokenを削除する
	if err := models.DeleteSSORefreshTokenByClientId(ctx, db, service.ClientID); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
