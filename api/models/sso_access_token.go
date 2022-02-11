package models

import (
	"context"

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

func (c *SSOAccessToken) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSOAccessToken", c.SSOAccessToken)
	return db.Put(ctx, key, c)
}
