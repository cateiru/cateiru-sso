package database

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Transaction struct {
	tx datastore.Transaction
}

func NewTransaction(ctx context.Context, db *Database) (*Transaction, error) {
	tx, err := db.Client.NewTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		tx: *tx,
	}, nil
}

// 取得
// err == datastore.ErrNoSuchEntity の場合は、リトライはしません
func (c *Transaction) Get(key *datastore.Key, entry interface{}) (bool, error) {
	err := c.tx.Get(key, entry)
	if err == datastore.ErrNoSuchEntity {
		return true, nil
	}
	return false, err
}

// 追加
func (c *Transaction) Put(key *datastore.Key, entry interface{}) error {
	if _, err := c.tx.Put(key, entry); err != nil {
		return err
	}
	return nil
}

// 削除
func (c *Transaction) Delete(key *datastore.Key) error {
	return c.tx.Delete(key)
}

// トランザクションをコミットする
// err != datastore.ErrConcurrentTransaction の場合はエラーを出す必要があります
func (c *Transaction) Commit() error {
	if _, err := c.tx.Commit(); err != nil {
		return err
	}
	return nil
}

// ロールバックする
func (c *Transaction) Rollback() error {
	return c.tx.Rollback()
}
