package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/src/database"
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

func DeleteOTPBufferPeriod(ctx context.Context, db *database.Database) error {
	query := datastore.NewQuery("OnetimePasswordBuffer")

	var entities []OnetimePasswordBuffer

	_, err := db.GetAll(ctx, query, &entities)
	if err != nil {
		return err
	}

	var keys []*datastore.Key

	for _, entity := range entities {
		if CheckExpired(&entity.Period) {
			key := database.CreateNameKey("OnetimePasswordBuffer", entity.Id)
			keys = append(keys, key)
		}
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteOTPBuffer(ctx context.Context, db *database.Database, id string) error {
	key := database.CreateNameKey("OnetimePasswordBuffer", id)
	return db.Delete(ctx, key)
}

func (c *OnetimePasswordBuffer) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("OnetimePasswordBuffer", c.Id)
	return db.Put(ctx, key, c)
}
