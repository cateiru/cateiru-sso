package info_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/user/info"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestChangeInfo(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	chaned := info.Request{
		FirstName: "New",
		LastName:  "taro",
		UserName:  "cannn",
		Theme:     "wwwww",
	}

	user, err := info.ChangeInfo(ctx, db, dummy.UserID, &chaned)
	require.NoError(t, err)

	require.Equal(t, user.FirstName, chaned.FirstName)
	require.Equal(t, user.LastName, chaned.LastName)
	require.Equal(t, chaned.UserName, chaned.UserName)
	require.Equal(t, user.Theme, chaned.Theme)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil && entity.FirstName == chaned.FirstName
	}, "")
}

func TestFailedUserName(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	old, err := dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	chaned := info.Request{
		UserName: "00",
	}

	_, err = info.ChangeInfo(ctx, db, dummy.UserID, &chaned)
	require.Error(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil && entity.FirstName == old.FirstName
	}, "変わっていない")
}

func TestAlreadyExistUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	u, err := dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	// ちゃんとDBに入っているか確認する
	goretry.Retry(t, func() bool {
		_u, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return _u != nil
	}, "")

	chaned := info.Request{
		UserName: u.UserName,
	}

	// ユーザはすでに存在しているのでエラーになる
	_, err = info.ChangeInfo(ctx, db, dummy.UserID, &chaned)
	require.Error(t, err)
}
