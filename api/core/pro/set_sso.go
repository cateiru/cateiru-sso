package pro

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/storage"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type SetRequest struct {
	ClientId string `json:"client_id"` // optional

	Name string `json:"name"`

	FromURL []string `json:"from_url"`
	ToURL   []string `json:"to_url"`

	AllowRoles []string `json:"allow_roles"`

	ChangeTokenSecert bool `json:"change_token_secret"`
}

func SetService(w http.ResponseWriter, r *http.Request) error {
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

	c := common.NewCert(w, r)
	if err := c.Login(ctx, db); err != nil {
		return err
	}
	userId := c.UserId

	// Proのみ
	if err := common.ProOnly(ctx, db, userId); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	var form SetRequest

	if err := net.GetJsonForm(w, r, &form); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	// URLが正しいかチェックする
	if len(form.FromURL) != 0 {
		if err := CheckURL(form.FromURL, true); err != nil {
			return err
		}
	}
	if len(form.ToURL) != 0 {
		if err := CheckURL(form.ToURL, true); err != nil {
			return err
		}
	}

	// client_id が定義されていない場合は新規作成する
	if len(form.ClientId) == 0 {
		newService := models.SSOService{
			ClientID:    utils.CreateID(30),
			TokenSecret: utils.CreateID(0),
			Name:        form.Name,
			ServiceIcon: "", // 最初は空

			FromUrl: form.FromURL,
			ToUrl:   form.ToURL,

			UserId: models.UserId{
				UserId: userId,
			},
		}

		if len(form.AllowRoles) != 0 {
			newService.AllowRoles = form.AllowRoles
		}

		if err := newService.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		net.ResponseOK(w, newService)
	} else {
		service, err := models.GetSSOServiceByClientId(ctx, db, form.ClientId)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		if service == nil {
			return status.NewBadRequestError(errors.New("service is not found")).Caller()
		}
		// 作成したユーザと違う場合はだめ
		if service.UserId.UserId != userId {
			return status.NewBadRequestError(errors.New("user failed")).Caller()
		}

		if len(form.Name) != 0 {
			service.Name = form.Name
		}
		if len(form.FromURL) != 0 {
			service.FromUrl = form.FromURL
		}
		if len(form.ToURL) != 0 {
			service.ToUrl = form.ToURL
		}
		if form.ChangeTokenSecert {
			service.TokenSecret = utils.CreateID(0)
		}
		if len(form.AllowRoles) != 0 {
			if len(form.AllowRoles) == 1 && form.AllowRoles[0] == "" {
				service.AllowRoles = []string{}
			} else {
				service.AllowRoles = form.AllowRoles
			}

		}

		if err := service.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		net.ResponseOK(w, service)
	}

	return nil
}

// sso serviceの画像をセットする
func SetImage(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	fileSrc, fileHeader, err := r.FormFile("image")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}
	defer fileSrc.Close()

	clientId := r.FormValue("client_id")

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

	// Proのみ
	if err := common.ProOnly(ctx, db, userId); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	contentType := fileHeader.Header.Get("Content-Type")

	if contentType != "image/png" {
		return status.NewBadRequestError(errors.New("content-type must be image/png")).Caller()
	}

	service, err := models.GetSSOServiceByClientId(ctx, db, clientId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if service == nil {
		return status.NewBadRequestError(errors.New("client id is null")).Caller()
	}
	// 作成したユーザと違う場合はだめ
	if service.UserId.UserId != userId {
		return status.NewBadRequestError(errors.New("user failed")).Caller()
	}

	// 画像保存
	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer s.Close()

	if err := s.WriteFile(ctx, []string{"sso"}, service.ClientID, fileSrc, contentType); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if len(service.ServiceIcon) == 0 {
		service.ServiceIcon = fmt.Sprintf("%s/%s/sso/%s", config.Defs.StorageURL, config.Defs.StorageBucket, service.ClientID)

		err = service.Add(ctx, db)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	net.ResponseOK(w, service)

	return nil
}

// URLが正しいかチェックする
func CheckURL(urls []string, allowDirect bool) error {

	pattan := `(https://[\w/:%#\$&\?\(\)~\.=\+\-]+|http://localhost)`

	if allowDirect {
		pattan = `(https://[\w/:%#\$&\?\(\)~\.=\+\-]+|http://localhost|direct)`
	}

	for _, url := range urls {
		match, err := regexp.MatchString(pattan, url)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		if !match {
			return status.NewBadRequestError(errors.New("url is not match")).Caller()
		}
	}

	return nil
}
