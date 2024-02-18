package src_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestFastlyTrust(t *testing.T) {
	trustOptions := src.FastlyTrust()

	require.Len(t, trustOptions, 22)
}
