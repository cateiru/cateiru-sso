package login

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type OTPRequest struct {
	Passcode string `json:"passcode"`
	OtpToken string `json:"otp_token"`
}

// OTPを入力してログインする
func OTPLoginHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	ctx := r.Context()

	var otpRequest OTPRequest
	if err := net.GetJsonForm(w, r, &otpRequest); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	ip := net.GetIPAddress(r)
	userAgent := net.GetUserAgent(r)

	login, err := LoginOTP(ctx, otpRequest.OtpToken, otpRequest.Passcode, ip, userAgent)
	if err != nil {
		return err
	}

	common.LoginSetCookie(w, login)

	return nil
}

func LoginOTP(ctx context.Context, id string, passcode string, ip string, userAgent string) (*common.LoginTokens, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	if len(id) == 0 || len(passcode) == 0 {
		return nil, status.NewBadRequestError(errors.New("incomplete form"))
	}

	buffer, err := models.GetOTPBufferByID(ctx, db, id)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// idが存在しない場合は400を返す
	if buffer == nil {
		return nil, status.NewBadRequestError(errors.New("entity not found")).Caller().AddCode(net.FailedLogin)
	}

	// bufferの有効期限切れの場合は400を返す
	if common.CheckExpired(&buffer.Period) {
		return nil, status.NewBadRequestError(err).Caller().AddCode(net.TimeOutError).AddCode(net.TimeOutError)
	}

	// （ないとは思うが）isLoginをチェックする
	if !buffer.IsLogin {
		return nil, status.NewInternalServerErrorError(errors.New("no IsLogin")).Caller()
	}

	cert, err := models.GetCertificationByUserID(ctx, db, buffer.UserId.UserId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if cert == nil {
		return nil, status.NewInternalServerErrorError(errors.New("entity not found")).Caller()
	}

	// OTPが設定されていない場合は400を返す
	if len(cert.OnetimePasswordSecret) == 0 {
		return nil, status.NewBadRequestError(errors.New("otp not set")).Caller().AddCode(net.FailedLogin)
	}

	ok, update := common.CheckOTP(passcode, cert, nil)
	// OTPが認証できない場合は400を返す
	if !ok {
		return nil, status.NewBadRequestError(errors.New("otp not varidated")).Caller().AddCode(net.FailedLogin)
	}

	// backupが更新された場合はDBを更新する
	if update {
		if err := cert.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller()
		}
	}

	if err := models.DeleteOTPBuffer(ctx, db, id); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// ログイントークンをセットする
	login, err := common.LoginByUserID(ctx, db, cert.UserId.UserId, ip, userAgent)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return login, nil
}
