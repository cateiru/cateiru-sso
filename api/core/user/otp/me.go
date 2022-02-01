package otp

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type OTPMeResponse struct {
	Enable bool `json:"enable"`
}

func OTPMeHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	userId, err := common.GetUserID(ctx, db, w, r)
	if err != nil {
		return err
	}

	cert, err := models.GetCertificationByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInsufficientStorageError(err).Caller()
	}

	isOTP := len(cert.OnetimePasswordSecret) != 0

	net.ResponseOK(w, OTPMeResponse{
		Enable: isOTP,
	})

	return nil
}
