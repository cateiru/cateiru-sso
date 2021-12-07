package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

// メールアドレスで、アカウント作成ログを取得する
// 身に覚えがない登録メールが送信されたという問い合わせが来た場合に、これを使用して送信したIPアドレスを取得する
func GetTryCreateAccountLogByMail(ctx context.Context, db *database.Database, mail string) ([]TryCreateAccountLog, error) {
	query := datastore.NewQuery("TryCreateAccountLog").Filter("targetMail =", mail)
	iter := db.Run(ctx, query)

	logs := []TryCreateAccountLog{}

	for {
		var entry TryCreateAccountLog

		_, err := iter.Next(&entry)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		logs = append(logs, entry)
	}

	return logs, nil
}

// IPアドレスで、アカウント作成ログを取得する
func GetTryCreateAccountLogByIP(ctx context.Context, db *database.Database, ip string) ([]TryCreateAccountLog, error) {
	query := datastore.NewQuery("TryCreateAccountLog").Filter("ip =", ip)
	iter := db.Run(ctx, query)

	logs := []TryCreateAccountLog{}

	for {
		var entry TryCreateAccountLog

		_, err := iter.Next(&entry)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		logs = append(logs, entry)
	}

	return logs, nil
}

// 要素をDatastoreに追加
func (c *TryCreateAccountLog) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("TryCreateAccountLog", c.LogId)

	return db.Put(ctx, key, c)
}
