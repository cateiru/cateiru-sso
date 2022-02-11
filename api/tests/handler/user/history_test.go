package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	historyType "github.com/cateiru/cateiru-sso/api/core/user/history"
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
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, historyServer(), true)
	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	// ログインしているやつの履歴
	history := &models.LoginHistory{
		AccessId:  dummy.AccessID,
		Date:      time.Now(),
		IpAddress: "192.168.0.1",
		UserAgent: "",

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = history.Add(ctx, db)
	require.NoError(t, err)

	for i := 0; 10 > i; i++ {
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
	}

	// ちゃんと入っているか一度確認する
	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(histores) == 11 && histores[0].IpAddress == "192.168.0.1"
	}, "Entityが10個ある")

	resp := s.Get(t, "/login")

	var histories []historyType.LoginHistory

	err = json.Unmarshal(tools.ConvertByteResp(resp), &histories)
	require.NoError(t, err)

	require.Len(t, histories, 11)

	isThisDevice := 0
	for _, history := range histories {
		if history.ThisDevice {
			isThisDevice += 1
		}
	}
	require.Equal(t, 1, isThisDevice)

	// --- limit指定

	respLimit := s.Get(t, "/login?limit=3")

	var historiesLimit []*historyType.LoginHistory

	err = json.Unmarshal(tools.ConvertByteResp(respLimit), &historiesLimit)
	require.NoError(t, err)

	require.Len(t, historiesLimit, 3)
}
