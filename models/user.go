package models

import (
	"time"

	"github.com/spf13/viper"

	"github.com/gofrs/uuid"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Пользователь.
type User struct {
	// Идентификатор пользователя.
	ID string `json:"id"`
	// Почта пользователя.
	Email string `json:"email"`
	// Пароль
	Password string `json:"-"`
	// JWT Refresh Token
	RefreshToken string `json:"-"`
	// 2fa secret
	Secret string `json:"-"`
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

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, 14)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) GenerateTokenPair() (map[string]string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Id:        id.String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "waybill-app-jwt",
		Subject:   u.ID,
	})

	token, err := claims.SignedString([]byte(viper.GetString("auth.signing_key")))
	if err != nil {
		return nil, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refresh.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = u.ID
	refreshClaims["exp"] = time.Now().Add(time.Hour * 168).Unix() //168 hours = 7 days = 1 week

	refreshToken, err := refresh.SignedString([]byte(viper.GetString("auth.signing_key")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"accessToken":  token,
		"refreshToken": refreshToken,
	}, nil
}
