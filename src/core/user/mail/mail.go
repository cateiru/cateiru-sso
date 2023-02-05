package mail

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/core/common"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type ResponseMail struct {
	Mail string `json:"mail"`
}

// ユーザのメールアドレスを取得する
func GetMailHandler(w http.ResponseWriter, r *http.Request) error {
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

	mail, err := common.GetMailByUserID(ctx, db, userId)
	if err != nil {
		return err
	}

	net.ResponseOK(w, ResponseMail{
		Mail: mail,
	})

	return nil
}
