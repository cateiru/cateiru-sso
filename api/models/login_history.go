package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

// userIdを指定して、ログインログを取得する
func GetAllLoginHistory(ctx context.Context, db *database.Database, userId string, limit ...interface{}) ([]LoginHistory, error) {
	limitNum := -1
	if len(limit) != 0 && limit != nil {
		if limitInt, ok := limit[0].(int); ok {
			limitNum = limitInt
		}
	}

	query := datastore.NewQuery("LoginHistory").Filter("userId =", userId).Limit(limitNum)
	var entries []LoginHistory

	if _, err := db.GetAll(ctx, query, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}

// ログインログを削除する
func DeleteAllLoginHistories(ctx context.Context, db *database.Database, userId string) error {
	query := datastore.NewQuery("LoginHistory").Filter("userId =", userId)
	var entries []LoginHistory

	keys, err := db.GetAll(ctx, query, &entries)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func (c *LoginHistory) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("LoginHistory", c.AccessId)
	return db.Put(ctx, key, c)
}
