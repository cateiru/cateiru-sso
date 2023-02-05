package config_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/stretchr/testify/require"
)

func TestSetConfig(t *testing.T) {
	t.Setenv("DEPLOY_MODE", "production")

	t.Setenv("PORT", "8080")

	t.Setenv("ADMIN_MAIL", "example@example.com")
	t.Setenv("ADMIN_PASSWORD", "password")

	t.Setenv("COOKIE_DOMAIN", "cookie_domain")
	t.Setenv("SITE_DOMAIN", "site_domain")
	t.Setenv("API_DOMAIN", "api_domain")

	t.Setenv("ISSUER", "test")

	t.Setenv("RECAPTCHA_SECRET", "secret")

	t.Setenv("MAILGUN_API_KEY", "mailgun")
	t.Setenv("MAIL_FROM_DOMAIN", "example.com")
	t.Setenv("SENDER_MAIL_ADDRESS", "info@example.com")

	t.Setenv("DATASTORE_PARENT_KEY", "parent_key")

	config.Init()

	require.Equal(t, config.Defs.DeployMode, "production")

	require.Equal(t, config.Defs.Port, "8080")

	require.Equal(t, config.Defs.AdminMail, "example@example.com")
	require.Equal(t, config.Defs.AdminPassword, "password")

	require.Equal(t, config.Defs.CookieDomain, "cookie_domain")
	require.Equal(t, config.Defs.SiteDomain, "site_domain")
	require.Equal(t, config.Defs.APIDomain, "api_domain")

	require.Equal(t, config.Defs.Issuer, "test")

	require.Equal(t, config.Defs.ReCaptchaSecret, "secret")

	require.Equal(t, config.Defs.MailgunAPIKey, "mailgun")
	require.Equal(t, config.Defs.MailFromDomain, "example.com")
	require.Equal(t, config.Defs.SenderMailAddress, "info@example.com")

	require.Equal(t, config.Defs.DatastoreParentKey, "parent_key")
}

func TestDeployMode(t *testing.T) {
	deploy := config.GetDeployMode()
	require.Equal(t, deploy, "develop")

	t.Setenv("DEPLOY_MODE", "production")

	deploy = config.GetDeployMode()
	require.Equal(t, deploy, "production")
}

func TestPort(t *testing.T) {
	port := config.GetPort()
	require.Equal(t, port, "3000")

	t.Setenv("PORT", "8080")

	port = config.GetPort()
	require.Equal(t, port, "8080")
}

func TestIssuer(t *testing.T) {
	issuer := config.GetIssuer()
	require.Equal(t, issuer, "TestIssuer")

	t.Setenv("ISSUER", "nya")

	issuer = config.GetIssuer()
	require.Equal(t, issuer, "nya")
}

func TestDatastore(t *testing.T) {
	key := config.GetDatastoreParentKey()
	require.Equal(t, key, "cateiru-sso")

	t.Setenv("DATASTORE_PARENT_KEY", "parent_key")

	key = config.GetDatastoreParentKey()
	require.Equal(t, key, "parent_key")
}
