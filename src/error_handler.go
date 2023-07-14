package src

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// カスタムユニークコード
const (
	ErrUniqueDefault = 0
	// reCAPTCHAに失敗した（BOT）場合
	ErrReCaptcha = 1
	// 該当Emailのアカウント登録のセッションが存在する
	ErrSessionExists = 2
	// 何らかの原因（メールアドレスがすでに登録されているなど）でアカウントが作成できない
	ErrImpossibleRegisterAccount = 3
	// リトライ回数を超えた
	ErrExceededRetry = 4
	// 有効期限切れ
	ErrExpired = 5
	// メール送信上限
	ErrEmailSendingLimit = 6
	// メール未認証
	ErrEmailNotVerified = 7
	// ログイン失敗
	ErrLoginFailed = 8
	// ログインに失敗したが別のアカウントでログインできる可能性がある場合
	ErrBeAbleToLoginWithAnotherAccount = 9
	// ユーザがいない
	ErrNotFoundUser = 10
	// パスワード登録していない
	ErrNoRegisteredPassword = 11
	// ユーザがすでに存在している
	ErrAlreadyExistUser = 12
	// 認証失敗
	ErrAuthenticationFailed = 13
	// 認証情報が無くなるため削除できない
	// 主にWebAuthnの削除で使用される
	ErrNoMoreAuthentication = 14
	// すでにログイン済み
	ErrAlreadyLoggedIn = 15
	// orgに入っていない
	ErrNoJoinedOrg = 16
)

// OIDCのエラーコード
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthError
const (
	// Authorization Server は処理を進めるためにいくつかの End-User interaction を必要とする.
	// Authentication Request 中の prompt パラメータが none であるが, End-User interaction のためのユーザーインターフェースの表示なしには Authentication Request が完了できない時にこのエラーが返される
	ErrInteractionRequired = "interaction_required"

	// Authorization Server は End-User の認証を必要とする. Authentication Request 中の prompt パラメータが none であるが,
	// End-User の認証のためのユーザーインターフェースの表示なしには Authentication Request が完了できない時にこのエラーが返される.
	ErrLoginRequired = "login_required"

	// End-User は Authorization Server にてセッションの選択を必要とされる (REQUIRED).
	// End-User は Authorization Server にて異なるアカウントで認証されているが, セッションを選択していないかもしれない (MAY).
	// Authentication Request 中の prompt パラメータが none であるが, 利用するセッションを選択するためのユーザーインターフェースの表示なしには Authentication Request が完了できない時にこのエラーが返される.
	ErrAccountSelectionRequired = "account_selection_required"

	// Authorization Server は End-User の同意を必要とする. Authentication Request 中の prompt パラメータが none であるが,
	// End-User の同意のためのユーザーインターフェースの表示なしには Authentication Request が完了できない時にこのエラーが返される.
	ErrConsentRequired = "consent_required"

	// Authorization Request 中の request_uri はエラーを返すか, 無効なデータを含む.
	ErrInvalidRequestURI = "invalid_request_uri"

	// request パラメータが無効な Request Object を含む.
	ErrInvalidRequestObject = "invalid_request_object"

	// OP は Section 6 にて定義されている request パラメータをサポートしていない.
	ErrRequestNotSupported = "request_not_supported"

	// OP は Section 6 にて定義されている request_uri パラメータをサポートしていない.
	ErrRequestURINotSupported = "request_uri_not_supported"

	// OP は Section 7.2.1 で定義されている registration パラメータをサポートしていない.
	ErrRegistrationNotSupported = "registration_not_supported"
)

type HTTPError struct {
	Code       int         `json:"-"`
	Message    interface{} `json:"message"`
	UniqueCode int         `json:"unique_code"`

	File string `json:"-"`
	Line int    `json:"-"`
}

type OIDCError struct {
	Code int `json:"-"`

	File string `json:"-"`
	Line int    `json:"-"`

	AuthenticationErrorResponse
}

