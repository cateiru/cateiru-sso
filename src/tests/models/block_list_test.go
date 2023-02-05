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
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestIPBlockListDB(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	config.TestInit(t)

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

	// 初回のみリトライする
	goretry.Retry(t, func() bool {
		result, err := models.GetBlockListByIP(ctx, db, ip)
		require.NoError(t, err)

		return result != nil
	}, "block ipのentryがある")

	result, err := models.GetBlockListByIP(ctx, db, "256.256.256.256")
	require.NoError(t, err)
	require.Nil(t, result, "block ipのentryはない")

	result2, err := models.GetAllBlocIP(ctx, db)
	require.NoError(t, err)
	require.True(t, len(result2) > 0)

	err = models.DeleteBlockIP(ctx, db, ip)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		result, err := models.GetBlockListByIP(ctx, db, ip)
		require.NoError(t, err)

		return result == nil
	}, "")
}

func TestMailBlockListDB(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	config.TestInit(t)

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

	// 初回のみリトライする
	goretry.Retry(t, func() bool {
		result, err := models.GetBlockListByMail(ctx, db, mail)
		require.NoError(t, err)

		return result != nil
	}, "block mailのentryがある")

	result, err := models.GetBlockListByMail(ctx, db, "example@example.com")
	require.NoError(t, err)
	require.Nil(t, result, "block mailのentryはない")

	result2, err := models.GetAllBlocMail(ctx, db)
	require.NoError(t, err)
	require.True(t, len(result2) > 0)

	err = models.DeleteBlockMail(ctx, db, mail)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		result, err := models.GetBlockListByMail(ctx, db, mail)
		require.NoError(t, err)

		return result == nil
	}, "")
}
