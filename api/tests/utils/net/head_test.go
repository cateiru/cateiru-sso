package net_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

type HeadResponse struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

func headServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HeaderHandler)

	return mux
}

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	ip := net.GetIPAddress(r)
	userAgent := net.GetUserAgent(r)

	body := HeadResponse{
		Ip:        ip,
		UserAgent: userAgent,
	}
	bodyJ, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bodyJ)
}

func TestHead(t *testing.T) {
	app := headServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	var body HeadResponse
	err = json.Unmarshal(buf.Bytes(), &body)
	require.NoError(t, err)

	t.Logf("IP: %s, User-Agent: %s", body.Ip, body.UserAgent)

	require.NotEmpty(t, body.Ip)
	require.NotEmpty(t, body.UserAgent)
}
