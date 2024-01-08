package src_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestUserinfoAuthentication(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: 認証できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)
		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(1 * time.Hour),
		}
		err = clientSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "Bearer "+sessionToken)

		c := m.Echo()

		returnsClientSession, err := h.UserinfoAuthentication(ctx, c)
		require.NoError(t, err)

		require.Equal(t, clientSession.ID, returnsClientSession.ID)
	})

	t.Run("失敗: 形式が違う", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)
		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(1 * time.Hour),
		}
		err = clientSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "hogehoge")

		c := m.Echo()

		_, err = h.UserinfoAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid code")
		require.Equal(t, m.Response().Header.Get("WWW-Authenticate"), "Bearer")
	})

	t.Run("失敗: ヘッダーが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)
		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(1 * time.Hour),
		}
		err = clientSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.UserinfoAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid code")
		require.Equal(t, m.Response().Header.Get("WWW-Authenticate"), "Bearer")
	})

	t.Run("失敗: ヘッダーがBasic", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)
		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(1 * time.Hour),
		}
		err = clientSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "Basic "+sessionToken)

		c := m.Echo()

		_, err = h.UserinfoAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid code")
		require.Equal(t, m.Response().Header.Get("WWW-Authenticate"), "Bearer")
	})

	t.Run("失敗: トークンが存在しない", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "Bearer invalid_token")

		c := m.Echo()

		_, err = h.UserinfoAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid code")
		require.Equal(t, m.Response().Header.Get("WWW-Authenticate"), "Bearer")
	})

	t.Run("失敗: トークンが有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)
		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(-1 * time.Hour), // 有効期限切れ
		}
		err = clientSession.Insert(ctx, h.DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "Bearer "+sessionToken)

		c := m.Echo()

		_, err = h.UserinfoAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid code")
		require.Equal(t, m.Response().Header.Get("WWW-Authenticate"), "Bearer")
	})
}

func TestResponseStandardClaims(t *testing.T) {

}
