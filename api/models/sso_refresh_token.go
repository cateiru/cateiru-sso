package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

func GetSSORefreshTokenByRefreshToken(ctx context.Context, db *database.Database, refresToken string) (*SSORefreshToken, error) {
	key := database.CreateNameKey("SSORefreshToken", refresToken)

	var entity SSORefreshToken

	isEmpty, err := db.Get(ctx, key, &entity)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return nil, nil
	}

	return &entity, nil
}

func DeleteSSORefreshTokenByClientId(ctx context.Context, db *database.Database, clientId string) error {
	query := datastore.NewQuery("SSORefreshToken").Filter("clientId =", clientId)

	var dummy []SSORefreshToken

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func (c *SSORefreshToken) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSORefreshToken", c.SSORefreshToken)

	return db.Put(ctx, key, c)
}
