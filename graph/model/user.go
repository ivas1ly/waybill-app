package model

// Пользователь.
type User struct {
	// Идентификатор пользователя.
	ID string `json:"id"`
	// Почта пользователя.
	Email    string `json:"email"`
	Password string `json:"-"`
	// Номер телефона.
	PhoneNumber string `json:"phoneNumber"`
	// Роль в сервисе.
	Role Role `json:"role"`
	// Путевые листы, созданные пользователем.
	Waybills []*Waybill `json:"waybills"`
}
