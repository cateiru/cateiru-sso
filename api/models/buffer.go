package models

import (
	"context"

	"github.com/cateiru/cateiru-sso/api/database"
)

// func GetCreateAccountBufferByBufferToken(ctx context.Context, db *database.Database, token string) (*CreateAccountBuffer, error) {

// }

func (c *CreateAccountBuffer) Add(ctx context.Context, db *database.Database) error {
	key := database.CreateNameKey("CreateAccountBuffer", c.BufferToken)

	return db.Put(ctx, key, c)
}
