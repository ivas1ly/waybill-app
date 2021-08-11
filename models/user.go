package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/spf13/viper"

	"github.com/gofrs/uuid"

	"github.com/golang-jwt/jwt/v4"
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
	// Дата создания пользователя.
	CreatedAt time.Time `json:"createdAt"`
	// Дата последнего обновления данных пользователя.
	UpdatedAt time.Time `json:"updatedAt"`
	// Soft-delete, дата удаления из БД.
	DeletedAt gorm.DeletedAt
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

func (u *User) CompareUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) GenerateTokenPair() (map[string]string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	//expiresAt := time.Now().Add(time.Minute * 30)
	accessExpiresAt := time.Now().Add(time.Hour * 24) //for dev only

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: accessExpiresAt.Unix(),
		Id:        id.String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "waybill-app",
		Subject:   u.ID,
	})

	token, err := claims.SignedString([]byte(viper.GetString("auth.signing_key")))
	if err != nil {
		return nil, err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refresh.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = u.ID
	refreshExpiresAt := time.Now().Add(time.Hour * 720)
	refreshClaims["exp"] = refreshExpiresAt.Unix() //720 hours = 30 days = 1 month

	refreshToken, err := refresh.SignedString([]byte(viper.GetString("auth.signing_key")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"accessToken":      token,
		"accessExpiresAt":  accessExpiresAt.Format(time.RFC3339),
		"refreshToken":     refreshToken,
		"refreshExpiresAt": refreshExpiresAt.Format(time.RFC3339),
	}, nil
}
