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

	result := c.DB.Model(&cars)
	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}

	if err := result.Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *CarsRepository) GetCarByID(id string) (*models.Car, error) {
	var car models.Car
	if err := c.DB.Where("%v = ?", id).First(&car).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func (c *CarsRepository) CreateCar(car *models.Car) (*models.Car, error) {
	if err := c.DB.Create(&car).Error; err != nil {
		return nil, err
	}
	return car, nil
}

func (c *CarsRepository) UpdateCar(id string, car *models.Car) (*models.Car, error) {
	if err := c.DB.Model(&car).Where("id = ?", id).Updates(models.Car{
		Brand:           car.Brand,
		Number:          car.Number,
		Fuel:            car.Fuel,
		Mileage:         car.Mileage,
		FuelConsumption: car.FuelConsumption,
		FuelRemaining:   car.FuelRemaining,
	}).Error; err != nil {
		return nil, err
	}
	return car, nil
}

func (c *CarsRepository) DeleteCar(id string) (string, error) {
	if err := c.DB.Where("id = ?", id).Delete(&models.Car{}).Error; err != nil {
		return "", err
	}
	return "Car successfully deleted.", nil
}
