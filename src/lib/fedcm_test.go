package lib_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestCheckFedCMHeaders(t *testing.T) {
	t.Run("成功", func(t *testing.T) {
		headers := http.Header{
			"Sec-Fetch-Dest": []string{"webidentity"},
		}

		require.True(t, lib.CheckFedCMHeaders(headers))
	})

	t.Run("失敗", func(t *testing.T) {
		headers := http.Header{
			"Sec-Fetch-Dest": []string{"document"},
		}

		require.False(t, lib.CheckFedCMHeaders(headers))
	})
}
