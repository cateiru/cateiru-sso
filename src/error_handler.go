package src

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// カスタムユニークコード
const (
	ErrUniqueDefault = 0
	ErrReCaptcha     = 1
)

type HTTPError struct {
	Code       int         `json:"-"`
	Message    interface{} `json:"message"`
	UniqueCode int         `json:"unique_code"`
	Internal   error       `json:"-"` // Stores the error returned by an external dependency

}

func NewHTTPError(code int, message ...any) *HTTPError {
	he := &HTTPError{Code: code, UniqueCode: ErrUniqueDefault, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func NewHTTPUniqueError(code int, unique int, message ...any) *HTTPError {
	he := &HTTPError{Code: code, UniqueCode: unique, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func (he *HTTPError) Error() string {
	m := fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
	if he.UniqueCode != ErrUniqueDefault {
		m = fmt.Sprintf("%s, unique=%d", m, he.UniqueCode)
	}

	if he.Internal == nil {
		return m
	}
	return fmt.Sprintf("%s, internal=%v", m, he.Internal)
}

// func (he *HTTPError) SetInternal(err error) *HTTPError {
// 	he.Internal = err
// 	return he
// }

// func (he *HTTPError) WithInternal(err error) *HTTPError {
// 	return &HTTPError{
// 		Code:     he.Code,
// 		Message:  he.Message,
// 		Internal: err,
// 	}
// }

// func (he *HTTPError) Unwrap() error {
// 	return he.Internal
// }

// カスタムエラーハンドラー
// 基本エラー時には、{"message": "error message"}のjsonを返す
// unique_codeもあれば表記
func CustomHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	code := http.StatusInternalServerError
	he, ok := err.(*HTTPError)
	if ok {
		code = he.Code
		c.JSON(code, he)
		return
	}
	echohe, eok := err.(*echo.HTTPError)
	if eok {
		code = echohe.Code
		c.JSON(code, echohe)
		return
	}
	c.String(code, err.Error())
}
