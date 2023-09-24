package lib

import (
	"strconv"
	"strings"
)

// サポートしているOIDCのスコープ
var AllowScopes = []string{
	"openid",
	"profile",
	"email",
}

type ResponseType string

const (
	// 不正な値の場合
	ResponseTypeInvalid ResponseType = ""

	// Authorization Code Flow
	ResponseTypeAuthorizationCode ResponseType = "code"

	// Implicit Flow
	ResponseTypeImplicit ResponseType = "id_token"

	// Hybrid Flow
	ResponseTypeHybrid ResponseType = "code id_token"
)

type ResponseMode string

const (
	ResponseModeQuery    ResponseMode = "query"
	ResponseModeFragment ResponseMode = "fragment"
	ResponseModeFormPost ResponseMode = "form_post"
)

type Display string

const (
	DisplayPage  Display = "page"
	DisplayPopup Display = "popup"
	DisplayTouch Display = "touch"
	DisplayWap   Display = "wap"
)

type Prompt string

const (
	PromptNone          Prompt = "none"
	PromptLogin         Prompt = "login"
	PromptConsent       Prompt = "consent"
	PromptSelectAccount Prompt = "select_account"
)

func ValidateScope(s string) bool {
	for _, allow := range AllowScopes {
		if allow == s {
			return true
		}
	}
	return false
}

func ValidateScopes(s string) ([]string, bool) {
	scopes := []string{}

	findOpenidScope := false
	for _, scope := range strings.Split(s, " ") {
		if ValidateScope(scope) {
			scopes = append(scopes, scope)
			if scope == "openid" {
				findOpenidScope = true
			}
		}
	}

	return scopes, findOpenidScope
}

func ValidateResponseType(s string) ResponseType {
	switch s {
	case "code":
		return ResponseTypeAuthorizationCode
	case "id_token", "id_token token":
		// id_token, id_token token は同じ扱いなので1つにまとめる
		return ResponseTypeImplicit
	case "code id_token", "code token", "code id_token token":
		return ResponseTypeHybrid
	}

	return ResponseTypeInvalid
}

// response_mode のバリデーション
// デフォルトは query
// https://openid.net/specs/oauth-v2-multiple-response-types-1_0.html
func ValidateResponseMode(s string) ResponseMode {
	switch s {
	case "query":
		return ResponseModeQuery
	case "fragment":
		return ResponseModeFragment
	case "form_post":
		return ResponseModeFormPost
	}

	return ResponseModeQuery
}

// 一旦page決め打ち
func ValidateDisplay(s string) Display {
	return DisplayPage
}

func ValidatePrompt(s string) Prompt {
	switch s {
	case "none":
		return PromptNone
	case "login":
		return PromptLogin
	case "consent":
		return PromptConsent
	case "select_account":
		return PromptSelectAccount
	}

	return PromptNone
}

// max-age
// OP によって明示的に認証されてからの経過時間の最大許容値 (秒)。指定しない場合や不正な値の場合は0
func ValidateMaxAge(s string) uint64 {
	if s == "" {
		return 0
	}

	n, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0
	}

	if n <= 0 {
		return 0
	}

	return n
}

func ValidateUiLocales(s string) []string {
	return []string{"ja_JP"}
}
