package database

import (
	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
)

type DriversRepository struct {
	DB *gorm.DB
}

func (d *DriversRepository) GetDrivers(limit, offset *int) ([]*models.Driver, error) {
	var drivers []*models.Driver

	result := d.DB.Model(&drivers)

	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}

	if err := result.Find(&drivers).Error; err != nil {
		return nil, err
	}

	return drivers, nil
}

func (d *DriversRepository) CreateDriver(driver *models.Driver) (*models.Driver, error) {
	if err := d.DB.Create(&driver).Error; err != nil {
		return nil, err
	}
	return driver, nil
}

func (d *DriversRepository) GetDriverByID(id string) (*models.Driver, error) {
	var driver models.Driver
	if err := d.DB.Where("id = ?", id).First(&driver).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (d *DriversRepository) Update(id string, driver *models.Driver) (*models.Driver, error) {
	if err := d.DB.Model(&driver).Where("id = ?").Updates(models.Driver{
		FirstName:  driver.FirstName,
		SecondName: driver.SecondName,
		Patronymic: driver.Patronymic,
		IsActive:   driver.IsActive,
	}).Error; err != nil {
		return nil, err
	}
	return driver, nil
}

func (d *DriversRepository) Delete(id string) (string, error) {
	if err := d.DB.Where("id = ?", id).Delete(&models.Driver{}).Error; err != nil {
		return "", err
	}
	return "Driver successfully deleted.", nil
}
