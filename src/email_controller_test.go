package src_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestNewEmail(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	require.Equal(t, e.Email, email2)
	require.Equal(t, e.Ip, ip)
	require.Equal(t, e.HasPreviewMode, false)
}

func TestRegisterEmailVerify(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	text, err := e.RegisterEmailVerify("123456")
	require.NoError(t, err)
	require.Equal(t, text, "")

	// プレビューモードにする
	e.HasPreviewMode = true

	text, err = e.RegisterEmailVerify("123456")
	require.NoError(t, err)
	require.NotEqual(t, text, "")
}

func TestResendRegisterEmailVerify(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	text, err := e.ResendRegisterEmailVerify("123456")
	require.NoError(t, err)
	require.Equal(t, text, "")

	// プレビューモードにする
	e.HasPreviewMode = true

	text, err = e.ResendRegisterEmailVerify("123456")
	require.NoError(t, err)
	require.NotEqual(t, text, "")
}

func TestUpdateEmail(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	text, err := e.UpdateEmail(email, "123456")
	require.NoError(t, err)
	require.Equal(t, text, "")

	// プレビューモードにする
	e.HasPreviewMode = true

	text, err = e.UpdateEmail(email, "123456")
	require.NoError(t, err)
	require.NotEqual(t, text, "")
}

func TestUpdatePassword(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	text, err := e.UpdatePassword("token", user.UserName)
	require.NoError(t, err)
	require.Equal(t, text, "")

	// プレビューモードにする
	e.HasPreviewMode = true

	text, err = e.UpdatePassword("token", user.UserName)
	require.NoError(t, err)
	require.NotEqual(t, text, "")
}

func TestInviteOrg(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	text, err := e.InviteOrg("token", "orgName", user.UserName)
	require.NoError(t, err)
	require.Equal(t, text, "")

	// プレビューモードにする
	e.HasPreviewMode = true

	text, err = e.InviteOrg("token", "orgName", user.UserName)
	require.NoError(t, err)
	require.NotEqual(t, text, "")
}

func TestTest(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	email := RandomEmail(t)
	email2 := RandomEmail(t)

	userData := &src.UserData{
		Browser:  "Chrome",
		OS:       "Linux",
		Device:   "",
		IsMobile: false,
	}
	ip := "172.0.0.1"

	user := RegisterUser(t, ctx, email)

	e := src.NewEmail(h.Sender, C, email2, userData, ip, &user)

	text, err := e.Test()
	require.NoError(t, err)
	require.Equal(t, text, "")

	// プレビューモードにする
	e.HasPreviewMode = true

	text, err = e.Test()
	require.NoError(t, err)
	require.NotEqual(t, text, "")
}
