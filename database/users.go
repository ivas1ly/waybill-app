package database

import (
	"fmt"

	"github.com/ivas1ly/waybill-app/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func (u *UsersRepository) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.First(&user, fmt.Sprintf("%v = ?", field), value).Error
	return &user, err
}

func (u *UsersRepository) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}

func (u *UsersRepository) CreateUser(user *models.User) (*models.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

func (u *UsersRepository) DeleteUser(user *models.User) (*models.User, error) {
	err := u.DB.Where("id = ?", user.ID).Delete(&user).Error
	return user, err
}

func (u *UsersRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := u.DB.Save(&user).Error
	return user, err
}
