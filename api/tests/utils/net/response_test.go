package net_test

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
	"github.com/stretchr/testify/require"
)

type ResponseOK struct {
	Status string `json:"status"`
}

// ref. https://orisano.hatenablog.com/entry/2020/08/01/222730
type dummyFailedJson struct {
	Dummy float64 `json:"dummy"`
}

func responseServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ok", testResponseOKHandler)
	mux.HandleFunc("/error", testResponseErrorHandler)
	mux.HandleFunc("/notfound", testNotfoundErrorHandler)
	mux.HandleFunc("/custom", testNotFoundErrorCustomCode)
	mux.HandleFunc("/internal", testInternalError)

	return mux
}

func testResponseOKHandler(w http.ResponseWriter, r *http.Request) {
	body := ResponseOK{
		Status: "OK",
	}

	net.ResponseOK(w, body)
}

func testResponseErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Dummy error")

	net.ResponseError(w, err)
}

func testNotfoundErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := status.NewNotFoundError(errors.New("Dummy error")).Wrap()

	net.ResponseError(w, err)
}

func testNotFoundErrorCustomCode(w http.ResponseWriter, r *http.Request) {
	net.ResponseErrorCustomCode(w, 404, errors.New("Dummy error"), 1)
}

func testInternalError(w http.ResponseWriter, r *http.Request) {
	dummy := dummyFailedJson{
		Dummy: math.Inf(0),
	}
	net.ResponseCustomStatus(w, 200, dummy)
}

func TestStatusOK(t *testing.T) {
	app := responseServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/ok")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 200, "ステータスコードが200ではない")

	var body ResponseOK
	json.Unmarshal(tools.ConvertByteResp(resp), &body)

	require.Equal(t, body.Status, "OK")
}

func TestInternalServerError(t *testing.T) {
	app := responseServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/error")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 500, "ステータスコードが500ではない")

	var body net.ErrorResponse
	json.Unmarshal(tools.ConvertByteResp(resp), &body)

	require.Equal(t, body.StatusCode, 500)
	require.Equal(t, body.Code, 1)
}

func TestNotfoundError(t *testing.T) {
	app := responseServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/notfound")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 404, "ステータスコードが404ではない")

	var body net.ErrorResponse
	json.Unmarshal(tools.ConvertByteResp(resp), &body)

	require.Equal(t, body.StatusCode, 404)
	require.Equal(t, body.Code, 1)
}

func TestCustomError(t *testing.T) {
	app := responseServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/custom")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 404, "ステータスコードが404ではない")

	var body net.ErrorResponse
	json.Unmarshal(tools.ConvertByteResp(resp), &body)

	require.Equal(t, body.StatusCode, 404)
	require.Equal(t, body.Code, 1)
}

func TestInternalError(t *testing.T) {
	app := responseServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/internal")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 500)

	var body net.ErrorResponse
	json.Unmarshal(tools.ConvertByteResp(resp), &body)

	require.Equal(t, body.StatusCode, 500)
	require.Equal(t, body.Code, 2)
}
