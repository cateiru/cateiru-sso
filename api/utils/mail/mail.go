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
	"errors"
	"os"

	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
	from := createFromAccount()
	to := mail.NewEmail(c.UserName, c.EmailAddress)
	message := mail.NewSingleEmail(from, c.Subject, to, c.PlainText, c.HTMLText)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	logging.Sugar.Debugf("Send email. to: %v, subj: %v, contetents: %v", c.EmailAddress, c.Subject, c.PlainText)

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	// APIから200以外が返ってくる場合はエラー
	if response.StatusCode != 200 {
		logging.Sugar.Errorf("Failed send email. status: %d, body: %v", response.StatusCode, response.Body)
		return errors.New("failed send email")
	}

	return nil
}

// 送信元のメールアドレスと名前を設定します
//
//	name: MAIL_FROM_NAME
//	email: MAIL_FROM_ADDRESS
func createFromAccount() *mail.Email {
	name := os.Getenv("MAIL_FROM_NAME")
	email := os.Getenv("MAIL_FROM_ADDRESS")
	return mail.NewEmail(name, email)
}
