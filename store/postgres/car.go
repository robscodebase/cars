package postgres

import (
	"cars/models"
	"cars/util"
	"fmt"
	"github.com/jinzhu/gorm"
)

func (p *Postgres) CreateCar(car *models.Car) (*models.Car, error) {
	if car == nil {
		return nil, ErrorInputIsNil
	}
	err := p.DB.Create(car).Error
	if err != nil {
		return nil, util.ErrorAndLog(err, "can't create car in postgres, error creating car")
	}
	return car, nil
}

func (p *Postgres) GetCar(vin string) (*models.Car, error) {
	if vin == "" {
		return nil, ErrorInputIsNil
	}
	car := &models.Car{}
	err := p.DB.Where("vin = ?", vin).First(car).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrorCarNotFound
		}
		return nil, util.ErrorAndLog(err, fmt.Sprintf("can't read car in postgres, error reading car with vin %s", vin))
	}
	return car, nil
}

func (p *Postgres) UpdateCar(car *models.Car) (*models.Car, error) {
	if car == nil {
		return nil, ErrorInputIsNil
	}
	err := p.DB.Save(car).Error
	if err != nil {
		return nil, util.ErrorAndLog(err, "can't update car in postgres, error updating car")
	}
	return car, nil
}

func (p *Postgres) DeleteCar(vin string) error {
	if vin == "" {
		return ErrorInputIsNil
	}
	err := p.DB.Where("vin = ?", vin).Delete(&models.Car{}).Error
	if err != nil {
		return util.ErrorAndLog(err, fmt.Sprintf("can't delete car in postgres, error deleting car with vin %s", vin))
	}
	return nil
}
