package src

import (
	"context"
	"database/sql"
)

// トランザクション組む
func TxDB(ctx context.Context, db *sql.DB, f func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
