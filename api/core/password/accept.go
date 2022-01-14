package password

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

type AccpetFortgetRequest struct {
	ForgetToken string `json:"forget_token"`
	NewPassword string `json:"new_password"`
}

func ForgetPasswordAcceptHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	var form AccpetFortgetRequest
	if err := net.GetJsonForm(w, r, &form); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	return ChangePWAccept(ctx, db, &form)
}

func ChangePWAccept(ctx context.Context, db *database.Database, form *AccpetFortgetRequest) error {
	buffer, err := models.GetPWForgetByToken(ctx, db, form.ForgetToken)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	// entityが無い = tokenが無効な場合は400を返す
	if buffer == nil {
		return status.NewBadRequestError(errors.New("pw forget is no entity")).Caller()
	}
	// 有効期限の場合は400を返す
	if common.CheckExpired(&buffer.Period) {
		return status.NewBadRequestError(errors.New("expired")).Caller()
	}

	cert, err := models.GetCertificationByMail(ctx, db, buffer.Mail)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	pw, err := secure.PWHash(form.NewPassword)
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	cert.Password = pw.Key
	cert.Salt = pw.Salt

	if err := cert.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if err := models.DeletePWForgetByToken(ctx, db, form.ForgetToken); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