func NewHTTPError(code int, message ...any) *HTTPError {
	_, file, line, _ := runtime.Caller(1)

	he := &HTTPError{
		Code:       code,
		UniqueCode: ErrUniqueDefault,
		Message:    http.StatusText(code),

		File: file,
		Line: line,
	}

	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func NewHTTPUniqueError(code int, unique int, message ...any) *HTTPError {
	_, file, line, _ := runtime.Caller(1)

	he := &HTTPError{
		Code:       code,
		UniqueCode: unique,
		Message:    http.StatusText(code),

		File: file,
		Line: line,
	}

	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func NewOIDCError(code string, message string, uri string, state string) *OIDCError {
	_, file, line, _ := runtime.Caller(1)

	oe := &OIDCError{
		Code: 400, // HTTPステータスコードは決め打ちで400で返す

		File: file,
		Line: line,

		AuthenticationErrorResponse: AuthenticationErrorResponse{
			Error:            code,
			ErrorDescription: message,
			ErrorURI:         uri,
			State:            state,
		},
	}

	return oe
}

func (he *HTTPError) Error() string {
	m := fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
	if he.UniqueCode != ErrUniqueDefault {
		m = fmt.Sprintf("%s, unique=%d", m, he.UniqueCode)
	}
	return m
}

func (oe *OIDCError) Error() string {
	m := fmt.Sprintf("code=%d, error=%v", oe.Code, oe.AuthenticationErrorResponse.Error)
	if oe.AuthenticationErrorResponse.ErrorURI != "" {
		m += fmt.Sprintf(", error_uri=%v", oe.AuthenticationErrorResponse.ErrorURI)
	}
	if oe.AuthenticationErrorResponse.State != "" {
		m += fmt.Sprintf(", state=%v", oe.AuthenticationErrorResponse.State)
	}
	return m
}

// カスタムエラーハンドラー
// 基本エラー時には、{"message": "error message"}のjsonを返す
// unique_codeもあれば表記
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	// `HTTPError` の場合
	he, ok := err.(*HTTPError)
	if ok {
		code = he.Code
		c.JSON(code, he)
		return
	}

	// `OIDCError` の場合
	oe, ok := err.(*OIDCError)
	if ok {
		code = oe.Code
		c.JSON(code, oe)
		return
	}

	// `echo.HTTPError` の場合
	echohe, eok := err.(*echo.HTTPError)
	if eok {
		code = echohe.Code
		c.JSON(code, echohe)
		return
	}

	c.String(code, err.Error())
}

func ErrorLog(v middleware.RequestLoggerValues) error {
	code := http.StatusInternalServerError
	line := 0
	file := ""

	he, ok := v.Error.(*HTTPError)
	if ok {
		line = he.Line
		file = he.File
		code = he.Code
	}

	oe, ok := v.Error.(*OIDCError)
	if ok {
		line = oe.Line
		file = oe.File
		code = oe.Code
	}

	echohe, ok := v.Error.(*echo.HTTPError)
	if ok {
		code = echohe.Code
	}

	// エラーコードが400番台の場合はInfo
	if code >= 400 && code < 500 {
		L.Info("request",
			zap.String("URI", v.URI),
			zap.String("method", v.Method),
			zap.Int("status", code),
			zap.String("host", v.Host),
			zap.String("response_time", time.Since(v.StartTime).String()),
			zap.String("ip", v.RemoteIP),
			zap.String("file", file),
			zap.Int("line", line),
			zap.String("error_message", v.Error.Error()),
		)
		return nil
	}
	L.Error("request",
		zap.String("URI", v.URI),
		zap.String("method", v.Method),
		zap.Int("status", code),
		zap.String("host", v.Host),
		zap.String("response_time", time.Since(v.StartTime).String()),
		zap.String("ip", v.RemoteIP),
		zap.String("file", file),
		zap.Int("line", line),
		zap.String("error_message", v.Error.Error()),
	)

	return nil
}
