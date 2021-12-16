package tools

import (
	"context"
	"fmt"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
)

type DummyUser struct {
	UserID string
	Mail   string
}

func NewDummyUser() *DummyUser {
	userID := utils.CreateID(30)
	mail := fmt.Sprintf("%s@example.com", utils.CreateID(5))

	return &DummyUser{
		UserID: userID,
		Mail:   mail,
	}
}

// ユーザを追加する
// (テスト用)
func (c *DummyUser) AddUserInfo(ctx context.Context, db *database.Database) (*models.User, error) {
	userInfo := &models.User{
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
		UserName:  "TestUserName",
		Theme:     "Dark",
		AvatarUrl: "",

		Role: []string{"user"},

		Mail: c.Mail,

		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := userInfo.Add(ctx, db); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// ユーザの認証情報を追加する
// (テスト用)
func (c *DummyUser) AddUserCert(ctx context.Context, db *database.Database) (*models.Certification, error) {
	certification := &models.Certification{
		AccountCreateDate: time.Now(),

		// アカウント作成後はOTPは設定しない
		// 設定ページから追加する
		OnetimePasswordSecret:  "",
		OnetimePasswordBackups: []string{},

		UserMailPW: models.UserMailPW{
			Mail:     c.Mail,
			Password: "passoword",
		},
		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := certification.Add(ctx, db); err != nil {
		return nil, err
	}

	return certification, nil
}

// session-tokenとrefresh-tokenをセットする
// テスト用
func (c *DummyUser) AddLoginToken(ctx context.Context, db *database.Database, now time.Time) (string, string, error) {
	sessionToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

	session := &models.SessionInfo{
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: now,
			PeriodHour: 6,
		},
		UserId: models.UserId{
			UserId: c.UserID,
		},
	}
	refresh := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: now,
			PeriodDay:  7,
		},
		UserId: models.UserId{
			UserId: c.UserID,
		},
	}

	if err := session.Add(ctx, db); err != nil {
		return "", "", err
	}
	if err := refresh.Add(ctx, db); err != nil {
		return "", "", err
	}

	return sessionToken, refreshToken, nil
}
