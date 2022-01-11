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

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func PasswordChangeHandler(w http.ResponseWriter, r *http.Request) error {
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

	userId, err := common.GetUserID(ctx, db, w, r)
	if err != nil {
		return err
	}

	var request ChangePasswordRequest

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	if err := ChangePassword(ctx, db, userId, &request); err != nil {
		return err
	}

	return nil
}

// パスワードを更新する
func ChangePassword(ctx context.Context, db *database.Database, userId string, form *ChangePasswordRequest) error {
	cert, err := models.GetCertificationByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if cert == nil {
		return status.NewInternalServerErrorError(errors.New("entity is not found")).Caller()
	}

	// old_passwordが違う場合は400を返す
	if !secure.ValidatePW(form.OldPassword, cert.Password, cert.Salt) {
		return status.NewBadRequestError(errors.New("password is not validate")).Caller()
	}

	newPassword, err := secure.PWHash(form.NewPassword)
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	cert.Password = newPassword.Key
	cert.Salt = newPassword.Salt

	if err := cert.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
