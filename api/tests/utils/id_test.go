package utils_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

// 生成されたUUIDが全部違うか
func TestUUID(t *testing.T) {
	uuids := []string{}
	length := 10

	for i := 0; length > i; i++ {
		uuids = append(uuids, utils.UUID())
	}

	for i, a := range uuids {
		for j := i + 1; length > j; j++ {
			require.NotEqual(t, a, uuids[j], "同じUUIDが存在する")
		}
	}
}

// 生成されたIDが全部違うか
func TestID(t *testing.T) {
	ids := []string{}
	length := 10

	for i := 0; length > i; i++ {
		ids = append(ids, utils.CreateID(0))
	}

	for i, a := range ids {
		for j := i + 1; length > j; j++ {
			require.NotEqual(t, a, ids[j], "同じIDが存在する")
		}
	}
}

// IDが設定したlengthになっているか
func TestIDLength(t *testing.T) {

	id := utils.CreateID(0)
	require.Equal(t, len(id), utils.MAX_ID_LENGTH, "生成されたIDの長さがおかしい")

	for i := 1; utils.MAX_ID_LENGTH >= i; i++ {
		id := utils.CreateID(i)
		require.Equal(t, len(id), i, "生成されたIDの長さがおかしい")
	}
}
