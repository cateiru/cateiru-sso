package src_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		u, err := src.RegisterUser(ctx, DB, email)
		require.NoError(t, err)

		require.Equal(t, u.Email, email)
		require.Len(t, u.UserName, 8)
	})

	t.Run("すでにEmailが存在している場合はエラー", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)

		RegisterUser(t, ctx, email)

		_, err := src.RegisterUser(ctx, DB, email)
		require.EqualError(t, err, "code=400, message=impossible register account, unique=3")
	})
}

func TestFindUserByUserNameOrEmail(t *testing.T) {
	ctx := context.Background()

	t.Run("成功: ユーザー名", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		user, err := src.FindUserByUserNameOrEmail(ctx, DB, u.UserName)
		require.NoError(t, err)

		require.Equal(t, user.ID, u.ID)
	})

	t.Run("成功: Email", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		user, err := src.FindUserByUserNameOrEmail(ctx, DB, u.Email)
		require.NoError(t, err)

		require.Equal(t, user.ID, u.ID)
	})

	t.Run("失敗", func(t *testing.T) {
		_, err := src.FindUserByUserNameOrEmail(ctx, DB, "aaaaaa")
		require.Error(t, err)
	})
}
