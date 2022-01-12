package main

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/routes"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func init() {
	config.Init()
}

func main() {
	mux := http.NewServeMux()
	h2s := &http2.Server{}

	routes.Routes(mux)

	server := &http.Server{
		Addr:    config.Defs.Address + ":" + config.Defs.Port,
		Handler: h2c.NewHandler(mux, h2s),
	}
	defer server.Close()

	logging.Sugar.Infof("Start server! addr: %v, port: %v", config.Defs.Address, config.Defs.Port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
