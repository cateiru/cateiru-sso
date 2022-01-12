package createaccount_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	bufferToken := utils.CreateID(20)

	buffer := models.CreateAccountBuffer{
		BufferToken: bufferToken,
		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
		UserMailPW: models.UserMailPW{
			Password: []byte("password"),
			Mail:     "example@example.com",
			Salt:     []byte(""),
		},
	}
	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	// メール認証がDBに格納されるまで待機
	goretry.Retry(t, func() bool {
		entry, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	user := createaccount.InfoRequestForm{
		FirstName: "名前",
		LastName:  "名字",
		UserName:  "cateiru",

		Theme:     "dark",
		AvatarUrl: "",
	}

	ip := "198.51.100.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	login, err := createaccount.InsertUserInfo(ctx, bufferToken, user, ip, userAgent)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		session, err := models.GetSessionToken(ctx, db, login.SessionToken)
		require.NoError(t, err)
		return session != nil
	}, "sessionTokenがある")

	session, err := models.GetSessionToken(ctx, db, login.SessionToken)
	require.NoError(t, err)
	require.NotNil(t, session)

	refresh, err := models.GetRefreshToken(ctx, db, login.RefreshToken)
	require.NoError(t, err)
	require.NotNil(t, refresh)

	userInfo, err := models.GetUserDataByUserID(ctx, db, session.UserId.UserId)
	require.NoError(t, err)
	require.Equal(t, userInfo.Mail, buffer.Mail, "メールアドレスが同じ")

	entryBuffer, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
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
