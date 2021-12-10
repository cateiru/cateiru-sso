package common

import (
	"context"
	"net/http"
	"os"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

type LoginTokens struct {
	SessionToken string
	RefreshToken string
}

// ユーザIDを設定し、新たにログインをします
func LoginByUserID(ctx context.Context, db *database.Database, userId string) (*LoginTokens, error) {
	// TODO
	return nil, nil
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
