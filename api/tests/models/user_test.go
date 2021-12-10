package models_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()
	userId := utils.CreateID(30)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userData := &models.User{
		FirstName: "あ",
		LastName:  "い",
		UserName:  "cateiru",

		Role:      []string{"user", "pro"},
		AvatarUrl: "",

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = userData.Add(ctx, db)
	require.NoError(t, err)

	// ----

	entry, err := models.GetUserDataByUserID(ctx, db, userId)
	require.NoError(t, err)
	require.NotNil(t, entry)

	require.Equal(t, entry.FirstName, "あ")
	require.Equal(t, entry.UserName, "cateiru")
}

func TestTXUser(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()
	userId := utils.CreateID(30)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userData := &models.User{
		FirstName: "あ",
		LastName:  "い",
		UserName:  "cateiru",

		Role:      []string{"user", "pro"},
		AvatarUrl: "",

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = userData.Add(ctx, db)
	require.NoError(t, err)

	// ---

	tx, err := database.NewTransaction(ctx, db)
	require.NoError(t, err)

	entry, err := models.GetUserDataTXByUserID(ctx, tx, userId)
	require.NoError(t, err)
	require.NotNil(t, entry)

	entry.LastName = "にゃあ"

	err = entry.AddTX(tx)
	require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)

	// ---

	entry, err = models.GetUserDataByUserID(ctx, db, userId)
	require.NoError(t, err)
	require.NotNil(t, entry)

	require.Equal(t, entry.LastName, "にゃあ")
}
