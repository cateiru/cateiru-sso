package createaccount_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/posener/wstest"
	"github.com/stretchr/testify/require"
)

func observerServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", observeHandler)

	return mux
}

func observeHandler(w http.ResponseWriter, r *http.Request) {
	if err := createaccount.MailVerifyObserve(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func TestObserve(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	form := &createaccount.PostForm{
		Mail:      fmt.Sprintf("%s@example.com", utils.CreateID(4)),
		ReCAPTCHA: "",
	}
	ip := "192.168.1.1"

	clientToken, err := createaccount.CreateTemporaryAccount(ctx, form, ip)
	require.NoError(t, err)

	////

	server := observerServer()

	d := wstest.NewDialer(server)

	c, resp, err := d.Dial(fmt.Sprintf("ws://whatever/?cct=%s", clientToken), nil)
	require.NoError(t, err)
	got, want := resp.StatusCode, http.StatusSwitchingProtocols
	require.Equal(t, got, want)

	go verifyMail(ctx, t, clientToken)

	// 受信待機
	var respm bool
	err = c.ReadJSON(&respm)
	require.NoError(t, err)

	// 返ってくる = メール認証が完了したためclient側からwsをcloseする
	err = c.Close()
	require.NoError(t, err)

	// response messageは`true`が返る
	require.True(t, respm)
}

func verifyMail(ctx context.Context, t *testing.T, clientToken string) {
	// 3秒間待機する: WSで待機するため
	time.Sleep(3 * time.Second)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	cert, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
	require.NoError(t, err)
	require.NotNil(t, cert)

	t.Logf("verify mailToken: %s", cert.MailToken)
	cert.Verify = true

	err = cert.Add(ctx, db)
	require.NoError(t, err)
}
