package oauth

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/oauth"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestTokenRequest(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	accessToken := utils.CreateID(0)

	access := models.SSOAccessToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: "",

		ClientID: utils.CreateID(0),

		RedirectURI: "https://example.com",

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 5,
		},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = access.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAccessTokenByAccessToken(ctx, db, accessToken)
		require.NoError(t, err)

		return entity != nil
	}, "")

	// -- 成功

	token := oauth.TokenRequest{
		GrantType:   "authorization_code",
		Code:        accessToken,
		RedirectUri: access.RedirectURI,
	}

	at, err := token.Required(ctx, db)
	require.NoError(t, err)

	require.Equal(t, at.ClientID, access.ClientID)

	// -- 失敗

	token = oauth.TokenRequest{
		GrantType:   "grant_type",
		Code:        accessToken,
		RedirectUri: access.RedirectURI,
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)

	// --

	token = oauth.TokenRequest{
		Code:        accessToken,
		RedirectUri: access.RedirectURI,
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)

	// --

	token = oauth.TokenRequest{
		GrantType:   "authorization_code",
		RedirectUri: access.RedirectURI,
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)

	// --

	token = oauth.TokenRequest{
		GrantType:   "authorization_code",
		Code:        "hoge",
		RedirectUri: access.RedirectURI,
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)

	// --

	token = oauth.TokenRequest{
		GrantType: "authorization_code",
		Code:      accessToken,
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)

	// --

	token = oauth.TokenRequest{
		GrantType:   "authorization_code",
		Code:        accessToken,
		RedirectUri: "https://cateiru.com",
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)

	// --

	accessToken = utils.CreateID(0)

	access = models.SSOAccessToken{
		SSOAccessToken:  accessToken,
		SSORefreshToken: "",

		ClientID: utils.CreateID(0),

		RedirectURI: "https://example.com",

		Period: models.Period{
			CreateDate:   time.Now().Add(time.Duration(-1 * time.Hour)), // 有効期限切れ
			PeriodMinute: 5,
		},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}

	err = access.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAccessTokenByAccessToken(ctx, db, accessToken)
		require.NoError(t, err)

		return entity != nil
	}, "")

	token = oauth.TokenRequest{
		GrantType:   "authorization_code",
		Code:        accessToken,
		RedirectUri: access.RedirectURI,
	}

	_, err = token.Required(ctx, db)
	require.Error(t, err)
}

func TestRefreshRequest(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	refreshToken := utils.CreateID(0)

	refresh := models.SSORefreshToken{
		SSOAccessToken:  "",
		SSORefreshToken: refreshToken,

		ClientID: utils.CreateID(0),

		RedirectURI: "https://example.com",

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodDay:  7,
		},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = refresh.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entity != nil
	}, "")

	// -- 成功

	refreshReq := oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     refresh.ClientID,
		ClientSecret: "hoge",
		RefreshToken: refreshToken,
		Scope:        []string{"openid"},
	}
	ref, err := refreshReq.Required(ctx, db)
	require.NoError(t, err)

	require.Equal(t, ref.ClientID, refresh.ClientID)

	// -- 失敗

	refreshReq = oauth.RefreshRequest{
		GrantType:    "hoge",
		ClientID:     refresh.ClientID,
		ClientSecret: "hoge",
		RefreshToken: refreshToken,
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		ClientID:     refresh.ClientID,
		ClientSecret: "hoge",
		RefreshToken: refreshToken,
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientSecret: "hoge",
		RefreshToken: refreshToken,
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     "",
		ClientSecret: "hoge",
		RefreshToken: refreshToken,
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     refresh.ClientID,
		RefreshToken: refreshToken,
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     refresh.ClientID,
		ClientSecret: "hoge",
		RefreshToken: "",
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     refresh.ClientID,
		ClientSecret: "hoge",
		Scope:        []string{"openid"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)

	// --

	refreshReq = oauth.RefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     refresh.ClientID,
		ClientSecret: "hoge",
		RefreshToken: refreshToken,
		Scope:        []string{"hoge"},
	}
	_, err = refreshReq.Required(ctx, db)
	require.Error(t, err)
}
