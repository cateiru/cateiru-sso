package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

func GetUserDataByUserID(ctx context.Context, db *database.Database, userId string) (*User, error) {
	key := database.CreateNameKey("User", userId)
	var entry User

	notExist, err := db.Get(ctx, key, &entry)
	if err != nil {
		return nil, err
	}

	// 要素が見つからない場合はnilを返す
	if notExist {
		return nil, nil
	}

	return &entry, nil
}

// TX modeで取得
func GetUserDataTXByUserID(ctx context.Context, db *database.Transaction, userId string) (*User, error) {
	key := database.CreateNameKey("User", userId)
	var entry User

	err := db.Get(key, &entry)

	// 要素が見つからない場合はnilを返す
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (c *User) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("User", c.UserId.UserId)
	return db.Put(ctx, key, c)
}

func (c *User) AddTX(db *database.Transaction) error {
	key := database.CreateNameKey("User", c.UserId.UserId)
	return db.Put(key, c)
}
