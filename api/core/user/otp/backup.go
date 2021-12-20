package otp

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

type ResponseBackups struct {
	Codes []string `json:"codes"`
}

// OTPのバックアップコードを返す
// ログインしているときのみ
func BackupHandler(w http.ResponseWriter, r *http.Request) error {
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

	// OTPが設定されていない場合は400を返す
	if len(userCert.OnetimePasswordBackups) == 0 {
		return nil, status.NewBadRequestError(errors.New("otp is not set")).Caller()
	}

	return userCert.OnetimePasswordBackups, nil
}
