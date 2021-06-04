package database

import (
	"github.com/99designs/gqlgen/example/scalars/model"
	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
)

type WaybillsRepository struct {
	DB *gorm.DB
}

func (w *WaybillsRepository) GetWaybills(limit, offset *int) ([]*models.Waybill, error) {
	var waybills []*models.Waybill

	result := w.DB.Find(&waybills).Order("id")

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

func (w *WaybillsRepository) Delete(waybill *models.Waybill) (*models.Waybill, error) {
	err := w.DB.Where("id = ?", waybill.ID).Delete(&waybill).Error
	return waybill, err
}

func (w *WaybillsRepository) GetWaybillsByUserID(user *model.User, limit, offset *int) ([]*models.Waybill, error) {
	var waybills []*models.Waybill
	result := w.DB.Find(&waybills).Where("user_id = ?", user.ID)

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

	return waybills, nil
}
