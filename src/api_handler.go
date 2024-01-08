package src

import (
	"net/http"
	"net/url"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JwksResponse struct {
	Keys jwk.Key `json:"keys"`
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
		Keys: keySet,
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
