package models

import (
	"time"
)

// ユーザIDのkv
type UserId struct {
	UserId string `datastore:"userId" json:"user_id"`
}

// メールアドレスとパスワード
// パスワードはハッシュ化する必要がある
type UserMailPW struct {
	Mail     string `datastore:"mail" json:"mail"`
	Password []byte `datastore:"password" json:"password"`
	Salt     []byte `datastore:"salt" json:"salt"`
}

// 認証テーブル
// OnetimePasswordSecret, OnetimePasswordBackupはOptional
// OTPはoptionalであるがアカウント登録時必須なため、実質admin userのログイン用
type Certification struct {
	AccountCreateDate time.Time `datastore:"accountCreateDate" json:"account_create_date"`

	OnetimePasswordSecret  string   `datastore:"onetimePasswordSecret,omitempty" json:"onetime_password_secret"`
	OnetimePasswordBackups []string `datastore:"onetimePasswordBackups,omitempty" json:"onetime_password_backups"`

	UserMailPW
	UserId
}

// パスコード再設定や、ワンタイムパスワード入力、ユーザ登録などのテーブルにおいて制限時間を設ける
//
// PeriodMinuteとPeriodHourはどちらか
type Period struct {
	CreateDate   time.Time `datastore:"createDate" json:"create_date"`
	PeriodMinute int       `datastore:"periodMinute,omitempty" json:"period_minute"`
	PeriodHour   int       `datastore:"periodHour,omitempty" json:"period_hour"`
	PeriodDay    int       `datastore:"periodDay,omitempty" json:"period_day"`
}

// メールアドレス認証用テーブル
type MailCertification struct {
	MailToken      string `datastore:"mailToken" json:"mail_token"`
	ClientToken    string `datastore:"clientToken" json:"client_token"`
	OpenNewWindow  bool   `datastore:"openNewWindow" json:"open_new_window"`
	Verify         bool   `datastore:"verify" json:"verify"`
	ChangeMailMode bool   `datastore:"changeMailMode" json:"change_mail_mode"`

	Period

	Mail string `datastore:"mail" json:"mail"`

	UserId string `datastore:"userId,omitempty" json:"user_id"` // Option
}

// パスワード忘れによる再登録用テーブル
type PWForget struct {
	ForgetToken string `datastore:"forgetToken" json:"forget_token"`
	Mail        string `datastore:"mail" json:"mail"`

	Period
}

// ワンタイムパスワード設定 & ログイン時一時保存用
//
//	- パスワード設定
//		OPTのトークンを発行した際に、そのトークンで生成したコードとこのidを送ることでOTPを設定できる。
//	- ログイン時
//		ログイン後、OTPが設定されている場合、このテーブルに格納し、Idをcookieに格納する。
//		その後、OTPを入力してもらい検証することでログインする。
type OnetimePasswordBuffer struct {
	Id string `datastore:"id" json:"id"`

	PublicKey string `datastore:"onetimePublicKey,omitempty" json:"onetime_public_key"`
	SecretKey string `datastore:"onetimeSecretKey" json:"onetime_secret_key"`

	IsLogin bool `datastore:"isLogin" json:"is_login"`

	Period
	UserId
}

// ユーザの情報
type User struct {
	FirstName         string `datastore:"firstName" json:"first_name"`
	LastName          string `datastore:"lastName" json:"last_name"`
	UserName          string `datastore:"userName" json:"user_name"`
	UserNameFormatted string `datastore:"userNameFormatted" json:"user_name_formatted"`

	Mail string `datastore:"mail" json:"mail"`

	Theme     string `datastore:"theme" json:"theme"`
	AvatarUrl string `datastore:"avatarUrl" json:"avatar_url"`

	Role []string `datastore:"role" json:"role"`

	UserId
}

// ユーザの権限
type Role struct {
	Role []string `datastore:"role" json:"role"`

	UserId
}

// ログイン履歴（個別）
// IsSSOとSSOPublicKeyはOptional
type LoginHistory struct {
	AccessId  string    `datastore:"accessId" json:"access_id"`
	Date      time.Time `datastore:"date" json:"date"`
	IpAddress string    `datastore:"ipAddress" json:"ip_address"`
	UserAgent string    `datastore:"userAgent" json:"user_agent"`

	UserId
}

// ログインしているSSO
type SSOLogins struct {
	SSORefreshTokens []string `datastore:"ssoRefreshTokens" json:"sso_refresh_tokens"`

	UserId
}

// CateiruSSOのセッション情報
type SessionInfo struct {
	SessionToken string `datastore:"sessionToken" json:"session_token"`

	AccessID string `datastore:"accessId" json:"access_id"`

	Period
	UserId
}

// CateiruSSOのリフレッシュトークン
type RefreshInfo struct {
	RefreshToken string `datastore:"refreshToken" json:"refresh_token"`
	SessionToken string `datastore:"sessionToken" json:"session_token"`

	AccessID string `datastore:"accessId" json:"access_id"`

	Period
	UserId
}

// SSO情報
// SessionTokenPeriod, RefreshTokenPeriodはOptional
type SSOService struct {
	ClientID string `datastore:"clientId" json:"client_id"`

	TokenSecret string `datastore:"tokenSecret" json:"token_secret"`

	Name        string `datastore:"name" json:"name"`
	ServiceIcon string `datastore:"serviceIcon" json:"service_icon"`

	FromUrl []string `datastore:"fromUrl" json:"from_url"`
	ToUrl   []string `datastore:"toUrl" json:"to_url"`

	RefreshTokenPeriod int `datastore:"refreshTokenPeriod,omitempty" json:"refresh_token_period"`

	UserId
}

type SSOAccessToken struct {
	SSOAccessToken string `datastore:"ssoAccessToken" json:"sso_access_token"`

	ClientID string `datastore:"clientId" json:"client_id"`

	Period
	UserId
}

// SSOのリフレッシュトークン
type SSORefreshToken struct {
	SSOAccessToken  string `datastore:"ssoAccessToken" json:"sso_access_token"`
	SSORefreshToken string `datastore:"ssoRefreshToken" json:"sso_refresh_token"`

	ClientID string `datastore:"clientId" json:"client_id"`

	Period
	UserId
}

// Workerのログ
type WorkerLog struct {
	RunId   string    `datastore:"runId" json:"run_id"`
	Status  int       `datastore:"status" json:"status"`
	Message string    `datastore:"message" json:"message"`
	RunDate time.Time `datastore:"runDate" json:"run_date"`
}

// IPアドレスのブロックリスト
// アカウント作成時、該当IPがブロックされていたら作成できない
type IPBlockList struct {
	IP string `datastore:"ip" json:"ip"`
}

// メールアドレスのブロックリスト
type MailBlockList struct {
	Mail string `datastore:"mail" json:"mail"`
}

// アカウント作成時のログ
// 悪意のあるユーザにスパムメールを送られた場合、このログからIPをブロックします
type TryCreateAccountLog struct {
	LogId      string    `datastore:"logId" json:"log_id"`
	IP         string    `datastore:"ip" json:"ip"`
	TryDate    time.Time `datastore:"tryDate" json:"try_date"`
	TargetMail string    `datastore:"targetMail" json:"target_mail"`
}

type SSOServiceLog struct {
	LogId      string    `datastore:"logId" json:"log_id"`
	AcceptDate time.Time `datastore:"acceptDate" json:"accept_date"`
	ClientID   string    `datastore:"clientId" json:"client_id"`

	UserId
}
