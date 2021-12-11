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
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestTryCreateAccountLog(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mail := fmt.Sprintf("%s@example.com", utils.CreateID(4))
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))

	entry := &models.TryCreateAccountLog{
		LogId:      utils.CreateID(10),
		IP:         ip,
		TryDate:    time.Now(),
		TargetMail: mail,
	}
	err = entry.Add(ctx, db)
	require.NoError(t, err)

	// メールアドレスが一緒なもう一つのログを追加する
	entry2 := &models.TryCreateAccountLog{
		LogId:      utils.CreateID(10),
		IP:         fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255)),
		TryDate:    time.Now(),
		TargetMail: mail,
	}
	err = entry2.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetTryCreateAccountLogByIP(ctx, db, ip)
		require.NoError(t, err)

		return len(logs) != 0
	}, "要素がある")

	logs, err := models.GetTryCreateAccountLogByIP(ctx, db, ip)
	require.NoError(t, err)
	require.Equal(t, len(logs), 1, "1つの要素が見つかった")
	require.Equal(t, logs[0].IP, ip, "IPが取得できる")
	require.Equal(t, logs[0].TargetMail, mail, "メールアドレスが取得できる")

	logs, err = models.GetTryCreateAccountLogByIP(ctx, db, "256.256.256.256")
	require.NoError(t, err)
	require.Equal(t, len(logs), 0, "なにも要素は見つからない")

	logs, err = models.GetTryCreateAccountLogByMail(ctx, db, mail)
	require.NoError(t, err)
	require.Equal(t, len(logs), 2, "2つの要素が見つかった")

	logs, err = models.GetTryCreateAccountLogByMail(ctx, db, "example@example.com")
	require.NoError(t, err)
	require.Equal(t, len(logs), 0, "なにも要素は見つからない")
}
