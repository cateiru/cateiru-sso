package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
)

func GetSSOServiceByClientId(ctx context.Context, db *database.Database, clientId string) (*SSOService, error) {
	key := database.CreateNameKey("SSOService", clientId)

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

func DeleteSSOServiceByClientId(ctx context.Context, db *database.Database, clientId string) error {
	key := database.CreateNameKey("SSOService", clientId)

	return db.Delete(ctx, key)
}

func (c *SSOService) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSOService", c.ClientID)

	return db.Put(ctx, key, c)
}
