package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

func GetSSOServiceLogsByUserId(ctx context.Context, db *database.Database, userId string) ([]SSOServiceLog, error) {
	query := datastore.NewQuery("SSOServiceLog").Filter("userId =", userId)
	iter := db.Run(ctx, query)

	logs := []SSOServiceLog{}

	for {
		var entity SSOServiceLog

		_, err := iter.Next(&entity)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		logs = append(logs, entity)
	}

	return logs, nil
}

func GetSSOServiceLogsByClientId(ctx context.Context, db *database.Database, clientId string) ([]SSOServiceLog, error) {
	query := datastore.NewQuery("SSOServiceLog").Filter("clientId =", clientId)
	iter := db.Run(ctx, query)

	logs := []SSOServiceLog{}

	for {
		var entity SSOServiceLog

		_, err := iter.Next(&entity)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		logs = append(logs, entity)
	}

	return logs, nil
}

func CountSSOServiceLogByClientId(ctx context.Context, db *database.Database, clientId string) (int, error) {
	query := datastore.NewQuery("SSOServiceLog").Filter("clientId =", clientId)

	return db.Count(ctx, query)
}

func DeleteSSOServiceLogByClientId(ctx context.Context, db *database.Database, clientId string) error {
	query := datastore.NewQuery("SSOServiceLog").Filter("clientId =", clientId)

	var dummy []SSOServiceLog

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func DeleteSSOServiceLogByUserId(ctx context.Context, db *database.Database, userId string) error {
	query := datastore.NewQuery("SSOServiceLog").Filter("userId =", userId)

	var dummy []SSOServiceLog

	keys, err := db.GetAll(ctx, query, &dummy)
	if err != nil {
		return err
	}

	return db.DeleteMulti(ctx, keys)
}

func (c *SSOServiceLog) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("SSOServiceLog", c.LogId)
	return db.Put(ctx, key, c)
}
