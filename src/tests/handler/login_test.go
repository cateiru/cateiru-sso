package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/login"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/stretchr/testify/require"
)

func loginServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/login/otp", handler.LoginOnetimeHandler)

	return mux
}

func TestLoginNoOTP(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser() // OTPは設定しない
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, loginServer(), true)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}

	s.Post(t, "/login", send)

	s.FindCookies(t, []string{"session-token", "refresh-token"})
}

func TestLoginOTP(t *testing.T) {
	config.TestInit(t)

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

	s := tools.NewTestServer(t, loginServer(), true)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}

	resp := s.Post(t, "/login", send)

	var response login.Response
	err = json.Unmarshal(tools.ConvertByteResp(resp), &response)
	require.NoError(t, err)

	require.True(t, response.IsOTP)

	// ---

	passcode, err := dummy.GenOTPCode()
	require.NoError(t, err)

	otpSend := login.OTPRequest{
		Passcode: passcode,
		OtpToken: response.OTPToken,
	}

	s.Post(t, "/login/otp", otpSend)

	// ---

	s.FindCookies(t, []string{"session-token", "refresh-token"})
}

func TestFailedPWNoOTP(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser() // OTPは設定しない
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, loginServer(), true)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "hoge",
	}

	reqForm, err := json.Marshal(send)
	require.NoError(t, err)

	resp, err := s.Client.Post(s.Server.URL+"/login", "application/json", bytes.NewBuffer(reqForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}

func TestFailedPWOTP(t *testing.T) {
	config.TestInit(t)

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

	s := tools.NewTestServer(t, loginServer(), true)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "hoge",
	}

	reqForm, err := json.Marshal(send)
	require.NoError(t, err)

	resp, err := s.Client.Post(s.Server.URL+"/login", "application/json", bytes.NewBuffer(reqForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}

func TestFailedOTP(t *testing.T) {
	config.TestInit(t)

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

	s := tools.NewTestServer(t, loginServer(), true)

	send := login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}

	resp := s.Post(t, "/login", send)

	var response login.Response
	err = json.Unmarshal(tools.ConvertByteResp(resp), &response)
	require.NoError(t, err)

	require.True(t, response.IsOTP)

	// ---

	otpSend := login.OTPRequest{
		Passcode: "123456",
		OtpToken: response.OTPToken,
	}

	reqForm, err := json.Marshal(otpSend)
	require.NoError(t, err)

	resp, err = s.Client.Post(s.Server.URL+"/login/otp", "application/json", bytes.NewBuffer(reqForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}
