package src

import (
	"net/url"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Mode string

	// オレオレ証明書設定
	SelfSignedCert bool

	// reCAPTCHAを使用するかどうか
	UseReCaptcha        bool
	ReCaptchaSecret     string
	ReCaptchaAllowScore float64

	// MySQLの設定
	DatabaseConfig *mysql.Config

	// APIのホスト
	Host *url.URL
	// サイトのホスト
	SiteHost *url.URL

	// 送信メール設定
	FromDomain        string
	MailgunSecret     string
	SenderMailAddress string

	// セッションの有効期限
	// アカウント登録時に使用するセッションの有効期限
	RegisterSessionPeriod time.Duration
}

var LocalConfig = &Config{
	Mode: "local",

	SelfSignedCert: true,
	// ローカル環境はreCAPTCHA使わない
	UseReCaptcha:        false,
	ReCaptchaSecret:     "",
	ReCaptchaAllowScore: 0,

	DatabaseConfig: &mysql.Config{
		DBName:               "cateiru-sso",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	Host: &url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	},
	SiteHost: &url.URL{
		Host:   "localhost:3000",
		Scheme: "http",
	},

	FromDomain:        "",
	MailgunSecret:     "",
	SenderMailAddress: "",

	RegisterSessionPeriod: 10 * time.Minute,
}

var CloudRunConfig = &Config{
	Mode: "cloudrun",

	SelfSignedCert:      false,
	UseReCaptcha:        true,
	ReCaptchaSecret:     "", // TODO
	ReCaptchaAllowScore: 50,

	DatabaseConfig: &mysql.Config{},

	Host: &url.URL{
		Host:   "api.sso.cateiru.com",
		Scheme: "https",
	},
	SiteHost: &url.URL{
		Host:   "sso.cateiru.com",
		Scheme: "https",
	},

	FromDomain:        "m.cateiru.com",
	MailgunSecret:     os.Getenv("MAILGUN_SECRET"),
	SenderMailAddress: "CateiruSSO <sso@m.cateiru.com>",

	RegisterSessionPeriod: 10 * time.Minute,
}

var TestConfig = &Config{
	Mode: "test",

	SelfSignedCert:      false,
	UseReCaptcha:        true, // mockするので問題なし
	ReCaptchaSecret:     "",
	ReCaptchaAllowScore: 50,

	DatabaseConfig: &mysql.Config{
		DBName:               "cateiru-sso-test",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	Host: &url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	},
	SiteHost: &url.URL{
		Host:   "localhost:3000",
		Scheme: "http",
	},

	FromDomain:        "m.cateiru.com",
	MailgunSecret:     "",
	SenderMailAddress: "CateiruSSO <sso@m.cateiru.com>",

	RegisterSessionPeriod: 10 * time.Minute,
}

func InitConfig(mode string) *Config {
	switch mode {
	case "test":
		return TestConfig
	case "local":
		return LocalConfig
	case "cloudrun":
		return CloudRunConfig
	default:
		return TestConfig
	}
}
