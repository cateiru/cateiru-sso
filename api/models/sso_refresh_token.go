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

func DeleteSSORefreshTokenByUserId(ctx context.Context, db *database.Database, userId string) error {
	query := datastore.NewQuery("SSORefreshToken").Filter("userId =", userId)

	var dummy []SSORefreshToken

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteSSORefreshTokenByRefreshToken(ctx context.Context, db *database.Database, refresh string) error {
	key := database.CreateNameKey("SSORefreshToken", refresh)

	return db.Delete(ctx, key)
}

func (c *SSORefreshToken) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSORefreshToken", c.SSORefreshToken)

	return db.Put(ctx, key, c)
}
