package createaccount_test

import (
	"context"
	"testing"
	"time"

	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	bufferToken := utils.CreateID(20)

	buffer := models.CreateAccountBuffer{
		BufferToken: bufferToken,
		VerifyPeriod: models.VerifyPeriod{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
		UserMailPW: models.UserMailPW{
			Password: "password",
			Mail:     "example@example.com",
		},
	}
	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	user := createaccount.InfoRequestForm{
		FirstName: "名前",
		LastName:  "名字",
		UserName:  "cateiru",

		Theme:     "dark",
		AvatarUrl: "",
	}

	login, err := createaccount.InsertUserInfo(ctx, bufferToken, user)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

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
}
