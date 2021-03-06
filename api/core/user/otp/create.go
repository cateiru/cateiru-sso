package otp

import (
	"context"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

type GetOTPTokenResponse struct {
	Id       string `json:"id"`
	OtpToken string `json:"otp_token"`
}

// OPTのトークンURLを取得する
func GetOTPTokenURL(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	c := common.NewCert(w, r)
	if err := c.Login(ctx, db); err != nil {
		return err
	}
	userId := c.UserId

	resp, err := GenerateOTPToken(ctx, db, userId)
	if err != nil {
		return err
	}

	net.ResponseOK(w, resp)

	return nil
}

func GenerateOTPToken(ctx context.Context, db *database.Database, userId string) (*GetOTPTokenResponse, error) {
	mail, err := common.GetMailByUserID(ctx, db, userId)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	otp, err := secure.NewOnetimePassword(mail)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	id := utils.CreateID(0)
	otpKey := otp.GetPublic()

	buffer := &models.OnetimePasswordBuffer{
		Id: id,

		PublicKey: otpKey,
		SecretKey: otp.GetSecret(),

		// ログイン用ではないためfalse
		IsLogin: false,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 5,
		},

		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err := buffer.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &GetOTPTokenResponse{
		Id:       id,
		OtpToken: otpKey,
	}, nil
}
