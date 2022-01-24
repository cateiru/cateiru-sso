package net

import (
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/rs/cors"
)

func CorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   url(),
		AllowCredentials: true,
		AllowedMethods: []string{
			"GET",
			"POST",
			"DELETE",
		},
	})
}

func url() []string {
	urls := []string{config.Defs.CookieDomain}

	if config.Defs.DeployMode != "production" {
		urls = append(urls, "http://localhost")
	}

	return urls
}
