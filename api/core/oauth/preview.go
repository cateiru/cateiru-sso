package oauth

import (
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type ResponsePerview struct {
	Name        string `datastore:"name" json:"name"`
	ServiceIcon string `datastore:"serviceIcon" json:"service_icon"`
}

// そのSSOサービスの存在可否をチェックしてサービス名、アイコンを返す
func ServicePreview(w http.ResponseWriter, r *http.Request) error {
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

	service, err, errCode := request.Required(ctx, db)
	if err != nil {
		return status.NewBadRequestError(err).Caller().AddCode(errCode)
	}

	// roleが設定している場合、そのユーザは対象のroleがあるかチェックする
	if len(service.AllowRoles) != 0 {
		roles, err := models.GetRoleByUserID(ctx, db, c.UserId)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
		if roles == nil {
			return status.NewInternalServerErrorError(errors.New("role empty")).Caller()
		}
		ok := false

		for _, role := range roles.Role {
			for _, requiredRole := range service.AllowRoles {
				if role == requiredRole {
					ok = true
					break
				}
			}
		}

		if !ok {
			return status.NewBadRequestError(errors.New("role")).Caller().AddCode(net.NoRole)
		}
	}

	net.ResponseOK(w, ResponsePerview{
		Name:        service.Name,
		ServiceIcon: service.ServiceIcon,
	})

	return nil
}
