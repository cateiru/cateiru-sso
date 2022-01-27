package common

import (
	"context"
	"errors"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
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
	sessionToken := utils.CreateID(0)
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
		return nil, status.NewInternalServerErrorError(err).Caller()
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
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if err := SetLoginHistory(ctx, db, userId, ip, userAgent); err != nil {
		return nil, err
	}

	return &LoginTokens{
		SessionToken: sessionToken,
		RefreshToken: refreshToken,
	}, nil
}

// session-tokenのcookieからuser idを取得します
// もし、session-tokenが無いor認証できない場合はrefresh-tokenを用いてログインを試みます
func GetUserID(ctx context.Context, db *database.Database, w http.ResponseWriter, r *http.Request) (string, error) {
	sessionToken, err := net.GetCookie(r, "session-token")
	if err != nil {
		return "", status.NewBadRequestError(err).Caller()
	}

	// session-tokenが存在しない場合、refresh-tokenからsession-tokenを作成する
	if sessionToken == "" {
		return LoginByCookie(ctx, db, w, r)
	}

	session, err := models.GetSessionToken(ctx, db, sessionToken)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	// sessionTokenが見つからない場合、refresh-tokenを使用してsession-tokenの作成を試みます
	if session == nil {
		return LoginByCookie(ctx, db, w, r)
	}

	// sessionTokenの有効期限が切れてしまっている場合、refresh-tokenを使用してsession-tokenの作成を試みます
	if CheckExpired(&session.Period) {
		return LoginByCookie(ctx, db, w, r)
	}

	return session.UserId.UserId, nil
}

// refresh-token cookieからログインします
// UserIDを返します
func LoginByCookie(ctx context.Context, db *database.Database, w http.ResponseWriter, r *http.Request) (string, error) {
	refreshToken, err := net.GetCookie(r, "refresh-token")
	if err != nil || len(refreshToken) == 0 {
		// cookieが存在しない、valueが存在しない場合は403を返す
		return "", status.NewForbiddenError(errors.New("cookie is not find")).Caller().AddCode(net.FailedLogin)
	}

	newSessionToken := utils.CreateID(0)
	newRefreshToken := utils.CreateID(0)

	var refresh *models.RefreshInfo

	for i := 0; 3 > i; i++ {
		tx, err := database.NewTransaction(ctx, db)
		if err != nil {
			return "", status.NewInternalServerErrorError(err).Caller()
		}

		refresh, err = models.GetRefreshTokenTX(tx, refreshToken)
		if err != nil {
			return "", status.NewInternalServerErrorError(err).Caller()
		}

		// refreshtokenが存在しない場合は、トランザクションをロールバック、該当cookieを削除して403を返す
		if refresh == nil {
			err = net.DeleteCookie(w, r, "refresh-token")
			if err != nil {
				logging.Sugar.Errorf("core/common/login.go line: 121. %s", err.Error())
			}
			// session-tokenがある場合は削除してしまう
			err = net.DeleteCookie(w, r, "session-token")
			if err != nil {
				logging.Sugar.Errorf("core/common/login.go line: 121. %s", err.Error())
			}

			return "", status.NewForbiddenError(errors.New("refresh token is not exist")).Caller()
		}

		// refresh-tokenが有効期限切れの場合は、トランザクションをロールバック、該当cookieを削除して403を返す
		if CheckExpired(&refresh.Period) {
			err = net.DeleteCookie(w, r, "refresh-token")
			if err != nil {
				logging.Sugar.Errorf("core/common/login.go line: 136. %s", err.Error())
			}
			// session-tokenがある場合は削除してしまう
			err = net.DeleteCookie(w, r, "session-token")
			if err != nil {
				logging.Sugar.Errorf("core/common/login.go line: 121. %s", err.Error())
			}

			return "", status.NewForbiddenError(errors.New("Expired")).Caller()
		}

		// session-tokenを削除する（ある場合は）
		err = models.DeleteSessionTokenTX(tx, refresh.SessionToken)
		if err != nil {
			return "", status.NewInternalServerErrorError(err).Caller()
		}

		// refresh-tokenを削除する
		err = models.DeleteRefreshTokenTX(tx, refresh.RefreshToken)
		if err != nil {
			return "", status.NewInternalServerErrorError(err).Caller()
		}

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
			return "", status.NewInternalServerErrorError(err).Caller()
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
			return "", status.NewInternalServerErrorError(err).Caller()
		}

		// 変更をコミットする
		err = tx.Commit()
		// コミットのエラーがErrConcurrentTransactionの場合はトライする
		// それ以外のエラーはthrowする
		if err != nil && err != datastore.ErrConcurrentTransaction {
			return "", status.NewInternalServerErrorError(err).Caller()
		}

		// 正常にコミットできればリトライループから抜ける
		if err == nil {
			break
		}

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
	if config.Defs.DeployMode == "production" {
		secure = true
	}
	// ブラウザ上でcookieを追加できるように、HttpOnlyはfalseにする
	cookie := net.NewCookie(config.Defs.CookieDomain, secure, http.SameSiteDefaultMode, false)

	sessionExp := net.NewSession()
	cookie.Set(w, "session-token", login.SessionToken, sessionExp)

	// リフレッシュトークンの期限は1週間
	refreshExp := net.NewCookieDayExp(7)
	cookie.Set(w, "refresh-token", login.RefreshToken, refreshExp)
}
