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

func TestValidateResponseType(t *testing.T) {
	t.Run("成功: Authorization Code Flow", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseType("code"), lib.ResponseTypeAuthorizationCode)
	})

	t.Run("成功: Implicit Flow", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseType("id_token"), lib.ResponseTypeImplicit)
		require.Equal(t, lib.ValidateResponseType("id_token token"), lib.ResponseTypeImplicit)
	})

	t.Run("成功: Hybrid Flow", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseType("code id_token"), lib.ResponseTypeHybrid)
		require.Equal(t, lib.ValidateResponseType("code token"), lib.ResponseTypeHybrid)
		require.Equal(t, lib.ValidateResponseType("code id_token token"), lib.ResponseTypeHybrid)
	})

	t.Run("失敗: 値が空", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseType(""), lib.ResponseTypeInvalid)
	})

	t.Run("失敗: 値が不正", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseType("aaaa"), lib.ResponseTypeInvalid)
	})
}

func TestValidateResponseMode(t *testing.T) {
	t.Run("成功: query", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseMode("query"), lib.ResponseModeQuery)
	})

	t.Run("成功: fragment", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseMode("fragment"), lib.ResponseModeFragment)
	})

	t.Run("成功: form_post", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseMode("form_post"), lib.ResponseModeFormPost)
	})

	t.Run("値が空の場合はqueryになる", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseMode(""), lib.ResponseModeQuery)
	})

	t.Run("値が不正の場合はqueryになる", func(t *testing.T) {
		require.Equal(t, lib.ValidateResponseMode("aaaa"), lib.ResponseModeQuery)
	})
}

func TestValidateDisplay(t *testing.T) {
	require.Equal(t, lib.ValidateDisplay("aaa"), lib.DisplayPage, "決め打ちでpageになる")
}

func TestValidatePrompt(t *testing.T) {
	t.Run("成功: none", func(t *testing.T) {
		require.Equal(t, lib.ValidatePrompt("none"), lib.PromptNone)
	})

	t.Run("成功: login", func(t *testing.T) {
		require.Equal(t, lib.ValidatePrompt("login"), lib.PromptLogin)
	})

	t.Run("成功: consent", func(t *testing.T) {
		require.Equal(t, lib.ValidatePrompt("consent"), lib.PromptConsent)
	})

	t.Run("成功: select_account", func(t *testing.T) {
		require.Equal(t, lib.ValidatePrompt("select_account"), lib.PromptSelectAccount)
	})

	t.Run("値が空の場合はconsentになる", func(t *testing.T) {
		require.Equal(t, lib.ValidatePrompt(""), lib.PromptConsent)
	})

	t.Run("値が不正の場合はconsentになる", func(t *testing.T) {
		require.Equal(t, lib.ValidatePrompt("aaaa"), lib.PromptConsent)
	})
}

func TestValidateMaxAge(t *testing.T) {

	t.Run("成功: max-ageを指定している", func(t *testing.T) {
		require.Equal(t, lib.ValidateMaxAge("100"), uint64(100))
	})

	t.Run("空の場合は0になる", func(t *testing.T) {
		require.Equal(t, lib.ValidateMaxAge(""), uint64(0))
	})

	t.Run("負の値の場合は0になる", func(t *testing.T) {
		require.Equal(t, lib.ValidateMaxAge("-1"), uint64(0))
	})

}

func TestValidateUiLocales(t *testing.T) {
	require.Equal(t, lib.ValidateUiLocales("hoge"), []string{"ja_JP"}, "決め打ちでja_JPになる")
}
