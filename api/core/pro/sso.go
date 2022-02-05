package pro

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func GetSSOHandler(w http.ResponseWriter, r *http.Request) error {
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
	userId := c.UserId

	// Pro以上のユーザのみ使用可
	if err := common.ProOnly(ctx, db, userId); err != nil {
		return err
	}

	entities, err := GetSSO(ctx, db, userId)
	if err != nil {
		return err
	}

	net.ResponseOK(w, entities)

	return nil
}

func GetSSO(ctx context.Context, db *database.Database, userId string) ([]models.SSOService, error) {
	entities, err := models.GetSSOServiceByUserID(ctx, db, userId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	if len(entities) == 0 {
		return nil, status.NewBadRequestError(errors.New("no defined sso")).Caller()
	}

	return entities, nil
}
