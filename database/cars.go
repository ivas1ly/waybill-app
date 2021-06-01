package database

import (
	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
)

type CarsRepository struct {
	DB *gorm.DB
}

func (c *CarsRepository) GetCars(limit, offset *int) ([]*models.Car, error) {
	var cars []*models.Car

	result := c.DB.Find(&cars).Order("id")

	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}
	err := result.Error
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *CarsRepository) GetCarByID(id string) (*models.Car, error) {
	var car models.Car
	err := c.DB.First(&car).Where("%v = ?", id).Error
	return &car, err
}

func (c *CarsRepository) Update(car *models.Car) (*models.Car, error) {
	err := c.DB.Save(&car).Error
	return car, err
}

func (c *CarsRepository) Delete(car *models.Car) (*models.Car, error) {
	err := c.DB.Where("id = ?", car.ID).Delete(&car).Error
	return car, err
}
