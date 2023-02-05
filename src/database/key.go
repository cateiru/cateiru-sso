package database

import (
	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/src/config"
)

// datasotoreのkeyを作成します
func CreateNameKey(tableName string, keyName string) *datastore.Key {
	return datastore.NameKey(tableName, keyName, createParentKey())
}

// 親レベルのKEYを作成します。
//
// DATASTORE_PARENT_KEYを使用します
func createParentKey() *datastore.Key {
	return datastore.NameKey(config.Defs.DatastoreParentKey, "default", nil)
}
