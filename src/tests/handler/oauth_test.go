package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/oauth"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func oauthServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/perview", handler.OAuthPreview)
	mux.HandleFunc("/login", handler.OAuthLogin)

	mux.HandleFunc("/token", handler.OAuthToken)
	mux.HandleFunc("/jwt", handler.OAuthJWTKey)

	return mux
}

func TestPerview(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	resp := s.Post(t, "/perview", req)

	var respBody oauth.ResponsePerview
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.Equal(t, respBody.Name, "test")
	require.Equal(t, respBody.ServiceIcon, "image")
}

func TestPerviewAllowRole(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		AllowRoles: []string{"test"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser().AddRole("test")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	resp := s.Post(t, "/perview", req)

	var respBody oauth.ResponsePerview
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.Equal(t, respBody.Name, "test")
	require.Equal(t, respBody.ServiceIcon, "image")
}

func TestPerviewAllowRoleFailed(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		AllowRoles: []string{"test"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser().AddRole("hoge") // 違うrole
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	reqForm, err := json.Marshal(req)
	require.NoError(t, err)

	resp, err := s.Client.Post(s.Server.URL+"/perview", "application/json", bytes.NewBuffer(reqForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}

func TestOauthLogin(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	resp := s.Post(t, "/login", req)

	var respBody oauth.LoginResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.NotEmpty(t, respBody.AccessToken)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAccessTokenByAccessToken(ctx, db, respBody.AccessToken)
		require.NoError(t, err)

		return entity != nil && entity.ClientID == clientId && entity.UserId.UserId == dummy.UserID && entity.RedirectURI == "https://example.com/login"
	}, "")

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(logs) == 1 && logs[0].ClientID == clientId
	}, "")
}

func TestOauthLoginRole(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		AllowRoles: []string{"hoge"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser().AddRole("hoge")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	resp := s.Post(t, "/login", req)

	var respBody oauth.LoginResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.NotEmpty(t, respBody.AccessToken)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAccessTokenByAccessToken(ctx, db, respBody.AccessToken)
		require.NoError(t, err)

		return entity != nil && entity.ClientID == clientId && entity.UserId.UserId == dummy.UserID && entity.RedirectURI == "https://example.com/login"
	}, "")

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByUserId(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(logs) == 1 && logs[0].ClientID == clientId
	}, "")
}

func TestOauthLoginFailedRole(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		AllowRoles: []string{"test"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser() // roleはない
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	reqForm, err := json.Marshal(req)
	require.NoError(t, err)

	resp, err := s.Client.Post(s.Server.URL+"/login", "application/json", bytes.NewBuffer(reqForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}

func TestToken(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	accessToken := models.SSOAccessToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: "",

		ClientID: clientId,

		RedirectURI: "https://example.com/login",

		Create: time.Now(),

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = accessToken.Add(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), false)

	// --- JWTのpublic keyもらう

	resp := s.Get(t, "/jwt")

	var jwtPublic oauth.JWTPublic
	err = json.Unmarshal(tools.ConvertByteResp(resp), &jwtPublic)
	require.NoError(t, err)

	require.NotEmpty(t, jwtPublic.PKCS8)

	// ---

	// token endpoint の認証はAuthorization headerを使う
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
			s.Server.URL, accessToken.SSOAccessToken, accessToken.RedirectURI), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", service.TokenSecret))
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err = s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var tokenRespBody oauth.TokenResponse

	err = json.Unmarshal(tools.ConvertByteResp(resp), &tokenRespBody)
	require.NoError(t, err)

	require.Equal(t, tokenRespBody.AccessToken, accessToken.SSOAccessToken)
	require.Equal(t, tokenRespBody.ExpiresIn, 10*60)
	require.NotEmpty(t, tokenRespBody.RefreshToken)

	IDToken := tokenRespBody.IDToken

	// JWT IDTokenを検証する

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(jwtPublic.PKCS8))
	require.NoError(t, err)

	token, err := jwt.Parse(IDToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return verifyKey, nil
	})
	require.NoError(t, err)

	require.True(t, token.Valid)

	require.NotEmpty(t, token.Raw)

	value := make(url.Values)
	value.Add("grant_type", "refresh_token")
	value.Add("client_id", clientId)
	value.Add("client_secret", service.TokenSecret)
	value.Add("refresh_token", tokenRespBody.RefreshToken)
	value.Add("scope", "openid")

	// refreshでAccessTokenを更新する
	resp, err = s.Client.PostForm(s.Server.URL+"/token", value)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var newTokenBody oauth.TokenResponse

	err = json.Unmarshal(tools.ConvertByteResp(resp), &newTokenBody)
	require.NoError(t, err)

	require.NotEqual(t, tokenRespBody.AccessToken, newTokenBody.AccessToken)
	require.NotEqual(t, tokenRespBody.RefreshToken, newTokenBody.RefreshToken)
}

func TestTokenError(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	accessToken := models.SSOAccessToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: "",

		ClientID: clientId,

		RedirectURI: "https://example.com/login",

		Create: time.Now(),

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}
	err = accessToken.Add(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), false)

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
			s.Server.URL, accessToken.SSOAccessToken, accessToken.RedirectURI), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Basic dummy")
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err := s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403)

	// ---

	req, err = http.NewRequest("GET",
		fmt.Sprintf("%s/token?code=%s&redirect_uri=%s",
			s.Server.URL, accessToken.SSOAccessToken, accessToken.RedirectURI), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", service.TokenSecret))
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err = s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)

	// ---

	req, err = http.NewRequest("GET",
		fmt.Sprintf("%s/token?grant_type=authorization_code&redirect_uri=%s",
			s.Server.URL, accessToken.RedirectURI), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", service.TokenSecret))
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err = s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)

	// ---

	req, err = http.NewRequest("GET",
		fmt.Sprintf("%s/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
			s.Server.URL, accessToken.SSOAccessToken, "hoge"), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", service.TokenSecret))
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err = s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)

	// ---

	req, err = http.NewRequest("GET",
		fmt.Sprintf("%s/token?grant_type=authorization_code&code=%s",
			s.Server.URL, accessToken.SSOAccessToken), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", service.TokenSecret))
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err = s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)

	// ---

	req, err = http.NewRequest("GET",
		fmt.Sprintf("%s/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
			s.Server.URL, "dummy", accessToken.RedirectURI), nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", service.TokenSecret))
	// req.SetBasicAuth("", service.TokenSecret)

	resp, err = s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}
