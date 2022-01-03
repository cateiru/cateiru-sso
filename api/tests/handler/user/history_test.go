package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func historyServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.UserHistoryHandler)

	return mux
}

func TestLoginHistory(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	for i := 0; 10 > i; i++ {
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
	}

	// ちゃんと入っているか一度確認する
	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(histores) == 10 && histores[0].IpAddress == "192.168.0.1"
	}, "Entityが10個ある")

	s := tools.NewTestServer(t, historyServer(), true)
	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	resp := s.Get(t, "/login")

	var histories []models.LoginHistory

	err = json.Unmarshal(tools.ConvertByteResp(resp), &histories)
	require.NoError(t, err)

	require.Len(t, histories, 10)

	// --- limit指定

	respLimit := s.Get(t, "/login?limit=3")

	var historiesLimit []models.LoginHistory

	err = json.Unmarshal(tools.ConvertByteResp(respLimit), &historiesLimit)
	require.NoError(t, err)

	require.Len(t, historiesLimit, 3)
}
