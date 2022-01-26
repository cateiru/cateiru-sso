package user_test

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func avatorServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/avator", handler.UserAvatorHandler)

	return mux
}

const LOGO_PATH = "logo.png"

func TestAvator(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, avatorServer(), true)
	defer s.Close()
	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("upload", LOGO_PATH)
	require.NoError(t, err)
	f, err := os.Open(LOGO_PATH)
	require.NoError(t, err)
	defer f.Close()
	_, err = io.Copy(part, f)
	require.NoError(t, err)
	writer.Close() // defer使わないでCloseしてしまう!! <--- ここ大事！数時間格闘した！！

	resp, err := s.Client.Post(s.Server.URL+"/avator", writer.FormDataContentType(), body)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- ちゃんとセットされているか確認する

	resp, err = s.Client.Get("http://" + os.Getenv("STORAGE_EMULATOR_HOST") + "/cateiru-sso/avator/" + dummy.UserID)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	require.NotEmpty(t, tools.ConvertByteResp(resp))

	// --- 削除する

	s.Delete(t, "/avator")

	// // --- ちゃんと削除されているか

	goretry.Retry(t, func() bool {
		resp, err = s.Client.Get("http://" + os.Getenv("STORAGE_EMULATOR_HOST") + "/cateiru-sso/avator/" + dummy.UserID)
		require.NoError(t, err)

		return resp.StatusCode == 404
	}, "削除されている")
}
