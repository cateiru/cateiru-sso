package src

import (
	"net/url"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Mode string

	// MySQLの設定
	DatabaseConfig *mysql.Config

	Host *url.URL
}

var TestConfig = &Config{}

func InitConfig(mode string) *Config {
	switch mode {
	default:
		return TestConfig
	}
}
