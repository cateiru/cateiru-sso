package src

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/go-sql-driver/mysql"
	"github.com/go-webauthn/webauthn/webauthn"
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

	// アカウント登録時に使用するセッションの有効期限
	RegisterSessionPeriod     time.Duration
	RegisterSessionRetryLimit uint8
	RegisterEmailSendLimit    uint8

	// webAuthn(passkeyの共通設定)
	// ref. https://github.com/go-webauthn/webauthn
	WebAuthnConfig        *webauthn.Config
	WebAuthnSessionPeriod time.Duration
	WebAuthnSessionCookie CookieConfig

	// ログインセッション
	SessionDBPeriod time.Duration
	SessionCookie   CookieConfig
	RefreshDBPeriod time.Duration
	RefreshCookie   CookieConfig
	// 現在ログインしているセッション
	LoginUserCookie CookieConfig
	// ログイン状態をJSで見れるようにするCookie
	LoginStateCookie CookieConfig

	// Passwordハッシュ設定
	Password *lib.Password
}

// Cookieの設定
// http.Cookieの一部
// Domainなどは別途設定するため存在しない
type CookieConfig struct {
	Name     string
	Secure   bool
	HttpOnly bool
	Path     string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	SameSite http.SameSite
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
		Loc:                  time.FixedZone("Asia/Tokyo", 9*60*60),
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

	RegisterSessionPeriod:     10 * time.Minute,
	RegisterSessionRetryLimit: 5,
	RegisterEmailSendLimit:    3,

	WebAuthnConfig: &webauthn.Config{
		RPDisplayName: "Cateiru SSO Local",
		RPID:          "localhost:3000",
		RPOrigins:     []string{"localhost:3000", "localhost:8080"},
	},
	WebAuthnSessionPeriod: 5 * time.Minute,
	WebAuthnSessionCookie: CookieConfig{
		Name:     "cateiru-sso-webauthn-session",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   0,
		SameSite: http.SameSiteDefaultMode,
	},

	SessionDBPeriod: 168 * time.Hour, // 7days
	SessionCookie: CookieConfig{
		Name:     "cateiru-sso-session",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   604800, // 7days
		SameSite: http.SameSiteDefaultMode,
	},
	RefreshDBPeriod: 720 * time.Hour, // 30days
	RefreshCookie: CookieConfig{
		Name:     "cateiru-sso-refresh",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},
	// 現在ログインしているセッション
	LoginUserCookie: CookieConfig{
		Name:     "cateiru-sso-users",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},
	// ログイン状態をJSで見れるようにするCookie
	LoginStateCookie: CookieConfig{
		Name:     "cateiru-sso-login-state",
		Secure:   false,
		HttpOnly: false,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},

	Password: &lib.Password{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	},
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

	RegisterSessionPeriod:     10 * time.Minute,
	RegisterSessionRetryLimit: 5,
	RegisterEmailSendLimit:    3,

	WebAuthnConfig: &webauthn.Config{
		RPDisplayName: "Cateiru SSO",
		RPID:          "sso.cateiru.com",
		RPOrigins:     []string{"sso.cateiru.com", "api.sso.cateiru.com"},
	},
	WebAuthnSessionPeriod: 5 * time.Minute,
	WebAuthnSessionCookie: CookieConfig{
		Name:     "cateiru-sso-webauthn-session",
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   0,
		SameSite: http.SameSiteDefaultMode,
	},

	SessionDBPeriod: 168 * time.Hour, // 7days
	SessionCookie: CookieConfig{
		Name:     "cateiru-sso-session",
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   604800, // 7days
		SameSite: http.SameSiteDefaultMode,
	},
	RefreshDBPeriod: 720 * time.Hour, // 30days
	RefreshCookie: CookieConfig{
		Name:     "cateiru-sso-refresh",
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},
	// 現在ログインしているセッション
	LoginUserCookie: CookieConfig{
		Name:     "cateiru-sso-users",
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},
	// ログイン状態をJSで見れるようにするCookie
	LoginStateCookie: CookieConfig{
		Name:     "cateiru-sso-login-state",
		Secure:   true,
		HttpOnly: false,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},

	Password: &lib.Password{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	},
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
		Loc:                  time.FixedZone("Asia/Tokyo", 9*60*60),
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

	RegisterSessionPeriod:     10 * time.Minute,
	RegisterSessionRetryLimit: 5,
	RegisterEmailSendLimit:    3,

	WebAuthnConfig: &webauthn.Config{
		RPDisplayName: "Cateiru SSO",
		RPID:          "localhost:3000",
		RPOrigins:     []string{"localhost:3000", "localhost:8080"},
	},
	WebAuthnSessionPeriod: 5 * time.Minute,
	WebAuthnSessionCookie: CookieConfig{
		Name:     "cateiru-sso-webauthn-session",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   0,
		SameSite: http.SameSiteDefaultMode,
	},

	SessionDBPeriod: 168 * time.Hour, // 7days
	SessionCookie: CookieConfig{
		Name:     "cateiru-sso-session",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   604800, // 7days
		SameSite: http.SameSiteDefaultMode,
	},
	RefreshDBPeriod: 720 * time.Hour, // 30days
	RefreshCookie: CookieConfig{
		Name:     "cateiru-sso-refresh",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},
	// 現在ログインしているセッション
	LoginUserCookie: CookieConfig{
		Name:     "cateiru-sso-users",
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},
	// ログイン状態をJSで見れるようにするCookie
	LoginStateCookie: CookieConfig{
		Name:     "cateiru-sso-login-state",
		Secure:   false,
		HttpOnly: false,
		Path:     "/",
		MaxAge:   2592000, // 30days
		SameSite: http.SameSiteDefaultMode,
	},

	Password: &lib.Password{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	},
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
