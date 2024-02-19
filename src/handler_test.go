package src_test

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	path, err := os.Getwd()
	require.NoError(t, err)

	newPath := filepath.Join(path, "..")

	h, err := src.NewHandler(DB, C, newPath)
	require.NoError(t, err)

	require.NotNil(t, h.DB)
	require.NotNil(t, h.C)
	require.NotNil(t, h.ReCaptcha)
}

func TestRootHandler(t *testing.T) {
	h := NewTestHandler(t)

	m, err := easy.NewMock("/", http.MethodGet, "")
	require.NoError(t, err)

	c := m.Echo()

	err = h.Root(c)
	require.NoError(t, err)

	require.Equal(t, c.Response().Status, http.StatusOK)
}

func TestDebugHandler(t *testing.T) {
	h := NewTestHandler(t)

	m, err := easy.NewMock("/debug", http.MethodGet, "")
	require.NoError(t, err)

	m.R.Header.Add("X-Forwarded-For", "203.0.113.1")

	c := m.Echo()

	err = h.DebugHandler(c)
	require.NoError(t, err)

	require.Equal(t, c.Response().Status, http.StatusOK)
	response := src.DebugResponse{}
	require.NoError(t, m.Json(&response))
	require.Equal(t, response.Mode, "test")
	require.Equal(t, response.XFF, "203.0.113.1")
	require.Equal(t, response.IPAddress, "203.0.113.1")
}

func TestParseUA(t *testing.T) {
	h := NewTestHandler(t)
	t.Run("UA-CH", func(t *testing.T) {
		r := http.Request{
			Header: http.Header{
				"User-Agent":         {""}, // UAはない
				"Sec-Ch-Ua":          {`"Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110"`},
				"Sec-Ch-Ua-Platform": {`"Windows"`},
				"Sec-Ch-Ua-Mobile":   {"?0"},
			},
		}

		d, err := h.ParseUA(&r)
		require.NoError(t, err)

		require.Equal(t, d.Browser, "Google Chrome")
		require.Equal(t, d.Device, "")
		require.Equal(t, d.OS, "Windows")
		require.False(t, d.IsMobile)
	})

	t.Run("UA", func(t *testing.T) {
		r := http.Request{
			Header: http.Header{
				"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"},
			},
		}

		d, err := h.ParseUA(&r)
		require.NoError(t, err)

		require.Equal(t, d.Browser, "Chrome")
		require.Equal(t, d.Device, "")
		require.Equal(t, d.OS, "Windows")
		require.False(t, d.IsMobile)
	})
}

func TestFormValues(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("成功: required", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			values, err := h.FormValues(c, key)
			if err != nil {
				return err
			}

			require.Len(t, values, 3)

			return nil
		}

		form := easy.NewMultipart()
		form.Insert("key_count", "3")
		form.Insert("key_0", "value_0")
		form.Insert("key_1", "value_1")
		form.Insert("key_2", "value_2")
		form.Insert("key_3", "value_3")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.NoError(t, err)
	})

	t.Run("成功: optional", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			values, err := h.FormValues(c, key, true)
			if err != nil {
				return err
			}

			require.Len(t, values, 0)

			return nil
		}

		form := easy.NewMultipart()

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.NoError(t, err)
	})

	t.Run("失敗: requiredなのに存在しない", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			_, err := h.FormValues(c, key)
			if err != nil {
				return err
			}

			return nil
		}

		form := easy.NewMultipart()

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.EqualError(t, err, "code=400, message=key_count is required")
	})

	t.Run("失敗: _countはあるけど_[index]がすべて存在しない", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			values, err := h.FormValues(c, key)
			if err != nil {
				return err
			}

			require.Len(t, values, 3)

			return nil
		}

		form := easy.NewMultipart()
		form.Insert("key_count", "3")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.EqualError(t, err, "code=400, message=key_0 is required")
	})

	t.Run("失敗: _countはあるけど_[index]が1つ存在しない", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			values, err := h.FormValues(c, key)
			if err != nil {
				return err
			}

			require.Len(t, values, 3)

			return nil
		}

		form := easy.NewMultipart()
		form.Insert("key_count", "3")
		form.Insert("key_0", "value_0")
		form.Insert("key_1", "value_1")
		form.Insert("key_3", "value_3")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.EqualError(t, err, "code=400, message=key_2 is required")
	})

	t.Run("失敗: _[index]はあるけど_countは無い", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			values, err := h.FormValues(c, key)
			if err != nil {
				return err
			}

			require.Len(t, values, 3)

			return nil
		}

		form := easy.NewMultipart()
		form.Insert("key_0", "value_0")
		form.Insert("key_1", "value_1")
		form.Insert("key_2", "value_2")
		form.Insert("key_3", "value_3")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.EqualError(t, err, "code=400, message=key_count is required")
	})

	t.Run("失敗: _[index]のどれかが空", func(t *testing.T) {
		key := "key"

		handler := func(c echo.Context) error {
			values, err := h.FormValues(c, key)
			if err != nil {
				return err
			}

			require.Len(t, values, 3)

			return nil
		}

		form := easy.NewMultipart()
		form.Insert("key_count", "3")
		form.Insert("key_0", "value_0")
		form.Insert("key_1", "value_1")
		form.Insert("key_2", "")
		form.Insert("key_3", "value_3")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		err = handler(c)
		require.EqualError(t, err, "code=400, message=key_2 is required")
	})
}

func TestQueryBodyParam(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("成功: URLのクエリパラメータから取得できる", func(t *testing.T) {
		m, err := easy.NewMock("/?key=value", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		values, err := h.QueryBodyParam(c)
		require.NoError(t, err)

		require.Equal(t, values.Get("key"), "value")
	})

	t.Run("成功: x-www-form-urlencodedのBodyから取得できる", func(t *testing.T) {
		query := url.Values{}
		query.Set("key", "value")

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		values, err := h.QueryBodyParam(c)
		require.NoError(t, err)

		require.Equal(t, values.Get("key"), "value")
	})

	t.Run("成功 X-WWW-Form-Urlencoded のように大文字の場合でも問題なく取得可能", func(t *testing.T) {
		query := url.Values{}
		query.Set("key", "value")

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()
		c.Request().Header.Set("Content-Type", "application/X-WWW-Form-Urlencoded")

		values, err := h.QueryBodyParam(c)
		require.NoError(t, err)

		require.Equal(t, values.Get("key"), "value")
	})

	t.Run("失敗: form-dataからは取得できない", func(t *testing.T) {
		form := easy.NewMultipart()
		form.Insert("key", "value")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		c := m.Echo()

		values, err := h.QueryBodyParam(c)
		require.NoError(t, err)

		require.Empty(t, values)
	})
}

func TestSaveOperationHistory(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("操作履歴が保存されている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.SaveOperationHistory(ctx, c, &u, 1)
		require.NoError(t, err)

		operationHistory, err := models.OperationHistories(
			models.OperationHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, operationHistory.Identifier, int8(1))
	})
}
