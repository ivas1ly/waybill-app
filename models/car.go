package models

// Машина.
type Car struct {
	// Уникальный идентификатор.
	ID string `json:"id"`
	// Автомобиль.
	Brand string `json:"brand"`
	// Гос. номер автомобиля.
	Number string `json:"number"`
	// Вид топлива.
	Fuel string `json:"fuel"`
	// Пробег.
	Mileage int `json:"mileage"`
	// Норма расхода.
	FuelConsumption float64 `json:"fuelConsumption"`
	// Остаток топлива.
	FuelRemaining float64 `json:"fuelRemaining"`
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
	Mileage int `json:"mileage"`
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
	Mileage *int `json:"mileage"`
	// Текущий остаток топлива.
	FuelRemaining *float64 `json:"fuelRemaining"`
	// Норма расхода топлива.
	FuelConsumption *float64 `json:"fuelConsumption"`
}
