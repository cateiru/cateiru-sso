package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

func GetAccessTokenByAccessToken(ctx context.Context, db *database.Database, token string) (*SSOAccessToken, error) {
	key := database.CreateNameKey("SSOAccessToken", token)

	var m SSOAccessToken

	isEmpty, err := db.Get(ctx, key, &m)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return nil, nil
	}

	return &m, nil
}

func DeleteAccessTokenByClientID(ctx context.Context, db *database.Database, clientID string) error {
	query := datastore.NewQuery("SSOAccessToken").Filter("clientId =", clientID)

	var dummy []SSOAccessToken

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteAccessTokenByUserId(ctx context.Context, db *database.Database, userId string) error {
	query := datastore.NewQuery("SSOAccessToken").Filter("userId =", userId)

	var dummy []SSOAccessToken

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteAccessTokenByAccessToken(ctx context.Context, db *database.Database, accessToken string) error {
	key := database.CreateNameKey("SSOAccessToken", accessToken)

	return db.Delete(ctx, key)
}

func (c *SSOAccessToken) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSOAccessToken", c.SSOAccessToken)
	return db.Put(ctx, key, c)
}
