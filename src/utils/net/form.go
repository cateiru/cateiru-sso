// POSTフォームの操作
//
// Example:
//	type Json struct {
//		Name string `json:"name"`
//	}
//
//	if CheckContentType(r) {
//		var json Json
//		if err := GetJsonForm(w, r, json); err != nil {
//			panic(err)
//		}
//		...
//	}
//
package net

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// HTTPヘッダのContent-Typeが`application/json`であるかをチェックする
func CheckContentType(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

// postのformのjsonをパースして内容を取得します
func GetJsonForm(w http.ResponseWriter, r *http.Request, obj interface{}) error {
	// To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	// Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return err
	}

	err = json.Unmarshal(body[:length], obj)
	if err != nil {
		return err
	}

	return nil
}
