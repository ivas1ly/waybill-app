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

// Создание нового водителя.
type NewDriver struct {
	// Имя водителя.
	FirstName string `json:"firstName"`
	// Фамилия водителя.
	SecondName string `json:"secondName"`
	// Отчество водителя.
	Patronymic *string `json:"patronymic"`
}

// Обновление данных водителя.
type UpdateDriver struct {
	// Имя водителя.
	FirstName *string `json:"firstName"`
	// Фамилия водителя.
	SecondName *string `json:"secondName"`
	// Отчество водителя.
	Patronymic *string `json:"patronymic"`
	// Есть открытый путевой лист или нет.
	IsActive *bool `json:"isActive"`
}
