package lib_test

import (
	"testing"
	"time"

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

func TestValidateOTPCode(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		otp := []string{
			"123424",
			"545456",
			"013853",
		}

		for _, o := range otp {
			require.True(t, lib.ValidateOTPCode(o))
		}
	})

	t.Run("失敗", func(t *testing.T) {
		values := map[string]string{
			"12453":   "5文字",
			"1234567": "6文字以上",
			"":        "空",
			"asdovk":  "アルファベット",
		}

		for _, v := range values {
			require.False(t, lib.ValidateOTPCode(v))
		}
	})
}

func TestValidateOS(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		values := map[string]string{
			"Windows": "Windows",
			"macOS":   "MacOS",
			"iOS":     "macOS",
			"iPadOS":  "macOS",
		}

		for os, currentOS := range values {
			require.True(t, lib.ValidateOS(os, currentOS))
		}
	})

	t.Run("失敗", func(t *testing.T) {
		values := map[string]string{
			"Windows": "",
			"macOS":   "Windows",
		}

		for os, currentOS := range values {
			require.False(t, lib.ValidateOS(os, currentOS))
		}
	})
}

func TestValidateGender(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		genders := []string{
			"0", "1", "2", "9",
		}

		for _, g := range genders {
			require.True(t, lib.ValidateGender(g))
		}
	})

	t.Run("失敗", func(t *testing.T) {
		genders := []string{
			"男性",
			"Man",
			"5",
		}

		for _, g := range genders {
			require.False(t, lib.ValidateGender(g))
		}
	})
}

func TestValidateBirthDate(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		birthDates := []string{
			"2000-10-01",
			"1999-03-10",
		}

		for _, birthDate := range birthDates {
			b, ok := lib.ValidateBirthDate(birthDate)
			require.True(t, ok)
			require.NotNil(t, b)
		}
	})

	t.Run("失敗: パース不可", func(t *testing.T) {
		birthDates := []string{
			"2000-01",
			"aaaaaaaa",
		}

		for _, birthDate := range birthDates {
			b, ok := lib.ValidateBirthDate(birthDate)
			require.False(t, ok)
			require.Nil(t, b)
		}
	})

	t.Run("失敗: 現在時刻より新しい", func(t *testing.T) {
		d := time.Now().Add(10 * 24 * time.Hour)
		birthDate := d.Format("2006-01-02")

		b, ok := lib.ValidateBirthDate(birthDate)
		require.False(t, ok)
		require.Nil(t, b)
	})
}

func TestValidateLocale(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		locales := []string{
			"ja-JP",
			"en-US",
			"sr-Cyrl-BA",
		}

		for _, l := range locales {
			require.True(t, lib.ValidateLocale(l), l)
		}
	})

	t.Run("存在しないロケールID", func(t *testing.T) {
		locales := []string{
			"xx-YY-ZZ", "ab-CD-EF", "mn-OP-QR", "ij-KL-MN", "uv-WX-YZ", "rs-TU-VW", "gh-JK-LM", "pq-RS-TU", "cd-EF-GH", "op-QR-ST",
		}

		for _, l := range locales {
			require.False(t, lib.ValidateLocale(l), l)
		}
	})
}

func TestValidateContentType(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		contentTypes := []string{
			"image/gif",
			"image/jpeg",
			"image/png",
			"image/webp",
		}

		for _, c := range contentTypes {
			require.True(t, lib.ValidateContentType(c), c)
		}
	})

	t.Run("失敗", func(t *testing.T) {
		contentTypes := []string{
			"image/x-jng",
			"image/bmp",
			"image/tiff",
			"image/svg+xml",
		}

		for _, c := range contentTypes {
			require.False(t, lib.ValidateContentType(c), c)
		}
	})
}
