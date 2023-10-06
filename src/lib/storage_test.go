package lib_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()
	s := lib.NewCloudStorage("test-oreore-me")
	os.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")

	t.Run("読み出し、書き込みが可能", func(t *testing.T) {
		body := "hoge"
		dir := "target"
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		filename := r + ".txt"
		// ファイル書き出しする
		err = s.Write(ctx, filepath.Join(dir, filename), strings.NewReader(body), "text/plain")
		require.NoError(t, err)

		time.Sleep(2 * time.Second)

		getBody, contentType, err := s.Read(ctx, filepath.Join(dir, filename))
		require.NoError(t, err)
		require.Equal(t, contentType, "text/plain")

		require.Equal(t, body, string(getBody))
	})

	t.Run("更新可能", func(t *testing.T) {
		dir := "target"
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		filename := r + ".txt"

		// ファイル書き出しする
		err = s.Write(ctx, filepath.Join(dir, filename), strings.NewReader("hogehoge"), "text/plain")
		require.NoError(t, err)

		time.Sleep(1 * time.Second)

		// ファイル書き出しする
		err = s.Write(ctx, filepath.Join(dir, filename), strings.NewReader("hugahuga"), "text/plain")
		require.NoError(t, err)

		time.Sleep(2 * time.Second)

		getBody, contentType, err := s.Read(ctx, filepath.Join(dir, filename))
		require.NoError(t, err)
		require.Equal(t, contentType, "text/plain")

		require.Equal(t, string(getBody), "hugahuga")
	})

	t.Run("存在しないファイルを読み出す", func(t *testing.T) {
		dir := "empty"
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		filename := r + ".txt"

		_, _, err = s.Read(ctx, filepath.Join(dir, filename))
		require.ErrorIs(t, err, storage.ErrObjectNotExist)
	})

	t.Run("Content-Typeが空", func(t *testing.T) {
		dir := "target"
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		filename := r + ".txt"

		err = s.Write(ctx, filepath.Join(dir, filename), strings.NewReader("hogehoge"), "")
		require.EqualError(t, err, "Content-Type is empty")
	})

}
