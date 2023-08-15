package main

import (
	"os"

	"github.com/cateiru/cateiru-sso/src"
)

// Set this variable at build time.
//
// Example:
//
//	go build  -ldflags="-X main.mode=prod"
var mode string = "local"

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Run backend server
	src.Main(mode, path)
}
