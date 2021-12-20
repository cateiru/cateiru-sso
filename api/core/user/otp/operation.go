package otp

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type OTPRequest struct {
	Type     string `json:"type"` // `enable` or `disable`
	Passcode string `json:"passcode"`
	Id       string `json:"id,omitempty"`
}

type SetOTPResponse struct {
	Backups []string `json:"backups"`
}

func OTPHandler(w http.ResponseWriter, r *http.Request) error {
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("request header is not application/json")).Caller()
	}

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

	var request OTPRequest

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(errors.New("parse not failed")).Caller()
	}

	switch request.Id {
	case "enable":
		// formにidがない場合400を返す
		if len(request.Id) == 0 {
			return status.NewBadRequestError(errors.New("parse not failed")).Caller()
		}
		backups, err := SetOTP(ctx, db, userId, request.Id, request.Passcode)
		if err != nil {
			return err
		}
		net.ResponseOK(w, backups)
		return nil
	case "disable":
		if err := DeleteOTP(ctx, db, userId, request.Passcode); err != nil {
			return err
		}
		return nil
	default:
		return status.NewBadRequestError(errors.New("parse not failed")).Caller()
	}
}

// OTPを設定します。
func SetOTP(ctx context.Context, db *database.Database, userId string, id string, passcode string) (*SetOTPResponse, error) {
	OTPBuffer, err := models.GetOTPBufferByID(ctx, db, id)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	if OTPBuffer == nil {
		return nil, status.NewBadRequestError(errors.New("entity not failed")).Caller()
	}

	// OTPを検証する
	// 検証が失敗した場合403を返す
	if !common.CheckOTP(passcode, nil, &OTPBuffer.SecretKey) {
		return nil, status.NewForbiddenError(errors.New("otp is not validated")).Caller()
	}

	userCert, err := models.GetCertificationByUserID(ctx, db, userId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	if userCert == nil {
		return nil, status.NewBadRequestError(errors.New("entity not failed")).Caller()
	}

	backups := []string{}
	for i := 0; 10 > i; i++ {
		backups = append(backups, utils.CreateID(10))
	}

	userCert.OnetimePasswordSecret = OTPBuffer.SecretKey
	userCert.OnetimePasswordBackups = backups

	if err := userCert.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &SetOTPResponse{
		Backups: backups,
	}, nil
}

// アカウントからOTPの設定を削除します
func DeleteOTP(ctx context.Context, db *database.Database, userId string, passcode string) error {
	userCert, err := models.GetCertificationByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if userCert == nil {
		return status.NewBadRequestError(errors.New("entity not failed")).Caller()
	}

	// OTPを検証する
	// 検証が失敗した場合403を返す
	if !common.CheckOTP(passcode, userCert, nil) {
		return status.NewForbiddenError(errors.New("otp is not validated")).Caller()
	}

	userCert.OnetimePasswordBackups = []string{}
	userCert.OnetimePasswordSecret = ""

	if err := userCert.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
