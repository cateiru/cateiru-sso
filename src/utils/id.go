// ID, UUIDを生成します。
//
// IDはUUIDをseedにしたSHA256のハッシュ値（長さ指定）を返します。
//
// Example:
//	uuid := UUID()
//
//	idMax := CreateID(0) // 長さ0だと最大長を返す
//	id10 := CreateID(10)
//
package utils

import (
	"github.com/google/uuid"
)

const MAX_ID_LENGTH = 64

// UUID生成
func UUID() string {
	return uuid.NewString()
}

// 指定行のIDを作成
// 最大 64文字
// 0を指定すると最大値
func CreateID(length int) string {
	hash := NewHash(UUID())

	result := hash.SHA256()

	if length == 0 {
		return result
	}
	return result[:length]
}
