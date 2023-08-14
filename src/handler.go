package src

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cateiru/cateiru-sso/src/lib"
	goclienthints "github.com/cateiru/go-client-hints/v2"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"
	"go.uber.org/zap"
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
	ReCaptcha lib.ReCaptchaInterface
	Sender    lib.SenderInterface
	WebAuthn  lib.WebAuthnInterface
	Session   SessionInterface
	Password  lib.PasswordInterface
	Storage   lib.CloudStorageInterface
	CDN       lib.CDNInterface
}

func NewHandler(db *sql.DB, config *Config) (*Handler, error) {
	reCaptcha := lib.NewReCaptcha(config.ReCaptchaSecret)

	fullpath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	s, err := lib.NewSender(filepath.Join(fullpath, `templates/*`), config.FromDomain, config.MailgunSecret, config.SenderMailAddress)
	if err != nil {
		return nil, err
	}

	var sender lib.SenderInterface = nil
	if config.SendMail {
		sender = s
	} else {
		sender = &SenderMock{
			Sender: s,
		}
	}

	webauthn, err := lib.NewWebAuthn(config.WebAuthnConfig)
	if err != nil {
		return nil, err
	}

	session := NewSession(config, db)

	storage := lib.NewCloudStorage(config.StorageBucketName)

	var cdn lib.CDNInterface = nil
	if config.UseCDN {
		_cdn, err := lib.NewCDN(config.FastlyApiToken)
		if err != nil {
			return nil, err
		}
		cdn = _cdn
	} else {
		cdn = &CDNMock{}
	}

	return &Handler{
		DB:        db,
		C:         config,
		ReCaptcha: reCaptcha,
		Sender:    sender,
		WebAuthn:  webauthn,
		Session:   session,
		Password:  config.Password,
		Storage:   storage,
		CDN:       cdn,
	}, nil
}

func (h *Handler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mode": h.C.Mode,
	})
}

// User-Agent または UA-CH からユーザ情報を取得する
// 最初、UA-CHの取得を試みる。もし、ブラウザが対応していない場合は
// User-Agentからユーザ情報を取得する
func (h *Handler) ParseUA(r *http.Request) (*UserData, error) {
	if goclienthints.IsSupportClientHints(&r.Header) {
		ch, err := goclienthints.Parse(&r.Header)
		if err != nil {
			return nil, NewHTTPError(http.StatusBadRequest, err)
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

// 複数のフォームを取得する
// `[key]_count`を取得し、その数だけ`[key]_0`から`[key]_[count-1]`までの値を取得する
func (h *Handler) FormValues(c echo.Context, key string, optional ...bool) ([]string, error) {
	optionalFlag := false
	if len(optional) > 0 {
		optionalFlag = optional[0]
	}

	keyName := fmt.Sprintf("%s_count", key)
	count := c.FormValue(keyName)
	if count == "" {
		if optionalFlag {
			return []string{}, nil
		}
		return nil, NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s is required", keyName))
	}
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return nil, NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s is invalid", keyName))
	}

	values := make([]string, countInt)
	for i := 0; i < countInt; i++ {
		itemKey := fmt.Sprintf("%s_%d", key, i)
		v := c.FormValue(itemKey)
		if v == "" {
			return nil, NewHTTPError(http.StatusBadRequest, itemKey+" is required")
		}
		values[i] = v
	}

	return values, nil
}

// ローカル環境ではメールを送信したくないのでモックする
type SenderMock struct {
	Sender *lib.Sender
}

func (s *SenderMock) Send(m *lib.MailBody) (string, string, error) {
	L.Debug("send mail",
		zap.String("email_address", m.EmailAddress),
		zap.String("subject", m.Subject),
		zap.Any("data", m.Data),
	)

	return "", "", nil
}

func (s *SenderMock) Preview(m *lib.MailBody) (string, error) {
	return s.Sender.Preview(m)
}

type CDNMock struct{}

func (c *CDNMock) Purge(url string) error {
	L.Debug("purge cdn",
		zap.String("url", url),
	)

	return nil
}
func (c *CDNMock) SoftPurge(url string) error {
	L.Debug("soft purge cdn",
		zap.String("url", url),
	)

	return nil
}
