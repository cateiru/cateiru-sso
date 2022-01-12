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
	"github.com/stretchr/testify/require"
)

// 一時的なアカウント作成（成功）
func TestSuccess(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	form := &createaccount.PostForm{
		Mail:       fmt.Sprintf("%s@example.com", utils.CreateID(4)),
		Password:   "test",
		ReCHAPTCHA: "",
	}
	ip := "192.168.1.1"

	clientCheckToken, err := createaccount.CreateTemporaryAccount(ctx, form, ip)
	require.NoError(t, err)

	// 有効期限を確認するため一旦sleepする
	time.Sleep(1 * time.Second)

	// 確認
	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	element, err := models.GetMailCertificationByCheckToken(ctx, db, clientCheckToken)
	require.NoError(t, err)

	require.NotEqual(t, len(element.MailToken), 0, "mail tokenが存在する")
	require.NotEqual(t, element.Verify, "まだメールアドレスは認証されていない")
	require.NotEqual(t, element.OpenNewWindow, "新しいウィンドウではない")
	require.NotEqual(t, element.ChangeMailMode, "メールアドレス変更ではない")

	require.NotEqual(t, element.Password, form.Password, "パスワードがハッシュ化されている")
	require.Equal(t, element.Mail, form.Mail, "メールアドレスがある")

	require.Equal(t, element.ClientCheckToken, clientCheckToken)

	now := time.Now()
	require.NotEqual(t, now.Sub(element.CreateDate), time.Duration(0), "ちゃんと作成日時が定義されている")
}
