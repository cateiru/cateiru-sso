package database

import (
	"os"

	"cloud.google.com/go/datastore"
)

// datasotoreのkeyを作成します
func CreateNameKey(tableName string, keyName string) *datastore.Key {
	return datastore.NameKey(tableName, keyName, createParentKey())
}

// 親レベルのKEYを作成します。
//
// DATASTORE_PARENT_KEYを使用します
func createParentKey() *datastore.Key {
	parentKey := os.Getenv("DATASTORE_PARENT_KEY")
	if len(parentKey) == 0 {
		parentKey = "cateiru-sso"
	}
	return datastore.NameKey(parentKey, "default", nil)
}
