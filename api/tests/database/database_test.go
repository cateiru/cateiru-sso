package database_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/iterator"
)

type sampleEntry struct {
	Text string `datastore:"text"`
}

type sampleEntrySecond struct {
	Id string `datastore:"id"`

	sampleEntry
}

// DBのアクセスをトライする
func waitDB(t *testing.T, f func() bool, message string) {
	// 合計: 18秒
	waitTimes := []int{1, 1, 2, 2, 3, 4, 5}
	successFlag := false
	id := utils.CreateID(5)

	for _, wait := range waitTimes {
		t.Logf("Wait DB sleep: %v s, id: %s", wait, id)
		time.Sleep(time.Duration(wait) * time.Second)
		result := f()

		if result {
			successFlag = true
			break
		}
	}

	if !successFlag {
		t.Fatalf("waitDB: %v", message)
	}
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

	query := datastore.NewQuery(tableName)

	waitDB(t, func() bool {
		numberOfEntry, err := client.Count(ctx, query)
		require.NoError(t, err, "Countできない")
		t.Logf("%v == %v", numberOfEntry, entries)
		return numberOfEntry == entries
	}, "カウントされた数が違う")

	returnEntries := []sampleEntry{}
	keys, err := client.GetAll(ctx, query, &returnEntries)
	require.NoError(t, err, "GetAllできない")
	require.Equal(t, len(returnEntries), entries, "GetAllしたentryの数が違う")

	err = client.DeleteMulti(ctx, keys)
	require.NoError(t, err, "削除できない")

	waitDB(t, func() bool {
		numberOfEntry, err := client.Count(ctx, query)
		require.NoError(t, err, "Countできない")
		t.Logf("%v == 0", numberOfEntry)
		return numberOfEntry == 0
	}, "削除できてない")

	client.Close()
}

func TestFindDB(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()
	tableName := utils.CreateID(5)
	numberOfEntry := 5
	targetIndex := 3

	client, err := database.NewDatabase(ctx)
	require.NoError(t, err, "Datastoreに接続できない")

	ids := []string{}
	for i := 0; numberOfEntry > i; i++ {
		ids = append(ids, utils.CreateID(5))
	}

	for index, id := range ids {
		key := datastore.NameKey(tableName, utils.CreateID(5), nil)
		client.Put(ctx, key, &sampleEntrySecond{
			Id: id,
			sampleEntry: sampleEntry{
				Text: strconv.Itoa(index),
			},
		})
	}

	query := datastore.NewQuery(tableName).Filter("id =", ids[targetIndex])

	waitDB(t, func() bool {
		entries := []sampleEntrySecond{}
		_, err := client.GetAll(ctx, query, &entries)
		require.NoError(t, err, "GetAllできない")

		t.Logf("Return number of entries: %d, entry id: %s", len(entries), entries[0].Id)

		return len(entries) == 1 && entries[0].Id == ids[targetIndex] && entries[0].Text == strconv.Itoa(targetIndex)

	}, "GetAllでFindできない")

	waitDB(t, func() bool {
		iter := client.Run(ctx, query)

		for {
			var entry sampleEntrySecond
			_, err := iter.Next(&entry)
			if err == iterator.Done {
				return false // 見つからなかった
			}
			require.NoError(t, err, "イテレータをNEXTできない")

			t.Logf("Find Iter. id: %v, value: %v", entry.Id, entry.Text)

			if entry.Id == ids[targetIndex] && entry.Text == strconv.Itoa(targetIndex) {
				return true
			}
		}

	}, "GetAllでFindできない")

	client.Close()
}
