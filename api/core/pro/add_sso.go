package pro

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type AddRequestForm struct {
	Name      string   `json:"name"`
	FromURL   []string `json:"from_url"`
	ToURL     []string `json:"to_url"`
	LoginOnly bool     `json:"login_only"`

	SessionTokenPeriod int `json:"session_token_period"`
	RefreshTokenPeriod int `json:"refresh_token_period"`
}

type AddResponse struct {
	PublicKey  string `json:"public_key"`
	SecretKey  string `json:"secret_key"`
	PrivateKey string `json:"private_key"`
}

func AddSSOHandler(w http.ResponseWriter, r *http.Request) error {
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

	// Pro以上のユーザのみ使用可
	if err := common.ProMoreOnly(ctx, db, userId); err != nil {
		return err
	}

	var form AddRequestForm
	if err := net.GetJsonForm(w, r, &form); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	resp, err := AddSSO(ctx, db, userId, &form)
	if err != nil {
		return err
	}

	net.ResponseOK(w, resp)

	return nil
}

// SSOを追加する
func AddSSO(ctx context.Context, db *database.Database, userId string, form *AddRequestForm) (*AddResponse, error) {
	if len(form.FromURL) == 0 || len(form.ToURL) == 0 {
		return nil, status.NewBadRequestError(errors.New("url is empty")).Caller()
	}
	// FromURLとToURLが同じではない and ToURL != 1の場合は400を返す
	//
	// FromURLが複数ある場合、そのIndexに対応したToURLに対してリダイレクトを行います。
	// もし、ToURLが1つである場合、複数のFromURLから1つのToURLにリダイレクトします。
	//
	// そのため、FromURLとToURLが同じであるか、違う場合はToURLが1つである必要があります。
	if len(form.ToURL) != 1 && len(form.FromURL) != len(form.ToURL) {
		return nil, status.NewBadRequestError(errors.New("url is failed")).Caller()
	}

	publicKey := utils.CreateID(0)
	secretKey := utils.CreateID(20)
	privateKey := utils.CreateID(0)

	service := models.SSOService{
		SSOPublicKey: publicKey,

		SSOSecretKey:  secretKey,
		SSOPrivateKey: privateKey,

		Name:      form.Name,
		FromUrl:   form.FromURL,
		ToUrl:     form.ToURL,
		LoginOnly: form.LoginOnly,

		SessionTokenPeriod: form.SessionTokenPeriod,
		RefreshTokenPeriod: form.RefreshTokenPeriod,

		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err := service.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &AddResponse{
		PublicKey:  publicKey,
		SecretKey:  secretKey,
		PrivateKey: privateKey,
	}, nil
}
