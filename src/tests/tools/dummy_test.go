package tools_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestDummyUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	userInfo, err := dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getUserInfo, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return getUserInfo != nil && getUserInfo.Mail == userInfo.Mail
	}, "正しくInfoがDBに格納できている")

	userCert, err := dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getUserCert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return getUserCert != nil && getUserCert.Mail == userCert.Mail
	}, "正しくCertがDBに格納できている")

	session, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getSession, err := models.GetSessionToken(ctx, db, session)
		require.NoError(t, err)

		getRefresh, err := models.GetRefreshToken(ctx, db, refresh)
		require.NoError(t, err)

		return getSession != nil && getRefresh != nil
	}, "session-token、refresh-tokenがDBにセットされている")
}
