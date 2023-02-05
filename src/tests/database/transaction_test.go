package database_test

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/utils"
	"github.com/stretchr/testify/require"
)

// トランザクションを使用したentryの更新
func TestTransactionSuccess(t *testing.T) {
	config.TestInit(t)

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

	isNoEntry, err := tx.Get(key, &e)
	require.NoError(t, err)
	require.False(t, isNoEntry)

	e.Text = "huga"

	err = tx.Put(key, &e)
	require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)

	var entry2 sampleEntry
	_, err = client.Get(ctx, key, &entry2)
	require.NoError(t, err)

	require.Equal(t, entry2.Text, "huga")

}

// TXで削除
func TestTransactionDelete(t *testing.T) {
	config.TestInit(t)

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
	for i := 0; 3 > i; i++ {
		tx, err := database.NewTransaction(ctx, client)
		require.NoError(t, err)

		err = tx.Delete(key)
		require.NoError(t, err)

		err = tx.Commit()
		if err != nil && err != datastore.ErrConcurrentTransaction {
			t.Fatal()
		}
		if err == nil {
			return
		}
	}

	var entry2 sampleEntry
	isNoEntity, err := client.Get(ctx, key, &entry2)
	require.NoError(t, err)
	require.True(t, isNoEntity)
}

// トランザクションのロールバック
func TestTransactionRollback(t *testing.T) {
	config.TestInit(t)

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

	isNoEntity, err := tx.Get(key, &e)
	require.NoError(t, err)
	require.False(t, isNoEntity)

	e.Text = "huga"

	err = tx.Put(key, &e)
	require.NoError(t, err)

	err = tx.Delete(key)
	require.NoError(t, err)

	err = tx.Rollback()
	require.NoError(t, err)

	var entry2 sampleEntry
	_, err = client.Get(ctx, key, &entry2)
	require.NoError(t, err)

	require.Equal(t, entry2.Text, "hoge")

}
