package models

import (
	"cars/util/random"
	"fmt"
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

type GetCarsRequest struct {
	VIN   string `json:"vin,omitempty"`
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
	Year  string `json:"year,omitempty"`
	Color string `json:"color,omitempty"`
}

func (r *GetCarsRequest) Validate() error {
	if r.VIN == "" && r.Make == "" && r.Model == "" && r.Year == "" && r.Color == "" {
		return fmt.Errorf("can't get car, all fields are empty, expected at least one field to be set, vin, make, model, year, or color")
	}
	return nil
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
