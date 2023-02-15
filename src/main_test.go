package src_test

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

var DB *sql.DB
var C *src.Config

// これをしないとテストが失敗するため追加している
// ref. https://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests
var _ = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")
var _ = flag.String("test.config", "", "Overrides the default config")

func TestMain(m *testing.M) {
	C = src.TestConfig

	ctx := context.Background()
	db, err := sql.Open("mysql", C.DatabaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	err = ResetDBTable(ctx, db)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	code := m.Run()
	os.Exit(code)
}

// テスト用にテーブルをクリアする
func ResetDBTable(ctx context.Context, db *sql.DB) error {
	rows, err := queries.Raw("show tables").QueryContext(ctx, db)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		table := ""
		if err := rows.Scan(&table); err != nil {
			return err
		}

		// SQLインジェクションの影響は無いためSprintfを使用している
		if _, err := queries.Raw(fmt.Sprintf("TRUNCATE TABLE %s", table)).Exec(db); err != nil {
			return err
		}
	}

	return nil
}
