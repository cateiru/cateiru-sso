package net_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	app := NewTestApp()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(server.URL + "/query?sample=hoge")
	require.NoError(t, err, "GETできない")
	require.Equal(t, resp.StatusCode, 200, "ステータスコードが200ではない")

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	require.Equal(t, buf.String(), "hoge", "正常にqueryを取得していない")
}
