package oauth

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

func JWTPublicHandler(w http.ResponseWriter, r *http.Request) error {
	public, err := GetPublicKey()
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	net.ResponseOK(w, public)

	return nil
}
