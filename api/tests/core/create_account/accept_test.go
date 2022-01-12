package createaccount_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/go-http-error/httperror"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestAcceptSuccess(t *testing.T) {
	config.TestInit(t)

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
			Password: []byte("password"),
			Salt:     []byte(""),
		},

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	// メール認証設定
	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	// メール認証がDBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	bufferToken, err := createaccount.AcceptVerify(ctx, clientCheckToken)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		buffer, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
		require.NoError(t, err)
		return buffer != nil
	}, "entryがある")

	buffer, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
	require.NoError(t, err)
	require.NotNil(t, buffer)
	require.Equal(t, buffer.Mail, "example@example.com")
}

// 認証が存在しない場合
func TestAcceptNoEntry(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	_, err := createaccount.AcceptVerify(ctx, "hogehoge")
	require.Error(t, err)

	httperr, ok := httperror.CastHTTPError(err)
	require.True(t, ok)
	require.Equal(t, httperr.StatusCode, 400, "400が帰る")
}

// 認証済みではない場合
func TestAcceptNotVerify(t *testing.T) {
	config.TestInit(t)

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
			Password: []byte("password"),
			Salt:     []byte(""),
		},

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	// メール認証設定
	err = mailVerify.Add(ctx, db)
	require.NoError(t, err)

	// メール認証がDBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	_, err = createaccount.AcceptVerify(ctx, clientCheckToken)
	require.Error(t, err)

	httperr, ok := httperror.CastHTTPError(err)
	require.True(t, ok)
	require.Equal(t, httperr.StatusCode, 403, "403が帰る")
}
