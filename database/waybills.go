package database

import (
	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WaybillsRepository struct {
	DB *gorm.DB
}

func (w *WaybillsRepository) GetWaybills(limit, offset *int) ([]*models.Waybill, error) {
	var waybills []*models.Waybill

	result := w.DB.Model(&waybills)

	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}
	err := result.Find(&waybills).Error
	if err != nil {
		return nil, err
	}

	return waybills, nil
}

func (w *WaybillsRepository) CreateWaybill(waybill *models.Waybill) (*models.Waybill, error) {
	err := w.DB.Create(&waybill).Error
	return waybill, err
}

func (w *WaybillsRepository) GetWaybillByID(id string) (*models.Waybill, error) {
	var waybill models.Waybill
	err := w.DB.First(&waybill).Where("id = ?", id).Error
	return &waybill, err
}

func (w *WaybillsRepository) Update(waybill *models.Waybill) (*models.Waybill, error) {
	err := w.DB.Save(&waybill).Error
	return waybill, err
}

func (w *WaybillsRepository) Delete(id string) (string, error) {
	if err := w.DB.Where("id = ?", id).Delete(&models.Waybill{}).Error; err != nil {
		return "", err
	}
	return "Record deleted from database.", nil
}

func (w *WaybillsRepository) GetWaybillsByUserID(id string, limit, offset *int) ([]*models.Waybill, error) {
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
