// Google reCAPCHAの検証を行う
//
// Based on: https://qiita.com/supertaihei02/items/fb15726fd603de7dcefb
package lib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const REACPTCHA_HOST = "https://www.google.com/recaptcha/api/siteverify"

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

func NewReCaptcha(secret string) *ReCaptcha {
	return &ReCaptcha{
		ServerName: REACPTCHA_HOST,
		Secret:     secret,
	}
}

// reCAPTCHAの検証を行う
func (c *ReCaptcha) ValidateOrder(token string, remoteIp string) (*RecaptchaResponse, error) {
	resp, err := http.PostForm(c.ServerName, url.Values{"secret": {c.Secret}, "remoteip": {remoteIp}, "response": {token}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}
	body := buf.Bytes()

	r := RecaptchaResponse{}
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
