package check

import (
	"context"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type ResponseCheckUserName struct {
	Exist bool `json:"exist"`
}

func CheckUserNameHandler(w http.ResponseWriter, r *http.Request) error {
	userName, err := net.GetQuery(r, "name")
	// クエリパラメータがない場合は400を返す
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	exist, err := CheckUsername(ctx, db, userName)
	if err != nil {
		return err
	}

	net.ResponseOK(w, ResponseCheckUserName{
		Exist: exist,
	})

	return nil
}

func CheckUsername(ctx context.Context, db *database.Database, userName string) (bool, error) {
	user, err := models.GetUserDataByUserName(ctx, db, userName)
	if err != nil {
		return false, status.NewInternalServerErrorError(err)
	}

	exist := false
	if user != nil {
		exist = true
	}

	return exist, nil
}
