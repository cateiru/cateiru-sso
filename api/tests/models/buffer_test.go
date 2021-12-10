package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

func TestBufferToken(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	bufferToken := utils.CreateID(0)

	buffer := &models.CreateAccountBuffer{
		BufferToken: bufferToken,
		VerifyPeriod: models.VerifyPeriod{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
		UserMailPW: models.UserMailPW{
			Mail:     "example@example.com",
			Password: "password",
		},
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	element, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
	require.NoError(t, err)

	require.NotNil(t, element)
	require.Equal(t, element.Mail, "example@example.com", "メールアドレスが同じ")

	err = models.DeleteCreateAccountBuffer(ctx, db, bufferToken)
	require.NoError(t, err)

	element, err = models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
	require.NoError(t, err)

	require.Nil(t, element)
}
