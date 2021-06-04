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
	ID string `json:"id" gorm:"type:uuid;size:255;uniqueIndex;not null;default:gen_random_uuid()"`
	// Почта пользователя.
	Email string `json:"email" gorm:"size:255;uniqueIndex;not null"`
	// Пароль
	Password string `json:"-" gorm:"size:255;not null"`
	// JWT Refresh Token
	RefreshToken string `json:"-" gorm:"size:255;not null"`
	// 2fa secret
	Secret string `json:"-" gorm:"size:255;not null"`
	// Роль в сервисе.
	Role Role `json:"role" gorm:"size:20;not null"`
	// Путевые листы, созданные пользователем.
	Waybills []*Waybill `json:"waybills" gorm:"foreignKey:UserID;"`
	// Дата создания пользователя.
	CreatedAt time.Time `json:"createdAt"`
	// Дата последнего обновления данных пользователя.
	UpdatedAt time.Time `json:"updatedAt"`
}

// Создание нового пользователя. Только администратор
type NewUser struct {
	// Почта пользователя. Почта не должна повторяться.
	Email string `json:"email"`
	// Роль пользователя в сервисе.
	Role *Role `json:"role"`
}

// Обновление данных пользователя.
type UpdateUser struct {
	// Почта пользователя.
	Email *string `json:"email"`
	// Пароль пользователя. Должен быть не менее 10 символов.
	Password *string `json:"password"`
}

type EditUser struct {
	// Почта пользователя.
	Email *string `json:"email"`
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

func (u *User) HashTOTP(totp string) error {
	byteTotp := []byte(totp)
	hashedTotp, err := bcrypt.GenerateFromPassword(byteTotp, 14)
	if err != nil {
		return err
	}

	u.Secret = string(hashedTotp)

	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) CompareTOTP(totp string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Secret), []byte(totp))
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
