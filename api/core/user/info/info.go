package info

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

type Request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Theme     string `json:"theme"`
}

func ChangeInfoHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	ctx := r.Context()

	var request Request

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	userId, err := common.GetUserID(ctx, db, w, r)
	if err != nil {
		return err
	}

	changedUser, err := ChangeInfo(ctx, db, userId, &request)
	if err != nil {
		return err
	}

	net.ResponseOK(w, changedUser)

	return nil
}

func ChangeInfo(ctx context.Context, db *database.Database, userId string, req *Request) (*models.User, error) {
	user, err := models.GetUserDataByUserID(ctx, db, userId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	isChange := false

	if err := ChangeFirstName(user, req.FirstName, &isChange); err != nil {
		return nil, err
	}
	if err := ChangeLastName(user, req.LastName, &isChange); err != nil {
		return nil, err
	}
	if err := ChangeUserName(user, req.UserName, &isChange); err != nil {
		return nil, err
	}
	if err := ChangeTheme(user, req.Theme, &isChange); err != nil {
		return nil, err
	}

	if isChange {
		if err := user.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller()
		}

		return user, nil
	}

	return nil, status.NewBadRequestError(errors.New("did not change anything")).Caller()
}

func ChangeFirstName(user *models.User, newFirstname string, isChange *bool) error {
	// 要素がからの場合はなにもしない
	if len(newFirstname) == 0 {
		return nil
	}

	// 変わっていない場合はなにもしない
	if user.FirstName == newFirstname {
		return nil
	}

	user.FirstName = newFirstname
	*isChange = true
	return nil
}

func ChangeLastName(user *models.User, newLastName string, isChange *bool) error {
	// 要素がからの場合はなにもしない
	if len(newLastName) == 0 {
		return nil
	}

	// 変わっていない場合はなにもしない
	if user.LastName == newLastName {
		return nil
	}

	user.LastName = newLastName
	*isChange = true
	return nil
}

func ChangeUserName(user *models.User, newUserName string, isChange *bool) error {
	// 要素がからの場合はなにもしない
	if len(newUserName) == 0 {
		return nil
	}

	// 変わっていない場合はなにもしない
	if user.UserName == newUserName {
		return nil
	}

	// ユーザ名がただしいか検証する
	if !utils.CheckUserName(newUserName) {
		return status.NewBadRequestError(errors.New("incorrect username")).Caller().AddCode(net.IncorrectUserName)
	}

	user.UserName = newUserName
	user.UserNameFormatted = utils.FormantUserName(newUserName)
	*isChange = true
	return nil
}

func ChangeTheme(user *models.User, newTheme string, isChange *bool) error {
	// 要素がからの場合はなにもしない
	if len(newTheme) == 0 {
		return nil
	}

	// 変わっていない場合はなにもしない
	if user.Theme == newTheme {
		return nil
	}

	user.Theme = newTheme
	*isChange = true
	return nil
}
