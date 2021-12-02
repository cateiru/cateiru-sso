// 返すレスポンスを統一します。
//
// Example:
//	err := errors.New("dummy error")
//	ResponseError(w, 500, err) // internal server error
//
//	WriteBody(w, response) // 200 OK
//
package net

import (
	"encoding/json"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/utils"
)

type AbstractResponse struct {
	// 独自ステータスコード
	//
	// 特殊な事情でエラーが起きた場合HTTP ステータスコード以外にこのコードを指定します。
	//
	//	0: 正常終了
	//	1: エラー
	//	2: ResponseError内でのエラー
	//
	Code int `json:"code"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrorID    string `json:"error_id"`

	AbstractResponse
}

// エラーをHTTPで返す
// レスポンスではerror idを返し、ログからそのIDを検索することでエラーメッセージを参照できる
func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	id := utils.CreateID(10)

	logging.Sugar.Errorf("HTTP ERROR. id: %v, message: %v", id, err.Error())

	body := ErrorResponse{
		StatusCode: statusCode,
		ErrorID:    id,
		AbstractResponse: AbstractResponse{
			Code: 1,
		},
	}

	ResponseOKCustomStatus(w, statusCode, body)
}

// カスタムに独自ステータスコードを決定し、エラーをHTTPで返す
// レスポンスではerror idを返し、ログからそのIDを検索することでエラーメッセージを参照できる
func ResponseErrorCustomCode(w http.ResponseWriter, statusCode int, err error, code int) {
	id := utils.CreateID(10)

	logging.Sugar.Errorf("HTTP ERROR. id: %v, message: %v, code: %v", id, err.Error(), code)

	body := ErrorResponse{
		StatusCode: statusCode,
		ErrorID:    id,
		AbstractResponse: AbstractResponse{
			Code: code,
		},
	}

	ResponseOKCustomStatus(w, statusCode, body)
}

// ステータスコード200で書き出す
func ResponseOK(w http.ResponseWriter, body interface{}) {
	ResponseOKCustomStatus(w, http.StatusOK, body)
}

// bodyをHTTP Responceに書き出す
func ResponseOKCustomStatus(w http.ResponseWriter, statusCode int, body interface{}) {
	bodyByte, err := json.Marshal(body)
	if err != nil {
		id := utils.CreateID(10)
		logging.Sugar.Errorf("WriteError error. id: %v error: %v", id, err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status_code": 500, "code": 2, "error_id": "` + id + `"}`))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(bodyByte)
}
