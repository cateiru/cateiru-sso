package createaccount_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

// 一時的なアカウント作成（成功）
func TestSuccess(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	form := &createaccount.PostForm{
		Mail:      fmt.Sprintf("%s@example.com", utils.CreateID(4)),
		ReCAPTCHA: "",
	}
	ip := "192.168.1.1"

	clientToken, err := createaccount.CreateTemporaryAccount(ctx, form, ip)
	require.NoError(t, err)

	// 有効期限を確認するため一旦sleepする
	time.Sleep(1 * time.Second)

	// 確認
	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	element := new(models.MailCertification)

	goretry.Retry(t, func() bool {
		element, err = models.GetMailCertificationByClientToken(ctx, db, clientToken)
		require.NoError(t, err)

		return element != nil
	}, "")

	require.NotEqual(t, len(element.MailToken), 0, "mail tokenが存在する")
	require.NotEqual(t, element.Verify, "まだメールアドレスは認証されていない")
	require.NotEqual(t, element.OpenNewWindow, "新しいウィンドウではない")
	require.NotEqual(t, element.ChangeMailMode, "メールアドレス変更ではない")

	require.Equal(t, element.Mail, form.Mail, "メールアドレスがある")

	require.Equal(t, element.ClientToken, clientToken)

	now := time.Now()
	require.NotEqual(t, now.Sub(element.CreateDate), time.Duration(0), "ちゃんと作成日時が定義されている")
}

func FailedMail(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	form := &createaccount.PostForm{
		Mail:      "aaaaaaa",
		ReCAPTCHA: "",
	}
	ip := "192.168.1.1"

	_, err := createaccount.CreateTemporaryAccount(ctx, form, ip)
	require.Error(t, err)
}
