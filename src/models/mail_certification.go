package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/src/database"
	"google.golang.org/api/iterator"
)

// MailTokenで取得します
func GetMailCertificationByMailToken(ctx context.Context, db *database.Database, mailToken string) (*MailCertification, error) {
	query := datastore.NewQuery("MailCertification").Filter("mailToken =", mailToken)

	iter := db.Run(ctx, query)

	var entry MailCertification
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

// ClientMailTokenで取得します
func GetMailCertificationByClientToken(ctx context.Context, db *database.Database, clientToken string) (*MailCertification, error) {
	query := datastore.NewQuery("MailCertification").Filter("clientToken =", clientToken)

	iter := db.Run(ctx, query)

	var entry MailCertification
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

func DeleteMailCertPeriod(ctx context.Context, db *database.Database) error {
	query := datastore.NewQuery("MailCertification")

	var entities []MailCertification

	_, err := db.GetAll(ctx, query, &entities)
	if err != nil {
		return err
	}

	var keys []*datastore.Key

	for _, entity := range entities {
		if CheckExpired(&entity.Period) {
			key := database.CreateNameKey("MailCertification", entity.MailToken)
			keys = append(keys, key)
		}
	}

	return db.DeleteMulti(ctx, keys)
}

// 削除
func DeleteMailCertification(ctx context.Context, db *database.Database, mailToken string) error {
	key := database.CreateNameKey("MailCertification", mailToken)
	return db.Delete(ctx, key)
}

// mailCertificationに要素を追加する
func (c *MailCertification) Add(ctx context.Context, db *database.Database) error {
	// MailTokenをkeyにする
	key := database.CreateNameKey("MailCertification", c.MailToken)

	return db.Put(ctx, key, c)
}
