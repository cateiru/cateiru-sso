// - メールアドレスの存在チェック
// - IPアドレスがブロックリストに存在するかチェック
package common

import (
	"context"
	"os"

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

// IPアドレス、メールアドレスがブロックリストに存在しているのかを確認する
// 存在している場合、trueが返る
func ChaeckBlock(ctx context.Context, db *database.Database, ip string, mail string) (bool, error) {
	resultIp, err := models.GetBlockListByIP(ctx, db, ip)
	if err != nil {
		return false, err
	}
	resultMail, err := models.GetBlockListByMail(ctx, db, mail)
	if err != nil {
		return false, err
	}

	if resultIp == nil && resultMail == nil {
		// IP、メールどちらもない場合（=allow）はfalse
		return false, nil
	}
	return true, nil
}

// メールアドレスがadminで定義したメールアドレスかをチェックします
func CheckAdminMail(mail string) bool {
	return os.Getenv("ADMIN_MAIL") == mail
}
