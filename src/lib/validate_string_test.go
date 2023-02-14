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
