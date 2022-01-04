package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

func GetSSOServiceByPublicKey(ctx context.Context, db *database.Database, publicKey string) (*SSOService, error) {
	key := database.CreateNameKey("SSOService", publicKey)

	var entity SSOService
	isEmpty, err := db.Get(ctx, key, &entity)
	if isEmpty {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func GetSSOServiceByUserID(ctx context.Context, db *database.Database, userID string) ([]SSOService, error) {
	query := datastore.NewQuery("SSOService").Filter("userId =", userID)

	var entities []SSOService

	_, err := db.GetAll(ctx, query, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (c *SSOService) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSOService", c.SSOPublicKey)

	return db.Put(ctx, key, c)
}
