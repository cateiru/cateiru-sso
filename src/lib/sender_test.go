package lib_test

import (
	"bytes"
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestNewSender(t *testing.T) {
	s, err := lib.NewSender("*.gtpl", "example.com", "secret", "<Text> m@example.com")
	require.NoError(t, err)

	require.Equal(t, s.FromDomain, "example.com")
	require.Equal(t, s.MailgunSecret, "secret")
	require.Equal(t, s.SenderMailAddress, "<Text> m@example.com")

	buff := new(bytes.Buffer)
	err = s.Template.ExecuteTemplate(buff, "test.gtpl", nil)
	require.NoError(t, err)

	require.Equal(t, buff.String(), "test\n")
}
