package src

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"go.uber.org/zap"
)

type EmailData struct {
	BrandName     string // ブランド名
	BrandUrl      string // ブランドのURL
	BrandImageUrl string // ブランドの画像URL
	BrandDomain   string // ブランドのドメイン

	Email string // 送信先メールアドレス
}

// アカウント作成のテンプレートデータ
type RegisterEmailVerifyTemplate struct {
	Code       string
	Expiration time.Time

	EmailData
}

// メールアドレス更新のテンプレートデータ
type UpdateEmailTemplate2 struct {
	OldEmail   string
	Code       string
	Expiration time.Time

	EmailData
}

// パスワード再登録テンプレートデータ
type AccountReRegisterPasswordTemplate2 struct {
	URL        string
	UserName   string
	Expiration time.Time

	EmailData
}

// 組織招待のテンプレートデータ
type InviteOrgSessionTemplate2 struct {
	URL                string
	Expiration         time.Time
	OrganizationName   string
	InvitationUserName string

	EmailData
}

type Email struct {
	S     lib.SenderInterface
	C     *Config
	Email string // 送信先メールアドレス

	EmailData *EmailData
	UserData  *UserData    // 送信したときのユーザーデータ
	Ip        string       // 送信したときのIP
	User      *models.User // 送信したときのユーザー nullable

	HasPreviewMode bool
}

func NewEmail(s lib.SenderInterface, c *Config, email string, userData *UserData, ip string, user *models.User) *Email {
	emailData := &EmailData{
		BrandName:     c.BrandName,
		BrandUrl:      c.SiteHost.String(),
		BrandImageUrl: "https://todo",
		BrandDomain:   c.SiteHost.Host,
		Email:         email,
	}

	return &Email{
		S:         s,
		C:         c,
		Email:     email,
		EmailData: emailData,
		UserData:  userData,
		Ip:        ip,
		User:      user,

		HasPreviewMode: false, // プレビューモードはOFF、ONにするには上書きする
	}
}

// アカウント登録時に送信するメール
func (e *Email) RegisterEmailVerify(code string) (string, error) {
	m := &lib.MailBody{
		EmailAddress: e.Email,
		Subject:      "メールアドレスの登録確認",
		Data: RegisterEmailVerifyTemplate{
			Code:       code,
			Expiration: time.Now().Add(e.C.RegisterSessionPeriod),
			EmailData:  *e.EmailData,
		},
		PlainTextFileName: "register.gtpl",
		HTMLTextFileName:  "register.html",
	}

	if e.HasPreviewMode {
		return e.S.Preview(m)
	}

	msg, id, err := e.S.Send(m)
	if err != nil {
		L.Error("register mail",
			zap.String("Email", e.Email),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", e.Ip),
			zap.String("Device", e.UserData.Device),
			zap.String("Browser", e.UserData.Browser),
			zap.String("OS", e.UserData.OS),
			zap.Bool("IsMobile", e.UserData.IsMobile),
		)
		return "", err
	}

	// メールを送信したのでログを出す
	L.Info("register mail",
		zap.String("Email", e.Email),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", e.Ip),
		zap.String("Device", e.UserData.Device),
		zap.String("Browser", e.UserData.Browser),
		zap.String("OS", e.UserData.OS),
		zap.Bool("IsMobile", e.UserData.IsMobile),
	)
	return "", nil
}

// アカウント登録時に送信するメールの再送メール
func (e *Email) ResendRegisterEmailVerify(code string) (string, error) {
	m := &lib.MailBody{
		EmailAddress: e.Email,
		Subject:      "【再送】メールアドレスの登録確認",
		Data: RegisterEmailVerifyTemplate{
			Code:       code,
			Expiration: time.Now().Add(e.C.RegisterSessionPeriod),
			EmailData:  *e.EmailData,
		},
		PlainTextFileName: "register.gtpl",
		HTMLTextFileName:  "register.html",
	}

	if e.HasPreviewMode {
		return e.S.Preview(m)
	}

	msg, id, err := e.S.Send(m)
	if err != nil {
		L.Error("resend register mail",
			zap.String("Email", e.Email),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", e.Ip),
			zap.String("Device", e.UserData.Device),
			zap.String("Browser", e.UserData.Browser),
			zap.String("OS", e.UserData.OS),
			zap.Bool("IsMobile", e.UserData.IsMobile),
		)
		return "", err
	}

	// メールを送信したのでログを出す
	L.Info("resend register mail",
		zap.String("Email", e.Email),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", e.Ip),
		zap.String("Device", e.UserData.Device),
		zap.String("Browser", e.UserData.Browser),
		zap.String("OS", e.UserData.OS),
		zap.Bool("IsMobile", e.UserData.IsMobile),
	)
	return "", nil
}

// メールアドレス更新
func (e *Email) UpdateEmail(oldEmail string, code string) (string, error) {
	m := &lib.MailBody{
		EmailAddress: e.Email,
		Subject:      "メールアドレスの確認して更新します",
		Data: UpdateEmailTemplate2{
			OldEmail:   oldEmail,
			Code:       code,
			Expiration: time.Now().Add(e.C.UpdateEmailSessionPeriod),

			EmailData: *e.EmailData,
		},
		PlainTextFileName: "update_email.gtpl",
		HTMLTextFileName:  "update_email.html",
	}

	if e.HasPreviewMode {
		return e.S.Preview(m)
	}

	msg, id, err := e.S.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", e.Email),
			zap.String("OldEmail", oldEmail),
			zap.String("UserID", e.User.ID),
			zap.String("UserName", e.User.UserName),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", e.Ip),
			zap.String("Device", e.UserData.Device),
			zap.String("Browser", e.UserData.Browser),
			zap.String("OS", e.UserData.OS),
			zap.Bool("IsMobile", e.UserData.IsMobile),
		)
		return "", err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("Email", e.Email),
		zap.String("OldEmail", oldEmail),
		zap.String("UserID", e.User.ID),
		zap.String("UserName", e.User.UserName),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", e.Ip),
		zap.String("Device", e.UserData.Device),
		zap.String("Browser", e.UserData.Browser),
		zap.String("OS", e.UserData.OS),
		zap.Bool("IsMobile", e.UserData.IsMobile),
	)

	return "", nil
}

