package storage_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/storage"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()
	str, err := storage.NewStorage(ctx, "test-bucket")
	require.NoError(t, err)

	body := []byte("hoge")
	dirs := []string{"target"}
	filename := "files.txt"

	// ファイル書き出しする
	err = str.WriteFile(ctx, dirs, filename, body)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		exist, err := str.FileExist(ctx, dirs, filename)
		require.NoError(t, err)

		return exist
	}, "存在している")

	getBody, err := str.ReadFile(ctx, dirs, filename)
	require.NoError(t, err)

	require.Equal(t, string(body), string(getBody))
}
