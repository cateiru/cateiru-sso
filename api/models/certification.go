package models

import (
	"context"

	"github.com/cateiru/cateiru-sso/api/database"
)

// メールアドレスから対象の認証情報を取得します
func GetCertification(ctx context.Context, db *database.Database, mail string) (*Certification, error) {
	// メールアドレスをkeyにする
	key := database.CreateNameKey("Certification", mail)

	entry := new(Certification)

	err := db.Get(ctx, key, entry)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

// certificationに要素を追加する
func (c *Certification) Add(ctx context.Context, db *database.Database) error {

	// メールアドレスをkeyにする
	key := database.CreateNameKey("Certification", c.Mail)

	return db.Put(ctx, key, c)
}
