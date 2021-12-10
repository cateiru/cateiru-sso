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

func (c *SessionInfo) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SessionInfo", c.SessionToken)
	return db.Put(ctx, key, c)
}
