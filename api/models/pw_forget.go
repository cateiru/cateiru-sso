package models

import (
	"context"

	"github.com/cateiru/cateiru-sso/api/database"
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

func (c *PWForget) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("PWForget", c.ForgetToken)

	return db.Put(ctx, key, c)
}
