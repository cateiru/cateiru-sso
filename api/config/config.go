package config

import (
	"os"
	"testing"
)

type Config struct {
	DeployMode string

	Port    string
	Address string

	AdminMail     string
	AdminPassword string

	CookieDomain string
	SiteDomain   string
	APIDomain    string

	CORS string

	Issuer string

	ReCaptchaSecret string

	MailgunAPIKey     string
	MailFromDomain    string
	SenderMailAddress string

	DatastoreParentKey string
}

var Defs Config

func Init() {
	Defs = Config{
		DeployMode: GetDeployMode(),

		Port:    GetPort(),
		Address: "0.0.0.0",

		AdminMail:     os.Getenv("ADMIN_MAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),

		CookieDomain: os.Getenv("COOKIE_DOMAIN"),
		SiteDomain:   os.Getenv("SITE_DOMAIN"),
		APIDomain:    os.Getenv("API_DOMAIN"),

		CORS: os.Getenv("CORS"),

		Issuer: GetIssuer(),

		ReCaptchaSecret: os.Getenv("RECAPTCHA_SECRET"),

		MailgunAPIKey:     os.Getenv("MAILGUN_API_KEY"),
		MailFromDomain:    os.Getenv("MAIL_FROM_DOMAIN"),
		SenderMailAddress: os.Getenv("SENDER_MAIL_ADDRESS"),

		DatastoreParentKey: GetDatastoreParentKey(),
	}
}

func TestInit(t *testing.T) {
	Init()

	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")
}

// envが設定されていない場合は、`develop`を指定します
func GetDeployMode() string {
	deployMode := os.Getenv("DEPLOY_MODE")

	if len(deployMode) == 0 {
		deployMode = "develop"
	}

	return deployMode
}

// envが設定されていない場合、3000を使用します
func GetPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	return port
}

func GetIssuer() string {
	issuer := os.Getenv("ISSUER")

	if len(issuer) == 0 {
		issuer = "TestIssuer"
	}

	return issuer
}

func GetDatastoreParentKey() string {
	parentKey := os.Getenv("DATASTORE_PARENT_KEY")
	if len(parentKey) == 0 {
		parentKey = "cateiru-sso"
	}

	return parentKey
}
