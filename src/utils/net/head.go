package net

import "net/http"

func GetIPAddress(r *http.Request) string {
	ip := r.Header.Get("x-forwarded-for")
	if ip != "" {
		return ip
	}
	return r.RemoteAddr
}

func GetUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}
