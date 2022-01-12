package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/user/otp"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
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

	s := tools.NewTestServer(t, otpServer(), true)
	s.AddSession(ctx, db, dummy)

	// ----

	resp := s.Get(t, "/otp")

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

	resp = s.Post(t, "/otp", send)

	var backups otp.SetOTPResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &backups)
	require.NoError(t, err)

	// ----

	resp = s.Get(t, "/otp/bup")

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
