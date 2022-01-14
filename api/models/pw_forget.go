package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/database"
	"google.golang.org/api/iterator"
)

func GetPWForgetByToken(ctx context.Context, db *database.Database, token string) (*PWForget, error) {
	key := database.CreateNameKey("PWForget", token)
	var entity PWForget

	isNotExist, err := db.Get(ctx, key, &entity)
	if err != nil {
		return nil, err
	}

	// 要素がない場合
	if isNotExist {
		return nil, nil
	}

	return &entity, nil
}

func GetPWForgetByMail(ctx context.Context, db *database.Database, mail string) ([]PWForget, error) {
	query := datastore.NewQuery("PWForget").Filter("mail =", mail)

	iter := db.Run(ctx, query)

	var entities []PWForget

	for {
		var entity PWForget
		_, err := iter.Next(&entity)
		// 要素がなにもない場合nilを返す
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func DeletePWForgetByToken(ctx context.Context, db *database.Database, token string) error {
	key := database.CreateNameKey("PWForget", token)
	return db.Delete(ctx, key)
}

func (c *PWForget) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("PWForget", c.ForgetToken)

	return db.Put(ctx, key, c)
}
