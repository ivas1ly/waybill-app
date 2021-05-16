package model

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
