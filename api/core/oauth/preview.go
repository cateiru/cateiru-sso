package oauth

import (
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type ResponsePerview struct {
	Name        string `datastore:"name" json:"name"`
	ServiceIcon string `datastore:"serviceIcon" json:"service_icon"`
}

// そのSSOサービスの存在可否をチェックしてサービス名、アイコンを返す
func ServicePrevier(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	var request Service

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(err).Caller()
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

	service, err := request.Required(ctx, db)
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	net.ResponseOK(w, ResponsePerview{
		Name:        service.Name,
		ServiceIcon: service.ServiceIcon,
	})

	return nil
}
