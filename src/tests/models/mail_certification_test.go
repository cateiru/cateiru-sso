package models_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestMailCertification(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mail := fmt.Sprintf("%s@example.com", utils.CreateID(4))

	mailToken := utils.CreateID(10)
	clientToken := utils.CreateID(10)

	entry := &models.MailCertification{
		MailToken:      mailToken,
		ClientToken:    clientToken,
		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: false,
		Mail:           mail,
		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}
	err = entry.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		result, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return result != nil
	}, "entryがある")

	result, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
	require.NoError(t, err)
	require.NotNil(t, result, "要素が見つかった")
	require.Equal(t, result.Mail, mail, "見つかった要素のメールアドレスが取得できる")

	result, err = models.GetMailCertificationByMailToken(ctx, db, "example@example.com")
	require.NoError(t, err)
	require.Nil(t, result, "要素が見つからない")

	result, err = models.GetMailCertificationByClientToken(ctx, db, clientToken)
	require.NoError(t, err)
	require.NotNil(t, result, "要素が見つかった")
	require.Equal(t, result.Mail, mail, "見つかった要素のメールアドレスが取得できる")

	result, err = models.GetMailCertificationByClientToken(ctx, db, "hoge")
	require.NoError(t, err)
	require.Nil(t, result, "要素が見つからない")

}

func TestDeletePeriod(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mailToken := utils.CreateID(10)

	entry := &models.MailCertification{
		MailToken:      mailToken,
		ClientToken:    utils.CreateID(0),
		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: false,
		Mail:           tools.NewDummyUser().Mail,
		Period: models.Period{
			CreateDate:   time.Now().Add(time.Duration(-1) * time.Hour),
			PeriodMinute: 30,
		},
	}
	err = entry.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		result, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return result != nil
	}, "entryがある")

	err = models.DeleteMailCertPeriod(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		result, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return result == nil
	}, "削除されている")
}
