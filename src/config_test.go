package src_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestInitConfig(t *testing.T) {
	modes := map[string]string{
		"test":     "test",
		"local":    "local",
		"cloudrun": "cloudrun",

		// other
		"hogehoge": "test",
	}

	for mode, configMode := range modes {
		t.Run(mode, func(t *testing.T) {
			c := src.InitConfig(mode)
			require.Equal(t, c.Mode, configMode)
		})
	}
}
