package lib_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const TEST_SECRET = "hogehoge123456"

func TestRecaotcha(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://www.google.com/recaptcha/api/siteverify",
		func(req *http.Request) (*http.Response, error) {
			token := req.FormValue("response")
			secret := req.FormValue("secret")
			remoteIp := req.FormValue("remoteip")

			body := new(lib.RecaptchaResponse)
			if TEST_SECRET != secret {
				body = &lib.RecaptchaResponse{
					Success:     false,
					Score:       0,
					Action:      "",
					ChallengeTS: time.Now(),
					Hostname:    "",
					ErrorCodes: []string{
						"error",
					},
				}
			} else {
				if token == "123123123" && remoteIp == "192.168.0.1" {
					body = &lib.RecaptchaResponse{
						Success:     true,
						Score:       100,
						Action:      "",
						ChallengeTS: time.Now(),
						Hostname:    "",
					}
				} else {
					// 判定 X
					body = &lib.RecaptchaResponse{
						Success:     false,
						Score:       0,
						Action:      "",
						ChallengeTS: time.Now(),
						Hostname:    "",
					}

				}
			}
			resp, err := httpmock.NewJsonResponse(200, body)
			if err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			return resp, nil
		})

	t.Run("secretが不正", func(t *testing.T) {
		r := lib.NewReCaptcha("aaaa")
		resp, err := r.ValidateOrder("123123123", "192.168.0.1")
		require.NoError(t, err)

		require.Equal(t, resp.Success, false)
		require.Equal(t, resp.ErrorCodes, []string{"error"})
	})

	t.Run("success", func(t *testing.T) {
		r := lib.NewReCaptcha(TEST_SECRET)
		resp, err := r.ValidateOrder("123123123", "192.168.0.1")
		require.NoError(t, err)

		require.Equal(t, resp.Success, true)
		require.Equal(t, resp.Score, float64(100))
	})

	t.Run("token, ipで失敗", func(t *testing.T) {
		r := lib.NewReCaptcha(TEST_SECRET)
		resp, err := r.ValidateOrder("1231231234", "192.168.0.1")
		require.NoError(t, err)

		require.Equal(t, resp.Success, false)
		require.Equal(t, resp.Score, float64(0))
	})

}
