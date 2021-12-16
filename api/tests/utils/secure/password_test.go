package secure_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/stretchr/testify/require"
)

func TestPasswordHash(t *testing.T) {
	rawPassword := "sdidf04o3nkasd-sda3q:zx"

	hashed := secure.PWHash(rawPassword)

	require.NotEmpty(t, hashed.Key)
	require.NotEmpty(t, hashed.Salt)

	isEqual := secure.ValidatePW(rawPassword, hashed.Key, hashed.Salt)

	require.True(t, isEqual)
}

// 同じパスワードでもSaltが違うためハッシュは違う値となる
func TestSomePasswordsHash(t *testing.T) {
	rawPassword := "q2oi30vifa:;q3o4-wq90scu9ewe@as"

	hashed1 := secure.PWHash(rawPassword)
	hashed2 := secure.PWHash(rawPassword)

	require.NotEqual(t, hashed1.Key, hashed2.Key)
	require.NotEqual(t, hashed1.Salt, hashed2.Salt)
}

func TestFailedPassword(t *testing.T) {
	rawPassword := "dffiwuri29as-d-1-180asdh]qwe1:"
	rawPassword2 := "sdp9r0w2as:dcd9c0"

	hashed := secure.PWHash(rawPassword)

	isEqual := secure.ValidatePW(rawPassword2, hashed.Key, hashed.Salt)

	require.False(t, isEqual)
}
