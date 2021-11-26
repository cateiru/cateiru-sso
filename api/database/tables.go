package database

import (
	"time"
)

// ユーザIDのkv
type userId struct {
	UserId string `datastore:"userId" json:"user_id"`
}

// メールアドレスとパスワード
// パスワードはハッシュ化する必要がある
type userMailPW struct {
	Mail     string `datastore:"mail" json:"mail"`
	Password string `datastore:"password" json:"password"`
}

// 認証テーブル
// OnetimePasswordKeyはOptional
type Certification struct {
	AccountCreateDate time.Time `datastore:"accountCreateDate" json:"account_create_date"`

	OnetimePasswordKey string `datastore:"onetimePasswordKey" json:"onetime_password_key"`

	userMailPW
	userId
}

// メールアドレス認証用テーブル
type MailCertification struct {
	MailToken     string    `datastore:"mailToken" json:"mail_token"`
	CreateDate    time.Time `datastore:"createDate" json:"create_date"`
	PeriodMinute  int       `datastore:"periodMinute" json:"period_minute"`
	OpenNewWindow bool      `datastore:"openNewWindow" json:"open_new_window"`
	Verify        bool      `datastore:"verify" json:"verify"`

	userMailPW
}

// ユーザの情報
type User struct {
	FirstName string `datastore:"firstName" json:"first_name"`
	LastName  string `datastore:"lastName" json:"last_name"`

	Role string
	Mail string `datastore:"mail" json:"mail"`

	Theme     string `datastore:"theme" json:"theme"`
	AvatarUrl string `datastore:"avatarUrl" json:"avatar_url"`

	userId
}

// ログイン履歴（個別）
type loginHistory struct {
	AccessId  string    `datastore:"accessId" json:"access_id"`
	Date      time.Time `datastore:"date" json:"date"`
	IpAddress string    `datastore:"ipAddress" json:"ip_address"`
}

// ユーザのログイン履歴
type UserLoginHistories struct {
	Histories []loginHistory `datastore:"histories" json:"histories"`

	userId
}

// ユーザの定義したSSOのpublic keys
type UserCreatedSSO struct {
	SsoPublickeys []string `datastore:"ssoPublickeys" json:"sso_publickeys"`

	userId
}

// セッションandリフレッシュトークン保管時に使うもの
type tokenInfo struct {
	CreateDate time.Time `datastore:"createDate" json:"create_date"`
	PeriodHour int       `datastore:"periodHour" json:"period_hour"`

	userId
}

// CateiruSSOのセッション情報
type SessionInfo struct {
	SessionToken string `datastore:"sessionToken" json:"session_token"`

	tokenInfo
}

// CateiruSSOのリフレッシュトークン
type RefreshInfo struct {
	RefreshToken string `datastore:"refreshToken" json:"refresh_token"`
	SessionToken string `datastore:"sessionToken" json:"session_token"`

	tokenInfo
}

// SSO情報
// SessionTokenPeriod, RefreshTokenPeriodはOptional
type SSOService struct {
	SSOPublicKey string `datastore:"ssoPublicKey" json:"sso_publickey"`

	SSOSecretKey  string `datastore:"ssoSecretKey" json:"sso_secretkey"`
	SSOPrivateKey string `datastore:"ssoPrivateKey" json:"sso_privatekey"`

	Name      string   `datastore:"name" json:"name"`
	FromUrl   []string `datastore:"fromUrl" json:"from_url"`
	ToUrl     []string `datastore:"toUrl" json:"to_url"`
	LoginOnly bool     `datastore:"loginOnly" json:"login_only"`

	SessionTokenPeriod int `datastore:"sessionTokenPeriod" json:"session_token_period"`
	RefreshTokenPeriod int `datastore:"refreshTokenPeriod" json:"refresh_token_period"`
}

// SSOのセッショントークン
type SSOSession struct {
	SSOSessionToken string `datastore:"ssoSessionToken" json:"sso_session_token"`

	tokenInfo
}

// SSOのリフレッシュトークン
type SSORefreshToken struct {
	SSOSessionToken string `datastore:"ssoSessionToken" json:"sso_session_token"`
	SSORefreshToken string `datastore:"ssoRefreshToken" json:"sso_refresh_token"`

	tokenInfo
}

// Workerのログ
type WorkerLog struct {
	RunId   string    `datastore:"runId" json:"run_id"`
	Status  int       `datastore:"status" json:"status"`
	Message string    `datastore:"message" json:"message"`
	RunDate time.Time `datastore:"runDate" json:"run_date"`
}
