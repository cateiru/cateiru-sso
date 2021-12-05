package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

// MailTokenで取得します
func GetMailCertificationByMailToken(ctx context.Context, db *database.Database, mailToken string) (*MailCertification, error) {
	query := datastore.NewQuery("MailCertification").Filter("mailToken =", mailToken)

	iter := db.Run(ctx, query)

	var entry MailCertification
	_, err := iter.Next(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// ClientMailTokenで取得します
func GetMailCertificationByCheckToken(ctx context.Context, db *database.Database, clientCheckToken string) (*MailCertification, error) {
	query := datastore.NewQuery("MailCertification").Filter("clientCheckToken =", clientCheckToken)

	iter := db.Run(ctx, query)

	var entry MailCertification
	_, err := iter.Next(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// mailCertificationに要素を追加する
func (c *MailCertification) Add(ctx context.Context, db *database.Database) error {
	// MailTokenをkeyにする
	key := database.CreateNameKey("MailCertification", c.MailToken)

	return db.Put(ctx, key, c)
}
