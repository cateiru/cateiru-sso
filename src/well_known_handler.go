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

// passkey-endpoint のレスポンス
// ref. https://github.com/ms-id-standards/MSIdentityStandardsExplainers/blob/main/PasskeyEndpointsWellKnownUrl/explainer.md
type PasskeyEndpointResponse struct {
	Enroll string `json:"enroll"`
	Mange  string `json:"manage"`
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

// FedCM の well-known レスポンス
func (h *Handler) WebIdentityHandler(c echo.Context) error {
	apiUrl := h.C.Host.String()

	providersUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}
	providersUrl.Path = "/fedcm/config.json"

	return c.JSON(http.StatusOK, &WebIdentityResponse{
		ProvidersUrl: providersUrl.String(),
	})
}

func (h *Handler) PasskeyEndpointHandler(c echo.Context) error {
	pageUrl := h.C.SiteHost.String()

	enroll, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	enroll.Path = "/settings"

	manage, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	manage.Path = "/settings"

	return c.JSON(http.StatusOK, &PasskeyEndpointResponse{
		Enroll: enroll.String(),
		Mange:  manage.String(),
	})
}

func (h *Handler) ChangePasswordHandler(c echo.Context) error {
	pageUrl := h.C.SiteHost.String()

	changePasswordUrl, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}
	changePasswordUrl.Path = "/forget_password"

	return c.Redirect(http.StatusFound, changePasswordUrl.String())
}
