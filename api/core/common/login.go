package common

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type LoginTokens struct {
	SessionToken string
	RefreshToken string
}

// ユーザIDを設定し、新たにログインをします
func LoginByUserID(ctx context.Context, db *database.Database, userId string, ip string, userAgent string) (*LoginTokens, error) {
	sessionToken := utils.CreateID(30)
	refreshToken := utils.CreateID(0)

	session := &models.SessionInfo{
		SessionToken: sessionToken,

		TokenInfo: models.TokenInfo{
			CreateDate: time.Now(),
			PeriodHour: 6,

			UserId: models.UserId{
				UserId: userId,
			},
		},
	}
	if err := session.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 40).Wrap()
	}

	refresh := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		TokenInfo: models.TokenInfo{
			CreateDate: time.Now(),
			PeriodHour: 168, // 24*7 = 168

			UserId: models.UserId{
				UserId: userId,
			},
		},
	}
	if err := refresh.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 40).Wrap()
	}

	// ログイン履歴を取る
	history := &models.LoginHistory{
		AccessId:     utils.CreateID(30),
		Date:         time.Now(),
		IpAddress:    ip,
		UserAgent:    userAgent,
		IsSSO:        false,
		SSOPublicKey: "",
	}
	if err := history.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 70).Wrap()
	}

	return &LoginTokens{
		SessionToken: sessionToken,
		RefreshToken: refreshToken,
	}, nil
}

// ログイン用のcookieをセットする
func LoginSetCookie(w http.ResponseWriter, login *LoginTokens) {
	// secure属性はproductionのみにする（テストが通らないため）
	secure := false
	if utils.DEPLOY_MODE == "production" {
		secure = true
	}
	// ブラウザ上でcookieを追加できるように、HttpOnlyはfalseにする
	cookie := net.NewCookie(os.Getenv("COOKIE_DOMAIN"), secure, http.SameSiteDefaultMode, false)

	sessionExp := net.NewSession()
	cookie.Set(w, "session-token", login.SessionToken, sessionExp)

	// リフレッシュトークンの期限は1週間
	refreshExp := net.NewCookieDayExp(7)
	cookie.Set(w, "refresh-token", login.RefreshToken, refreshExp)
}
