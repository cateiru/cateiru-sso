// メールを送信します。
//
// メール本文は、テンプレートファイルを指定することができます。
//
// Example:
//	sendUserName := "test"
//	sendEmail := "example@example.com"
//	subject := "test"
//	templateFilePath := "test" // read test.html and test.tmpl
//
//	type Contents struct {
//		Title: string
//		Body:  string
//	}
//
//	element := Contents{Title: "test", Body: "dummy body"}
//
//	mail, err := NewMail(sendUserName, sendEmail, subject).AddContentsFromTemplate(templateFilePath, element)
//	err := mail.Send()
//
package mail

import (
	"os"

	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/mailgun/mailgun-go"
)

type Mail struct {
	UserName     string
	EmailAddress string
	Subject      string

	PlainText string
	HTMLText  string
}

// メール送信設定
func NewMail(userName string, email string, subject string) *Mail {
	return &Mail{
		UserName:     userName,
		EmailAddress: email,
		Subject:      subject,

		PlainText: "",
		HTMLText:  "",
	}
}

// メール本文を追加する
func (c *Mail) AddContents(plain string, html string) *Mail {
	c.PlainText = plain
	c.HTMLText = html

	return c
}

// テンプレートファイルを指定してメール本文を作成
func (c *Mail) AddContentsFromTemplate(tempName string, elements interface{}) (*Mail, error) {
	var err error

	c.HTMLText, err = Template(tempName+".html", elements)
	if err != nil {
		return nil, err
	}
	c.PlainText, err = Template(tempName+".tmpl", elements)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// メールをAPI経由で送信する
func (c *Mail) Send() error {
	domain := os.Getenv("MAIL_FROM_DOMAIL")
	secret := os.Getenv("MAILGUN_API_KEY")
	sender := os.Getenv("SENDER_MAIL_ADDRESS")

	mg := mailgun.NewMailgun(domain, secret)
	message := mg.NewMessage(sender, c.Subject, c.PlainText, c.EmailAddress)
	message.SetHtml(c.HTMLText)

	resp, id, err := mg.Send(message)
	if err != nil {
		return err
	}

	logging.Sugar.Infof("Send email. ID: %s, Resp: %s", id, resp)

	return nil
}
