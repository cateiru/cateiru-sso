package models

import (
	"context"

	"github.com/cateiru/cateiru-sso/api/database"
)

func GetOTPBufferByID(ctx context.Context, db *database.Database, id string) (*OnetimePasswordBuffer, error) {
	key := database.CreateNameKey("OnetimePasswordBuffer", id)
	var entity OnetimePasswordBuffer

	isNotExist, err := db.Get(ctx, key, &entity)
	if err != nil {
		return nil, err
	}
	// 要素がない場合
	if isNotExist {
		return nil, nil
	}

	return &entity, nil
}

func DeleteOTPBuffer(ctx context.Context, db *database.Database, id string) error {
	key := database.CreateNameKey("OnetimePasswordBuffer", id)
	return db.Delete(ctx, key)
}

func (c *OnetimePasswordBuffer) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("OnetimePasswordBuffer", c.Id)
	return db.Put(ctx, key, c)
}
