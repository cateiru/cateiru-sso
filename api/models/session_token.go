package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

func GetSessionTokenByUserId(ctx context.Context, db *database.Database, userId string) ([]SessionInfo, error) {
	query := datastore.NewQuery("SessionInfo").Filter("userId =", userId)

	iter := db.Run(ctx, query)

	var entries []SessionInfo

	for {
		var entry SessionInfo
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

func GetSessionToken(ctx context.Context, db *database.Database, sessionToken string) (*SessionInfo, error) {
	key := database.CreateNameKey("SessionInfo", sessionToken)
	var entry SessionInfo
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

func GetSessionTokenTX(tx *database.Transaction, sessionToken string) (*SessionInfo, error) {
	key := database.CreateNameKey("SessionInfo", sessionToken)
	var entry SessionInfo
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

func DeleteSessionTokenTX(tx *database.Transaction, sessionToken string) error {
	key := database.CreateNameKey("SessionInfo", sessionToken)
	return tx.Delete(key)
}

func DeleteSessionToken(ctx context.Context, db *database.Database, sessionToken string) error {
	key := database.CreateNameKey("SessionInfo", sessionToken)
	return db.Delete(ctx, key)
}

func (c *SessionInfo) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SessionInfo", c.SessionToken)
	return db.Put(ctx, key, c)
}

func (c *SessionInfo) AddTX(tx *database.Transaction) error {
	key := database.CreateNameKey("SessionInfo", c.SessionToken)
	return tx.Put(key, c)
}

func DeleteSessionByUserId(ctx context.Context, db *database.Database, userId string) error {
	query := datastore.NewQuery("SessionInfo").Filter("userId =", userId)

	var dummy []SessionInfo

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteSessionTokenPeriod(ctx context.Context, db *database.Database) error {
	query := datastore.NewQuery("SessionInfo")

	var entities []SessionInfo

	_, err := db.GetAll(ctx, query, &entities)
	if err != nil {
		return err
	}

	var keys []*datastore.Key

	for _, entity := range entities {
		if CheckExpired(&entity.Period) {
			key := database.CreateNameKey("SessionInfo", entity.SessionToken)
			keys = append(keys, key)
		}
	}

	return db.DeleteMulti(ctx, keys)
}
