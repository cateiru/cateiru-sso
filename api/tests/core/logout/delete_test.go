package logout_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/logout"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)

	// --- ユーザ情報を定義する

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	// ログイン履歴
	history := &models.LoginHistory{
		AccessId:     utils.CreateID(20),
		Date:         time.Now(),
		IpAddress:    "192.168.0.1",
		IsSSO:        false,
		SSOPublicKey: "",
		UserAgent:    "",

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = history.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(entity) == 1
	}, "要素が入った")

	// --- 削除する

	err = logout.Delete(ctx, db, dummy.UserID)
	require.NoError(t, err)

	// --- チェックする

	goretry.Retry(t, func() bool {
		info, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return info == nil && cert == nil
	}, "ユーザの認証情報が消えている")

	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(histores) == 0
	}, "ユーザのログイン履歴が消えている")
}
