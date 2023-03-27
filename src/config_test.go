package src_test

import (
	"fmt"
	"reflect"
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

func TestConfig(t *testing.T) {
	configs := []*src.Config{
		src.LocalConfig,
		src.TestConfig,
		src.CloudRunConfig,
	}

	// テストを除外する項目
	// os.Getenvなどを使用しているやつ
	// cloudrunのみ
	ignore := []string{
		"ReCaptchaSecret",
		"MailgunSecret",
		"FastlyApiToken",
	}

	for _, c := range configs {
		t.Run(c.Mode, func(t *testing.T) {

			rv := reflect.ValueOf(*c)
			rt := rv.Type()
			for i := 0; i < rt.NumField(); i++ {
				field := rt.Field(i)
				kind := field.Type.Kind()
				value := rv.FieldByName(field.Name)

				if c.Mode == "cloudrun" {
					skip := false
					for _, i := range ignore {
						if i == field.Name {
							skip = true
							break
						}
					}
					if skip {
						continue
					}
				}

				switch kind.String() {
				case "string":
					// フィールドがstringの場合、空文字列じゃないことを確認する
					require.NotEqual(t, value.String(), "", fmt.Sprintf("mode: %s, field: %s", c.Mode, field.Name))
				case "ptr":
					require.False(t, value.IsNil(), fmt.Sprintf("mode: %s, field: %s", c.Mode, field.Name))
				case "struct":
					require.NotNil(t, value, fmt.Sprintf("mode: %s, field: %s", c.Mode, field.Name))
				}
			}
		})
	}
}
