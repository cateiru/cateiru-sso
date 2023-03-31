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

	// メールを送信するかどうか
	SendMail bool

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

	// OTPセッションの有効期限
	OTPSessionPeriod time.Duration
	// OTP登録セッションの有効期限
	OTPRegisterSessionPeriod time.Duration
	// OTPのリトライ回数の上限
	OTPRetryLimit uint8
	// OTP登録セッションのリトライ回数上限
	OTPRegisterLimit uint8
	// OTPのissuer
	OTPIssuer      string
	OTPBackupCount uint8

	// パスワード再設定の有効期限
	ReregistrationPasswordSessionPeriod time.Duration
	// パスワード再設定のレコード削除有効期限
	// この有効期限が切れるまで再度メールを送信することができない
	ReregistrationPasswordSessionClearPeriod time.Duration

	// メールアドレス更新のセッション有効期限
	UpdateEmailSessionPeriod time.Duration
	UpdateEmailRetryCount    uint8

	// Internal エンドポイントのBasic Auth
	InternalBasicAuthUserName string
	InternalBasicAuthPassword string

	// CDNのホスト
	CDNHost        *url.URL
	FastlyApiToken string

	StorageBucketName string

	// CloudStorageのエミュレータホスト
	StorageEmulatorHost struct {
		Value   string
		IsValid bool
	}

	// ---- Clientの設定

	ClientSessionPeriod time.Duration
	ClientRefreshPeriod time.Duration
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
	ReCaptchaSecret:     "secret",
	ReCaptchaAllowScore: 0,

	DatabaseConfig: &mysql.Config{
		DBName:               "cateiru-sso",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306", // APIから接続するのでlocalhost
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  time.FixedZone("Asia/Tokyo", 9*60*60),
		AllowNativePasswords: true,
	},
	Host: &url.URL{
		Host:   "sso.cateiru.test:8080",
		Scheme: "http",
	},
	SiteHost: &url.URL{
		Host:   "sso.cateiru.test:3000",
		Scheme: "http",
	},

	FromDomain:        "m.cateiru.test",
	MailgunSecret:     "secret",
	SenderMailAddress: "CateiruSSO <sso@m.cateiru.test>",

	SendMail: false,

	RegisterSessionPeriod:     10 * time.Minute,
	RegisterSessionRetryLimit: 5,
	RegisterEmailSendLimit:    3,

	WebAuthnConfig: &webauthn.Config{
		RPDisplayName: "Cateiru SSO Local",
		RPID:          "sso.cateiru.test",
		RPOrigins:     []string{"sso.cateiru.test:8080", "sso.cateiru.test:3000"},
		Timeouts: webauthn.TimeoutsConfig{
			Login: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Second * 60,
				TimeoutUVD: time.Second * 60,
			},
			Registration: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Second * 60,
				TimeoutUVD: time.Second * 60,
			},
		},
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

	OTPSessionPeriod:         5 * time.Minute,
	OTPRegisterSessionPeriod: 5 * time.Minute,
	OTPRetryLimit:            5,
	OTPRegisterLimit:         3,
	OTPIssuer:                "CateiruSSO",
	OTPBackupCount:           10,

	ReregistrationPasswordSessionPeriod:      5 * time.Minute,
	ReregistrationPasswordSessionClearPeriod: 1 * time.Hour,

	UpdateEmailSessionPeriod: 5 * time.Minute,
	UpdateEmailRetryCount:    5,

	InternalBasicAuthUserName: "user",
	InternalBasicAuthPassword: "password",

	CDNHost: &url.URL{
		Host:   "sso.cateiru.test:4443",
		Scheme: "http",
	},
	FastlyApiToken: "token",

	StorageBucketName: "cateiru-sso",

	StorageEmulatorHost: struct {
		Value   string
		IsValid bool
	}{
		Value:   "localhost:4443",
		IsValid: true,
	},

	ClientSessionPeriod: 1 * time.Hour,  // 1hour
	ClientRefreshPeriod: 24 * time.Hour, // 1days
}

var CloudRunConfig = &Config{
	Mode: "cloudrun",

	SelfSignedCert:      false,
	UseReCaptcha:        true,
	ReCaptchaSecret:     os.Getenv("RECAPTCHA_SECRET"),
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

	SendMail: true,

	RegisterSessionPeriod:     10 * time.Minute,
	RegisterSessionRetryLimit: 5,
	RegisterEmailSendLimit:    3,

	WebAuthnConfig: &webauthn.Config{
		RPDisplayName: "Cateiru SSO",
		RPID:          "sso.cateiru.com",
		RPOrigins:     []string{"sso.cateiru.com", "api.sso.cateiru.com"},
		Timeouts: webauthn.TimeoutsConfig{
			Login: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Second * 60,
				TimeoutUVD: time.Second * 60,
			},
			Registration: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Second * 60,
				TimeoutUVD: time.Second * 60,
			},
		},
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

	OTPSessionPeriod:         5 * time.Minute,
	OTPRegisterSessionPeriod: 5 * time.Minute,
	OTPRetryLimit:            5,
	OTPRegisterLimit:         3,
	OTPIssuer:                "CateiruSSO",
	OTPBackupCount:           10,

	ReregistrationPasswordSessionPeriod:      5 * time.Minute,
	ReregistrationPasswordSessionClearPeriod: 1 * time.Hour,

	UpdateEmailSessionPeriod: 5 * time.Minute,
	UpdateEmailRetryCount:    5,

	InternalBasicAuthUserName: "user",
	InternalBasicAuthPassword: "password",

	CDNHost: &url.URL{
		Host:   "cdn.sso.cateiru.com",
		Scheme: "https",
	},
	FastlyApiToken: os.Getenv("FASTLY_API_TOKEN"),

	StorageBucketName: "cateiru-sso",

	StorageEmulatorHost: struct {
		Value   string
		IsValid bool
	}{
		Value:   "",
		IsValid: false,
	},

	ClientSessionPeriod: 1 * time.Hour,  // 1hour
	ClientRefreshPeriod: 24 * time.Hour, // 1days
}

var TestConfig = &Config{
	Mode: "test",

	SelfSignedCert:      false,
	UseReCaptcha:        true, // mockするので問題なし
	ReCaptchaSecret:     "secret",
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
	MailgunSecret:     "secret",
	SenderMailAddress: "CateiruSSO <sso@m.cateiru.com>",

	SendMail: true,

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

	OTPSessionPeriod:         5 * time.Minute,
	OTPRegisterSessionPeriod: 5 * time.Minute,
	OTPRetryLimit:            5,
	OTPRegisterLimit:         3,
	OTPIssuer:                "CateiruSSO",
	OTPBackupCount:           10,

	ReregistrationPasswordSessionPeriod:      5 * time.Minute,
	ReregistrationPasswordSessionClearPeriod: 1 * time.Hour,

	UpdateEmailSessionPeriod: 5 * time.Minute,
	UpdateEmailRetryCount:    5,

	InternalBasicAuthUserName: "user",
	InternalBasicAuthPassword: "password",

	CDNHost: &url.URL{
		Host:   "localhost:4000",
		Scheme: "http",
	},
	FastlyApiToken: "token",

	StorageBucketName: "test-cateiru-sso",

	StorageEmulatorHost: struct {
		Value   string
		IsValid bool
	}{
		Value:   "localhost:4443",
		IsValid: true,
	},

	ClientSessionPeriod: 1 * time.Hour,  // 1hour
	ClientRefreshPeriod: 24 * time.Hour, // 1days
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
