package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/login"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/stretchr/testify/require"
)

func loginServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/login/otp", handler.LoginOnetimeHandler)

	return mux
}

func TestLoginNoOTP(t *testing.T) {
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

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}
	form, err := json.Marshal(send)
	require.NoError(t, err)

	resp, err := client.Post(server.URL+"/login", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var sessionToken string
	var refreshToken string
	for _, cookie := range jar.Cookies(url) {
		if cookie.Name == "session-token" {
			sessionToken = cookie.Value
		} else if cookie.Name == "refresh-token" {
			refreshToken = cookie.Value
		}
	}
	require.NotEmpty(t, sessionToken)
	require.NotEmpty(t, refreshToken)
}

func TestLoginOTP(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}
	form, err := json.Marshal(send)
	require.NoError(t, err)

	resp, err := client.Post(server.URL+"/login", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var response login.Response
	err = json.Unmarshal(tools.ConvertByteResp(resp), &response)
	require.NoError(t, err)

	require.True(t, response.IsOTP)

	// ---

	passcode, err := dummy.GenOTPCode()
	require.NoError(t, err)

	otpSend := login.OTPRequest{
		Passcode: passcode,
	}
	form, err = json.Marshal(otpSend)
	require.NoError(t, err)

	resp, err = client.Post(server.URL+"/login/otp", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// ---

	var sessionToken string
	var refreshToken string
	for _, cookie := range jar.Cookies(url) {
		if cookie.Name == "session-token" {
			sessionToken = cookie.Value
		} else if cookie.Name == "refresh-token" {
			refreshToken = cookie.Value
		}
	}
	require.NotEmpty(t, sessionToken)
	require.NotEmpty(t, refreshToken)
}
