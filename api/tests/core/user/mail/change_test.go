package mail_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/user/mail"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/iterator"
)

func TestChangeRequestMail(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	err = mail.ChangeMail(ctx, db, dummy.Mail, dummy.UserID)
	require.NoError(t, err)

	query := datastore.NewQuery("MailCertification").Filter("mail =", dummy.Mail)
	goretry.Retry(t, func() bool {
		iter := db.Run(ctx, query)
		var entity models.MailCertification
		_, err := iter.Next(&entity)

		// 要素がなにもない場合
		if err == iterator.Done {
			return false
		}
		require.NoError(t, err)

		return entity.ChangeMailMode && entity.UserId == dummy.UserID
	}, "dbに要素がある")
}

func TestVerifyChangeMail(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	newMail := tools.NewDummyUser().Mail
	mailToken := utils.CreateID(20)

	// ---

	mailVerify := &models.MailCertification{
		MailToken:   mailToken,
		ClientToken: utils.CreateID(0), // 使わないが一応keyを指定しておく

		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: true, // メールアドレス変更なので

		UserMailPW: models.UserMailPW{
			Mail: newMail,
		},

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},

		UserId: dummy.UserID,
	}

	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	// --- 変更する

	goretry.Retry(t, func() bool {
		err = mail.VerifyNewMail(ctx, db, mailToken, dummy.UserID)
		if err != nil {
			t.Log(err)
			return false
		}
		return true
	}, "")

	// --- 確認する

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && cert.Mail == newMail
	}, "certのメールアドレスが変更されている")

	goretry.Retry(t, func() bool {
		info, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return info != nil && info.Mail == newMail
	}, "userInfoのメールアドレスが変更されている")
}
