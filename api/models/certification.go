package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

// メールアドレスから対象の認証情報を取得します
func GetCertificationByMail(ctx context.Context, db *database.Database, mail string) (*Certification, error) {
	query := datastore.NewQuery("Certification").Filter("mail =", mail)
	iter := db.Run(ctx, query)

	var entry Certification
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

// ユーザIDから対象の認証情報を取得します
func GetCertificationByUserID(ctx context.Context, db *database.Database, userId string) (*Certification, error) {
	query := datastore.NewQuery("Certification").Filter("userId =", userId)
	iter := db.Run(ctx, query)

	var entry Certification
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

func DeleteCertificationByUserId(ctx context.Context, db *database.Database, userId string) error {
	key := database.CreateNameKey("Certification", userId)
	return db.Delete(ctx, key)
}

// certificationに要素を追加する
func (c *Certification) Add(ctx context.Context, db *database.Database) error {

	// ユーザIDをkeyにする
	key := database.CreateNameKey("Certification", c.UserId.UserId)

	return db.Put(ctx, key, c)
}
