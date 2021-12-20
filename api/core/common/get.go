package common

import (
	"context"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/go-http-error/httperror/status"
)

// ユーザIDからメールアドレスを取得する
func GetMailByUserID(ctx context.Context, db *database.Database, userID string) (string, error) {
	userInfo, err := models.GetUserDataByUserID(ctx, db, userID)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	if userInfo == nil {
		return "", status.NewBadRequestError(err).Caller()
	}

	return userInfo.Mail, nil
}
