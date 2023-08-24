package integration_tests

import (
	"cars/controller"
	"cars/store"
	tests2 "cars/util/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

const carTestAmount = 100

// IntegrationTestsRun runs all integration tests
// Postgres and Redis must be running
func TestIntegrationRun(t *testing.T) {
	tests2.SetTestEnvVars()
	router, store := startServer(t)
	server := httptest.NewServer(router)
	defer server.Close()
	client := server.Client()
	t.Run("Test Suite", func(t *testing.T) {
		t.Run("Cars", func(t *testing.T) {
			t.Parallel()
			t.Run("CreateCars", func(t *testing.T) {
				sharedTestCreateCars(t, server, client, store)
			})
			t.Run("GetCars", func(t *testing.T) {
				sharedTestGetCars(t, server, client, store)
			})
			t.Run("UpdateCars", func(t *testing.T) {
				sharedTestUpdateCars(t, server, client, store)
			})
			t.Run("DeleteCars", func(t *testing.T) {
				sharedTestDeleteCars(t, server, client, store)
			})
		})

	})
}

func startServer(t *testing.T) (*gin.Engine, *store.Store) {
	s, err := store.NewStore()
	require.NoError(t, err)
	require.NotNil(t, s)
	return controller.InitServer(s), s
}
