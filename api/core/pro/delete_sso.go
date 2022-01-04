package pro

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

func DeleteSSOHandler(w http.ResponseWriter, r *http.Request) error {
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

	// Pro以上のユーザのみ使用可
	if err := common.ProMoreOnly(ctx, db, userId); err != nil {
		return nil
	}

	publicKey, err := net.GetQuery(r, "key")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	return DeleteSSO(ctx, db, publicKey, userId)
}

func DeleteSSO(ctx context.Context, db *database.Database, publicKey string, userId string) error {
	sso, err := models.GetSSOServiceByPublicKey(ctx, db, publicKey)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if sso == nil {
		return status.NewBadRequestError(errors.New("publickey is failed")).Caller()
	}

	// UserIDが違う = SSOを作成した人が違う場合は400を返す
	if userId != sso.UserId.UserId {
		return status.NewBadRequestError(errors.New("userid is failed")).Caller()
	}

	err = models.DeleteSSOServiceByPublicKey(ctx, db, publicKey)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	return nil
}
