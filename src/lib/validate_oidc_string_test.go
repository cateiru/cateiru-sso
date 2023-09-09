package lib_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestValidateScope(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		scopes := []string{
			"openid",
			"profile",
			"email",
		}

		for _, s := range scopes {
			require.True(t, lib.ValidateScope(s), s)
		}
	})

	t.Run("失敗", func(t *testing.T) {
		scopes := []string{
			"address",
			"phone",
			"offline_access",
		}

		for _, s := range scopes {
			require.False(t, lib.ValidateScope(s), s)
		}
	})
}
