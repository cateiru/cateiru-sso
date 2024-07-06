package lib_test

import (
	"fmt"
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	p := &lib.Password{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}

	t.Run("ハッシュ化して戻せる", func(t *testing.T) {
		password := "password"
		hashed, salt, err := p.HashPassword(password)
		require.NoError(t, err)

		require.NotEqual(t, password, hashed)
		require.NotEqual(t, salt, hashed)

		t.Run("同じPWの場合は認証できる", func(t *testing.T) {
			password2 := "password"
			result := p.VerifyPassword(password2, hashed, salt)
			require.True(t, result)
		})

		t.Run("違うPWの場合は認証できない", func(t *testing.T) {
			password2 := "aaaa"
			result := p.VerifyPassword(password2, hashed, salt)
			require.False(t, result)
		})
	})

	t.Run("KeyLen", func(t *testing.T) {
		password := "password"
		for i := 5; 40 > i; i += 10 {
			t.Run(fmt.Sprintf("len: %d", i), func(t *testing.T) {
				pp := &lib.Password{
					Time:    p.Time,
					Memory:  p.Memory,
					Threads: p.Threads,
					KeyLen:  uint32(i),
				}
				hashed, salt, err := pp.HashPassword(password)
				require.NoError(t, err)

				require.Len(t, hashed, i)
				require.Len(t, salt, i)
			})
		}
	})
}
