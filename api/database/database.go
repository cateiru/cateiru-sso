package database

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Database struct {
	client *datastore.Client
}

// Datastoreのclientを作成する
// projectIDは環境変数 `DATASTORE_PROJECT_ID` で設定する必要があります
func NewDatabase(ctx context.Context) (*Database, error) {
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
	}, nil
}

// Datastoreのclientを閉じる
func (c *Database) Close() {
	c.client.Close()
}

// keyを指定してdatastoreのentryを1つ取得する
func (c *Database) Get(ctx context.Context, key *datastore.Key, entity interface{}) error {
	return c.client.Get(ctx, key, entity)
}

// queryに一致するdatastoreのentryをすべて取得する
func (c *Database) GetAll(ctx context.Context, query *datastore.Query, entities interface{}) ([]*datastore.Key, error) {
	return c.client.GetAll(ctx, query, entities)
}

// keyを指定してentryをdatastoreに追加する
func (c *Database) Put(ctx context.Context, key *datastore.Key, entry interface{}) error {
	if _, err := c.client.Put(ctx, key, entry); err != nil {
		return err
	}

	return nil
}

// 指定したqueryの数を返します
func (c *Database) Count(ctx context.Context, query *datastore.Query) (int, error) {
	return c.client.Count(ctx, query)
}

// queryを実行します
func (c *Database) Run(ctx context.Context, query *datastore.Query) *datastore.Iterator {
	return c.client.Run(ctx, query)
}

// 指定したkeyのentryを複数削除
func (c *Database) DeleteMulti(ctx context.Context, key []*datastore.Key) error {
	return c.client.DeleteMulti(ctx, key)
}

// 指定したkeyのentryを削除
func (c *Database) Delete(ctx context.Context, key *datastore.Key) error {
	return c.client.Delete(ctx, key)
}
