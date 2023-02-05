package createaccount_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/common"
	createaccount "github.com/cateiru/cateiru-sso/src/core/create_account"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	clientToken := utils.CreateID(20)

	buffer := models.MailCertification{
		MailToken:      utils.CreateID(20),
		ClientToken:    clientToken,
		OpenNewWindow:  false,
		Verify:         true,
		ChangeMailMode: false,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},

		Mail: dummy.Mail,
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	// メール認証がDBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	user := createaccount.InfoRequestForm{
		ClientToken: clientToken,

		FirstName: "名前",
		LastName:  "名字",
		UserName:  utils.CreateID(10),

		Theme:     "dark",
		AvatarUrl: "",

		Password: "password",
	}

	ip := "198.51.100.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	c := &common.Cert{
		Ip: ip,
		Request: &http.Request{
			Header: http.Header{
				"User-Agent": {userAgent},
			},
		},
	}

	userInfo, err := createaccount.InsertUserInfo(ctx, user, c)
	require.NoError(t, err)

	require.Equal(t, userInfo.FirstName, "名前")

	goretry.Retry(t, func() bool {
		session, err := models.GetSessionToken(ctx, db, c.SessionToken)
		require.NoError(t, err)
		return session != nil
	}, "sessionTokenがある")

	session, err := models.GetSessionToken(ctx, db, c.SessionToken)
	require.NoError(t, err)
	require.NotNil(t, session)

	refresh, err := models.GetRefreshToken(ctx, db, c.RefreshToken)
	require.NoError(t, err)
	require.NotNil(t, refresh)

	userInfo, err = models.GetUserDataByUserID(ctx, db, session.UserId.UserId)
	require.NoError(t, err)
	require.Equal(t, userInfo.Mail, buffer.Mail, "メールアドレスが同じ")

	entryBuffer, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
	require.NoError(t, err)
	require.Nil(t, entryBuffer, "bufferは削除されているためnilである")

	goretry.Retry(t, func() bool {
		histories, err := models.GetAllLoginHistory(ctx, db, session.UserId.UserId)
		require.NoError(t, err)
		return len(histories) == 1
	}, "")
	histories, err := models.GetAllLoginHistory(ctx, db, session.UserId.UserId)
	require.NoError(t, err)
	require.Equal(t, len(histories), 1)
	require.Equal(t, histories[0].IpAddress, ip)
}

func TestInfoUnauthenticated(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	clientToken := utils.CreateID(20)

	buffer := models.MailCertification{
		MailToken:      utils.CreateID(20),
		ClientToken:    clientToken,
		OpenNewWindow:  false,
		Verify:         false, // メールアドレス未認証にする
		ChangeMailMode: false,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},

		Mail: dummy.Mail,
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	// メール認証がDBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	user := createaccount.InfoRequestForm{
		FirstName: "名前",
		LastName:  "名字",
		UserName:  "cateiru",

		Theme:     "dark",
		AvatarUrl: "",

		Password: "password",
	}

	ip := "198.51.100.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	c := &common.Cert{
		Ip: ip,
		Request: &http.Request{
			Header: http.Header{
				"User-Agent": {userAgent},
			},
		},
	}

	_, err = createaccount.InsertUserInfo(ctx, user, c)
	require.Error(t, err)
}

func TestFailedUserName(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	clientToken := utils.CreateID(20)

	buffer := models.MailCertification{
		MailToken:      utils.CreateID(20),
		ClientToken:    clientToken,
		OpenNewWindow:  false,
		Verify:         true,
		ChangeMailMode: false,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},

		Mail: dummy.Mail,
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	// メール認証がDBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	user := createaccount.InfoRequestForm{
		FirstName: "名前",
		LastName:  "名字",
		UserName:  "あいうえ",

		Theme:     "dark",
		AvatarUrl: "",

		Password: "password",
	}

	ip := "198.51.100.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	c := &common.Cert{
		Ip: ip,
		Request: &http.Request{
			Header: http.Header{
				"User-Agent": {userAgent},
			},
		},
	}

	_, err = createaccount.InsertUserInfo(ctx, user, c)
	require.Error(t, err)
}
