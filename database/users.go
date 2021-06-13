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

func (u *UsersRepository) GetAllUsers(limit, offset *int) ([]*models.User, error) {
	var users []*models.User

	result := u.DB.Model(&users)
	if limit != nil {
		result.Limit(*limit)
	}
	if offset != nil {
		result.Offset(*offset)
	}

	if err := result.Find(&users).Error; err != nil {
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

func (u *UsersRepository) DeleteUser(id string) (string, error) {
	if err := u.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return "", err
	}
	return "User successfully deleted.", nil
}

func (u *UsersRepository) UpdateUser(id string, user *models.User) (*models.User, error) {

	if err := u.DB.Model(&user).Where("id = ?", id).Updates(models.User{
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsersRepository) UpdateRefresh(user *models.User, refreshToken string) (*models.User, error) {
	if err := u.DB.Model(&user).Update("refresh_token", refreshToken).Error; err != nil {
		return nil, err
	}
	return user, nil
}
