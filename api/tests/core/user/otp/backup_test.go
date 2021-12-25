package otp_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/user/otp"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestBackup(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)

	cert, err := dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		backups, err := otp.GetBackupcodes(ctx, db, dummy.UserID)
		require.NoError(t, err)

		if len(cert.OnetimePasswordBackups) == 0 {
			return false
		}

		for _, e := range cert.OnetimePasswordBackups {
			flag := false
			for _, e2 := range backups {
				if e == e2 {
					flag = true
				}
			}

			if !flag {
				return false
			}
		}

		return true
	}, "otp backupsが取得できる")
}
