package src

import (
	"net/url"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Mode string

	// MySQLの設定
	DatabaseConfig *mysql.Config

	// APIのホスト
	Host *url.URL
	// サイトのホスト
	SiteHost *url.URL
}

var LocalConfig = &Config{
	Mode: "local",

	DatabaseConfig: &mysql.Config{
		DBName:               "cateiru-sso",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	Host: &url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	},
	SiteHost: &url.URL{
		Host:   "localhost:3000",
		Scheme: "http",
	},
}

var CloudRunConfig = &Config{
	Mode: "cloudrun",
}

var TestConfig = &Config{
	Mode: "test",

	DatabaseConfig: &mysql.Config{
		DBName:               "cateiru-sso-test",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	Host: &url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	},
	SiteHost: &url.URL{
		Host:   "localhost:3000",
		Scheme: "http",
	},
}

func InitConfig(mode string) *Config {
	switch mode {
	case "test":
		return TestConfig
	case "local":
		return LocalConfig
	case "cloudrun":
		return CloudRunConfig
	default:
		return TestConfig
	}
}
