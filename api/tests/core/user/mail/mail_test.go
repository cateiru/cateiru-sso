package mail_test

import (
	"context"
	"testing"

	core_mail "github.com/cateiru/cateiru-sso/api/core/user/mail"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestMail(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		mail, err := core_mail.GetMail(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return mail == dummy.Mail
	}, "メールアドレスが取得できる")
}
