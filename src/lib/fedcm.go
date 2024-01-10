package lib

import (
	"net/http"
)

func CheckFedCMHeaders(headers http.Header) bool {
	secFetchDest := headers.Get("Sec-Fetch-Dest")

	return secFetchDest == "webidentity"
}
