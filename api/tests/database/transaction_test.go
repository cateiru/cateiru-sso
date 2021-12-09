package database_test

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

// トランザクションを使用したentryの更新
func TestTransactionSuccess(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()
	tableName := utils.CreateID(5)

	client, err := database.NewDatabase(ctx)
	require.NoError(t, err, "Datastoreに接続できない")

	key := datastore.NameKey(tableName, utils.CreateID(5), nil)

	// 最初の要素追加
	entry := &sampleEntry{
		Text: "hoge",
	}

	err = client.Put(ctx, key, entry)
	require.NoError(t, err, "Putできない")

	// トランザクションを開始し、Textの内容を変える
	tx, err := database.NewTransaction(ctx, client)
	require.NoError(t, err)

	var e sampleEntry

	err = tx.Get(key, &e)
	require.NoError(t, err)

	e.Text = "huga"

	err = tx.Put(key, &e)
	require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)

	var entry2 sampleEntry
	err = client.Get(ctx, key, &entry2)
	require.NoError(t, err)

	require.Equal(t, entry2.Text, "huga")

}

// トランザクションのロールバック
func TestTransactionRollback(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()
	tableName := utils.CreateID(5)

	client, err := database.NewDatabase(ctx)
	require.NoError(t, err, "Datastoreに接続できない")

	key := datastore.NameKey(tableName, utils.CreateID(5), nil)

	// 最初の要素追加
	entry := &sampleEntry{
		Text: "hoge",
	}

	err = client.Put(ctx, key, entry)
	require.NoError(t, err, "Putできない")

	// トランザクションを開始し、Textの内容を変える
	tx, err := database.NewTransaction(ctx, client)
	require.NoError(t, err)

	var e sampleEntry

	err = tx.Get(key, &e)
	require.NoError(t, err)

	e.Text = "huga"

	err = tx.Put(key, &e)
	require.NoError(t, err)

	err = tx.Rollback()
	require.NoError(t, err)

	var entry2 sampleEntry
	err = client.Get(ctx, key, &entry2)
	require.NoError(t, err)

	require.Equal(t, entry2.Text, "hoge")

}
