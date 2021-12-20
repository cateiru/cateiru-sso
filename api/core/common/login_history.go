package common

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/go-http-error/httperror/status"
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

// ログイン履歴をセットします
func SetLoginHistory(ctx context.Context, db *database.Database, userId string, ip string, userAgent string) error {
	userAgentInfo, err := UserAgentToJson(userAgent)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// ログイン履歴を取る
	history := &models.LoginHistory{
		AccessId:     utils.CreateID(0),
		Date:         time.Now(),
		IpAddress:    ip,
		UserAgent:    string(userAgentInfo),
		IsSSO:        false,
		SSOPublicKey: "",
		UserId: models.UserId{
			UserId: userId,
		},
	}
	if err := history.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
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
