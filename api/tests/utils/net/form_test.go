package net_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// POSTが正常にでき、application/jsonでjson fromを正しくパースできているか
func TestPostForm(t *testing.T) {
	app := NewTestApp()
	server := httptest.NewServer(app)
	defer server.Close()

	sendForm := `{"name": "hoge"}`

	resp, err := http.Post(server.URL+"/form", "application/json", bytes.NewBuffer([]byte(sendForm)))
	require.NoError(t, err, "postに失敗した")
	require.Equal(t, resp.StatusCode, 200, "200で返ってきてない")

	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	require.Equal(t, buf.String(), "hoge", "正しくPOST formを取得できてない")
}
