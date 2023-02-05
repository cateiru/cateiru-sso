package pro_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/pro"
	"github.com/stretchr/testify/require"
)

func TestCheckURL(t *testing.T) {
	urls := map[string]bool{
		"https://cateiru.com":              true,
		"https://uie.jp":                   true,
		"https://www.example.com":          true,
		"https://cateiru.com/login":        true,
		"https://cateiru.com/":             true,
		"http://localhost":                 true,
		"http://192.168.3.10":              true, // IPで検証したい話題がある
		"http://192.168.3.10/hogehoge/aaa": true,
		"http://192.168.3.10:3000/aaa":     true,
		"http://0.0.0.0":                   true,
		"https://8.8.8.8":                  true,
		"http://0.0.0.":                    false,
		"http://example.com":               false,
		"https://":                         false,
		"ftp:":                             false,
		"aaaaaa":                           false,
		"direct":                           false,
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
	urls := map[string]bool{
		"https://cateiru.com":              true,
		"https://uie.jp":                   true,
		"https://www.example.com":          true,
		"https://cateiru.com/login":        true,
		"https://cateiru.com/":             true,
		"http://localhost":                 true,
		"http://192.168.3.10":              true, // IPで検証したい話題がある
		"http://192.168.3.10/hogehoge/aaa": true,
		"http://192.168.3.10:3000/aaa":     true,
		"http://0.0.0.0":                   true,
		"https://8.8.8.8":                  true,
		"http://0.0.0.":                    false,
		"http://example.com":               false,
		"https://":                         false,
		"ftp:":                             false,
		"aaaaaa":                           false,
		"direct":                           true,
	}

	for url, result := range urls {
		if result {
			require.NoError(t, pro.CheckURL([]string{url}, true), url)
		} else {
			require.Error(t, pro.CheckURL([]string{url}, true), url)
		}
	}
}
