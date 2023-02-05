package storage_test

import (
	"context"
	"strings"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/storage"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()
	str, err := storage.NewStorage(ctx, config.Defs.StorageBucket)
	require.NoError(t, err)
	defer str.Close()

	body := "hoge"
	dirs := []string{"target"}
	filename := "files.txt"

	// ファイル書き出しする
	err = str.WriteFile(ctx, dirs, filename, strings.NewReader(body), "text/plain")
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		exist, err := str.FileExist(ctx, dirs, filename)
		require.NoError(t, err)

		return exist
	}, "存在している")

	getBody, contentType, err := str.ReadFile(ctx, dirs, filename)
	require.NoError(t, err)
	require.Equal(t, contentType, "text/plain")

	require.Equal(t, body, string(getBody))
}
