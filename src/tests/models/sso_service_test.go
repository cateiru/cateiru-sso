package models_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestSSOService(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientIds := []string{utils.CreateID(0), utils.CreateID(0)}
	dummy := tools.NewDummyUser()

	// 複数追加する
	for _, publicKey := range clientIds {
		entity := models.SSOService{
			ClientID: publicKey,

			TokenSecret: utils.CreateID(0),

			Name:        "Test",
			ServiceIcon: "",
			FromUrl:     []string{"https://example.com/login"},
			ToUrl:       []string{"https://example.com/login/redirect"},

			AllowRoles: []string{"hoge"},

			UserId: models.UserId{
				UserId: dummy.UserID,
			},
		}

		err = entity.Add(ctx, db)
		require.NoError(t, err)
	}

	goretry.Retry(t, func() bool {
		a, err := models.GetSSOServiceByClientId(ctx, db, clientIds[0])
		require.NoError(t, err)

		return a != nil && a.UserId.UserId == dummy.UserID
	}, "ちゃんと格納できて取得できる")

	goretry.Retry(t, func() bool {
		entities, err := models.GetSSOServiceByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(entities) == 2
	}, "UserIDをkeyにして複数取得できる")

	err = models.DeleteSSOServiceByClientId(ctx, db, clientIds[1])
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entities, err := models.GetSSOServiceByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(entities) == 1
	}, "削除できている")
}
