package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

// IPを指定して、ブラックリストを取得
func GetBlockListByIP(ctx context.Context, db *database.Database, ip string) (*IPBlockList, error) {
	query := datastore.NewQuery("IPBlockList").Filter("ip =", ip)

	iter := db.Run(ctx, query)

	var entry IPBlockList
	_, err := iter.Next(&entry)
	// 要素がなにもない場合nilを返す
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// メールアドレスを指定して、ブロックリストを取得
func GetBlockListByMail(ctx context.Context, db *database.Database, mail string) (*MailBlockList, error) {
	query := datastore.NewQuery("MailBlockList").Filter("mail =", mail)

	iter := db.Run(ctx, query)

	var entry MailBlockList
	_, err := iter.Next(&entry)
	// 要素がなにもない場合nilを返す
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// IPのブラックリストを追加
func (c *IPBlockList) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("IPBlockList", c.IP)

	return db.Put(ctx, key, c)
}

// メールアドレスのブラックリストを追加
func (c *MailBlockList) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("MailBlockList", c.Mail)

	return db.Put(ctx, key, c)
}
