package lib

import (
	"bytes"
	"html/template"

	"github.com/mailgun/mailgun-go"
)

type Sender struct {
	// 使用するテンプレートのディレクトリ
	Template *template.Template

	FromDomain        string
	MailgunSecret     string
	SenderMailAddress string
}

type MailBody struct {
	EmailAddress string
	Subject      string

	Data any

	PlainTextFileName string
	HTMLTextFileName  string
}

func NewSender(pattern string, fromDomain string, mailgunSecret string, senderMailAddress string) (*Sender, error) {
	template, err := template.New("templates").ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	return &Sender{
		Template: template,

		FromDomain:        fromDomain,
		MailgunSecret:     mailgunSecret,
		SenderMailAddress: senderMailAddress,
	}, nil
}

// MailGun経由でメールを送信する
// 戻り値は、(message, id, error)です。
// message、idはエラーが発生しなかった場合にのみ設定される
func (s *Sender) Send(m *MailBody) (string, string, error) {

	plainTextBuff := new(bytes.Buffer)
	if err := s.Template.ExecuteTemplate(plainTextBuff, m.PlainTextFileName, m.Data); err != nil {
		return "", "", err
	}
	htmlBuff := new(bytes.Buffer)
	if err := s.Template.ExecuteTemplate(htmlBuff, m.HTMLTextFileName, m.Data); err != nil {
		return "", "", err
	}

	mg := mailgun.NewMailgun(s.FromDomain, s.MailgunSecret)
	message := mg.NewMessage(s.SenderMailAddress, m.Subject, plainTextBuff.String(), m.EmailAddress)
	message.SetHtml(htmlBuff.String())

	return mg.Send(message)
}
