package pro_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/pro"
	"github.com/stretchr/testify/require"
)

func TestCheckURL(t *testing.T) {
	urls := map[string]bool{
		"https://cateiru.com":       true,
		"https://uie.jp":            true,
		"https://www.example.com":   true,
		"https://cateiru.com/login": true,
		"https://cateiru.com/":      true,
		"localhost":                 true,
		"http://example.com":        false,
		"https://":                  false,
		"ftp:":                      false,
		"aaaaaa":                    false,
		"direct":                    false,
	}

	for url, result := range urls {
		if result {
			require.NoError(t, pro.CheckURL([]string{url}, false), url)
		} else {
			require.Error(t, pro.CheckURL([]string{url}, false), url)
		}
	}
}

func TestCheckURLAllowDirect(t *testing.T) {
	require.NoError(t, pro.CheckURL([]string{"direct"}, true))
}
