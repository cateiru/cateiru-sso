package src

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"

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

	var sender lib.SenderInterface = nil
	if config.SendMail {
		fullpath, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		s, err := lib.NewSender(filepath.Join(fullpath, `templates/*.html`), config.FromDomain, config.MailgunSecret, config.SenderMailAddress)
		if err != nil {
			return nil, err
		}
		sender = s
	} else {
		sender = &SenderMock{}
	}

	webauthn, err := lib.NewWebAuthn(config.WebAuthnConfig)
	if err != nil {
		return nil, err
	}

	session := NewSession(config, db)

	storage := lib.NewCloudStorage(config.StorageBucketName)
	cdn, err := lib.NewCDN(config.FastlyApiToken)
	if err != nil {
		return nil, err
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

// ローカル環境ではメールを送信したくないのでモックする
type SenderMock struct{}

func (s *SenderMock) Send(m *lib.MailBody) (string, string, error) {
	L.Info("send mail",
		zap.String("email_address", m.EmailAddress),
		zap.String("subject", m.Subject),
		zap.Any("data", m.Data),
	)

	return "", "", nil
}
