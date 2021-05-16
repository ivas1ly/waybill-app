// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Результат проверки Access Token и пользователь. для которого он был создан.
type AuthResponse struct {
	Response string `json:"response"`
	User     *User  `json:"user"`
}

type EditUser struct {
	// Почта пользователя.
	Email *string `json:"email"`
	// Номер телефона.
	PhoneNumber *string `json:"phoneNumber"`
	// Роль пользователя.
	Role *Role `json:"role"`
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
	MileageStart *int `json:"mileageStart"`
	// Показания спидометра при заезде.
	MileageEnd *int `json:"mileageEnd"`
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

// Вход в сервис обработки путевых листов.
type Login struct {
	// Почта пользователя.
	Email string `json:"email"`
	// Пароль пользователя.
	Password string `json:"password"`
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

// Создание нового водителя.
type NewDriver struct {
	// Имя водителя.
	FirstName string `json:"firstName"`
	// Фамилия водителя.
	SecondName string `json:"secondName"`
	// Отчество водителя.
	Patronymic *string `json:"patronymic"`
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

// Создание нового путевого листа.
type NewWaybill struct {
	// Остаток топлива при выезде.
	FuelRemaining float64 `json:"fuelRemaining"`
	// Дата и время создания путевого листа.
	DateStart *time.Time `json:"dateStart"`
}

// Refresh Token для получения нового Access Token и Refresh Token.
type RefreshToken struct {
	Response string `json:"response"`
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

// Обновление данных пользователя.
type UpdateUser struct {
	// Почта пользователя.
	Email *string `json:"email"`
	// Номер телефона.
	PhoneNumber *string `json:"phoneNumber"`
	// Пароль пользователя. Должен быть не менее 10 символов.
	Password *string `json:"password"`
}

// Обновление существующего путевого листа водителем.
type UpdateWaybill struct {
	// Показания спидометра при заезде.
	MileageEnd int `json:"mileageEnd"`
	// Остаток топлива при заезде.
	FuelRemaining float64 `json:"fuelRemaining"`
}

// Роли сервиса.
type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleMechanic Role = "MECHANIC"
	RoleDriver   Role = "DRIVER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleMechanic,
	RoleDriver,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleMechanic, RoleDriver:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
