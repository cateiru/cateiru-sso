package common

import (
	"context"
	"errors"
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

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodHour: 6,
		},

		UserId: models.UserId{
			UserId: userId,
		},
	}
	if err := session.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 40).Wrap()
	}

	refresh := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  7,
		},

		UserId: models.UserId{
			UserId: userId,
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
		UserId: models.UserId{
			UserId: userId,
		},
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

// cookieからログインします

// UserIDを返します
func LoginByCookie(ctx context.Context, db *database.Database, w http.ResponseWriter, r *http.Request) (string, error) {
	refreshToken, err := net.GetCookie(r, "refresh-token")
	if err != nil || len(refreshToken) == 0 {
		// cookieが存在しない、valueが存在しない場合は403を返す
		return "", status.NewForbiddenError(errors.New("cookie is not found")).Caller(
			"core/create_account/info.go", 36)
	}

	tx, err := database.NewTransaction(ctx, db)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 99).Wrap()
	}

	refresh, err := models.GetRefreshTokenTX(tx, refreshToken)
	if err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			err = errors.New(err.Error() + rerr.Error())
		}
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 104).Wrap()
	}

	// refreshtokenが存在しない場合は400を返す
	if refresh == nil {
		err = tx.Rollback()
		if err == nil {
			err = errors.New("refresh token is not exist")
		}

		return "", status.NewBadRequestError(err).Caller(
			"core/common/login.go", 111).Wrap()
	}

	// refresh-tokenが有効期限切れの場合は400を返す
	if CheckExpired(&refresh.Period) {
		err = tx.Rollback()
		if err == nil {
			err = errors.New("Expired")
		}

		return "", status.NewBadRequestError(err).Caller(
			"core/common/login.go", 111).AddCode(net.TimeOutError).Wrap()
	}

	// session-tokenを削除する（ある場合は）
	err = models.DeleteSessionTokenTX(tx, refresh.SessionToken)
	if err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			err = errors.New(err.Error() + rerr.Error())
		}
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 131).Wrap()
	}

	// refresh-tokenを削除する
	err = models.DeleteRefreshTokenTX(tx, refresh.RefreshToken)
	if err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			err = errors.New(err.Error() + rerr.Error())
		}
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 131).Wrap()
	}

	newSessionToken := utils.CreateID(30)
	newRefreshToken := utils.CreateID(0)

	// 新しいsession-tokenを作成する
	session := &models.SessionInfo{
		SessionToken: newSessionToken,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodHour: 6,
		},

		UserId: refresh.UserId,
	}
	if err := session.AddTX(tx); err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			err = errors.New(err.Error() + rerr.Error())
		}
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 131).Wrap()
	}

	// 新しいrefresh-tokenを作成する
	newRefresh := &models.RefreshInfo{
		RefreshToken: newRefreshToken,
		SessionToken: newSessionToken,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  7,
		},

		UserId: refresh.UserId,
	}
	if err := newRefresh.AddTX(tx); err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			err = errors.New(err.Error() + rerr.Error())
		}
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/common/login.go", 131).Wrap()
	}

	// cookieを上書き
	// 同じkeyでcookieを設定すれば上書きされるはず
	login := &LoginTokens{
		SessionToken: newSessionToken,
		RefreshToken: newRefreshToken,
	}
	LoginSetCookie(w, login)

	return refresh.UserId.UserId, nil
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
