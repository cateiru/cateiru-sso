package admin_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func adminMailCertServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/cert", handler.AdminMailCertLog)

	return mux
}

func TestMailCert(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	log := models.TryCreateAccountLog{
		LogId:      utils.CreateID(0),
		IP:         "172.0.0.1",
		TryDate:    time.Now(),
		TargetMail: "example@example.com",
	}
	err = log.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser().AddRole("admin")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	s := tools.NewTestServer(t, adminMailCertServer(), true)
	defer s.Close()

	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	resp := s.Get(t, "/cert")

	var logs []models.TryCreateAccountLog

	err = json.Unmarshal(tools.ConvertByteResp(resp), &logs)
	require.NoError(t, err)

	require.True(t, len(logs) > 0)
}