// パスワード更新
func (e *Email) UpdatePassword(token string, userName string) (string, error) {
	u, err := url.Parse(e.C.SiteHost.String())
	if err != nil {
		return "", err
	}

	u.Path = "/forget_password/reregister"

	query := u.Query()
	query.Set("token", token)
	query.Set("email", e.Email)
	u.RawQuery = query.Encode()

	m := &lib.MailBody{
		EmailAddress: e.Email,
		Subject:      "パスワードを再設定してください",
		Data: AccountReRegisterPasswordTemplate2{
			URL:        u.String(),
			UserName:   userName,
			Expiration: time.Now().Add(e.C.ReregistrationPasswordSessionPeriod),

			EmailData: *e.EmailData,
		},
		PlainTextFileName: "forget_reregistration_password.gtpl",
		HTMLTextFileName:  "forget_reregistration_password.html",
	}

	if e.HasPreviewMode {
		return e.S.Preview(m)
	}

	msg, id, err := e.S.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", e.Email),
			zap.String("UserID", e.User.ID),
			zap.String("UserName", e.User.UserName),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", e.Ip),
			zap.String("Device", e.UserData.Device),
			zap.String("Browser", e.UserData.Browser),
			zap.String("OS", e.UserData.OS),
			zap.Bool("IsMobile", e.UserData.IsMobile),
		)
		return "", err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("Email", e.Email),
		zap.String("UserID", e.User.ID),
		zap.String("UserName", e.User.UserName),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", e.Ip),
		zap.String("Device", e.UserData.Device),
		zap.String("Browser", e.UserData.Browser),
		zap.String("OS", e.UserData.OS),
		zap.Bool("IsMobile", e.UserData.IsMobile),
	)

	return "", nil
}

func (e *Email) InviteOrg(token string, orgName string, InvitationUserName string) (string, error) {
	u, err := url.Parse(e.C.SiteHost.String())
	if err != nil {
		return "", err
	}

	u.Path = "/register"

	query := u.Query()
	query.Set("invite_token", token)
	query.Set("email", e.Email)
	u.RawQuery = query.Encode()

	m := &lib.MailBody{
		EmailAddress: e.Email,
		Subject:      fmt.Sprintf("%sに招待されています", orgName),
		Data: InviteOrgSessionTemplate2{
			URL:                u.String(),
			Expiration:         time.Now().Add(e.C.InviteOrgSessionPeriod),
			OrganizationName:   orgName,
			InvitationUserName: InvitationUserName,

			EmailData: *e.EmailData,
		},
		PlainTextFileName: "invite_org.gtpl",
		HTMLTextFileName:  "invite_org.html",
	}

	if e.HasPreviewMode {
		return e.S.Preview(m)
	}

	msg, id, err := e.S.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", e.Email),
			zap.String("Subject", m.Subject),
			zap.String("OrgName", orgName),
			zap.String("InvitationUserName", InvitationUserName),
			zap.Error(err),
			zap.String("IP", e.Ip),
			zap.String("Device", e.UserData.Device),
			zap.String("Browser", e.UserData.Browser),
			zap.String("OS", e.UserData.OS),
			zap.Bool("IsMobile", e.UserData.IsMobile),
		)
		return "", err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("Email", e.Email),
		zap.String("Subject", m.Subject),
		zap.String("OrgName", orgName),
		zap.String("InvitationUserName", InvitationUserName),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", e.Ip),
		zap.String("Device", e.UserData.Device),
		zap.String("Browser", e.UserData.Browser),
		zap.String("OS", e.UserData.OS),
		zap.Bool("IsMobile", e.UserData.IsMobile),
	)

	return "", nil
}

func (e *Email) Test() (string, error) {
	m := &lib.MailBody{
		EmailAddress:      e.Email,
		Subject:           "test",
		Data:              e.EmailData,
		PlainTextFileName: "test.gtpl",
		HTMLTextFileName:  "test.html",
	}

	if e.HasPreviewMode {
		return e.S.Preview(m)
	}

	msg, id, err := e.S.Send(m)
	if err != nil {
		L.Error("mail",
			zap.String("Email", e.Email),
			zap.String("Subject", m.Subject),
			zap.Error(err),
			zap.String("IP", e.Ip),
			zap.String("Device", e.UserData.Device),
			zap.String("Browser", e.UserData.Browser),
			zap.String("OS", e.UserData.OS),
			zap.Bool("IsMobile", e.UserData.IsMobile),
		)
		return "", err
	}

	// メールを送信したのでログを出す
	L.Info("mail",
		zap.String("Email", e.Email),
		zap.String("Subject", m.Subject),
		zap.String("MailGunMessage", msg),
		zap.String("MailGunID", id),
		zap.String("IP", e.Ip),
		zap.String("Device", e.UserData.Device),
		zap.String("Browser", e.UserData.Browser),
		zap.String("OS", e.UserData.OS),
		zap.Bool("IsMobile", e.UserData.IsMobile),
	)

	return "", nil
}
