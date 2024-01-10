package src

import (
	"net/http"
	"net/url"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JwksResponse struct {
	Keys []jwk.Key `json:"keys"`
}

// FedCM の well-known レスポンス
// ref. https://developer.mozilla.org/en-US/docs/Web/API/FedCM_API#provide_a_well-known_file
type WebIdentityResponse struct {
	ProvidersUrl string `json:"provider_urls"`
}

// ref. https://fedidcg.github.io/FedCM/#dictdef-identityproviderapiconfig
type FedCMConfigResponse struct {
	AccountsEndpoint       string               `json:"accounts_endpoint"`
	ClientMetadataEndpoint string               `json:"client_metadata_endpoint"`
	IdAssertionEndpoint    string               `json:"id_assertion_endpoint"`
	DisconnectEndpoint     string               `json:"disconnect_endpoint,omitempty"`
	Branding               *FedCMConfigBranding `json:"branding,omitempty"`
}

// ref. https://fedidcg.github.io/FedCM/#dictdef-identityproviderbranding
type FedCMConfigBranding struct {
	BackgroundColor string             `json:"background_color,omitempty"`
	Color           string             `json:"color,omitempty"`
	Name            string             `json:"name,omitempty"`
	Icons           []FedCMConfigIcons `json:"icons"`
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

// OpenID Connect Discovery 1.0 incorporating errata set 1 で定義されている、 `.well-known/openid-configuration` のエンドポイント
// ref. https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationRequest
func (h *Handler) ApiOpenidConfigurationHandler(c echo.Context) error {

	apiUrl := h.C.Host.String()
	pageUrl := h.C.SiteHost.String()

	authorizationEndpointUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	authorizationEndpointUrl.Path = "/oidc"

	tokenEndpointUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}
	tokenEndpointUrl.Path = "/v2/token"

	userinfoEndpointUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}
	userinfoEndpointUrl.Path = "/v2/userinfo"

	jwksUri, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	jwksUri.Path = "/.well-known/jwks.json"

	registrationEndpointUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}
	registrationEndpointUrl.Path = "/v2/register"

	policyUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	policyUrl.Path = "/policy"

	configuration := OpenidConfiguration{
		Issuer:                pageUrl,
		AuthorizationEndpoint: authorizationEndpointUrl.String(),
		TokenEndpoint:         tokenEndpointUrl.String(),
		UserinfoEndpoint:      userinfoEndpointUrl.String(),
		JwksUri:               jwksUri.String(),
		RegistrationEndpoint:  registrationEndpointUrl.String(),
		ScopesSupported:       lib.AllowScopes,
		// 一旦、Authorization Code Flow のみ対応
		ResponseTypesSupported: []string{
			"code",
		},
		// TODO
		AcrValuesSupported: []string{},
		// XXX: これでいいのか？
		SubjectTypesSupported: []string{
			"public",
			"pairwise",
		},
		// TODO: 後で検討
		IdTokenSigningAlgValuesSupported: []string{
			"RS256",
		},
		TokenEndpointAuthMethodsSupported: []string{
			"client_secret_basic",
			"client_secret_post",
		},
		ClaimsSupported: []string{
			"iss",
			"sub",
			"aud",
			"exp",
			"iat",
			"auth_time",
			"nonce",
			"acr",
			"azp",

			"name",
			"given_name",
			"family_name",
			"middle_name",
			"nickname",
			"preferred_username",
			"picture",
			"email",
			"email_verified",
			"gender",
			"birthdate",
			"zoneinfo",
			"locale",
			"updated_at",
		},
		ServiceDocumentation: "https://github.com/cateiru/cateiru-sso",
		ClaimsLocalesSupported: []string{
			"ja-JP",
		},
		UiLocalesSupported: []string{
			"ja-JP",
		},
		OpPolicyUri: policyUrl.String(),
		OpTosUri:    policyUrl.String(),
	}

	return c.JSON(200, configuration)
}

// JSON Web Key Set Endpoint
// ref. https://openid-foundation-japan.github.io/rfc7517.ja.html#JWKSetParamReg
func (h *Handler) JwksJsonHandler(c echo.Context) error {
	keySet, err := lib.JsonWebKeys(h.C.JWTPublicKeyFilePath, "RS256", "sig", h.C.JWTKid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &JwksResponse{
		Keys: []jwk.Key{
			keySet,
		},
	})
}

// TODO: テスト
// OIDC Token Endpoint
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenEndpoint
func (h *Handler) TokenEndpointHandler(c echo.Context) error {
	ctx := c.Request().Context()

	// レスポンスのキャッシュを無効化
	// ref. https://openid-foundation-japan.github.io/rfc6749.ja.html#token-response
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Pragma", "no-cache")

	// 認証
	client, err := h.ClientAuthentication(ctx, c)
	if err != nil {
		return err
	}

	param, err := h.QueryBodyParam(c)
	if err != nil {
		return err
	}

	grantType := param.Get("grant_type")
	formattedGrantType := lib.ValidateTokenEndpointGrantType(grantType)

	switch formattedGrantType {
	case lib.TokenEndpointGrantTypeAuthorizationCode:
		return h.TokenEndpointAuthorizationCode(ctx, c, client)

	case lib.TokenEndpointGrantTypeRefreshToken:
		return h.TokenEndpointRefreshToken(ctx, c, client)

	default:
		return NewOIDCError(http.StatusBadRequest, ErrTokenUnsupportedGrantType, "unsupported grant type", "", "")

	}
}

// TODO: テスト
// OIDC Userinfo Endpoint
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#UserInfoRequest
func (h *Handler) UserinfoEndpointHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientSession, err := h.UserinfoAuthentication(ctx, c)
	if err != nil {
		return err
	}

	response, err := h.ResponseStandardClaims(ctx, clientSession.ClientID, clientSession.UserID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// FedCM の well-known レスポンス
func (h *Handler) WebIdentityHandler(c echo.Context) error {
	// キャッシュしたいのでサイトURL
	pageUrl := h.C.SiteHost.String()

	providersUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}

	providersUrl.Path = "/api/fedcm/config.json"

	return c.JSON(http.StatusOK, &WebIdentityResponse{
		ProvidersUrl: providersUrl.String(),
	})
}

// FedCM の設定レスポンス
func (h *Handler) FedCMConfigHandler(c echo.Context) error {
	pageUrl := h.C.SiteHost.String()
	apiUrl := h.C.Host.String()

	accountsEndpointUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}
	accountsEndpointUrl.Path = "/v2/fedcm/accounts_list"

	// メタデータはキャッシュしたいのでページURL
	clientMetadataEndpointUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	clientMetadataEndpointUrl.Path = "/api/fedcm/client_metadata"

	idAssertionEndpointUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}
	idAssertionEndpointUrl.Path = "/v2/fedcm/id_assertion"

	return c.JSON(http.StatusOK, &FedCMConfigResponse{
		AccountsEndpoint:       accountsEndpointUrl.String(),
		ClientMetadataEndpoint: clientMetadataEndpointUrl.String(),
		IdAssertionEndpoint:    idAssertionEndpointUrl.String(),
		Branding: &FedCMConfigBranding{
			BackgroundColor: h.C.BrandBackgroundColor,
			Color:           h.C.BrandColor,
			Name:            h.C.BrandName,
			Icons:           []FedCMConfigIcons{}, // TODO: アイコン埋める
		},
	})
}

// FedCM のログイン可能なアカウントリストを取得する
func (h *Handler) FedCMAccountsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.Session.LoggedInAccounts(ctx, c.Cookies())
	if err != nil {
		return err
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
