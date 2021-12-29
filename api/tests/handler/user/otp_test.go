package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/user/otp"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	_otp "github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

func otpServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/otp", handler.UserOnetimePWHandler)
	mux.HandleFunc("/otp/bup", handler.UserOnetimePWBackupHandler)

	return mux
}

func TestOTP(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser() // OTPは設定しない
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)
	session, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	app := otpServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	tools.SetCookie(jar, "session-token", session, exp, url)
	tools.SetCookie(jar, "refresh-token", refresh, exp, url)

	// ----

	resp, err := client.Get(server.URL + "/otp")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var otpToken otp.GetOTPTokenResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &otpToken)
	require.NoError(t, err)

	w, err := _otp.NewKeyFromURL(otpToken.OtpToken)
	require.NoError(t, err, "publicを解析できない")

	code, err := totp.GenerateCode(w.Secret(), time.Now().UTC())
	require.NoError(t, err)

	// ----

	send := otp.OTPRequest{
		Type:     "enable",
		Passcode: code,
		Id:       otpToken.Id,
	}
	form, err := json.Marshal(send)
	require.NoError(t, err)

	resp, err = client.Post(server.URL+"/otp", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var backups otp.SetOTPResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &backups)
	require.NoError(t, err)

	// ----

	resp, err = client.Get(server.URL + "/otp/bup")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var getBackups otp.ResponseBackups
	err = json.Unmarshal(tools.ConvertByteResp(resp), &getBackups)
	require.NoError(t, err)

	// ----

	for _, code := range backups.Backups {
		flag := false
		for _, getCode := range getBackups.Codes {
			if getCode == code {
				flag = true
			}
		}
		if !flag {
			t.Fatal("OTPが無いやつがある")
		}
	}
}
