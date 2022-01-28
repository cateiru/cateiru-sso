package login

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

type RequestFrom struct {
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	ReCAPTCHA string `json:"re_captcha"`
}

type Response struct {
	IsOTP    bool   `json:"is_otp"`
	OTPToken string `json:"otp_token"`
}

type LoginState struct {
	Response

	common.LoginTokens
}

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	ctx := r.Context()

	var request RequestFrom
	err := net.GetJsonForm(w, r, &request)
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	ip := net.GetIPAddress(r)
	userAgent := net.GetUserAgent(r)

	loginState, err := Login(ctx, &request, ip, userAgent)
	if err != nil {
		return err
	}

	if !loginState.IsOTP {
		// OTPが設定されていない場合
		// ログイントークンをcookieにセットする
		common.LoginSetCookie(w, &loginState.LoginTokens)
	}

	net.ResponseOK(w, loginState.Response)

	return nil
}

// メールアドレスとパスワードでログインをする（試みる）
// もし、OTPが設定されている場合はパスコードの入力を求めます
//
// TODO: admin userの設定
func Login(ctx context.Context, form *RequestFrom, ip string, userAgent string) (*LoginState, error) {
	// reCAPTCHA
	if config.Defs.DeployMode == "production" {
		isOk, err := secure.NewReCaptcha().Validate(form.ReCAPTCHA, ip)
		if err != nil {
			return nil, err
		}
		// reCAPTCHAが認証できなかった場合、400を返す
		if !isOk {
			return nil, status.NewBadRequestError(errors.New("reCAPTCHA is failed")).Caller().AddCode(net.BotError)
		}
	}

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	cert, err := models.GetCertificationByMail(ctx, db, form.Mail)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// メールアドレスが登録されていない = アカウントが存在しない場合は400を返す
	if cert == nil {
		if common.CheckAdminMail(form.Mail) {
			// メールアドレスがadminであるとき
			userId, err := LoginAdmin(ctx, db, form)
			if err != nil {
				return nil, err
			}
			// ログイントークンをセットする
			login, err := common.LoginByUserID(ctx, db, userId, ip, userAgent)
			if err != nil {
				return nil, status.NewInternalServerErrorError(err).Caller()
			}

			return &LoginState{
				Response: Response{
					IsOTP: false, // OTPはセットされていないためfalse
				},
				LoginTokens: *login,
			}, nil

		} else {
			return nil, status.NewBadRequestError(
				errors.New("account is not found")).Caller().AddCode(net.AccountNoExist)
		}
	}

	// OTPが設定されている場合
	if len(cert.OnetimePasswordSecret) != 0 {
		id := utils.CreateID(0)
		otpBuffer := &models.OnetimePasswordBuffer{
			Id:      id,
			IsLogin: true,

			Period: models.Period{
				CreateDate:   time.Now(),
				PeriodMinute: 10,
			},

			UserId: cert.UserId,
		}

		if err = otpBuffer.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller()
		}

		return &LoginState{
			Response: Response{
				IsOTP:    true,
				OTPToken: id,
			},
		}, nil
	}

	// パスワードを検証
	// パスワードが違う場合は400を返す
	if !secure.ValidatePW(form.Password, cert.Password, cert.Salt) {
		return nil, status.NewBadRequestError(errors.New("no validate password")).Caller()
	}

	// ログイントークンをセットする
	login, err := common.LoginByUserID(ctx, db, cert.UserId.UserId, ip, userAgent)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &LoginState{
		Response: Response{
			IsOTP: false, // OTPはセットされていないためfalse
		},
		LoginTokens: *login,
	}, nil
}

// Adminでログインする
// 初回のみで、その後は通常と同じアカウントと同じ方法でログインします
func LoginAdmin(ctx context.Context, db *database.Database, form *RequestFrom) (string, error) {
	// パスワード検証
	if form.Password != config.Defs.AdminPassword {
		return "", status.NewBadRequestError(errors.New("admin pw is not validate")).Caller()
	}

	hashedPW, err := secure.PWHash(form.Password)
	if err != nil {
		return "", status.NewBadRequestError(err).Caller()
	}

	userID := utils.CreateID(30)

	newCert := &models.Certification{
		AccountCreateDate: time.Now(),
		UserMailPW: models.UserMailPW{
			Mail:     form.Mail,
			Password: hashedPW.Key,
			Salt:     hashedPW.Salt,
		},
		UserId: models.UserId{
			UserId: userID,
		},
	}
	if err := newCert.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	newUser := &models.User{
		FirstName: "Admin",
		LastName:  "User",
		UserName:  "admin",

		Mail: form.Mail,

		// TODO: 初期値設定する
		Theme:     "",
		AvatarUrl: "",

		UserId: models.UserId{
			UserId: userID,
		},
	}

	if err := newUser.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	role := &models.Role{
		Role: []string{"user", "pro", "admin"},

		UserId: models.UserId{
			UserId: userID,
		},
	}

	if err := role.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	return userID, nil
}
