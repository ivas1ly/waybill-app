package model

// Водитель.
type Driver struct {
	// Уникальный идентификатор
	ID string `json:"id"`
	// Фамилия Имя Отчество
	Fio string `json:"fio"`
	// Есть ли сейчас открытый путевой лист или нет.
	IsActive bool `json:"isActive"`
}
