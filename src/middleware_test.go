package src_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestBasicAuthMiddleware(t *testing.T) {
	handler := func(c echo.Context) error {
		return nil
	}
	t.Run("ユーザーパスワードが一致する場合は通る", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		// user:password
		m.R.Header.Set("Authorization", "Basic dXNlcjpwYXNzd29yZA==")

		e := m.Echo()

		middleware := src.BasicAuthMiddleware(C)
		err = middleware(handler)(e)
		require.NoError(t, err)
	})

	t.Run("ユーザーパスワードが一致しない場合は401", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		// hoge:dummy
		m.R.Header.Set("Authorization", "Basic aG9nZTpkdW1teQ==")

		e := m.Echo()

		middleware := src.BasicAuthMiddleware(C)
		err = middleware(handler)(e)
		require.EqualError(t, err, "code=401, message=Unauthorized")
	})
}

func TestCSRFMiddleware(t *testing.T) {
	handler := func(c echo.Context) error {
		return nil
	}

	t.Run("Sec-Fetch-Siteが same-origin の場合は通る", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Sec-Fetch-Site", "same-origin")

		e := m.Echo()

		h := src.CSRFMiddleware(handler)
		err = h(e)
		require.NoError(t, err)
	})

	t.Run("Sec-Fetch-Siteが same-site の場合は通る", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Sec-Fetch-Site", "same-site")

		e := m.Echo()

		h := src.CSRFMiddleware(handler)
		err = h(e)
		require.NoError(t, err)
	})

	t.Run("Sec-Fetch-Siteがcross-siteの場合は403", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Sec-Fetch-Site", "cross-site")

		e := m.Echo()

		h := src.CSRFMiddleware(handler)
		err = h(e)
		require.EqualError(t, err, "code=403, message=CSRF Error")
	})
}

func TestFedCMMiddleware(t *testing.T) {
	handler := func(c echo.Context) error {
		return nil
	}

	t.Run("Sec-Fetch-Dest が webidentity 以外の場合エラー", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Sec-Fetch-Dest", "invalid")

		e := m.Echo()

		h := src.FedCMMiddleware(handler)
		err = h(e)
		require.EqualError(t, err, "code=401, message=Unauthorized")
	})

	t.Run("Sec-Fetch-Dest が webidentity の場合通る", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Sec-Fetch-Dest", "webidentity")

		e := m.Echo()

		h := src.FedCMMiddleware(handler)
		err = h(e)
		require.NoError(t, err)
	})
}

func TestCacheMiddleware(t *testing.T) {
	handler := func(c echo.Context) error {
		return nil
	}

	t.Run("CacheTypeNoCache の場合はCache-Control: no-cache", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		e := m.Echo()

		h := src.CacheMiddleware(src.CacheTypeNoCache)
		err = h(handler)(e)
		require.NoError(t, err)

		require.Equal(t, "no-store", e.Response().Header().Get("Cache-Control"))
	})

	t.Run("CacheType15Min の場合はCache-Control: public", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		e := m.Echo()

		h := src.CacheMiddleware(src.CacheType15Min)
		err = h(handler)(e)
		require.NoError(t, err)

		require.Equal(t, "public, s-maxage=900", e.Response().Header().Get("Cache-Control"))
	})

	t.Run("CacheType1Hour の場合はCache-Control: public", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		e := m.Echo()

		h := src.CacheMiddleware(src.CacheType1Hour)
		err = h(handler)(e)
		require.NoError(t, err)

		require.Equal(t, "public, s-maxage=3600", e.Response().Header().Get("Cache-Control"))
	})
}
