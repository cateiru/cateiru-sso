package otp

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/src/core/common"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/cateiru/cateiru-sso/src/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

type ResponseBackups struct {
	Codes []string `json:"codes"`
}

// OTPのバックアップコードを返す
// ログインしているときのみ
func BackupHandler(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	ctx := r.Context()

	password := r.Form.Get("password")

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

	cert, err := models.GetCertificationByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// パスワードを検証する
	if !secure.ValidatePW(password, cert.Password, cert.Salt) {
		return status.NewBadRequestError(errors.New("no validate password")).Caller().AddCode(net.FailedLogin)
	}

	codes, err := GetBackupcodes(ctx, db, userId)
	if err != nil {
		return err
	}

	net.ResponseOK(w, ResponseBackups{
		Codes: codes,
	})

	return nil
}

// OTPのバックアップコードを返す
func GetBackupcodes(ctx context.Context, db *database.Database, userId string) ([]string, error) {
	userCert, err := models.GetCertificationByUserID(ctx, db, userId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	if userCert == nil {
		return nil, status.NewBadRequestError(errors.New("entity is not found")).Caller()
	}

	// OTPが設定されていない場合は400を返す
	if len(userCert.OnetimePasswordBackups) == 0 {
		return nil, status.NewBadRequestError(errors.New("otp is not set")).Caller()
	}

	return userCert.OnetimePasswordBackups, nil
}
