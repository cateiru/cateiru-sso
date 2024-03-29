package src

import (
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// ref. https://developers.google.com/privacy-sandbox/3pcd/fedcm-developer-guide?hl=ja
type FedCMConfigResponse struct {
	// アカウント リスト エンドポイントの URL。
	AccountsEndpoint string `json:"accounts_endpoint"`
	// クライアント メタデータ エンドポイントの URL。
	ClientMetadataEndpoint string `json:"client_metadata_endpoint,omitempty"`

	// ID アサーション エンドポイントの URL。
	IdAssertionEndpoint string `json:"id_assertion_endpoint"`

	RevocationEndpoint string `json:"revocation_endpoint,omitempty"`

	// ユーザーが IdP にログインするためのログインページの URL。
	LoginUrl  string `json:"login_url,omitempty"`
	SignInUrl string `json:"signin_url,omitempty"`

	Branding *FedCMConfigBranding `json:"branding,omitempty"`
}

// ref. https://fedidcg.github.io/FedCM/#dictdef-identityproviderbranding
type FedCMConfigBranding struct {
	BackgroundColor string             `json:"background_color,omitempty"`
	Color           string             `json:"color,omitempty"`
	Name            string             `json:"name,omitempty"`
	Icons           []FedCMConfigIcons `json:"icons,omitempty"`
}

type FedCMConfigIcons struct {
	Url  string `json:"url"`
	Size uint64 `json:"size,omitempty"`
}

// ref. https://fedidcg.github.io/FedCM/#dictdef-identityprovideraccountlist
type FedCMAccountsResponse struct {
	Accounts []FedCMAccount `json:"accounts"`
}

type FedCMClientMetadataResponse struct {
	PrivacyPolicyUrl  string `json:"privacy_policy_url,omitempty"`
	TermsOfServiceUrl string `json:"terms_of_service_url,omitempty"`
}

// ref. https://fedidcg.github.io/FedCM/#dictdef-identityprovideraccount
type FedCMAccount struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	GivenName       string   `json:"given_name,omitempty"`
	Picture         string   `json:"picture,omitempty"`
	ApprovedClients []string `json:"approved_clients,omitempty"`
	LoginHints      []string `json:"login_hints,omitempty"`
	DomainHints     []string `json:"domain_hints,omitempty"`
}

type FedCMIdAssertionResponse struct {
	Token string `json:"token"`
}

// FedCM の設定レスポンス
func (h *Handler) FedCMConfigHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &FedCMConfigResponse{
		AccountsEndpoint: "/fedcm/accounts",

		ClientMetadataEndpoint: "/fedcm/client_metadata",

		IdAssertionEndpoint: "/fedcm/id_assertion",

		LoginUrl:  "/fedcm/login",
		SignInUrl: "/fedcm/login",

		RevocationEndpoint: "/fedcm/revocation", // TODO: 未実装

		Branding: &FedCMConfigBranding{
			BackgroundColor: h.C.BrandColor,
			Color:           h.C.BrandTextColor,
			Name:            h.C.BrandName,
			// TODO: アイコン埋める
		},
	})
}

// サインイン（ログイン）URLにリダイレクトする
func (h *Handler) FedCMSignInHandler(c echo.Context) error {
	pageUrl := h.C.SiteHost.String()

	loginUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	loginUrl.Path = "/login"

	return c.Redirect(http.StatusFound, loginUrl.String())
}

// FedCM のログイン可能なアカウントリストを取得する
func (h *Handler) FedCMAccountsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.Session.LoggedInAccounts(ctx, c.Cookies())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return NewHTTPError(http.StatusUnauthorized, "no logged in accounts")
	}

	accounts := make([]FedCMAccount, len(users))

	for i, user := range users {
		accounts[i] = FedCMAccount{
			ID:        user.ID,
			Name:      user.UserName,
			Email:     user.Email,
			GivenName: user.GivenName.String,
			Picture:   user.Avatar.String,
		}
	}

	return c.JSON(http.StatusOK, &FedCMAccountsResponse{
		Accounts: accounts,
	})
}

// FedCM のクライアントメタデータを返す
func (h *Handler) FedCMClientMetadataHandler(c echo.Context) error {
	pageUrl := h.C.SiteHost.String()

	privacyPolicyUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	privacyPolicyUrl.Path = "/policy" // TODO: プライバシーポリシーのページ作ったら見直す

	termsUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	termsUrl.Path = "/terms" // TODO: 利用規約のページ作ったら見直す

	return c.JSON(http.StatusOK, &FedCMClientMetadataResponse{
		PrivacyPolicyUrl:  privacyPolicyUrl.String(),
		TermsOfServiceUrl: termsUrl.String(),
	})
}

// FedCM の認証
// 返すtokenは一旦OIDCの code にする
// TODO: https://developers.google.com/privacy-sandbox/3pcd/fedcm-developer-guide?hl=ja#error-response に従ってエラーレスポンスを返す
func (h *Handler) FedCMIdAssertionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientId := c.FormValue("client_id")
	nonce := c.FormValue("nonce")
	userId := c.FormValue("account_id")

	// 一旦使わないのでコメントアウト
	// disclosureTextShown := c.FormValue("disclosure_text_shown")

	if userId == "" {
		return NewHTTPError(http.StatusBadRequest, "account_id is required")
	}
	if clientId == "" {
		return NewHTTPError(http.StatusBadRequest, "client_id is required")
	}

	user, err := h.Session.GetUserFromUserIDAndCookie(ctx, c.Cookies(), userId)
	if err != nil {
		return err
	}

	// クライアントの存在チェック
	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewHTTPError(http.StatusBadRequest, "client_id is invalid")
	}
	if err != nil {
		return err
	}

	code, err := lib.RandomStr(63)
	if err != nil {
		return err
	}

	oauthSession := models.OauthSession{
		Code:   code,
		UserID: user.ID,

		ClientID: client.ClientID,

		Nonce:    null.NewString(nonce, nonce != ""),
		AuthTime: time.Now(),

		Period: time.Now().Add(h.C.OauthLoginSessionPeriod),
	}
	if err := oauthSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
		return err
	}

	if err := h.SaveOperationHistory(ctx, c, user, 30); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &FedCMIdAssertionResponse{
		Token: code,
	})
}
