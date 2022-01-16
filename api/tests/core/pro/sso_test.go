package pro_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/pro"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestGetSSO(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")

	form := pro.AddRequestForm{
		Name:      "TestSSO",
		FromURL:   []string{"https://example.com/login"},
		ToURL:     []string{"https://example.com/login/redirect"},
		LoginOnly: true,

		SessionTokenPeriod: 10,
		RefreshTokenPeriod: 60,
	}

	tokens, err := pro.AddSSO(ctx, db, dummy.UserID, &form)
	require.NoError(t, err)

	// --- 取得する

	goretry.Retry(t, func() bool {
		services, err := pro.GetSSO(ctx, db, dummy.UserID)
		if err != nil {
			t.Log(err)
			return false
		}

		return len(services) == 1 && services[0].PublicKey == tokens.PublicKey
	}, "取得できる")

	// --- 削除する

	err = pro.DeleteSSO(ctx, db, tokens.PublicKey, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		_, err := pro.GetSSO(ctx, db, dummy.UserID)
		return err != nil
	}, "削除されている")
}

func TestURL(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")

	// --- FromURL = 2, ToURL = 1

	form := pro.AddRequestForm{
		Name:      "TestSSO",
		FromURL:   []string{"https://example.com/login", "https://example.com/hogehoge"},
		ToURL:     []string{"https://example.com/login/redirect"},
		LoginOnly: true,

		SessionTokenPeriod: 10,
		RefreshTokenPeriod: 60,
	}

	_, err = pro.AddSSO(ctx, db, dummy.UserID, &form)
	require.NoError(t, err)

	// --- FromURL = 3, ToURL = 3

	form2 := pro.AddRequestForm{
		Name:      "TestSSO",
		FromURL:   []string{"https://example.com/login", "https://example.com/hogehoge", "https://example.com/nya"},
		ToURL:     []string{"https://example.com/login/redirect", "https://example.com/", "https://example.com/sdsd"},
		LoginOnly: true,

		SessionTokenPeriod: 10,
		RefreshTokenPeriod: 60,
	}

	_, err = pro.AddSSO(ctx, db, dummy.UserID, &form2)
	require.NoError(t, err)

	// --- FromURL = 0, ToURL = 0  failed

	form3 := pro.AddRequestForm{
		Name:      "TestSSO",
		FromURL:   []string{},
		ToURL:     []string{},
		LoginOnly: true,

		SessionTokenPeriod: 10,
		RefreshTokenPeriod: 60,
	}

	_, err = pro.AddSSO(ctx, db, dummy.UserID, &form3)
	require.Error(t, err)

	// --- FromURL = 1, ToURL = 2 failed

	form4 := pro.AddRequestForm{
		Name:      "TestSSO",
		FromURL:   []string{"https://example.com/login"},
		ToURL:     []string{"https://example.com/login/redirect", "https://example.com/"},
		LoginOnly: true,

		SessionTokenPeriod: 10,
		RefreshTokenPeriod: 60,
	}

	_, err = pro.AddSSO(ctx, db, dummy.UserID, &form4)
	require.Error(t, err)
}
