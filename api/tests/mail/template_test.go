package mail_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils/mail"
	"github.com/stretchr/testify/require"
)

const TEMPFILE_PATH = "test"

type Element struct {
	Title string
	Body  string
}

// メール送信用のテンプレートパースのテスト
func TestTemplate(t *testing.T) {
	text := Element{
		Title: "sample title",
		Body:  "sample body",
	}

	mail := mail.NewMail("test_user", "example@example.com", "test title")

	_, err := mail.AddContentsFromTemplate(TEMPFILE_PATH, text)
	require.NoError(t, err, "テンプレートのパースに失敗した")

	t.Log(mail.HTMLText)
	t.Log(mail.PlainText)

	require.NotEqual(t, len(mail.HTMLText), 0, "HTMLTextがある")
	require.NotEqual(t, len(mail.PlainText), 0, "PlainTextがある")

}
