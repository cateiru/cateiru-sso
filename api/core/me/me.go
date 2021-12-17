package me

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

// 現在ログインしているユーザ情報を返す
func MeHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller("core/me/me.go", 16).Wrap()
	}
	defer db.Close()

	userId, err := common.GetUserID(ctx, db, w, r)
	if err != nil {
		return err
	}

	userInfo, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller("core/me/me.go", 27).Wrap()
	}

	if userInfo == nil {
		return status.NewBadRequestError(err).Caller(
			"core/me/me.go", 35)
	}

	net.ResponseOK(w, userInfo)

	return nil
}
