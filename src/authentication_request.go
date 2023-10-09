package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

// https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthRequest
type AuthenticationRequest struct {
	Scopes []string

	// コードフローを決定する値
	// Authorization Code Flow の場合は `code`
	ResponseType lib.ResponseType

	// レスポンスが返される Redirection URI
	RedirectUri *url.URL

	// リクエストとコールバックの間で維持されるランダムな値
	// SPAなのでjs側で保持する予定。そのためこの値はサーバー側で使う予定はない
	State null.String

	// パラメータを返す手段
	ResponseMode lib.ResponseMode

	// Client セッションと ID Token を紐づける文字列であり、リプレイアタック対策に用いられる
	Nonce null.String

	// Authorization Server が認証および同意のためのユーザーインタフェースを End-User にどのように表示するかを指定するための ASCII 値
	// Authorization Server は User Agent の機能を検知して適切な表示を行うようにしても良い
	//
	// - page: Authorization Server は認証および同意 UI を User Agent の全画面に表示すべきである (SHOULD). display パラメータが指定されていない場合, この値がデフォルトとなる
	// - popup: Authorization Server は認証および同意 UI を User Agent のポップアップウィンドウに表示すべきである (SHOULD). User Agent のポップアップウィンドウはログインダイアログに適切なサイズで, 親ウィンドウ全体を覆うことのないようにすべきである
	// - touch: Authorization Server は認証および同意 UI をタッチインタフェースを持つデバイスに適した形で表示すべきである (SHOULD)
	// - wap: Authorization Server は認証および同意 UI を "feature phone" に適した形で表示すべきである (SHOULD)
	Display lib.Display

	// Authorization Server が End-User に再認証および同意を再度要求するかどうか指定するための, スペース区切りの ASCII 文字列のリスト. 以下の値が定義されている
	// - none: Authorization Server はいかなる認証および同意 UI をも表示してはならない
	// - login: Authorization Server は End-User を再認証するべきである
	// - consent: Authorization Server は Client にレスポンスを返す前に End-User に同意を要求するべきである
	// - select_account: Authorization Server は End-User にアカウント選択を促すべきである
	Prompts []lib.Prompt

	// Authentication Age の最大値. End-User が OP によって明示的に認証されてからの経過時間の最大許容値 (秒)
	MaxAge uint64

	// ロケール。一旦これはja_JPのみを想定するが、将来的には他の言語も対応すると想定してサーバにも持ってくる
	UiLocales   []string
	IdTokenHint null.String
	LoginHint   null.String
	AcrValues   null.String

	Client *models.Client

	AllowRules []*models.ClientAllowRule
}

// preview で返す用のもの
type PublicAuthenticationRequest struct {
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
}

// プレビュー用のレスポンスを返す
func (a *AuthenticationRequest) GetPreviewResponse(ctx context.Context, db *sql.DB) (*PublicAuthenticationRequest, error) {

	orgName := null.NewString("", false)
	orgImage := null.NewString("", false)

	if a.Client.OrgID.Valid {
		// orgは見つからないことはないはずなので、見つからなかったら500エラーにする
		org, err := models.Organizations(
			models.OrganizationWhere.ID.EQ(a.Client.OrgID.String),
		).One(ctx, db)
		if err != nil {
			return nil, err
		}
		orgName = null.NewString(org.Name, true)
		orgImage = org.Image
	}

	// userは見つからないことはないはずなので、見つからなかったら500エラーにする
	user, err := models.Users(
		models.UserWhere.ID.EQ(a.Client.OwnerUserID),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	return &PublicAuthenticationRequest{
		ClientId:          a.Client.ClientID,
		ClientName:        a.Client.Name,
		ClientDescription: a.Client.Description,
		Image:             a.Client.Image,

		OrgName:       orgName,
		OrgImage:      orgImage,
		OrgMemberOnly: a.Client.OrgMemberOnly,

		Scopes:       a.Scopes,
		RedirectUri:  a.RedirectUri.String(),
		ResponseType: string(a.ResponseType),

		RegisterUserName:  user.UserName,
		RegisterUserImage: user.Avatar,
	}, nil
}

// ユーザーが認証可能かチェックする
func (a *AuthenticationRequest) CheckUserAuthenticationPossible(ctx context.Context, db *sql.DB, user *models.User) (bool, error) {
	// ルールが存在しない場合はすべてが認証可能
	if len(a.AllowRules) == 0 {
		return true, nil
	}

	for _, rule := range a.AllowRules {
		// ユーザーが一致している場合
		if rule.UserID.Valid && rule.UserID.String == user.ID {
			return true, nil
		}

		// メールドメインが後方一致している場合
		if rule.EmailDomain.Valid && strings.HasSuffix(user.Email, rule.EmailDomain.String) {
			return true, nil
		}
	}

	return false, nil
}

// Authentication Request を取得する
// RFCではGETかPOSTでx-www-form-urlencodedでリクエストを送るとあるが、
// これはjs側で対応するのでjs - サーバ間は multipart/form-data で送る
func (h *Handler) NewAuthenticationRequest(ctx context.Context, c echo.Context) (*AuthenticationRequest, error) {

	// Request Type
	requestType := c.FormValue("response_type")
	validatedResponseType := lib.ValidateResponseType(requestType)
	if validatedResponseType == lib.ResponseTypeInvalid {
		return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "request_type is invalid", "", "")
	}

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
	// リファラーが設定されている場合はチェックする
	// ホストのみのチェック
	if len(dbReferer) != 0 {
		ok := false
		for _, referrer := range dbReferer {
			if referrer.Host == refererUrl.Host {
				ok = true
				break
			}
		}
		if !ok {
			return nil, NewOIDCError(http.StatusBadRequest, ErrInvalidRequestURI, "referer is invalid", "", "")
		}
	}

	return &AuthenticationRequest{
		Scopes:       enableScopes,
		ResponseType: validatedResponseType,
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

		Client:     client,
		AllowRules: allowRules,
	}, nil
}