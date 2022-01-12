package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestBufferToken(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	bufferToken := utils.CreateID(0)

	buffer := &models.CreateAccountBuffer{
		BufferToken: bufferToken,
		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
		UserMailPW: models.UserMailPW{
			Mail:     "example@example.com",
			Password: []byte("password"),
			Salt:     []byte(""),
		},
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		element, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
		require.NoError(t, err)

		return element != nil
	}, "entryがある")

	element, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
	require.NoError(t, err)

	require.NotNil(t, element)
	require.Equal(t, element.Mail, "example@example.com", "メールアドレスが同じ")

	err = models.DeleteCreateAccountBuffer(ctx, db, bufferToken)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		element, err = models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
		require.NoError(t, err)

		return element == nil
	}, "bufferTokenが消えている")
}
