package admin

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/logging"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-error/httperror/status"
)

func WorkerHandler(w http.ResponseWriter, r *http.Request) error {
	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		return status.NewBadRequestError(errors.New("authorization heder required")).Caller()
	}
	authSplitted := strings.Split(auth, " ")
	if authSplitted[0] != "Basic" && len(authSplitted[1]) == 0 {
		return status.NewBadRequestError(errors.New("authorization heder must be basic")).Caller()
	}

	if authSplitted[1] != config.Defs.WorkerPassword {
		return status.NewForbiddenError(errors.New("")).Caller()
	}

	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	logging.Sugar.Info("Run worker")

	return Worker(ctx, db)
}

func Worker(ctx context.Context, db *database.Database) error {

	if err := models.DeleteMailCertPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteOTPBufferPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeletePWForgetPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteRefreshTokenPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteSessionTokenPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteAccessTokenPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteSSORefreshTokenPeriod(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
