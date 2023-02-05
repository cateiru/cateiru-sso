package admin

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

type RoleRequest struct {
	Action string `json:"action"` // `enable`, `disable`
	Role   string `json:"role"`
	UserId string `json:"user_id"`
}

func remove(s []string, t string) ([]string, error) {
	index := 0
	exist := false

	for i, v := range s {
		if t == v {
			exist = true
			index = i
			break
		}
	}

	if !exist {
		return nil, errors.New("target is not exist")
	}

	return s[:index+copy(s[index:], s[index+1:])], nil
}

func AdminRoleHand(w http.ResponseWriter, r *http.Request) error {
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

	// Adminのユーザのみ使用可
	if err := common.AdminOnly(ctx, db, userId); err != nil {
		return err
	}

	var form RoleRequest

	if err := net.GetJsonForm(w, r, &form); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	// 自分自身は変更できない
	if form.UserId == userId {
		return status.NewBadRequestError(errors.New("you can't delete yourself")).Caller()
	}

	switch form.Action {
	case "enable":
		if err := AddRole(ctx, db, form.Role, form.UserId); err != nil {
			return err
		}
	case "disable":
		if err := DeleteRole(ctx, db, form.Role, form.UserId); err != nil {
			return err
		}
	default:
		return status.NewBadRequestError(errors.New("no action")).Caller()
	}

	return nil
}

func AddRole(ctx context.Context, db *database.Database, role string, userId string) error {
	user, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return status.NewBadRequestError(errors.New("user is empty")).Caller()
	}

	user.Role = append(user.Role, role)

	if err := user.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	roleEntity, err := models.GetRoleByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if roleEntity == nil {
		return status.NewBadRequestError(errors.New("role is empty")).Caller()
	}

	roleEntity.Role = append(roleEntity.Role, role)

	if err := roleEntity.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

func DeleteRole(ctx context.Context, db *database.Database, role string, userId string) error {
	user, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if user == nil {
		return status.NewBadRequestError(errors.New("user is empty")).Caller()
	}

	user.Role, err = remove(user.Role, role)
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	if err := user.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	roleEntity, err := models.GetRoleByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if roleEntity == nil {
		return status.NewBadRequestError(errors.New("role is empty")).Caller()
	}

	roleEntity.Role, err = remove(roleEntity.Role, role)
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	if err := roleEntity.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
