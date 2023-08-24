package integration_tests

import (
	"cars/config"
	"cars/controller"
	"cars/models"
	"cars/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"log"
	"net/http/httptest"
	"testing"
)

const carTestAmount = 100

// IntegrationTestsRun runs all integration tests
// Postgres and Redis must be running
func TestIntegrationRun(t *testing.T) {
	err := config.Init("../.env")
	require.NoError(t, err)
	r, s := startServer(t)
	if config.Get().Postgres.RunMigrations == "true" {
		err = s.Postgres.DB.AutoMigrate(
			&models.Car{},
		).Error
		if err != nil {
			log.Fatalf("main failed to run migrations: %v", err)
		}
	}
	server := httptest.NewServer(r)
	defer server.Close()
	client := server.Client()
	t.Run("Test Suite", func(t *testing.T) {
		t.Run("Cars", func(t *testing.T) {
			t.Parallel()
			t.Run("CreateCars", func(t *testing.T) {
				sharedTestCreateCars(t, server, client, s)
			})
			t.Run("GetCars", func(t *testing.T) {
				sharedTestGetCars(t, server, client, s)
			})
			t.Run("UpdateCars", func(t *testing.T) {
				sharedTestUpdateCars(t, server, client, s)
			})
			t.Run("DeleteCars", func(t *testing.T) {
				sharedTestDeleteCars(t, server, client, s)
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
