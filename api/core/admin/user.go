package admin

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/core/logout"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
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

	id, emptyErr := net.GetQuery(r, "id")

	if emptyErr != nil {
		users, err := models.GetAllUsers(ctx, db)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		net.ResponseOK(w, users)
	} else {
		user, err := models.GetUserDataByUserID(ctx, db, id)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		net.ResponseOK(w, []models.User{*user})
	}

	return nil
}

func DeleteUserHand(w http.ResponseWriter, r *http.Request) error {
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

	targetUserId, err := net.GetQuery(r, "id")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	// 自分自身を削除はできない
	logging.Sugar.Info(targetUserId, userId)
	if targetUserId == userId {
		return status.NewBadRequestError(errors.New("you can't delete yourself")).Caller()
	}

	if err := DeleteUserSessions(ctx, db, targetUserId); err != nil {
		return err
	}

	if err := logout.Delete(ctx, db, targetUserId); err != nil {
		return err
	}

	return nil
}

func DeleteUserSessions(ctx context.Context, db *database.Database, userId string) error {
	if err := models.DeleteSessionByUserId(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteRefreshByUserId(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	return nil
}
