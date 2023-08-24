package integration_tests

import (
	"cars/models"
	"cars/store"
	tests2 "cars/util/tests"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func sharedTestCreateCars(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store) {
	cars := models.RandomCars(carTestAmount)
	t.Parallel()
	for _, car := range cars {
		t.Run("CreateCar", func(t *testing.T) {
			sharedTestCreateCar(t, server, client, store, car)
		})
	}
}

func sharedTestCreateCar(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store, car *models.Car) {
	req := tests2.SetAuthRequest(t, server, store.Redis, http.MethodPost, "/cars", car)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
}

func sharedTestGetCars(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store) {
	cars := models.RandomCars(carTestAmount)
	t.Parallel()
	for _, car := range cars {
		t.Run("GetCars", func(t *testing.T) {
			sharedTestGetCar(t, server, client, store, car)
		})
	}
}

func sharedTestGetCar(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store, car *models.Car) {
	sharedTestCreateCar(t, server, client, store, car)
	endpoint := fmt.Sprint("/cars/", car.VIN)
	req := tests2.SetAuthRequest(t, server, store.Redis, http.MethodGet, endpoint, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func sharedTestUpdateCars(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store) {
	cars := models.RandomCars(carTestAmount)
	t.Parallel()
	for _, car := range cars {
		t.Run("UpdateCars", func(t *testing.T) {
			sharedTestUpdateCar(t, server, client, store, car)
		})
	}
}

func sharedTestUpdateCar(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store, car *models.Car) {
	sharedTestCreateCar(t, server, client, store, car)
	endpoint := fmt.Sprint("/cars")
	car.Color = "blue"
	req := tests2.SetAuthRequest(t, server, store.Redis, http.MethodPut, endpoint, car)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	sharedTestGetCarColor(t, server, client, store, car)
}

func sharedTestGetCarColor(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store, car *models.Car) {
	endpoint := fmt.Sprint("/cars/", car.VIN)
	req := tests2.SetAuthRequest(t, server, store.Redis, http.MethodGet, endpoint, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var carResponse models.Car
	err = json.NewDecoder(resp.Body).Decode(&carResponse)
	require.NoError(t, err)
	require.Equal(t, car.Color, carResponse.Color)
}

func sharedTestDeleteCars(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store) {
	cars := models.RandomCars(carTestAmount)
	t.Parallel()
	for _, car := range cars {
		t.Run("DeleteCars", func(t *testing.T) {
			sharedTestDeleteCar(t, server, client, store, car)
		})
	}
}

func sharedTestDeleteCar(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store, car *models.Car) {
	sharedTestCreateCar(t, server, client, store, car)
	endpoint := fmt.Sprint("/cars/", car.VIN)
	req := tests2.SetAuthRequest(t, server, store.Redis, http.MethodDelete, endpoint, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	sharedTestGetCarNotFound(t, server, client, store, car)
}

func sharedTestGetCarNotFound(t *testing.T, server *httptest.Server, client *http.Client, store *store.Store, car *models.Car) {
	endpoint := fmt.Sprint("/cars/", car.VIN)
	req := tests2.SetAuthRequest(t, server, store.Redis, http.MethodGet, endpoint, nil)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
