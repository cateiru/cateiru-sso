package lib_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestRandomStr(t *testing.T) {
	t.Run("指定した文字数のランダム文字列が生成される", func(t *testing.T) {
		for i := 5; 100 > i; i++ {
			value, err := lib.RandomStr(i)
			require.NoError(t, err)

			require.Len(t, value, i)
		}
	})

	t.Run("10文字100パターン", func(t *testing.T) {
		values := make([]string, 100)

		for i := 0; 100 > i; i++ {
			value, err := lib.RandomStr(10)
			require.NoError(t, err)
			values[i] = value
		}

		for i := 0; 100 > i; i++ {
			for j := i + 1; 100 > j; j++ {
				require.NotEqual(t, values[i], values[j])
			}
		}
	})
}
