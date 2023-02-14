package src

import (
	"database/sql"
	"net/http"

	"github.com/cateiru/cateiru-sso/src/lib"
	goclienthints "github.com/cateiru/go-client-hints/v2"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"
)

type UserData struct {
	Browser  string
	OS       string
	Device   string
	IsMobile bool
}

type Handler struct {
	DB        *sql.DB
	C         *Config
	ReCaptcha *lib.ReCaptcha
}

func NewHandler(db *sql.DB, config *Config) *Handler {
	reCaptcha := lib.NewReCaptcha(config.ReCaptchaSecret)

	return &Handler{
		DB:        db,
		C:         config,
		ReCaptcha: reCaptcha,
	}
}

func (h *Handler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mode": h.C.Mode,
	})
}

// User-Agent または UA-CH からユーザ情報を取得する
// 最初、UA-CHの取得を試みる。もし、ブラウザが対応していない場合は
// User-Agentからユーザ情報を取得する
func ParseUA(r *http.Request) (*UserData, error) {
	if goclienthints.IsSupportClientHints(&r.Header) {
		ch, err := goclienthints.Parse(&r.Header)
		if err != nil {
			return nil, lib.NewHTTPError(http.StatusBadRequest, err)
		}

		return &UserData{
			Browser:  ch.Brand.Brand,
			OS:       string(ch.Platform),
			Device:   "",
			IsMobile: ch.IsMobile,
		}, nil
	}

	ua := useragent.Parse(r.UserAgent())

	return &UserData{
		Browser:  ua.Name,
		OS:       ua.OS,
		Device:   ua.Device,
		IsMobile: ua.Mobile,
	}, nil
}
