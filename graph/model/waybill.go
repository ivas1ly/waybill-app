package model

import "time"

type Waybill struct {
	// Идентификатор путевого листа.
	ID string `json:"id"`
	// Идентификатор пользователя.
	UserID string `json:"userID"`
	// Идентификатор водителя.
	DriverID string `json:"driverID"`
	// Идентификатор машины.
	CarID string `json:"carID"`
	// Дата и время создания путевого листа.
	DateStart time.Time `json:"dateStart"`
	// Дата и время закрытия путевого листа.
	DateEnd time.Time `json:"dateEnd"`
	// Показания спидометра при выезде.
	MileageStart int `json:"mileageStart"`
	// Показания спидометра при заезде.
	MileageEnd *int `json:"mileageEnd"`
	// Заправлено топлива.
	FuelFill float64 `json:"fuelFill"`
	// Расход топлива по факту
	FuelConsumptionFact float64 `json:"fuelConsumptionFact"`
	// Остаток топлива при выезде.
	FuelRemainingStart float64 `json:"fuelRemainingStart"`
	// Остаток топлива при заезде.
	FuelRemainingEnd *float64 `json:"fuelRemainingEnd"`
	// Возможность редактировать путевой лист.
	IsActive bool `json:"isActive"`
	// Водитель, к которому относится путевой лист.
	Driver *Driver `json:"driver"`
	// Пользователь, создавший путевой лист.
	User *User `json:"user"`
	// Машина.
	Car *Car `json:"car"`
}
