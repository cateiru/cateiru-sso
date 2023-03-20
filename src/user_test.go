package src_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestUserMeHandler(t *testing.T) {

	t.Run("成功: ユーザ情報を取得できる", func(t *testing.T) {})
}

func TestUserUpdateHandler(t *testing.T) {

	t.Run("成功: ユーザ情報を更新できる", func(t *testing.T) {})

	t.Run("失敗: 空", func(t *testing.T) {})
}

func TestUserUpdateSettingHandler(t *testing.T) {

	t.Run("成功: 設定を更新できる", func(t *testing.T) {})

	t.Run("失敗: 空", func(t *testing.T) {})
}

func TestUserBrandHandler(t *testing.T) {
	t.Run("成功: ブランドを取得できる", func(t *testing.T) {
		t.Run("ブランドが指定されている", func(t *testing.T) {})

		t.Run("ブランドは設定されていない", func(t *testing.T) {})
	})
}

func TestUserUpdateEmailHandler(t *testing.T) {
	t.Run("成功: メールアドレスに更新メールが送られる", func(t *testing.T) {})

	t.Run("失敗: メールアドレスが空", func(t *testing.T) {})

	t.Run("失敗: メールアドレスはすでに別のユーザが使用している", func(t *testing.T) {})

	t.Run("失敗: メールアドレスが不正", func(t *testing.T) {})
}

func TestUserUpdateEmailRegisterHandler(t *testing.T) {

	t.Run("成功: メールアドレスを更新できる", func(t *testing.T) {})

	t.Run("失敗: セッションが無い", func(t *testing.T) {})

	t.Run("失敗: セッションが不正", func(t *testing.T) {})

	t.Run("失敗: セッションの有効期限切れ", func(t *testing.T) {})
}

func TestUserAvatarHandler(t *testing.T) {
	t.Run("成功: アバターを新規作成できる", func(t *testing.T) {})

	t.Run("成功: アバターを更新できる", func(t *testing.T) {})

	t.Run("失敗: 画像が指定されていない", func(t *testing.T) {})
}

func TestUserDeleteAvatarHandler(t *testing.T) {
	t.Run("成功: アバターが削除されている", func(t *testing.T) {})
}

// TODO: クライアントの実装してから
func TestUserLogoutClient(t *testing.T) {
	t.Run("成功: 指定したクライアントからログアウトできる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが空", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: そもそもクライアントIDのクライアントにログインしていない", func(t *testing.T) {})
}

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
