package lib_test

import (
	"bytes"
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

type TestTemplateData struct {
	Text string
}

func TestNewSender(t *testing.T) {
	s, err := lib.NewSender("templates/*", "example.com", "secret", "<Text> m@example.com")
	require.NoError(t, err)

	require.Equal(t, s.FromDomain, "example.com")
	require.Equal(t, s.MailgunSecret, "secret")
	require.Equal(t, s.SenderMailAddress, "<Text> m@example.com")

	buff := new(bytes.Buffer)
	err = s.Template.ExecuteTemplate(buff, "test.gtpl", nil)
	require.NoError(t, err)

	require.Equal(t, buff.String(), "test\n")
}

func TestPreview(t *testing.T) {
	s, err := lib.NewSender("templates/*", "example.com", "secret", "<Text> m@example.com")
	require.NoError(t, err)

	m := &lib.MailBody{
		EmailAddress: "test@example.test",
		Subject:      "subject",

		Data: TestTemplateData{
			Text: "OK",
		},

		PlainTextFileName: "test.gtpl",
		HTMLTextFileName:  "test.html",
	}

	text, err := s.Preview(m)
	require.NoError(t, err)

	require.Equal(t, text, "<div>OK</div>\n")
}
