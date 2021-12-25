package login

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

type RequestFrom struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type Response struct {
	IsOTP bool   `json:"is_otp"`
	OTPId string `json:"otp_id"`
}

type LoginState struct {
	IsOTP bool
	OTPId string

	common.LoginTokens
}

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
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

	if loginState.IsOTP {
		// OTPが設定している場合

		// secure属性はproductionのみにする（テストが通らないため）
		secure := false
		if utils.DEPLOY_MODE == "production" {
			secure = true
		}
		// ブラウザ上でcookieを追加できるように、HttpOnlyはfalseにする
		cookie := net.NewCookie(os.Getenv("COOKIE_DOMAIN"), secure, http.SameSiteDefaultMode, false)

		sessionExp := net.NewSession()
		cookie.Set(w, "otp-token", loginState.OTPId, sessionExp)

		resp := Response{
			OTPId: loginState.OTPId,
			IsOTP: true,
		}
		net.ResponseOK(w, resp)
	} else {
		// OTPは設定されていない場合
		// ログイントークンをcookieにセットする
		common.LoginSetCookie(w, &loginState.LoginTokens)
	}

	return nil
}

// メールアドレスとパスワードでログインをする（試みる）
// もし、OTPが設定されている場合はパスコードの入力を求めます
//
// TODO: admin userの設定
func Login(ctx context.Context, form *RequestFrom, ip string, userAgent string) (*LoginState, error) {
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
				IsOTP:       false, // OTPはセットされていないためfalse
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
			Id:        id,
			SecretKey: cert.OnetimePasswordSecret,
			Backups:   cert.OnetimePasswordBackups,

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
			IsOTP: true,
			OTPId: id,
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
		IsOTP:       false, // OTPはセットされていないためfalse
		LoginTokens: *login,
	}, nil
}

// Adminでログインする
// 初回のみで、その後は通常と同じアカウントと同じ方法でログインします
func LoginAdmin(ctx context.Context, db *database.Database, form *RequestFrom) (string, error) {
	// パスワード検証
	if form.Password != os.Getenv("ADMIN_PASSWORD") {
		return "", status.NewBadRequestError(errors.New("admin pw is not validate")).Caller()
	}

	hashedPW := secure.PWHash(form.Password)
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

		Role: []string{"user", "pro", "admin"},

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

	return userID, nil
}
