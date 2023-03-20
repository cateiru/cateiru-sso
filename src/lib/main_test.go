package lib_test

import (
	"flag"
	"os"
	"testing"
)

// これをしないとテストが失敗するため追加している
// ref. https://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests
var _ = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")
var _ = flag.String("test.config", "", "Overrides the default config")

func TestMain(m *testing.M) {
	// ローカルのCloud Storage接続
	os.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")

	flag.Parse()

	code := m.Run()
	os.Exit(code)
}
