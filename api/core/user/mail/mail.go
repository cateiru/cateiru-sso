package mail

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

type ResponseMail struct {
	Mail string `json:"mail"`
}

// ユーザのメールアドレスを取得する
func GetMailHandler(w http.ResponseWriter, r *http.Request) error {
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

	mail, err := GetMail(ctx, db, userId)
	if err != nil {
		return err
	}

	net.ResponseOK(w, ResponseMail{
		Mail: mail,
	})

	return nil
}

// userIDからメールアドレスを取得する
func GetMail(ctx context.Context, db *database.Database, userId string) (string, error) {
	userInfo, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}
	if userInfo == nil {
		return "", status.NewInternalServerErrorError(errors.New("user info is empty")).Caller()
	}

	return userInfo.Mail, nil
}
