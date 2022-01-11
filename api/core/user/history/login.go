package history

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func UserLoginHistoryHandler(w http.ResponseWriter, r *http.Request) error {
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

	// limitを設定する
	// 例: ?limit=10
	// もし設定しない場合はすべての要素を返します
	limitInt := -1
	limit, err := net.GetQuery(r, "limit")
	if err == nil {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return status.NewBadRequestError(err).Caller()
		}
	}

	histories, err := LoginHistories(ctx, db, userId, limitInt)
	if err != nil {
		return err
	}

	net.ResponseOK(w, histories)
	return nil
}

// ログイン履歴を取得する
func LoginHistories(ctx context.Context, db *database.Database, userId string, limit int) ([]models.LoginHistory, error) {
	histories, err := models.GetAllLoginHistory(ctx, db, userId, limit)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// 要素がない場合は400を返す
	if len(histories) == 0 {
		return nil, status.NewBadRequestError(errors.New("entity is failed")).Caller()
	}

	return histories, nil
}