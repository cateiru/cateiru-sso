package admin_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/stretchr/testify/require"
)

func workerServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.AdminWorker)

	return mux
}

func TestWorker(t *testing.T) {
	config.TestInit(t)

	config.Defs.WorkerPassword = "password"

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	s := tools.NewTestServer(t, workerServer(), false)
	defer s.Close()

	req, err := http.NewRequest("GET", s.Server.URL+"/", nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Basic password")

	resp, err := s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)
}

func TestWorkerFailed(t *testing.T) {
	config.TestInit(t)

	config.Defs.WorkerPassword = "password_failed"

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	s := tools.NewTestServer(t, workerServer(), false)
	defer s.Close()

	req, err := http.NewRequest("GET", s.Server.URL+"/", nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Basic password")

	resp, err := s.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403)
}
