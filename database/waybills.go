package database

import (
	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WaybillsRepository struct {
	DB *gorm.DB
}

func (w *WaybillsRepository) GetAllWaybills(limit, offset *int) ([]*models.Waybill, error) {
	var waybills []*models.Waybill

	result := w.DB.Model(&waybills)
	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}

	if err := result.Preload("Driver").Preload("Car").Preload("User").Preload(clause.Associations).Find(&waybills).Error; err != nil {
		return nil, err
	}

	return waybills, nil
}

func (w *WaybillsRepository) CreateWaybill(waybill *models.Waybill) (*models.Waybill, error) {
	if err := w.DB.Create(&waybill).Preload("Driver").Preload("Car").Preload("User").Preload(clause.Associations).Error; err != nil {
		return nil, err
	}
	return waybill, nil
}

func (w *WaybillsRepository) GetWaybillByID(id string) (*models.Waybill, error) {
	var waybill models.Waybill
	if err := w.DB.Where("id = ?", id).First(&waybill).Preload("Driver").Preload("Car").Preload("User").Preload(clause.Associations).Error; err != nil {
		return nil, err
	}
	return &waybill, nil
}

func (w *WaybillsRepository) UpdateWaybill(id string, waybill *models.Waybill) (*models.Waybill, error) {
	if err := w.DB.Model(&waybill).Where("id = ?", id).Updates(models.Waybill{
		UserID:              waybill.UserID,
		DriverID:            waybill.DriverID,
		CarID:               waybill.CarID,
		DateStart:           waybill.DateStart,
		DateEnd:             waybill.DateEnd,
		MileageStart:        waybill.MileageStart,
		MileageEnd:          waybill.MileageEnd,
		FuelFill:            waybill.FuelFill,
		FuelConsumptionFact: waybill.FuelConsumptionFact,
		FuelRemainingStart:  waybill.FuelRemainingStart,
		FuelRemainingEnd:    waybill.FuelRemainingEnd,
		IsActive:            waybill.IsActive,
	}).Preload("Driver").Preload("Car").Preload("User").Preload(clause.Associations).Error; err != nil {
		return nil, err
	}
	return waybill, nil
}

func (w *WaybillsRepository) DeleteWaybill(id string) (string, error) {
	if err := w.DB.Where("id = ?", id).Delete(&models.Waybill{}).Error; err != nil {
		return "", err
	}
	return "Record deleted from database.", nil
}

func (w *WaybillsRepository) GetAllWaybillsByUserID(id string, limit, offset *int) ([]*models.Waybill, error) {
	var waybills []*models.Waybill

	result := w.DB.Model(&waybills)
	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}

	if err := result.Where("user_id = ?", id).Preload("Driver").Preload("Car").Preload("User").Preload(clause.Associations).Find(&waybills).Error; err != nil {
		return nil, err
	}

	return waybills, nil
}
