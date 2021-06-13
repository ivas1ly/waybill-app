package models

import (
	"time"

	"gorm.io/gorm"
)

// Машина.
type Car struct {
	// Уникальный идентификатор.
	ID string `json:"id" gorm:"type:uuid;size:255;uniqueIndex;not null;default:gen_random_uuid()"`
	// Автомобиль.
	Brand string `json:"brand" gorm:"size:255;not null;"`
	// Гос. номер автомобиля.
	Number string `json:"number" gorm:"size:15;"`
	// Вид топлива.
	Fuel string `json:"fuel" gorm:"size:15;"`
	// Пробег.
	Mileage float64 `json:"mileage" gorm:"type:integer;not null;"`
	// Норма расхода.
	FuelConsumption float64 `json:"fuelConsumption;not null;"`
	// Остаток топлива.
	FuelRemaining float64 `json:"fuelRemaining;not null;"`
	// Дата создания машины.
	CreatedAt time.Time `json:"createdAt"`
	// Дата последнего обновления данных машины.
	UpdatedAt time.Time `json:"updatedAt"`
	// Soft-delete, дата удаления из БД.
	DeletedAt gorm.DeletedAt
}

// Создать новую машину.
type NewCar struct {
	// Название машины (бренд и модель).
	Brand string `json:"brand"`
	// Гос. номер машины.
	Number string `json:"number"`
	// Тип топлива для заправки.
	Fuel string `json:"fuel"`
	// Текущий пробег машины.
	Mileage float64 `json:"mileage"`
	// Текущий остаток топлива.
	FuelRemaining float64 `json:"fuelRemaining"`
	// Норма расхода топлива.
	FuelConsumption float64 `json:"fuelConsumption"`
}

// Обновление данных машины.
type UpdateCar struct {
	// Название машины (бренд и модель).
	Brand *string `json:"brand"`
	// Гос. номер машины.
	Number *string `json:"number"`
	// Топливо для заправки.
	Fuel *string `json:"fuel"`
	// Текущий пробега машины.
	Mileage *float64 `json:"mileage"`
	// Текущий остаток топлива.
	FuelRemaining *float64 `json:"fuelRemaining"`
	// Норма расхода топлива.
	FuelConsumption *float64 `json:"fuelConsumption"`
}
