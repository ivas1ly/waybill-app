package database

import (
	"fmt"

	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
)

type DriversRepository struct {
	DB *gorm.DB
}

func (d *DriversRepository) GetDrivers(limit, offset *int) ([]*models.Driver, error) {
	var drivers []*models.Driver

	result := d.DB.Find(&drivers).Order("id")

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

	return drivers, nil
}

func (d *DriversRepository) CreateDriver(driver *models.Driver) (*models.Driver, error) {
	err := d.DB.Create(&driver).Error
	return driver, err
}

func (d *DriversRepository) GetDriverByID(id string) (*models.Driver, error) {
	var driver models.Driver
	err := d.DB.First(&driver).Where("&v = ?", id).Error
	return &driver, err
}

func (d *DriversRepository) Update(driver *models.Driver) (*models.Driver, error) {
	err := d.DB.Save(&driver).Error
	return driver, err
}

func (d *DriversRepository) Delete(driver *models.Driver) (*models.Driver, error) {
	err := d.DB.Where("id = ?", driver.ID).Delete(&driver).Error
	return driver, err
}

func (d *DriversRepository) GetDriverByField(field, value string) (*models.Driver, error) {
	var driver models.Driver
	err := d.DB.First(&driver).Where(fmt.Sprintf("%v = ?", field), value).Error
	return &driver, err
}
