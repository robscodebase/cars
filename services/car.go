package services

import (
	"cars/models"
	"cars/store"
)

func CreateCar(store *store.Store, car *models.Car) (*models.Car, error) {
	return store.Postgres.CreateCar(car)
}

func GetCar(store *store.Store, vin string) (*models.Car, error) {
	return store.Postgres.GetCar(vin)
}

func UpdateCar(store *store.Store, car *models.Car) (*models.Car, error) {
	return store.Postgres.UpdateCar(car)
}

func DeleteCar(store *store.Store, vin string) error {
	return store.Postgres.DeleteCar(vin)
}
