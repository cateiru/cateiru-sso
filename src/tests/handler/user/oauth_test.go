package user_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/user"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func oauthServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/oauth", handler.UserOAuthHandler)

	return mux
}

func TestGetOAuth(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	// SSO サービス x2
	service1 := models.SSOService{
		ClientID:    utils.CreateID(30),
		TokenSecret: utils.CreateID(0),

		Name:        "test1",
		ServiceIcon: "",

		FromUrl: []string{},
		ToUrl:   []string{},

		UserId: models.UserId{
			UserId: utils.CreateID(0), // 違うユーザ
		},
	}
	err = service1.Add(ctx, db)
	require.NoError(t, err)

	service2 := models.SSOService{
		ClientID:    utils.CreateID(30),
		TokenSecret: utils.CreateID(0),

		Name:        "test2",
		ServiceIcon: "",

		FromUrl: []string{},
		ToUrl:   []string{},

		UserId: models.UserId{
			UserId: utils.CreateID(0), // 違うユーザ2
		},
	}
	err = service2.Add(ctx, db)
	require.NoError(t, err)

	// ユーザはそれぞれ、1: 2回, 2: 1回ログインしてログがある
	log11 := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service1.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	log12 := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service1.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = log11.Add(ctx, db)
	require.NoError(t, err)
	err = log12.Add(ctx, db)
	require.NoError(t, err)

	log21 := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service2.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = log21.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(logs) == 3
	}, "")

	// ----

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	resp := s.Get(t, "/oauth")

	var reqBody []user.SSOLoginLog
	err = json.Unmarshal(tools.ConvertByteResp(resp), &reqBody)
	require.NoError(t, err)

	require.Len(t, reqBody, 2) // 実際、ログは3つあるがclient idをkeyにすると2つ

	allLogItems := len(reqBody[0].Logs) + len(reqBody[1].Logs)
	require.Equal(t, allLogItems, 3)
}

func TestDeleteService(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	// SSO サービス x2
	service1 := models.SSOService{
		ClientID:    utils.CreateID(30),
		TokenSecret: utils.CreateID(0),

		Name:        "test1",
		ServiceIcon: "",

		FromUrl: []string{},
		ToUrl:   []string{},

		UserId: models.UserId{
			UserId: utils.CreateID(0), // 違うユーザ
		},
	}
	err = service1.Add(ctx, db)
	require.NoError(t, err)

	service2 := models.SSOService{
		ClientID:    utils.CreateID(30),
		TokenSecret: utils.CreateID(0),

		Name:        "test2",
		ServiceIcon: "",

		FromUrl: []string{},
		ToUrl:   []string{},

		UserId: models.UserId{
			UserId: utils.CreateID(0), // 違うユーザ2
		},
	}
	err = service2.Add(ctx, db)
	require.NoError(t, err)

	// ユーザはそれぞれ、1: 2回, 2: 1回ログインしてログがある
	log11 := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service1.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	log12 := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service1.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = log11.Add(ctx, db)
	require.NoError(t, err)
	err = log12.Add(ctx, db)
	require.NoError(t, err)

	log21 := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   service2.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = log21.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(logs) == 3
	}, "")

	// ----

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	s.Delete(t, fmt.Sprintf("/oauth?id=%s", service1.ClientID))

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(logs) == 1
	}, "削除できている")

}
