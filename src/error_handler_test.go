package src_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/stretchr/testify/require"
)

func TestHTTPError(t *testing.T) {
	t.Run("NewHTTPError", func(t *testing.T) {
		e := src.NewHTTPError(http.StatusOK, "text")

		require.Equal(t, e.Code, http.StatusOK)
		require.Equal(t, e.Message, "text")
		require.Equal(t, e.UniqueCode, 0)

		require.Equal(t, e.Error(), "code=200, message=text")
	})

	t.Run("newHTTPError error", func(t *testing.T) {
		err := errors.New("error message")

		e := src.NewHTTPError(http.StatusBadRequest, err)

		require.Equal(t, e.Error(), "code=400, message=error message")
	})

	t.Run("NewHTTPUniqueError", func(t *testing.T) {
		e := src.NewHTTPUniqueError(http.StatusOK, 10, "text")

		require.Equal(t, e.Code, http.StatusOK)
		require.Equal(t, e.Message, "text")
		require.Equal(t, e.UniqueCode, 10)

		require.Equal(t, e.Error(), "code=200, message=text, unique=10")
	})
}
