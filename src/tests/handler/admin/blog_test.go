package admin_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/admin"
	createaccount "github.com/cateiru/cateiru-sso/src/core/create_account"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func adminBlock() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", handler.CreateHandler)
	mux.HandleFunc("/ban", handler.AdminBanHandler)

	return mux
}

func TestBlockMail(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	blockMail := tools.NewDummyUser().Mail

	dummy := tools.NewDummyUser().AddRole("admin")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	s := tools.NewTestServer(t, adminBlock(), true)
	defer s.Close()

	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	form := createaccount.PostForm{
		Mail:      blockMail,
		ReCAPTCHA: "",
	}

	s.Post(t, "/create", &form) // blockしてないので通る

	// -- blockする

	ban := admin.BanRequest{
		Mail: blockMail,
	}

	s.Post(t, "/ban", &ban)

	resp := s.Get(t, "/ban?mode=mail")

	var respBody []models.MailBlockList

	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.True(t, len(respBody) > 0)

	// もう一回やるとブロックされている

	fromJ, err := json.Marshal(form)
	require.NoError(t, err)

	resp, err = s.Client.Post(s.Server.URL+"/create", "application/json", bytes.NewBuffer(fromJ))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403)

	s.Delete(t, fmt.Sprintf("/ban?mode=mail&element=%s", blockMail))

	s.Post(t, "/create", &form) // 削除したので通る
}

func TestBlockIP(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	blockIP := utils.CreateID(10)

	dummy := tools.NewDummyUser().AddRole("admin")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	s := tools.NewTestServer(t, adminBlock(), true)
	defer s.Close()

	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	ban := admin.BanRequest{
		IP: blockIP,
	}

	s.Post(t, "/ban", &ban)

	resp := s.Get(t, "/ban?mode=ip")

	var respBody []models.IPBlockList

	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	include := false
	for _, element := range respBody {
		if element.IP == blockIP {
			include = true
			break
		}
	}
	require.True(t, include)

	s.Delete(t, fmt.Sprintf("/ban?mode=ip&element=%s", blockIP))

	resp = s.Get(t, "/ban?mode=ip")

	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	include = false
	for _, element := range respBody {
		if element.IP == blockIP {
			include = true
			break
		}
	}
	require.False(t, include)

}
