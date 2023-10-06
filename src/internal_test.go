package src_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

// TODO
func TestInternalAvatarHandler(t *testing.T) {
	// ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功", func(t *testing.T) {
		m, err := easy.NewMock("/:key/:id", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()
		c.SetParamNames("key", "id")
		c.SetParamValues("test", "image")

		err = h.InternalAvatarHandler(c)
		require.NoError(t, err)

		cacheControl := m.Response().Header.Get("Cache-Control")
		require.Equal(t, "s-maxage=31536000, stale-while-revalidate", cacheControl)

		contentType := m.Response().Header.Get("Content-Type")
		require.Equal(t, "image/png", contentType)
	})

	t.Run("失敗: keyが空", func(t *testing.T) {
		m, err := easy.NewMock("/:key/:id", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()
		c.SetParamNames("key", "id")
		c.SetParamValues("", "image")

		err = h.InternalAvatarHandler(c)
		require.EqualError(t, err, "code=400, message=key and id are required")
	})

	t.Run("失敗: idが空", func(t *testing.T) {
		m, err := easy.NewMock("/:key/:id", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()
		c.SetParamNames("key", "id")
		c.SetParamValues("test", "")

		err = h.InternalAvatarHandler(c)
		require.EqualError(t, err, "code=400, message=key and id are required")
	})
}
