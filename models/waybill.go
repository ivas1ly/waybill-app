package models

import (
	"time"

	"gorm.io/gorm"
)

type Waybill struct {
	// Идентификатор путевого листа.
	ID string `json:"id" gorm:"type:uuid;primaryKey;size:255;uniqueIndex;not null;default:gen_random_uuid()"`
	// Идентификатор пользователя.
	UserID string `json:"userID" gorm:"type:uuid;primaryKey;size:255;not null;"`
	// Идентификатор водителя.
	DriverID string `json:"driverID" gorm:"type:uuid;primaryKey;size:255;not null;"`
	// Идентификатор машины.
	CarID string `json:"carID" gorm:"type:uuid;primaryKey;size:255;not null;"`
	// Дата и время создания путевого листа.
	DateStart time.Time `json:"dateStart" gorm:"not null;"`
	// Дата и время закрытия путевого листа.
	DateEnd time.Time `json:"dateEnd"`
	// Показания спидометра при выезде.
	MileageStart float64 `json:"mileageStart" gorm:"type:integer;not null;"`
	// Показания спидометра при заезде.
	MileageEnd *float64 `json:"mileageEnd" gorm:"type:integer;"`
	// Заправлено топлива.
	FuelFill *float64 `json:"fuelFill"`
	// Расход топлива по факту
	FuelConsumptionFact float64 `json:"fuelConsumptionFact"`
	// Остаток топлива при выезде.
	FuelRemainingStart float64 `json:"fuelRemainingStart"`
	// Остаток топлива при заезде.
	FuelRemainingEnd *float64 `json:"fuelRemainingEnd"`
	// Возможность редактировать путевой лист.
	IsActive bool `json:"isActive" gorm:"not null;default:false"`
	// Водитель, к которому относится путевой лист.
	Driver *Driver `json:"driver"`
	// Пользователь, создавший путевой лист.
	User *User `json:"user"`
	// Машина.
	Car *Car `json:"car"`
	// Дата создания путевого листа.
	CreatedAt time.Time `json:"createdAt"`
	// Дата последнего обновления данных в путевом листе.
	UpdatedAt time.Time `json:"updatedAt"`
	// Soft-delete, дата удаления из БД.
	DeletedAt gorm.DeletedAt
}

// Создание нового путевого листа.
type NewWaybill struct {
	// Идентификатор водителя.
	DriverID string `json:"driverID"`
	// Идентификатор машины.
	CarID string `json:"carID"`
	// Остаток топлива при выезде.
	FuelRemaining float64 `json:"fuelRemaining"`
	// Дата и время создания путевого листа.
	DateStart *time.Time `json:"dateStart"`
}

// Обновление существующего путевого листа водителем.
type UpdateWaybill struct {
	// Заправлено топлива.
	FuelFill float64 `json:"fuelFill"`
	// Показания спидометра при заезде.
	MileageEnd float64 `json:"mileageEnd"`
	// Расход топлива по факту
	FuelConsumptionFact float64 `json:"fuelConsumptionFact"`
	// Дата и время закрытия путевого листа.
	DateEnd *time.Time `json:"dateEnd"`
}

// Редактирование путевого листа. Только для механика.
type EditWaybill struct {
	// Идентификатор пользователя.
	UserID *string `json:"userID"`
	// Идентификатор водителя.
	DriverID *string `json:"driverID"`
	// Идентификатор машины.
	CarID *string `json:"carID"`
	// Дата и время создания путевого листа.
	DateStart *time.Time `json:"dateStart"`
	// Дата и время закрытия путевого листа.
	DateEnd *time.Time `json:"dateEnd"`
	// Показания спидометра при выезде.
	MileageStart *float64 `json:"mileageStart"`
	// Показания спидометра при заезде.
	MileageEnd *float64 `json:"mileageEnd"`
	// Заправлено топлива.
	FuelFill *float64 `json:"fuelFill"`
	// Расход топлива по факту
	FuelConsumptionFact *float64 `json:"fuelConsumptionFact"`
	// Остаток топлива при выезде.
	FuelRemainingStart *float64 `json:"fuelRemainingStart"`
	// Остаток топлива при заезде.
	FuelRemainingEnd *float64 `json:"fuelRemainingEnd"`
	// Возможность редактировать путевой лист.
	IsActive *bool `json:"isActive"`
}
