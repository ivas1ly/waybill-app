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

// Создание нового пользователя. Только администратор
type NewUser struct {
	// Почта пользователя. Почта не должна повторяться.
	Email string `json:"email"`
	// Номер телефона.
	PhoneNumber string `json:"phoneNumber"`
	// Роль пользователя в сервисе.
	Role *Role `json:"role"`
}

// Обновление данных пользователя.
type UpdateUser struct {
	// Почта пользователя.
	Email *string `json:"email"`
	// Номер телефона.
	PhoneNumber *string `json:"phoneNumber"`
	// Пароль пользователя. Должен быть не менее 10 символов.
	Password *string `json:"password"`
}

type EditUser struct {
	// Почта пользователя.
	Email *string `json:"email"`
	// Номер телефона.
	PhoneNumber *string `json:"phoneNumber"`
	// Роль пользователя.
	Role *Role `json:"role"`
}
