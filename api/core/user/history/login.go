package history

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type LoginHistory struct {
	ThisDevice    bool      `json:"this_device"`
	IsLogout      bool      `json:"is_logout"`
	LastLoginDate time.Time `json:"last_login_date"`

	models.LoginHistory
}

func UserLoginHistoryHandler(w http.ResponseWriter, r *http.Request) error {
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

	var newHistories []LoginHistory

	refreshes, err := models.GetRefreshTokenByUserId(ctx, db, userId)
	if err != nil {
		return status.NewInsufficientStorageError(err).Caller()
	}

	for _, history := range histories {
		write := false
		for _, refresh := range refreshes {
			if history.AccessId == refresh.AccessID {
				thisDevice := false
				if history.AccessId == c.AccessID {
					thisDevice = true
				}

				newHistories = append(newHistories, LoginHistory{
					ThisDevice:    thisDevice,
					IsLogout:      false,
					LastLoginDate: refresh.CreateDate,

					LoginHistory: history,
				})
				write = true
				break
			}
		}
		if !write {
			newHistories = append(newHistories, LoginHistory{
				ThisDevice:   false,
				IsLogout:     true,
				LoginHistory: history,
			})
		}
	}

	net.ResponseOK(w, newHistories)
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
