package models

import (
	"context"

	"github.com/cateiru/cateiru-sso/src/database"
)

func GetRoleByUserID(ctx context.Context, db *database.Database, userId string) (*Role, error) {
	key := database.CreateNameKey("Role", userId)

	var role Role

	isEmpty, err := db.Get(ctx, key, &role)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return nil, nil
	}

	return &role, nil
}

func DeleteRoleByUserID(ctx context.Context, db *database.Database, userId string) error {
	key := database.CreateNameKey("Role", userId)

	return db.Delete(ctx, key)
}

func (c *Role) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("Role", c.UserId.UserId)

	return db.Put(ctx, key, c)
}
