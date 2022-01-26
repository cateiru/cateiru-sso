package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/storage"
	"github.com/cateiru/go-http-error/httperror/status"
)

// ユーザのアバターを設定する
func AvatorSetHandler(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	fileSrc, fileHeader, err := r.FormFile("upload")
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer fileSrc.Close()

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

	contentType := fileHeader.Header.Get("Content-Type")

	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer s.Close()

	if err := s.WriteFile(ctx, []string{"avator"}, userId, fileSrc, contentType); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// user_infoのavatorを書き換える（空の場合）
	// avatorのURLはユーザごとに一意であり、画像が変わってもURLは変わらない必要があります
	user, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return status.NewInternalServerErrorError(errors.New("user data is empty")).Caller()
	}
	if user.AvatarUrl == "" {
		user.AvatarUrl = fmt.Sprintf("https://%s/%s/avator/%s", config.Defs.StorageDomain, config.Defs.StorageBucket, userId)
		if err := user.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	return nil
}

// ユーザのアバターを削除する
func DeleteAvatorHandler(w http.ResponseWriter, r *http.Request) error {
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

	user, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return status.NewInternalServerErrorError(errors.New("user data is empty")).Caller()
	}

	// アバターが設定されていない場合は400を返す
	if user.AvatarUrl == "" {
		return status.NewBadRequestError(errors.New("avator is null"))
	}

	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	exist, err := s.FileExist(ctx, []string{"avator"}, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	// storageに存在しない場合は400を返す
	if !exist {
		return status.NewBadRequestError(errors.New("avator not found in gcs")).Caller()
	}

	if err := s.Delete(ctx, []string{"avator"}, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
