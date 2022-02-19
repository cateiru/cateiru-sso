package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

func GetRefreshTokenByUserId(ctx context.Context, db *database.Database, userId string) ([]RefreshInfo, error) {
	query := datastore.NewQuery("RefreshInfo").Filter("userId =", userId)

	iter := db.Run(ctx, query)

	var entries []RefreshInfo

	for {
		var entry RefreshInfo
		_, err := iter.Next(&entry)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func GetRefreshTokenBySessionToken(ctx context.Context, db *database.Database, sessionToken string) (*RefreshInfo, error) {
	query := datastore.NewQuery("RefreshInfo").Filter("sessionToken =", sessionToken)

	iter := db.Run(ctx, query)

	var entry RefreshInfo

	_, err := iter.Next(&entry)
	// entryがない場合はnilを返す
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func GetRefreshToken(ctx context.Context, db *database.Database, refreshToken string) (*RefreshInfo, error) {
	key := database.CreateNameKey("RefreshInfo", refreshToken)
	var entry RefreshInfo
	isNotExist, err := db.Get(ctx, key, &entry)
	if err != nil {
		return nil, err
	}
	// entryが存在しない場合はnilを返す
	if isNotExist {
		return nil, nil
	}

	return &entry, nil
}

func GetRefreshTokenTX(tx *database.Transaction, refreshToken string) (*RefreshInfo, error) {
	key := database.CreateNameKey("RefreshInfo", refreshToken)
	var entry RefreshInfo
	isNotExist, err := tx.Get(key, &entry)
	if err != nil {
		return nil, err
	}
	// entryが存在しない場合はnilを返す
	if isNotExist {
		return nil, nil
	}

	return &entry, nil
}

func DeleteRefreshTokenPeriod(ctx context.Context, db *database.Database) error {
	query := datastore.NewQuery("RefreshInfo")

	var entities []RefreshInfo

	_, err := db.GetAll(ctx, query, &entities)
	if err != nil {
		return err
	}

	var keys []*datastore.Key

	for _, entity := range entities {
		if CheckExpired(&entity.Period) {
			key := database.CreateNameKey("RefreshInfo", entity.RefreshToken)
			keys = append(keys, key)
		}
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteRefreshTokenTX(tx *database.Transaction, refreshToken string) error {
	key := database.CreateNameKey("RefreshInfo", refreshToken)
	return tx.Delete(key)
}

func DeleteRefreshToken(ctx context.Context, db *database.Database, refreshToken string) error {
	key := database.CreateNameKey("RefreshInfo", refreshToken)
	return db.Delete(ctx, key)
}

func (c *RefreshInfo) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("RefreshInfo", c.RefreshToken)
	return db.Put(ctx, key, c)
}

func (c *RefreshInfo) AddTX(tx *database.Transaction) error {
	key := database.CreateNameKey("RefreshInfo", c.RefreshToken)
	return tx.Put(key, c)
}

func DeleteRefreshByUserId(ctx context.Context, db *database.Database, userId string) error {
	query := datastore.NewQuery("RefreshInfo").Filter("userId =", userId)

	var dummy []RefreshInfo

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}
