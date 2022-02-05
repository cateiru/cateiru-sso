package admin

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func AllUsersHand(w http.ResponseWriter, r *http.Request) error {
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

	users, err := models.GetAllUsers(ctx, db)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	net.ResponseOK(w, users)

	return nil
}
