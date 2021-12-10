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
	"github.com/stretchr/testify/require"
)

func TestAcceptSuccess(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mailToken := utils.CreateID(20)
	clientCheckToken := utils.CreateID(20)

	mailVerify := &models.MailCertification{
		MailToken:        mailToken,
		ClientCheckToken: clientCheckToken,

		OpenNewWindow: false,
		// 認証済みとする
		Verify:         true,
		ChangeMailMode: false,

		UserMailPW: models.UserMailPW{
			Mail:     "example@example.com",
			Password: "password",
		},

		VerifyPeriod: models.VerifyPeriod{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	// メール認証設定
	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	bufferToken, err := createaccount.AcceptVerify(ctx, clientCheckToken)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	buffer, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
	require.NoError(t, err)
	require.NotNil(t, buffer)
	require.Equal(t, buffer.Mail, "example@example.com")
}

// 認証が存在しない場合
func TestAcceptNoEntry(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	_, err := createaccount.AcceptVerify(ctx, "hogehoge")
	require.Error(t, err)

	httperr, ok := httperror.CastHTTPError(err)
	require.True(t, ok)
	require.Equal(t, httperr.StatusCode, 400, "400が帰る")
}

// 認証済みではない場合
func TestAcceptNotVerify(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mailToken := utils.CreateID(20)
	clientCheckToken := utils.CreateID(20)

	mailVerify := &models.MailCertification{
		MailToken:        mailToken,
		ClientCheckToken: clientCheckToken,

		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: false,

		UserMailPW: models.UserMailPW{
			Mail:     "example@example.com",
			Password: "password",
		},

		VerifyPeriod: models.VerifyPeriod{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	// メール認証設定
	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	_, err = createaccount.AcceptVerify(ctx, clientCheckToken)
	require.Error(t, err)

	httperr, ok := httperror.CastHTTPError(err)
	require.True(t, ok)
	require.Equal(t, httperr.StatusCode, 403, "403が帰る")
}
