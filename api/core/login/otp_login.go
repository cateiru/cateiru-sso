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

	id, err := net.GetCookie(r, "otp-token")
	if err != nil {
		// cookieが設定されていない場合は400を返す
		return status.NewBadRequestError(err).Caller()
	}

	// cookieを削除する
	err = net.DeleteCookie(w, r, "otp-token")
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	ip := net.GetIPAddress(r)
	userAgent := net.GetUserAgent(r)

	login, err := LoginOTP(ctx, id, otpRequest.Passcode, ip, userAgent)
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

	buffer, err := models.GetOTPBufferByID(ctx, db, id)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// idが存在しない場合は400を返す
	if buffer == nil {
		return nil, status.NewBadRequestError(errors.New("entity not found")).Caller()
	}

	// bufferの有効期限切れの場合は400を返す
	if common.CheckExpired(&buffer.Period) {
		return nil, status.NewBadRequestError(err).Caller()
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
		return nil, status.NewBadRequestError(errors.New("otp not set")).Caller()
	}

	ok, update := common.CheckOTP(passcode, cert, nil)
	// OTPが認証できない場合は400を返す
	if !ok {
		return nil, status.NewBadRequestError(errors.New("otp not varidated")).Caller()
	}
	// backupが更新された場合はDBを更新する
	if update {
		if err := cert.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller()
		}
	}

	// ログイントークンをセットする
	login, err := common.LoginByUserID(ctx, db, cert.UserId.UserId, ip, userAgent)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return login, nil
}
