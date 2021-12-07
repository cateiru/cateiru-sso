package models_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

func TestIPBlockListDB(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))

	block := &models.IPBlockList{
		IP: ip,
	}
	err = block.Add(ctx, db)
	require.NoError(t, err)

	result, err := models.GetBlockListByIP(ctx, db, ip)
	require.NoError(t, err)
	require.NotNil(t, &result, "block ipのentryがある")

	result, err = models.GetBlockListByIP(ctx, db, "256.256.256.256")
	require.NoError(t, err)
	require.Nil(t, result, "block ipのentryはない")
}

func TestMailBlockListDB(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mail := fmt.Sprintf("%s@example.com", utils.CreateID(4))

	block := &models.MailBlockList{
		Mail: mail,
	}
	err = block.Add(ctx, db)
	require.NoError(t, err)

	result, err := models.GetBlockListByIP(ctx, db, mail)
	require.NoError(t, err)
	require.NotNil(t, &result, "block mailのentryがある")

	result, err = models.GetBlockListByIP(ctx, db, "example@example.com")
	require.NoError(t, err)
	require.Nil(t, result, "block mailのentryはない")
}
