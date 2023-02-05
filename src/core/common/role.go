package common

import (
	"context"
	"errors"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/go-http-error/httperror/status"
)

// Pro以上のユーザのみ
func ProOnly(ctx context.Context, db *database.Database, userId string) error {
	return findRole(ctx, db, userId, []string{"pro"})
}

func AdminOnly(ctx context.Context, db *database.Database, userId string) error {
	return findRole(ctx, db, userId, []string{"admin"})
}

// roleを見る
func findRole(ctx context.Context, db *database.Database, userId string, targetRoles []string) error {
	role, err := models.GetRoleByUserID(ctx, db, userId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if role == nil {
		return status.NewBadRequestError(err).Caller()
	}

	for _, role := range role.Role {
		for _, target := range targetRoles {
			if role == target {
				return nil
			}
		}
	}

	return status.NewBadRequestError(errors.New("not authorized")).Caller()
}
