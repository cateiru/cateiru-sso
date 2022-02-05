package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

func GetAllUsers(ctx context.Context, db *database.Database) ([]User, error) {
	query := datastore.NewQuery("User")

	var entities []User

	if _, err := db.GetAll(ctx, query, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

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

func GetUserDataByUserName(ctx context.Context, db *database.Database, userName string) (*User, error) {
	query := datastore.NewQuery("User").Filter("userNameFormatted =", userName)

	iter := db.Run(ctx, query)

	var entry User
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

// TX modeで取得
func GetUserDataTXByUserID(db *database.Transaction, userId string) (*User, error) {
	key := database.CreateNameKey("User", userId)
	var entry User

	isNoEntry, err := db.Get(key, &entry)

	// 要素が見つからない場合はnilを返す
	if isNoEntry {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// ユーザ情報を削除
func DeleteUserDataByUserID(ctx context.Context, db *database.Database, userId string) error {
	key := database.CreateNameKey("User", userId)
	return db.Delete(ctx, key)
}

func (c *User) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("User", c.UserId.UserId)
	return db.Put(ctx, key, c)
}

func (c *User) AddTX(db *database.Transaction) error {
	key := database.CreateNameKey("User", c.UserId.UserId)
	return db.Put(key, c)
}
