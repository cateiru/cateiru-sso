// Google reCAPCHAの検証を行う
//
// Based on: https://qiita.com/supertaihei02/items/fb15726fd603de7dcefb
package secure

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/go-http-error/httperror/status"
)

type ReCaptcha struct {
	ServerName string
	Secret     string
}

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func NewReCaptcha() *ReCaptcha {
	return &ReCaptcha{
		ServerName: "https://www.google.com/recaptcha/api/siteverify",
		Secret:     config.Defs.ReCaptchaSecret,
	}
}

// reCAPTCHAの検証を行い、結果をboolで返す
// 0.5を閾値とする
func (c *ReCaptcha) Validate(token string, remoteIp string) (bool, error) {
	result, err := c.ValidateOrder(token, remoteIp)
	if err != nil {
		return false, status.NewInternalServerErrorError(err).Caller()
	}
	if !result.Success {
		return false, status.NewBadRequestError(errors.New("reCAPTCHA token is failed")).Caller()
	}

	return result.Score > 0.5, nil
}

// reCAPTCHAの検証を行う
func (c *ReCaptcha) ValidateOrder(token string, remoteIp string) (*RecaptchaResponse, error) {
	resp, err := http.PostForm(c.ServerName, url.Values{"secret": {c.Secret}, "remoteip": {remoteIp}, "response": {token}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := RecaptchaResponse{}
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	logging.Sugar.Debugf(
		"reCAPTCHA validate. Token: %s Status: %v, Score: %v, Action: %s, ErrCode: %v",
		token, r.Success, r.Score, r.Action, r.ErrorCodes)

	return &r, nil
}
