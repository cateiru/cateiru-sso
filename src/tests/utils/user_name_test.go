package utils_test

import (
	"strings"
	"testing"

	"github.com/cateiru/cateiru-sso/src/utils"
	"github.com/stretchr/testify/require"
)

var USER_NAMES = map[string]bool{
	"hoge":    true,
	"cateiru": true,
	"fuga123": true,
	"abc":     true,
	"GAAA":    true,
	"True123": true,
	"nil":     true,

	"a":     false,
	"as":    false,
	"2d":    false,
	"あいうえ":  false,
	"tataあ": false,
}

func TestCheckUserName(t *testing.T) {
	for text, isOk := range USER_NAMES {
		result := utils.CheckUserName(text)
		require.Equal(t, result, isOk, text)
	}
}

func TestFormatUserName(t *testing.T) {
	for text, _ := range USER_NAMES {
		formatted := utils.FormantUserName(text)
		require.Equal(t, formatted, strings.ToLower(text))
	}
}
