package domain

import (
	"github.com/ivas1ly/waybill-app/database"
	"github.com/ivas1ly/waybill-app/models"
)

type Domain struct {
	UsersRepository database.UsersRepository
}

func NewDomain(usersRepository database.UsersRepository) *Domain {
	return &Domain{UsersRepository: usersRepository}
}

type Roles interface {
	HasRole(user *models.User) models.Role
}

func checkRole(r Roles, user *models.User) models.Role {
	return r.HasRole(user)
}
