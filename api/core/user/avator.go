package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/storage"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type SetAvatarResp struct {
	Url string `json:"url"`
}

// ユーザのアバターを設定する
func AvatarSetHandler(w http.ResponseWriter, r *http.Request) error {
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

	logging.Sugar.Info(contentType)
	if contentType != "image/png" {
		return status.NewBadRequestError(errors.New("content-type must be image/png")).Caller()
	}

	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer s.Close()

	if err := s.WriteFile(ctx, []string{"avatar"}, userId, fileSrc, contentType); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// user_infoのavatarを書き換える（空の場合）
	// avatarのURLはユーザごとに一意であり、画像が変わってもURLは変わらない必要があります
	user, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return status.NewInternalServerErrorError(errors.New("user data is empty")).Caller()
	}
	if user.AvatarUrl == "" {
		user.AvatarUrl = fmt.Sprintf("%s/%s/avatar/%s", config.Defs.StorageURL, config.Defs.StorageBucket, userId)
		if err := user.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	net.ResponseOK(w, SetAvatarResp{
		Url: user.AvatarUrl,
	})

	return nil
}

// ユーザのアバターを削除する
func DeleteAvatarHandler(w http.ResponseWriter, r *http.Request) error {
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
		return status.NewBadRequestError(errors.New("avatar is null"))
	}

	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	exist, err := s.FileExist(ctx, []string{"avatar"}, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	// storageに存在しない場合は400を返す
	if !exist {
		return status.NewBadRequestError(errors.New("avatar not found in gcs")).Caller()
	}

	if err := s.Delete(ctx, []string{"avatar"}, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
