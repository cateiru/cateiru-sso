// - メールアドレスの存在チェック
// - IPアドレスがブロックリストに存在するかチェック
package common

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/logging"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
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

// ユーザ名が存在するかチェックする
func CheckUsername(ctx context.Context, db *database.Database, userName string) (bool, error) {
	user, err := models.GetUserDataByUserName(ctx, db, userName)
	if err != nil {
		return false, status.NewInternalServerErrorError(err).Caller()
	}

	exist := false
	if user != nil {
		exist = true
	}

	return exist, nil
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
	return config.Defs.AdminMail == mail
}

// 有効期限切れかどうかを調べます。
// 期限切れの場合、trueが返る
func CheckExpired(entry *models.Period) bool {
	now := time.Now()
	var periodTime time.Time

	// 有効期限の日時
	if entry.PeriodMinute != 0 {
		periodTime = entry.CreateDate.Add(time.Duration(entry.PeriodMinute) * time.Minute)
	} else if entry.PeriodHour != 0 {
		periodTime = entry.CreateDate.Add(time.Duration(entry.PeriodHour) * time.Hour)
	} else if entry.PeriodDay != 0 {
		periodTime = entry.CreateDate.Add(time.Duration(entry.PeriodDay*24) * time.Hour)
	} else {
		// hour, minuteどちらも定義されていない場合、createDateで比較する = かならずtrueが返る
		periodTime = entry.CreateDate
	}

	// 有効期限より前に今の時間がある場合Falseを返す
	return now.After(periodTime)
}

// OTPが正しいかをチェックします
func CheckOTP(passcode string, cert *models.Certification, secret *string) (bool, bool) {
	if cert == nil {
		return secure.ValidateOnetimePassword(passcode, *secret), false
	}

	if secure.ValidateOnetimePassword(passcode, cert.OnetimePasswordSecret) {
		return true, false
	}

	for index, backup := range cert.OnetimePasswordBackups {
		if backup == passcode {
			// バックアップコードは1回使用したら使えなくなる
			cert.OnetimePasswordBackups = append(
				cert.OnetimePasswordBackups[:index], cert.OnetimePasswordBackups[index+1:]...)
			return true, true
		}
	}

	return false, false
}
