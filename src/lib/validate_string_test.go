package lib_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestValidateEmail(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		emails := []string{
			"test@example.com",
			"hoge@test.test",
			"123aaa@cateiru.test",
			"aaa@aaa.bbb.test",
		}

		for _, e := range emails {
			require.True(t, lib.ValidateEmail(e), e)
		}
	})

	t.Run("失敗", func(t *testing.T) {
		emails := []string{
			"aaaa",
			"123123",
			"",
		}

		for _, e := range emails {
			require.False(t, lib.ValidateEmail(e))
		}
	})
}

func TestValidatePassword(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		passwords := []string{
			`ePA<pi>glgQa?E_`,
			`_8ph.ND12D(\lc2`,
			`$]4.!<a;LPO'pQ3`,
			`L)V$rQ-BxKo7i#x`,
			`;%*aoK[{J$M0Xmv`,
			"wFOHS5io2B3d3dw",
			"TXP6qXb4ERWfKVL",
			"lcc3ln5P0i3jyYI",
			"kb6mQeJHIIndv40",
			"raxTPN2fhTFhudc",
		}

		for _, p := range passwords {
			require.True(t, lib.ValidatePassword(p))
		}
	})

	t.Run("失敗", func(t *testing.T) {
		passwords := map[string]string{
			"aaaaaaaaaaaaaaa": "繰り返しの文字",
			"abc123;":         "13文字以下",
			"日本語ああああああああああああああああああああ": "ascii以外",
		}

		for p, message := range passwords {
			require.False(t, lib.ValidatePassword(p), message)
		}
	})
}

func TestValidateUsername(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		names := []string{
			"aaaaaa",
			"aa123",
			"qawsedrftgyhuji",
			"cateiru",
			"aaa_123",
		}

		for _, n := range names {
			require.True(t, lib.ValidateUsername(n), n)
		}
	})

	t.Run("失敗", func(t *testing.T) {
		names := []string{
			"as",
			"",
			"a",
			"cateiru--",
			"qawsedrftgyhujik",
		}

		for _, n := range names {
			require.False(t, lib.ValidateUsername(n), n)
		}
	})
}
