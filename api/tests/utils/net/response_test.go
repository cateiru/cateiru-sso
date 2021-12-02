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

func TestStatusOK(t *testing.T) {
	app := NewTestApp()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/ok")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 200, "ステータスコードが200ではない")

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	var body ResponseOK
	json.Unmarshal(buf.Bytes(), &body)

	require.Equal(t, body.Status, "OK")
}

func TestStatusError(t *testing.T) {
	app := NewTestApp()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/error")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 500, "ステータスコードが500ではない")

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	var body net.ErrorResponse
	json.Unmarshal(buf.Bytes(), &body)

	require.Equal(t, body.StatusCode, 500)
	require.Equal(t, body.Code, 1)
}
