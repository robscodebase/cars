package models

import (
	"cars/util/random"
)

type Car struct {
	VIN   string `json:"vin" gorm:"column:vin;primary_key;not null;unique;index"`
	Make  string `json:"make" gorm:"column:make;index"`
	Model string `json:"model" gorm:"column:model;index"`
	Year  string `json:"year" gorm:"column:year;index"`
	Color string `json:"color" gorm:"column:color;index"`
}

func NewCar(vin, make, model, color, year string) *Car {
	return &Car{
		VIN:   vin,
		Make:  make,
		Model: model,
		Year:  year,
		Color: color,
	}
}

type CreateCarRequest struct {
	VIN   string `json:"vin" binding:"required"`
	Make  string `json:"make" binding:"required"`
	Model string `json:"model" binding:"required"`
	Year  string `json:"year" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type GetCarRequest struct {
	VIN string `param:"id" binding:"required"`
}

func RandomCars(i int) []*Car {
	cars := make([]*Car, i)
	for j := 0; j < i; j++ {
		cars[j] = RandomCar()
	}
	return cars
}

func RandomCar() *Car {
	return NewCar(
		random.String(),
		random.String(),
		random.String(),
		random.String(),
		random.String(),
	)
}
