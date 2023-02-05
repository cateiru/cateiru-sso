package admin

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/core/common"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func MailCertLogHandler(w http.ResponseWriter, r *http.Request) error {
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

	// Adminのユーザのみ使用可
	if err := common.AdminOnly(ctx, db, userId); err != nil {
		return err
	}

	logs, err := models.GetAllTryCreateAccountLog(ctx, db)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	net.ResponseOK(w, logs)

	return nil
}
