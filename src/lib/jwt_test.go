package lib_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

const JWT_PUBLIC_KEY_PATH = "jwt/jwt.pkcs8"

func TestJsonWebKeys(t *testing.T) {
	resp, err := lib.JsonWebKeys(JWT_PUBLIC_KEY_PATH, "RS256", "sig", "test")
	require.NoError(t, err)

	snaps.MatchSnapshot(t, resp)
}
