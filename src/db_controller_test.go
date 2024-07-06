package src_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestTxDB(t *testing.T) {
	ctx := context.Background()

	t.Run("トランザクションをcommitできる", func(t *testing.T) {
		email := RandomEmail(t)
		id := ulid.Make()

		userName, err := lib.RandomStr(8)
		require.NoError(t, err)

		err = src.TxDB(ctx, DB, func(tx *sql.Tx) error {
			user := models.User{
				ID:       id.String(),
				Email:    email,
				UserName: userName,
			}

			err := user.Insert(ctx, tx, boil.Infer())
			require.NoError(t, err)

			return nil
		})
		require.NoError(t, err)

		user, err := models.Users(models.UserWhere.Email.EQ(email)).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, user)
	})

	t.Run("トランザクションをロールバックできる", func(t *testing.T) {
		email := RandomEmail(t)
		id := ulid.Make()

		userName, err := lib.RandomStr(8)
		require.NoError(t, err)

		err = src.TxDB(ctx, DB, func(tx *sql.Tx) error {
			user := models.User{
				ID:       id.String(),
				Email:    email,
				UserName: userName,
			}

			err := user.Insert(ctx, tx, boil.Infer())
			require.NoError(t, err)

			// ロールバックさせるためにエラー出す
			return errors.New("rollback")
		})
		require.Error(t, err)

		user, err := models.Users(models.UserWhere.Email.EQ(email)).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, user)
	})
}
