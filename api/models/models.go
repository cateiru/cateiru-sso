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
	Password string `datastore:"password" json:"password"`
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
type VerifyPeriod struct {
	CreateDate   time.Time `datastore:"createDate" json:"create_date"`
	PeriodMinute int       `datastore:"periodMinute" json:"period_minute"`
}

// メールアドレス認証用テーブル
type MailCertification struct {
	MailToken        string `datastore:"mailToken" json:"mail_token"`
	ClientCheckToken string `datastore:"clientCheckToken" json:"client_check_token"`
	OpenNewWindow    bool   `datastore:"openNewWindow" json:"open_new_window"`
	Verify           bool   `datastore:"verify" json:"verify"`
	ChangeMailMode   bool   `datastore:"changeMailMode" json:"change_mail_mode"`

	VerifyPeriod
	UserMailPW
}

// メールアドレスの認証が済んでいるが、名前、その他ユーザ設定が完了してないユーザのデータの一時保管場所
type CreateAccountBuffer struct {
	BufferToken string `datastore:"bufferToken" json:"buffer_token"`

	VerifyPeriod
	UserMailPW
}

// パスワード忘れによる再登録用テーブル
type PWForget struct {
	ForgetToken string `datastore:"forgetToken" json:"forget_token"`
	Mail        string `datastore:"mail" json:"mail"`

	VerifyPeriod
}

// ワンタイムパスワード設定用
type OnetimePassword struct {
	PublicKey string `datastore:"onetimePublicKey" json:"onetime_public_key"`
	SecretKey string `datastore:"onetimeSecretKey" json:"onetime_secret_key"`

	VerifyPeriod
	UserId
}

// ログイン時、メアドとPWを入力後、ワンタイムパスワードが求められる場合のテーブル
type OnetimePasswordValidate struct {
	OnetimeToken          string `datastore:"onetimeToken" json:"onetime_token"`
	OnetimePasswordSecret string `datastore:"onetimePasswordSecret" json:"onetime_password_secret"`

	VerifyPeriod
	UserId
}

// ユーザの情報
type User struct {
	FirstName string `datastore:"firstName" json:"first_name"`
	LastName  string `datastore:"lastName" json:"last_name"`
	UserName  string `datastore:"userName" json:"user_name"`

	Role []string `datastore:"role" json:"role"`
	Mail string   `datastore:"mail" json:"mail"`

	Theme     string `datastore:"theme" json:"theme"`
	AvatarUrl string `datastore:"avatarUrl" json:"avatar_url"`

	UserId
}

// ログイン履歴（個別）
// IsSSOとSSOPublicKeyはOptional
type LoginHistory struct {
	AccessId     string    `datastore:"accessId" json:"access_id"`
	Date         time.Time `datastore:"date" json:"date"`
	IpAddress    string    `datastore:"ipAddress" json:"ip_address"`
	UserAgent    string    `datastore:"userAgent" json:"user_agent"`
	IsSSO        bool      `datastore:"isSSO,omitempty" json:"is_sso"`
	SSOPublicKey string    `datastore:"ssoPublicKey,omitempty" json:"sso_publickey"`

	UserId
}

// ログインしているSSO
type SSOLogins struct {
	SSORefreshTokens []string `datastore:"ssoRefreshTokens" json:"sso_refresh_tokens"`

	UserId
}

// セッションandリフレッシュトークン保管時に使うもの
type TokenInfo struct {
	CreateDate time.Time `datastore:"createDate" json:"create_date"`
	PeriodHour int       `datastore:"periodHour" json:"period_hour"`

	UserId
}

// CateiruSSOのセッション情報
type SessionInfo struct {
	SessionToken string `datastore:"sessionToken" json:"session_token"`

	TokenInfo
}

// CateiruSSOのリフレッシュトークン
type RefreshInfo struct {
	RefreshToken string `datastore:"refreshToken" json:"refresh_token"`
	SessionToken string `datastore:"sessionToken" json:"session_token"`

	TokenInfo
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

	SessionTokenPeriod int `datastore:"sessionTokenPeriod,omitempty" json:"session_token_period"`
	RefreshTokenPeriod int `datastore:"refreshTokenPeriod,omitempty" json:"refresh_token_period"`

	UserId
}

// SSOのセッショントークン
type SSOSession struct {
	SSOSessionToken string `datastore:"ssoSessionToken" json:"sso_session_token"`

	TokenInfo
}

// SSOのリフレッシュトークン
type SSORefreshToken struct {
	SSOSessionToken string `datastore:"ssoSessionToken" json:"sso_session_token"`
	SSORefreshToken string `datastore:"ssoRefreshToken" json:"sso_refresh_token"`

	TokenInfo
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
