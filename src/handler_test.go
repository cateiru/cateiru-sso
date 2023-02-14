package src_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := src.NewHandler(DB, C)

	require.NotNil(t, h.DB)
	require.NotNil(t, h.C)
	require.NotNil(t, h.ReCaptcha)
}

func TestParseUA(t *testing.T) {
	t.Run("UA-CH", func(t *testing.T) {
		r := http.Request{
			Header: http.Header{
				"User-Agent":         {""}, // UAはない
				"Sec-Ch-Ua":          {`"Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110"`},
				"Sec-Ch-Ua-Platform": {`"Windows"`},
				"Sec-Ch-Ua-Mobile":   {"?0"},
			},
		}

		d, err := src.ParseUA(&r)
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

		d, err := src.ParseUA(&r)
		require.NoError(t, err)

		require.Equal(t, d.Browser, "Chrome")
		require.Equal(t, d.Device, "")
		require.Equal(t, d.OS, "Windows")
		require.False(t, d.IsMobile)
	})
}
