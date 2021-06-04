package models

import "time"

// Водитель.
type Driver struct {
	// Уникальный идентификатор
	ID string `json:"id" gorm:"type:uuid;primaryKey;size:255;uniqueIndex;not null;default:gen_random_uuid()"`
	// Имя водителя.
	FirstName string `json:"firstName" gorm:"size:255;not null"`
	// Фамилия водителя.
	SecondName string `json:"secondName" gorm:"size:255;not null"`
	// Отчество водителя.
	Patronymic *string `json:"patronymic" gorm:"size:255;"`
	// Есть ли сейчас открытый путевой лист или нет.
	IsActive bool `json:"isActive" gorm:"not null;default:false"`
	// Дата создания водителя.
	CreatedAt time.Time `json:"createdAt"`
	// Дата последнего обновления данных водителя.
	UpdatedAt time.Time `json:"updatedAt"`
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
