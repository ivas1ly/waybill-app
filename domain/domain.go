package domain

import (
	"errors"

	"github.com/ivas1ly/waybill-app/database"
	"github.com/ivas1ly/waybill-app/models"
)

var (
	ErrBadCredentials  = errors.New("email/password combination don't work")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UsersRepository    database.UsersRepository
	WaybillsRepository database.WaybillsRepository
	DriversRepository  database.DriversRepository
	CarsRepository     database.CarsRepository
}

func NewDomain(usersRepository database.UsersRepository,
	waybillsRepository database.WaybillsRepository,
	driversRepository database.DriversRepository, carsRepository database.CarsRepository) *Domain {
	return &Domain{
		UsersRepository:    usersRepository,
		WaybillsRepository: waybillsRepository,
		DriversRepository:  driversRepository,
		CarsRepository:     carsRepository,
	}
}

type Roles interface {
	HasRole(user *models.User) models.Role
}

func checkRole(r Roles, user *models.User) models.Role {
	return r.HasRole(user)
}
