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

	t.Run("NewOIDCError", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			e := src.NewOIDCError(http.StatusBadRequest, src.ErrLoginRequired, "message", "uri", "state")

			require.Equal(t, e.Code, http.StatusBadRequest)
			require.Equal(t, e.AuthenticationErrorResponse.Error, src.ErrLoginRequired)
			require.Equal(t, e.AuthenticationErrorResponse.ErrorDescription, "message")
			require.Equal(t, e.AuthenticationErrorResponse.ErrorURI, "uri")
			require.Equal(t, e.AuthenticationErrorResponse.State, "state")

			require.Equal(t, e.Error(), "code=400, error=login_required, message=message, error_uri=uri, state=state")
		})

		t.Run("no uri", func(t *testing.T) {
			e := src.NewOIDCError(http.StatusBadRequest, src.ErrLoginRequired, "message", "", "state")
			require.Equal(t, e.Error(), "code=400, error=login_required, message=message, state=state")
		})

		t.Run("no state", func(t *testing.T) {
			e := src.NewOIDCError(http.StatusBadRequest, src.ErrLoginRequired, "message", "uri", "")
			require.Equal(t, e.Error(), "code=400, error=login_required, message=message, error_uri=uri")
		})

		t.Run("no uri and state", func(t *testing.T) {
			e := src.NewOIDCError(http.StatusBadRequest, src.ErrLoginRequired, "message", "", "")
			require.Equal(t, e.Error(), "code=400, error=login_required, message=message")
		})
	})
}
