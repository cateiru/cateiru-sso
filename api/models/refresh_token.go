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

func (c *RefreshInfo) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("RefreshInfo", c.RefreshToken)
	return db.Put(ctx, key, c)
}