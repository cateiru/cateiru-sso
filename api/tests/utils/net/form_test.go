package net_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

type PostForm struct {
	Name string `json:"name"`
}

func formServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testPostFormHandler)

	return mux
}

// postのjson formを取得し、name要素を返す
func testPostFormHandler(w http.ResponseWriter, r *http.Request) {
	if !net.CheckContentType(r) {
		w.WriteHeader(http.StatusBadRequest)
		logging.Sugar.Error("content-type is not application/json")
		return
	}
	var postForm PostForm

	if err := net.GetJsonForm(w, r, &postForm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Sugar.Error(err.Error())
		return
	}

	w.Write([]byte(postForm.Name))
}

// POSTが正常にでき、application/jsonでjson fromを正しくパースできているか
func TestPostForm(t *testing.T) {
	app := formServer()
	server := httptest.NewServer(app)
	defer server.Close()

	sendForm := `{"name": "hoge"}`

	resp, err := http.Post(server.URL+"/", "application/json", bytes.NewBuffer([]byte(sendForm)))
	require.NoError(t, err, "postに失敗した")
	require.Equal(t, resp.StatusCode, 200, "200で返ってきてない")

	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	require.Equal(t, buf.String(), "hoge", "正しくPOST formを取得できてない")
}
