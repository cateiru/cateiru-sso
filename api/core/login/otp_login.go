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

	c := common.NewCert(w, r).AddUser()

	if err := LoginOTP(ctx, otpRequest.OtpToken, otpRequest.Passcode, c); err != nil {
		return err
	}

	c.SetCookie()

	return nil
}

func LoginOTP(ctx context.Context, id string, passcode string, c *common.Cert) error {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	if len(id) == 0 || len(passcode) == 0 {
		return status.NewBadRequestError(errors.New("incomplete form"))
	}

	buffer, err := models.GetOTPBufferByID(ctx, db, id)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// idが存在しない場合は400を返す
	if buffer == nil {
		return status.NewBadRequestError(errors.New("entity not found")).Caller().AddCode(net.FailedLogin)
	}

	// bufferの有効期限切れの場合は400を返す
	if common.CheckExpired(&buffer.Period) {
		return status.NewBadRequestError(err).Caller().AddCode(net.TimeOutError).AddCode(net.TimeOutError)
	}

	// （ないとは思うが）isLoginをチェックする
	if !buffer.IsLogin {
		return status.NewInternalServerErrorError(errors.New("no IsLogin")).Caller()
	}

	cert, err := models.GetCertificationByUserID(ctx, db, buffer.UserId.UserId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if cert == nil {
		return status.NewInternalServerErrorError(errors.New("entity not found")).Caller()
	}

	// OTPが設定されていない場合は400を返す
	if len(cert.OnetimePasswordSecret) == 0 {
		return status.NewBadRequestError(errors.New("otp not set")).Caller().AddCode(net.FailedLogin)
	}

	ok, update := common.CheckOTP(passcode, cert, nil)
	// OTPが認証できない場合は400を返す
	if !ok {
		// 総当り対策で、4回間違えると認証不可になる
		if buffer.FailedCount < 3 {
			buffer.FailedCount += 1

			if err := buffer.Add(ctx, db); err != nil {
				return status.NewInternalServerErrorError(err).Caller()
			}
		} else {
			if err := models.DeleteOTPBuffer(ctx, db, id); err != nil {
				return status.NewInternalServerErrorError(err).Caller()
			}
		}

		return status.NewBadRequestError(errors.New("otp not varidated")).Caller().AddCode(net.FailedOTP)
	}

	// backupが更新された場合はDBを更新する
	if update {
		if err := cert.Add(ctx, db); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	if err := models.DeleteOTPBuffer(ctx, db, id); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// ログイントークンをセットする
	if err := c.NewLogin(ctx, db, cert.UserId.UserId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
