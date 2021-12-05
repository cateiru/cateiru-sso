package net

import (
	"errors"
	"net/http"
)

// URLのクエリパラメータを取得します
func GetQuery(r *http.Request, key string) (string, error) {
	query := r.URL.Query().Get(key)

	if len(query) == 0 {
		return "", errors.New("query is empty")
	}

	return query, nil
}
