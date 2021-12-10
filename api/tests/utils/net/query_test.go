package net_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

func queryServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testURLQuery)

	return mux
}

func testURLQuery(w http.ResponseWriter, r *http.Request) {
	query, err := net.GetQuery(r, "sample")
	if err != nil {
		net.ResponseError(w, err)
		return
	}

	w.Write([]byte(query))
}

func TestQuery(t *testing.T) {
	app := queryServer()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/?sample=hoge")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 200, "ステータスコードが200ではない")

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	require.Equal(t, buf.String(), "hoge", "正常にqueryを取得していない")
}
