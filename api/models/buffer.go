package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

func GetCreateAccountBufferByBufferToken(ctx context.Context, db *database.Database, token string) (*CreateAccountBuffer, error) {
	query := datastore.NewQuery("CreateAccountBuffer").Filter("bufferToken =", token)
	iter := db.Run(ctx, query)

	var entry CreateAccountBuffer
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

func (c *CreateAccountBuffer) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("CreateAccountBuffer", c.BufferToken)

	return db.Put(ctx, key, c)
}
