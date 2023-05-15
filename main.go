package main

import "github.com/cateiru/cateiru-sso/src"

// Set this variable at build time.
//
// Example:
//
//	go build  -ldflags="-X main.mode=prod"
var mode string = "local"

func main() {
	// Run backend server
	src.Main(mode)
}
