package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/cateiru/cateiru-sso/api/routes"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var addr string
var port string

func init() {
	_port := os.Getenv("PORT")

	if len(_port) == 0 {
		port = ":3000"
	} else {
		port = strings.Join([]string{":", _port}, "")
	}

	addr = "0.0.0.0"
}

func main() {
	mux := http.NewServeMux()
	h2s := &http2.Server{}

	mux = routes.Routes(mux)

	server := &http.Server{
		Addr:    strings.Join([]string{addr, port}, ""),
		Handler: h2c.NewHandler(mux, h2s),
	}

	logrus.Infof("Start server! addr: %v, port: %v", addr, port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
