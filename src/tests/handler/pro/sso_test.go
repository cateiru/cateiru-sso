package pro_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/pro"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func ssoServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.ProSSOHandler)
	mux.HandleFunc("/image", handler.ProSSOImage)

	return mux
}

const LOGO_PATH = "icon.png"

func TestSSO(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.SetRequest{
		Name:    "Test",
		FromURL: []string{"https://example.com/login"},
		ToURL:   []string{"https://example.com/login/redirect"},
	}

	resp := s.Post(t, "/", form)

	var keys models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys)
	require.NoError(t, err)

	require.NotEmpty(t, keys.ClientID)
	require.Equal(t, keys.Name, form.Name)
	require.NotEmpty(t, keys.TokenSecret)

	// --- SSO一覧を取得する

	resp = s.Get(t, "/")

	var sso []pro.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 1)
	require.Equal(t, sso[0].UserId.UserId, dummy.UserID)
	require.Equal(t, sso[0].LoginCount, 0)

	// --- SSOを削除する

	s.Delete(t, fmt.Sprintf("/?id=%s", keys.ClientID))

	// --- もう一度一覧を取得する（削除されたか確認する）

	resp = s.Get(t, "/")

	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 0)
}

func TestNoProUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.SetRequest{
		Name:    "Test",
		FromURL: []string{"https://example.com/login"},
		ToURL:   []string{"https://example.com/login/redirect"},
	}

	requestForm, err := json.Marshal(form)
	require.NoError(t, err)

	resp, err := s.Client.Post(s.Server.URL+"/", "application/json", bytes.NewBuffer(requestForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}

func TestCount(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.SetRequest{
		Name:    "Test",
		FromURL: []string{"https://example.com/login"},
		ToURL:   []string{"https://example.com/login/redirect"},
	}

	resp := s.Post(t, "/", form)

	var keys models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys)
	require.NoError(t, err)

	// -- そのSSOでログインする

	log := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   keys.ClientID,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = log.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSSOServiceLogsByClientId(ctx, db, keys.ClientID)
		require.NoError(t, err)

		return len(entity) != 0
	}, "")

	resp = s.Get(t, "/")

	var sso []pro.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 1)
	require.Equal(t, sso[0].UserId.UserId, dummy.UserID)
	require.Equal(t, sso[0].LoginCount, 1)
}

func TestChangeServiceInfo(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.SetRequest{
		Name:    "Test",
		FromURL: []string{"https://example.com/login"},
		ToURL:   []string{"https://example.com/login/redirect"},
	}

	resp := s.Post(t, "/", form)

	var keys models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys)
	require.NoError(t, err)

	changeForm := pro.SetRequest{
		ClientId: keys.ClientID,
		FromURL:  []string{"https://cateiru.com"},

		AllowRoles: []string{"test"},

		ChangeTokenSecert: true,
	}

	resp = s.Post(t, "/", changeForm)

	var keys2 models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys2)
	require.NoError(t, err)

	require.Equal(t, keys.Name, keys2.Name)
	require.Equal(t, keys.ToUrl[0], keys2.ToUrl[0])

	require.NotEqual(t, keys.TokenSecret, keys2.TokenSecret)
	require.NotEqual(t, keys.FromUrl[0], keys2.FromUrl[0])

	require.Equal(t, keys2.AllowRoles[0], "test")

	changeForm2 := pro.SetRequest{
		ClientId: keys.ClientID,
		ToURL:    []string{"https://cateiru.com/login"},

		AllowRoles: []string{""},

		Name: "new",
	}

	resp = s.Post(t, "/", changeForm2)

	var keys3 models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys3)
	require.NoError(t, err)

	require.NotEqual(t, keys2.Name, keys3.Name)
	require.NotEqual(t, keys2.ToUrl[0], keys3.ToUrl[0])

	require.Equal(t, keys2.TokenSecret, keys3.TokenSecret)
	require.Equal(t, keys2.FromUrl[0], keys3.FromUrl[0])
	require.Len(t, keys3.AllowRoles, 0)

}

func TestSetImage(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.SetRequest{
		Name:    "Test",
		FromURL: []string{"https://example.com/login"},
		ToURL:   []string{"https://example.com/login/redirect"},
	}

	resp := s.Post(t, "/", form)

	var keys models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys)
	require.NoError(t, err)

	// --- 画像を追加する

	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "image/png")
	mh.Set("Content-Disposition", `form-data; name="image"; filename="icon.png"`)

	mh2 := make(textproto.MIMEHeader)
	mh2.Set("Content-Type", "text/plain")
	mh2.Set("Content-Disposition", `form-data; name="client_id";`)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreatePart(mh)
	require.NoError(t, err)
	f, err := os.Open(LOGO_PATH)
	require.NoError(t, err)
	defer f.Close()
	_, err = io.Copy(part, f)
	require.NoError(t, err)

	part, err = writer.CreatePart(mh2)
	require.NoError(t, err)
	part.Write([]byte(keys.ClientID))

	writer.Close()

	resp, err = s.Client.Post(s.Server.URL+"/image", writer.FormDataContentType(), body)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var respBody models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.NotEmpty(t, respBody.ServiceIcon)

	resp = s.Get(t, "/")

	var sso []pro.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 1)
	require.Equal(t, sso[0].UserId.UserId, dummy.UserID)
	require.Equal(t, sso[0].ServiceIcon, respBody.ServiceIcon)

	// storage に画像がセットされているか確認する
	resp, err = s.Client.Get("http://" + os.Getenv("STORAGE_EMULATOR_HOST") + "/cateiru-sso/sso/" + keys.ClientID)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	require.NotEmpty(t, tools.ConvertByteResp(resp))
}

func TestSSOAllowRole(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.SetRequest{
		Name:       "Test",
		FromURL:    []string{"https://example.com/login"},
		ToURL:      []string{"https://example.com/login/redirect"},
		AllowRoles: []string{"test"},
	}

	resp := s.Post(t, "/", form)

	var keys models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys)
	require.NoError(t, err)

	require.NotEmpty(t, keys.ClientID)
	require.Equal(t, keys.Name, form.Name)
	require.NotEmpty(t, keys.TokenSecret)

	// --- SSO一覧を取得する

	resp = s.Get(t, "/")

	var sso []pro.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 1)
	require.Equal(t, sso[0].UserId.UserId, dummy.UserID)
	require.Equal(t, sso[0].LoginCount, 0)
	require.Equal(t, sso[0].AllowRoles[0], "test")

	// --- SSOを削除する

	s.Delete(t, fmt.Sprintf("/?id=%s", keys.ClientID))

	// --- もう一度一覧を取得する（削除されたか確認する）

	resp = s.Get(t, "/")

	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 0)
}
