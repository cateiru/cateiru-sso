package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type PreviewResponse struct {
	ClientId          string      `json:"client_id"`
	ClientName        string      `json:"client_name"`
	ClientDescription null.String `json:"client_description"`
	Image             null.String `json:"image"`

	OrgName       null.String `json:"org_name"`
	OrgImage      null.String `json:"org_image"`
	OrgMemberOnly bool        `json:"org_member_only"`

	Scopes       []string `json:"scopes"`
	RedirectUri  string   `json:"redirect_uri"`
	ResponseType string   `json:"response_type"`

	RegisterUserName  string      `json:"register_user_name"`
	RegisterUserImage null.String `json:"register_user_image"`

	Prompts []lib.Prompt `json:"prompts"`

	LoginSession *LoginSession `json:"login_session,omitempty"`
}

type LoginSession struct {
	LoginSessionToken string    `json:"login_session_token"`
	LimitDate         time.Time `json:"limit_date"`
}

type OauthResponse struct {
	RedirectUrl string `json:"redirect_url"`
	Code        string `json:"code"`
}

type OidcRequest interface {
	// プレビュー用のレスポンスを返す
	GetPreviewResponse(ctx context.Context, loginSessionPeriod time.Duration, db *sql.DB, sessionToken string) (*PreviewResponse, error)

	// ログインが必要な場合のセッションを返す
	GetLoginSession(ctx context.Context, period time.Duration, db *sql.DB) (*LoginSession, error)

	// ユーザーが認証可能かチェックする
	CheckUserAuthenticationPossible(ctx context.Context, db *sql.DB, user *models.User) (bool, error)

	Submit(ctx context.Context, db *sql.DB) (*OauthResponse, error)
	Cancel(ctx context.Context, db *sql.DB) (*OauthResponse, error)
}

// OIDC の Request を取得する
// RFCではGETかPOSTでx-www-form-urlencodedでリクエストを送るとあるが、
// これはjs側で対応するのでjs - サーバ間は multipart/form-data で送る
func (h *Handler) NewOidcRequest(ctx context.Context, c echo.Context) (OidcRequest, error) {
	// Client ID
	clientId := c.FormValue("client_id")
	if clientId == "" {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "client_id is required", "", "")
	}
	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "client_id is invalid", "", "")
	}
	if err != nil {
		return nil, err
	}

	// State
	state := c.FormValue("state")

	// Response Mode
	responseMode := c.FormValue("response_mode")
	validatedResponseMode := lib.ValidateResponseMode(responseMode)

	// Nonce
	nonce := c.FormValue("nonce")

	// Display
	display := c.FormValue("display")
	validatedDisplay := lib.ValidateDisplay(display)

	// Prompt
	prompt := c.FormValue("prompt")
	validatedPrompt := []lib.Prompt{}
	for _, p := range strings.Split(prompt, " ") {
		validatedPrompt = append(validatedPrompt, lib.ValidatePrompt(p))
	}

	// MaxAge
	maxAge := c.FormValue("max_age")
	validatedMaxAge := lib.ValidateMaxAge(maxAge)

	// UI Locales
	uiLocales := c.FormValue("ui_locales")
	validateUiLocales := lib.ValidateUiLocales(uiLocales)

	// ID Token Hint
	idTokenHint := c.FormValue("id_token_hint")

	// Login Hint
	loginHint := c.FormValue("login_hint")

	// ACR Values
	acrValues := c.FormValue("acr_values")

	// Scope
	scope := c.FormValue("scope")
	if scope == "" {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "scope is required", "", "")
	}
	validatedScopes, scopesOk := lib.ValidateScopes(scope)
	if !scopesOk {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "scope is invalid", "", "")
	}

	clientScopes, err := models.ClientScopes(
		models.ClientScopeWhere.ClientID.EQ(client.ClientID),
		models.ClientScopeWhere.Scope.IN(validatedScopes),
	).All(ctx, h.DB)
	if err != nil {
		return nil, err
	}

	enableScopes := []string{}
	for _, clientScope := range clientScopes {
		enableScopes = append(enableScopes, clientScope.Scope)
	}

	// Redirect URI
	redirectUri := c.FormValue("redirect_uri")
	if redirectUri == "" {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "redirect_uri is required", "", "")
	}
	parsedRedirectUri, redirectUriOk := lib.ValidateURL(redirectUri)
	if !redirectUriOk {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "redirect_uri is invalid", "", "")
	}

	existRedirect, err := models.ClientRedirects(
		models.ClientRedirectWhere.ClientID.EQ(client.ClientID),
		models.ClientRedirectWhere.URL.EQ(redirectUri),
	).Exists(ctx, h.DB)
	if err != nil {
		return nil, err
	}
	if !existRedirect {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "redirect_uri is invalid", "", "")
	}

	allowRules := []*models.ClientAllowRule{}

	if client.IsAllow {
		allowRules, err = models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(client.ClientID),
		).All(ctx, h.DB)
		if err != nil {
			return nil, err
		}
	}

	refererUrl, err := url.Parse(c.Request().Referer())
	if err != nil {
		return nil, err
	}

	dbReferer, err := models.ClientReferrers(
		models.ClientReferrerWhere.ClientID.EQ(client.ClientID),
	).All(ctx, h.DB)
	if err != nil {
		return nil, err
	}
	referrerHost := ""
	// リファラーが設定されている場合はチェックする
	// ホストのみのチェック
	if len(dbReferer) != 0 {
		ok := false
		for _, referrer := range dbReferer {
			if referrer.Host == refererUrl.Host {
				referrerHost = referrer.Host
				ok = true
				break
			}
		}
		if !ok {
			return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "referer is invalid", "", "")
		}
	}

	// Response Type
	responseType := c.FormValue("response_type")
	validatedResponseType := lib.ValidateResponseType(responseType)
	if validatedResponseType == lib.ResponseTypeInvalid {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "request_type is invalid", "", "")
	}

	// response_type によって分ける
	switch validatedResponseType {
	case lib.ResponseTypeAuthorizationCode:
		return &AuthenticationRequest{
			Scopes:       enableScopes,
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri:  parsedRedirectUri,
			State:        null.NewString(state, state != ""),
			ResponseMode: validatedResponseMode,
			Nonce:        null.NewString(nonce, nonce != ""),
			Display:      validatedDisplay,
			Prompts:      validatedPrompt,
			MaxAge:       validatedMaxAge,
			UiLocales:    validateUiLocales,

			IdTokenHint: null.NewString(idTokenHint, idTokenHint != ""),
			LoginHint:   null.NewString(loginHint, loginHint != ""),
			AcrValues:   null.NewString(acrValues, acrValues != ""),

			Client:      client,
			AllowRules:  allowRules,
			RefererHost: referrerHost,
		}, nil
	default:
		// 一旦 Authorization Code Flow のみ対応
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "request_type is invalid", "", "")
	}
}
