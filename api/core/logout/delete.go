package logout

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/go-http-error/httperror/status"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	// session-tokenとrefresh-tokenのcookieとDBを削除する
	// (logoutと同じ処理)
	userId, err := Logout(ctx, db, w, r)
	if err != nil {
		return err
	}

	if len(userId) != 0 {
		return Delete(ctx, db, userId)
	}
	return status.NewInternalServerErrorError(errors.New("no set userID")).Caller()
}

// アカウントを削除する
func Delete(ctx context.Context, db *database.Database, userId string) error {
	// ユーザの認証情報
	if err := models.DeleteCertificationByUserId(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// ユーザの基本情報
	if err := models.DeleteUserDataByUserID(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// ユーザのログイン履歴
	if err := models.DeleteAllLoginHistories(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// TODO: 削除するユーザに紐付けられているSSOの情報や履歴なども削除する

	return nil
}
