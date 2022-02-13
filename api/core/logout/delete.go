package logout

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/storage"
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
	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

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

	// ユーザのロール
	if err := models.DeleteRoleByUserID(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// アバター
	exist, err := s.FileExist(ctx, []string{"avatar"}, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if exist {
		if err := s.Delete(ctx, []string{"avatar"}, userId); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	if err := DeleteSSOService(ctx, db, s, userId); err != nil {
		return err
	}

	if err := DeleteSSOLogin(ctx, db, userId); err != nil {
		return err
	}

	// TODO: 削除するユーザに紐付けられているSSOの情報や履歴なども削除する

	return nil
}

// Proアカウントでそのユーザが管理しているsso serviceを削除する
func DeleteSSOService(ctx context.Context, db *database.Database, s *storage.Storage, userId string) error {
	services, err := models.GetSSOServiceByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	for _, service := range services {
		// そのserviceのログインログを削除する
		if err := models.DeleteSSOServiceLogByClientId(ctx, db, service.ClientID); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		// AccessTokenを削除する（今ログインを試みている場合など）
		if err := models.DeleteAccessTokenByClientID(ctx, db, service.ClientID); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		// SSO Refresh tokenを削除する
		if err := models.DeleteSSORefreshTokenByClientId(ctx, db, service.ClientID); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		exist, err := s.FileExist(ctx, []string{"sso"}, service.ClientID)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		if exist {
			if err := s.Delete(ctx, []string{"sso"}, service.ClientID); err != nil {
				return status.NewInternalServerErrorError(err).Caller()
			}
		}

		// serviceを削除する
		if err := models.DeleteSSOServiceByClientId(ctx, db, service.ClientID); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}
	return nil
}

func DeleteSSOLogin(ctx context.Context, db *database.Database, userId string) error {
	if err := models.DeleteSSOServiceLogByUserId(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if err := models.DeleteSSORefreshTokenByUserId(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if err := models.DeleteAccessTokenByUserId(ctx, db, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
