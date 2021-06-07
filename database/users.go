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
	if err := u.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepository) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}

func (u *UsersRepository) GetUsers(limit, offset *int) ([]*models.User, error) {
	var users []*models.User

	result := u.DB.Model(&users)
	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}
	err := result.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UsersRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsersRepository) DeleteUser(user *models.User) (*models.User, error) {
	err := u.DB.Where("id = ?", user.ID).Delete(&user).Error
	return user, err
}

func (u *UsersRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := u.DB.Save(&user).Error
	return user, err
}
