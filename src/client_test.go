package src_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

func TestClientHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: client_idを指定するとそのクライアントを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ClientID, clientId)
	})

	t.Run("成功: client_idを指定しないと自分のすべてのクライアントを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterClient(t, ctx, &u, "openid", "profile")
		RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := []src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)
	})

	t.Run("失敗: client_idが存在しない値", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		noExistClientId, err := lib.RandomStr(32)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", noExistClientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: client_idが指定するクライアントが自分のものではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		clientId, _ := RegisterClient(t, ctx, &u2, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})
}

func TestClientCreateHandler(t *testing.T) {
	t.Run("成功: クライアントを新規作成できる", func(t *testing.T) {})

	t.Run("成功: スコープを複数設定して新規作成", func(t *testing.T) {})

	t.Run("成功: 画像を設定して新規作成", func(t *testing.T) {})

	t.Run("失敗: promptの値が不正", func(t *testing.T) {})

	t.Run("失敗: スコープの値が不正", func(t *testing.T) {})

	t.Run("失敗: クライアントの作成上限が超えている", func(t *testing.T) {})
}

func TestClientUpdateHandler(t *testing.T) {
	t.Run("成功: クライアントを更新できる", func(t *testing.T) {})

	t.Run("成功: スコープはすべて置き換わる", func(t *testing.T) {})

	t.Run("成功: シークレットが更新できる", func(t *testing.T) {})

	t.Run("成功: 画像を更新", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})

	t.Run("失敗: promptの値が不正", func(t *testing.T) {})

	t.Run("失敗: スコープの値が不正", func(t *testing.T) {})
}

func TestClientDeleteHandler(t *testing.T) {
	t.Run("成功: クライアントを削除できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})
}

func TestClientDeleteImageHandler(t *testing.T) {
	t.Run("成功: 画像を削除できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})
}

func TestClientAllowUserHandler(t *testing.T) {
	t.Run("成功: ルールを取得できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})
}

func TestClientAddAllowUserHandler(t *testing.T) {
	t.Run("成功: ルールにユーザーIDを指定して追加できる", func(t *testing.T) {})

	t.Run("成功: ルールにメールアドレスのドメインを指定して追加できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})

	t.Run("失敗: user_idとemail_domainどちらも指定しない", func(t *testing.T) {})

	t.Run("失敗: user_idとemail_domainどちらも指定してしまっている", func(t *testing.T) {})
}

func TestClientDeleteAllowUserHandler(t *testing.T) {
	t.Run("成功: ルールからIDを指定して削除できる", func(t *testing.T) {})

	t.Run("失敗: idが不正", func(t *testing.T) {})

	t.Run("失敗: idが空", func(t *testing.T) {})

	t.Run("失敗: そのルールのクライアントのオーナーではない", func(t *testing.T) {})
}

func TestClientLoginUsersHandler(t *testing.T) {
	// TODO
}
