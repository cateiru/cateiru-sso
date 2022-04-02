package common

import (
	"encoding/json"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/utils/net"
	ua "github.com/mileusna/useragent"
)

type UserAgent struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
	Device    string `json:"device"`
	Mobile    bool   `json:"mobile"`
	Tablet    bool   `json:"access_id"`
	Desktop   bool   `json:"desktop"`
	Bot       bool   `json:"bot"`
	URL       string `json:"url"`
	String    string `json:"string"`
}

func ParseUserData(r *http.Request) ([]byte, error) {
	if len(r.Header.Get("Sec-Ch-Ua")) != 0 {
		return UACHToJson(r)
	} else {
		userAgent := net.GetUserAgent(r)
		return UserAgentToJson(userAgent)
	}
}

func UACHToJson(r *http.Request) ([]byte, error) {
	ch := r.Header.Get("Sec-Ch-Ua")
	mobile := r.Header.Get("Sec-Ch-Ua-Mobile")
	device := r.Header.Get("Sec-Ch-Ua-Platform")

	isMobile := false
	isDeskTop := false

	if mobile == "?1" {
		isMobile = true
	} else if mobile == "?0" {
		isDeskTop = true
	}

	if len(device) == 0 {
		device = "Unknown"
	} else if device[0] == '"' {
		device = device[1 : len(device)-1]
	}

	converted := &UserAgent{
		Device:  device,
		Mobile:  isMobile,
		Desktop: isDeskTop,
		String:  ch,
	}

	return json.Marshal(converted)
}

// userAgentを解析し、json形式で返します
func UserAgentToJson(userAgent string) ([]byte, error) {
	parsed := ua.Parse(userAgent)

	converted := &UserAgent{
		Name:      parsed.Name,
		Version:   parsed.Version,
		OS:        parsed.OS,
		OSVersion: parsed.OSVersion,
		Device:    parsed.Device,
		Mobile:    parsed.Mobile,
		Tablet:    parsed.Tablet,
		Desktop:   parsed.Desktop,
		Bot:       parsed.Bot,
		URL:       parsed.URL,
		String:    parsed.String,
	}

	return json.Marshal(converted)
}

// UserAgentToJsonのjsonを構造体に戻します
func ParseUserAgentJson(target []byte) (*UserAgent, error) {
	var userAgent UserAgent

	if err := json.Unmarshal(target, &userAgent); err != nil {
		return nil, err
	}

	return &userAgent, nil
}
