package common

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"google.golang.org/api/iterator"
)

// certification内に同じメールアドレスが存在するかチェックします
// 存在している場合、そのメールアドレスを使用しているユーザがいることになります
//
// 同じメールアドレスで複数のアカウントを持つことはできません
func CheckExistMail(ctx context.Context, db *database.Database, mail string) (bool, error) {
	logging.Sugar.Debugf("check exist mail: %s", mail)
	query := datastore.NewQuery("Certification").Filter("mail =", mail)

	iter := db.Run(ctx, query)

	var entry models.Certification
	_, err := iter.Next(&entry)
	// 要素が見つからない場合
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}