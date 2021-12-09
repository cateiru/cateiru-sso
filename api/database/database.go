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
func (c *Database) Get(ctx context.Context, key *datastore.Key, entity interface{}) error {
	err := c.Client.Get(ctx, key, entity)
	if err == datastore.ErrNoSuchEntity {
		return nil
	}
	return err
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

// // 要素を変更します
// // Transactionを使用して、変更中は他クライアントによるDatastoreアクセスをロックします
// func (c *Database) Change(ctx context.Context, key *datastore.Key, entry struct, f func(entry interface{}) error, retries int) error {
// 	err := *new(error)

// 	// 失敗した場合、retriesで指定した分リトライします
// 	for i := 0; retries > i; i++ {
// 		tx, err := c.client.NewTransaction(ctx)
// 		if err != nil {
// 			break
// 		}

// 		if err := tx.Get(key, &entry); err != datastore.ErrNoSuchEntity {
// 			break
// 		}

// 		if err := f(&entry); err != nil {
// 			break
// 		}

// 		if _, err := tx.Put(key, &entry); err != nil {
// 			break
// 		}

// 		// commit
// 		if _, err = tx.Commit(); err != datastore.ErrConcurrentTransaction {
// 			break
// 		}
// 	}

// 	return err
