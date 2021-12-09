package database

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Database struct {
	Client *datastore.Client
}

// Datastoreのclientを作成する
// projectIDは環境変数 `DATASTORE_PROJECT_ID` で設定する必要があります
func NewDatabase(ctx context.Context) (*Database, error) {
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		return nil, err
	}

	return &Database{
		Client: client,
	}, nil
}

// Datastoreのclientを閉じる
func (c *Database) Close() {
	c.Client.Close()
}

// keyを指定してdatastoreのentryを1つ取得する
//
// 要素が見つからない場合、trueを返します
func (c *Database) Get(ctx context.Context, key *datastore.Key, entity interface{}) (bool, error) {
	err := c.Client.Get(ctx, key, entity)
	if err == datastore.ErrNoSuchEntity {
		return true, nil
	}
	return false, err
}

// queryに一致するdatastoreのentryをすべて取得する
func (c *Database) GetAll(ctx context.Context, query *datastore.Query, entities interface{}) ([]*datastore.Key, error) {
	keys, err := c.Client.GetAll(ctx, query, entities)
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	return keys, err
}

// keyを指定してentryをdatastoreに追加する
func (c *Database) Put(ctx context.Context, key *datastore.Key, entry interface{}) error {
	if _, err := c.Client.Put(ctx, key, entry); err != nil {
		return err
	}

	return nil
}

// 指定したqueryの数を返します
func (c *Database) Count(ctx context.Context, query *datastore.Query) (int, error) {
	return c.Client.Count(ctx, query)
}

// queryを実行します
func (c *Database) Run(ctx context.Context, query *datastore.Query) *datastore.Iterator {
	return c.Client.Run(ctx, query)
}

// 指定したkeyのentryを複数削除
func (c *Database) DeleteMulti(ctx context.Context, key []*datastore.Key) error {
	return c.Client.DeleteMulti(ctx, key)
}

// 指定したkeyのentryを削除
func (c *Database) Delete(ctx context.Context, key *datastore.Key) error {
	return c.Client.Delete(ctx, key)
}
