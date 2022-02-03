package common

import (
	"context"
	"errors"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type Cert struct {
	Writer  http.ResponseWriter
	Request *http.Request

	SessionToken string
	RefreshToken string

	UserId string

	AccessID string

	Ip        string
	UserAgent string
}

func NewCert(w http.ResponseWriter, r *http.Request) *Cert {
	return &Cert{
		Writer:  w,
		Request: r,
	}
}

func (c *Cert) AddUser() *Cert {
	c.Ip = net.GetIPAddress(c.Request)
	c.UserAgent = net.GetUserAgent(c.Request)

	return c
}

// 新しくログインする
func (c *Cert) NewLogin(ctx context.Context, db *database.Database, userId string) error {
	c.SessionToken = utils.CreateID(0)
	c.RefreshToken = utils.CreateID(0)
	c.AccessID = utils.CreateID(0)
	c.UserId = userId

	if err := c.saveSessionToken(ctx, db); err != nil {
		return err
	}
	if err := c.saveRefreshToken(ctx, db); err != nil {
		return err
	}

	// ログイン履歴を保存する
	if err := c.setLoginHistory(ctx, db); err != nil {
		return err
	}

	return nil
}

func (c *Cert) Login(ctx context.Context, db *database.Database) error {
	sessionToken, err := net.GetCookie(c.Request, "session-token")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	// session-tokenが存在しない場合、refresh-tokenからsession-tokenを作成する
	if sessionToken == "" {
		return c.RefreshLogin(ctx, db)
	}

	session, err := models.GetSessionToken(ctx, db, sessionToken)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// sessionTokenが見つからない場合、refresh-tokenを使用してsession-tokenの作成を試みます
	if session == nil {
		return c.RefreshLogin(ctx, db)
	}

	// sessionTokenの有効期限が切れてしまっている場合、refresh-tokenを使用してsession-tokenの作成を試みます
	if CheckExpired(&session.Period) {
		return c.RefreshLogin(ctx, db)
	}

	c.UserId = session.UserId.UserId
	c.AccessID = session.AccessID

	return nil
}

// refresh tokenからログインする
func (c *Cert) RefreshLogin(ctx context.Context, db *database.Database) error {
	refreshToken, err := net.GetCookie(c.Request, "refresh-token")
	if err != nil || len(refreshToken) == 0 {
		// cookieが存在しない、valueが存在しない場合は403を返す
		return status.NewForbiddenError(errors.New("cookie is not find")).Caller().AddCode(net.FailedLogin)
	}

	newSessionToken := utils.CreateID(0)
	newRefreshToken := utils.CreateID(0)

	var refresh *models.RefreshInfo

	for i := 0; 3 > i; i++ {
		tx, err := database.NewTransaction(ctx, db)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		refresh, err = models.GetRefreshTokenTX(tx, refreshToken)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		// refreshtokenが存在しない場合は、トランザクションをロールバック、該当cookieを削除して403を返す
		if refresh == nil {
			err = c.deleteCookie()
			if err != nil {
				return err
			}

			return status.NewForbiddenError(errors.New("refresh token is not exist")).Caller()
		}

		// refresh-tokenが有効期限切れの場合は、トランザクションをロールバック、該当cookieを削除して403を返す
		if CheckExpired(&refresh.Period) {
			err = c.deleteCookie()
			if err != nil {
				return err
			}

			return status.NewForbiddenError(errors.New("Expired")).Caller()
		}

		// 過去のトークンは削除する
		if err := c.deleteTokenTx(tx, refresh.SessionToken, refresh.RefreshToken); err != nil {
			return err
		}

		// 新しいsession-tokenを作成する
		session := &models.SessionInfo{
			SessionToken: newSessionToken,

			AccessID: refresh.AccessID,

			Period: models.Period{
				CreateDate: time.Now(),
				PeriodHour: 6,
			},

			UserId: refresh.UserId,
		}
		if err := session.AddTX(tx); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		// 新しいrefresh-tokenを作成する
		newRefresh := &models.RefreshInfo{
			RefreshToken: newRefreshToken,
			SessionToken: newSessionToken,

			AccessID: refresh.AccessID,

			Period: models.Period{
				CreateDate: time.Now(),
				PeriodDay:  7,
			},

			UserId: refresh.UserId,
		}
		if err := newRefresh.AddTX(tx); err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}

		// 変更をコミットする
		err = tx.Commit()
		// コミットのエラーがErrConcurrentTransactionの場合はトライする
		// それ以外のエラーはthrowする
		if err != nil && err != datastore.ErrConcurrentTransaction {
			return status.NewInternalServerErrorError(err).Caller()
		}

		c.UserId = refresh.UserId.UserId
		c.AccessID = refresh.AccessID

		// 正常にコミットできればリトライループから抜ける
		if err == nil {
			break
		}

	}

	c.SetCookie()

	return nil
}

// ログイン用のcookieをセットする
func (c *Cert) SetCookie() {
	// secure属性はproductionのみにする（テストが通らないため）
	secure := false
	if config.Defs.DeployMode == "production" {
		secure = true
	}
	// ブラウザ上でcookieを追加できるように、HttpOnlyはfalseにする
	cookie := net.NewCookie(config.Defs.CookieDomain, secure, http.SameSiteDefaultMode, false)

	sessionExp := net.NewSession()
	cookie.Set(c.Writer, "session-token", c.SessionToken, sessionExp)

	// リフレッシュトークンの期限は1週間
	refreshExp := net.NewCookieDayExp(7)
	cookie.Set(c.Writer, "refresh-token", c.RefreshToken, refreshExp)
}

func (c *Cert) saveSessionToken(ctx context.Context, db *database.Database) error {
	session := &models.SessionInfo{
		SessionToken: c.SessionToken,

		AccessID: c.AccessID,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodHour: 6,
		},

		UserId: models.UserId{
			UserId: c.UserId,
		},
	}
	if err := session.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

func (c *Cert) saveRefreshToken(ctx context.Context, db *database.Database) error {
	refresh := &models.RefreshInfo{
		RefreshToken: c.RefreshToken,
		SessionToken: c.SessionToken,

		AccessID: c.AccessID,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  7,
		},

		UserId: models.UserId{
			UserId: c.UserId,
		},
	}
	if err := refresh.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

func (c *Cert) deleteCookie() error {
	if err := net.DeleteCookie(c.Writer, c.Request, "refresh-token"); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// session-tokenがある場合は削除してしまう
	if err := net.DeleteCookie(c.Writer, c.Request, "session-token"); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	return nil
}

func (c *Cert) deleteTokenTx(tx *database.Transaction, sessionToken string, refreshToken string) error {
	// session-tokenを削除する（ある場合は）
	if err := models.DeleteSessionTokenTX(tx, sessionToken); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// refresh-tokenを削除する
	if err := models.DeleteRefreshTokenTX(tx, refreshToken); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

// ログイン履歴を保存する
func (c *Cert) setLoginHistory(ctx context.Context, db *database.Database) error {
	userAgentInfo, err := UserAgentToJson(c.UserAgent)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// ログイン履歴を取る
	history := &models.LoginHistory{
		AccessId:     c.AccessID,
		Date:         time.Now(),
		IpAddress:    c.Ip,
		UserAgent:    string(userAgentInfo),
		IsSSO:        false,
		SSOPublicKey: "",
		UserId: models.UserId{
			UserId: c.UserId,
		},
	}
	if err := history.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
