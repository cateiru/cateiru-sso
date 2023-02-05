package logout_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/logout"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/storage"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
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
		AccessId:  utils.CreateID(20),
		Date:      time.Now(),
		IpAddress: "192.168.0.1",
		UserAgent: "",

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

	s, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	require.NoError(t, err)

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

	goretry.Retry(t, func() bool {
		role, err := models.GetRoleByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return role == nil
	}, "ユーザのロールが消えている")

	goretry.Retry(t, func() bool {
		exist, err := s.FileExist(ctx, []string{"avatar"}, dummy.UserID)
		require.NoError(t, err)

		return !exist
	}, "ユーザのアバターが消えている")

	// TODO: sso log, access token, refresh, imageのチェックも追加する
	goretry.Retry(t, func() bool {
		services, err := models.GetSSOServiceByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(services) == 0
	}, "ユーザが定義したSSOが消えている")

	// TODO: access token, refreshのチェックも追加する
	goretry.Retry(t, func() bool {
		logins, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(logins) == 0
	}, "ユーザがログインしているSSOが消えている")
}

func TestDeleteSSOService(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")

	service := models.SSOService{
		ClientID:    utils.CreateID(30),
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "",

		FromUrl: []string{"http://cateiru.com"},
		ToUrl:   []string{"https://example.com"},

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	log := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service.ClientID,

		UserId: models.UserId{
			UserId: utils.CreateID(0), // 違うユーザ
		},
	}
	err = log.Add(ctx, db)
	require.NoError(t, err)

	accessToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

	access := models.SSOAccessToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: refreshToken,

		ClientID: service.ClientID,

		RedirectURI: "https://example.com",

		Create: time.Now(),

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		UserId: log.UserId,
	}
	err = access.Add(ctx, db)
	require.NoError(t, err)

	refresh := models.SSORefreshToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: refreshToken,

		ClientID: service.ClientID,

		RedirectURI: "https://example.com",

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  7,
		},

		UserId: log.UserId,
	}
	err = refresh.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entity != nil
	}, "最後の要素が格納されている")

	// ---

	err = logout.Delete(ctx, db, dummy.UserID)
	require.NoError(t, err)

	// --- チェックする

	goretry.Retry(t, func() bool {
		services, err := models.GetSSOServiceByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, service.ClientID)
		require.NoError(t, err)

		accessTokenG, err := models.GetAccessTokenByAccessToken(ctx, db, accessToken)
		require.NoError(t, err)

		refreshTokenG, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return len(services) == 0 && len(logs) == 0 && accessTokenG == nil && refreshTokenG == nil
	}, "ユーザが定義したSSOが消えている")
}

func TestDeleteMyLoginSSO(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	service := models.SSOService{
		ClientID:    utils.CreateID(30),
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "",

		FromUrl: []string{"http://cateiru.com"},
		ToUrl:   []string{"https://example.com"},

		UserId: models.UserId{
			UserId: utils.CreateID(0), // 違うユーザ
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	log := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = log.Add(ctx, db)
	require.NoError(t, err)

	accessToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

	access := models.SSOAccessToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: refreshToken,

		ClientID: service.ClientID,

		RedirectURI: "https://example.com",

		Create: time.Now(),

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		UserId: log.UserId,
	}
	err = access.Add(ctx, db)
	require.NoError(t, err)

	refresh := models.SSORefreshToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: refreshToken,

		ClientID: service.ClientID,

		RedirectURI: "https://example.com",

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  7,
		},

		UserId: log.UserId,
	}
	err = refresh.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entity != nil
	}, "最後の要素が格納されている")

	// ---

	err = logout.Delete(ctx, db, dummy.UserID)
	require.NoError(t, err)

	// --- チェックする

	goretry.Retry(t, func() bool {
		serviceA, err := models.GetSSOServiceByClientId(ctx, db, service.ClientID) // これは消えていない
		require.NoError(t, err)

		logs, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		accessTokenG, err := models.GetAccessTokenByAccessToken(ctx, db, accessToken)
		require.NoError(t, err)

		refreshTokenG, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return serviceA != nil && len(logs) == 0 && accessTokenG == nil && refreshTokenG == nil
	}, "ユーザが定義したSSOが消えている")
}
