package pro_test

import "net/http"

func ssoServer() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
