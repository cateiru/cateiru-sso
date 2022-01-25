package models_test

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()
	userId := utils.CreateID(30)
	userName := "cateiru"

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userData := &models.User{
		FirstName: "あ",
		LastName:  "い",
		UserName:  userName,

		AvatarUrl: "",

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = userData.Add(ctx, db)
	require.NoError(t, err)

	// ----

	goretry.Retry(t, func() bool {
		entry, err := models.GetUserDataByUserID(ctx, db, userId)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	entry, err := models.GetUserDataByUserID(ctx, db, userId)
	require.NoError(t, err)
	require.NotNil(t, entry)

	require.Equal(t, entry.FirstName, "あ")
	require.Equal(t, entry.UserName, "cateiru")

	entity, err := models.GetUserDataByUserName(ctx, db, userName)
	require.NoError(t, err)
	require.NotNil(t, entity)

	require.Equal(t, entity.FirstName, "あ")
	require.Equal(t, entity.UserName, "cateiru")

}

func TestTXUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()
	userId := utils.CreateID(30)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userData := &models.User{
		FirstName: "あ",
		LastName:  "い",
		UserName:  "cateiru",

		AvatarUrl: "",

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = userData.Add(ctx, db)
	require.NoError(t, err)

	// ---

	var entry *models.User
	for i := 0; 3 > i; i++ {
		tx, err := database.NewTransaction(ctx, db)
		require.NoError(t, err)

		entry, err = models.GetUserDataTXByUserID(tx, userId)
		require.NoError(t, err)
		require.NotNil(t, entry)

		entry.LastName = "にゃあ"

		err = entry.AddTX(tx)
		require.NoError(t, err)

		err = tx.Commit()
		if err != nil && err != datastore.ErrConcurrentTransaction {
			t.Fatal()
		}
		if err == nil {
			return
		}
	}

	// ---

	entry, err = models.GetUserDataByUserID(ctx, db, userId)
	require.NoError(t, err)
	require.NotNil(t, entry)

	require.Equal(t, entry.LastName, "にゃあ")
}

func TestDeleteUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	// 実際に格納されているか確認する
	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "格納された")

	// 削除する
	err = models.DeleteUserDataByUserID(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity == nil
	}, "削除された")
}
