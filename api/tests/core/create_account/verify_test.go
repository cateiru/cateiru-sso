package createaccount_test

import (
	"context"
	"testing"
	"time"

	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/go-http-error/httperror"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestVerifySuccess(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mailToken := utils.CreateID(20)

	mailVerify := &models.MailCertification{
		MailToken:        mailToken,
		ClientCheckToken: "hugahuga",

		OpenNewWindow:  true,
		Verify:         false,
		ChangeMailMode: false,

		UserMailPW: models.UserMailPW{
			Mail:     "example@example.com",
			Password: "password",
		},

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	// メール認証設定
	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	// DBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	res, err := createaccount.CreateVerify(ctx, mailToken)
	require.NoError(t, err)

	require.NotNil(t, res)
	require.True(t, res.IsKeepThisPage)

	goretry.Retry(t, func() bool {
		element, err := models.GetCreateAccountBufferByBufferToken(ctx, db, res.BufferToken)
		require.NoError(t, err)

		return element != nil && element.Mail == "example@example.com"
	}, "メールアドレスが同じ")
}

// 認証が存在しない場合
func TestVerifyNotExist(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	_, err := createaccount.CreateVerify(ctx, "hogehoge")
	require.Error(t, err)

	httperr, ok := httperror.CastHTTPError(err)
	require.True(t, ok)
	require.Equal(t, httperr.StatusCode, 400, "400が帰る")
}

// 既に認証済みの場合
func TestVerifyAlreadyDone(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mailToken := utils.CreateID(20)

	mailVerify := &models.MailCertification{
		MailToken:        mailToken,
		ClientCheckToken: "hugahuga",

		OpenNewWindow: true,
		// 既に認証済み
		Verify:         true,
		ChangeMailMode: false,

		UserMailPW: models.UserMailPW{
			Mail:     "example@example.com",
			Password: "password",
		},

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	// メール認証設定
	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	// DBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	_, err = createaccount.CreateVerify(ctx, mailToken)
	require.Error(t, err)

	httperr, ok := httperror.CastHTTPError(err)
	require.True(t, ok)
	require.Equal(t, httperr.StatusCode, 400, "400が帰る")
}
