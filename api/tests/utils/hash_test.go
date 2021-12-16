package utils_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

// 同じseedのハッシュは同じ値
func TestSameHash(t *testing.T) {
	seed := "SAMPLE"

	hash1 := utils.NewHash(seed).SHA256()
	hash2 := utils.NewHash(seed).SHA256()

	require.Equal(t, hash1, hash2, "同じSEEDのhash値が違う")
}

// 違うseedのハッシュは違う値
func TestDiffrentHash(t *testing.T) {
	hash1 := utils.NewHash("SAMPLE").SHA256()
	hash2 := utils.NewHash("SOMPLE").SHA256()

	require.NotEqual(t, hash1, hash2, "違うSEEDのhash値が同じ")
}

// 同じseedのハッシュ(byte)は同じ値
func TestSameHashByte(t *testing.T) {
	seed := "SAMPLE"

	hash1 := utils.NewHash(seed).SHA256Byte()
	hash2 := utils.NewHash(seed).SHA256Byte()

	require.Equal(t, hash1, hash2, "同じSEEDのhash(byte)値が違う")
}

// 違うseedのハッシュ(byte)は違う値
func TestDiffrentHashByte(t *testing.T) {
	hash1 := utils.NewHash("SAMPLE").SHA256Byte()
	hash2 := utils.NewHash("SOMPLE").SHA256Byte()

	require.NotEqual(t, hash1, hash2, "違うSEEDのhash(byte)値が同じ")
}
