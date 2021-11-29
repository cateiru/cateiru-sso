package database_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

type sampleEntry struct {
	Text string `datastore:"text"`
}

// データベース（Cloud datastore）の接続、Put、Getを試す
func TestConnectDB(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	client, err := database.NewDatabase(ctx)
	require.NoError(t, err, "Datastoreに接続できない")

	key := datastore.NameKey("SampleTable", utils.CreateID(5), nil)

	entry := &sampleEntry{
		Text: "hoge",
	}

	err = client.Put(ctx, key, entry)
	require.NoError(t, err, "Putできない")

	// 非同期でDBに追加するため追加される（であろう）まで待機する
	time.Sleep(1 * time.Second)

	returnedEntry := new(sampleEntry)
	err = client.Get(ctx, key, returnedEntry)
	require.NoError(t, err, "Getできない")

	require.Equal(t, entry.Text, returnedEntry.Text, "返ってきた値が違う")

	err = client.Delete(ctx, key)
	require.NoError(t, err, "Deleteできない")

	// 非同期でDBに追加するため追加される（であろう）まで待機する
	time.Sleep(1 * time.Second)

	returnedEntry = new(sampleEntry)
	err = client.Get(ctx, key, returnedEntry)
	require.Error(t, err, "削除できてない")

	client.Close()
}

// CountとGetAll、DeleteMultiのテスト
func TestMultiEntryDB(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()
	tableName := utils.CreateID(5)
	entries := 5

	client, err := database.NewDatabase(ctx)
	require.NoError(t, err, "Datastoreに接続できない")

	entry := &sampleEntry{
		Text: "hogehoge",
	}

	for i := 0; entries > i; i++ {
		key := datastore.NameKey(tableName, utils.CreateID(5), nil)
		err = client.Put(ctx, key, entry)
		require.NoError(t, err, "Putできない")
	}

	// 非同期でDBに追加するため追加される（であろう）まで待機する
	time.Sleep(3 * time.Second)

	query := datastore.NewQuery(tableName)

	numberOfEntry, err := client.Count(ctx, query)
	require.NoError(t, err, "Countできない")
	require.Equal(t, numberOfEntry, entries, "カウントされた数が違う")

	returnEntries := []sampleEntry{}
	keys, err := client.GetAll(ctx, query, &returnEntries)
	require.NoError(t, err, "GetAllできない")
	require.Equal(t, len(returnEntries), entries, "GetAllしたentryの数が違う")

	err = client.DeleteMulti(ctx, keys)
	require.NoError(t, err, "削除できない")

	// 非同期でDBに追加するため追加される（であろう）まで待機する
	time.Sleep(3 * time.Second)

	numberOfEntry, err = client.Count(ctx, query)
	require.NoError(t, err, "Countできない")
	require.Equal(t, numberOfEntry, 0, "削除できてない")

	client.Close()
}
