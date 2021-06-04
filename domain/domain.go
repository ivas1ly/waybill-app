package domain

import (
	"github.com/ivas1ly/waybill-app/database"
	"github.com/ivas1ly/waybill-app/models"
	"go.uber.org/zap"
)

type Domain struct {
	Logger             *zap.Logger
	UsersRepository    database.UsersRepository
	WaybillsRepository database.WaybillsRepository
	DriversRepository  database.DriversRepository
	CarsRepository     database.CarsRepository
}

func NewDomain(logger *zap.Logger, usersRepository database.UsersRepository,
	waybillsRepository database.WaybillsRepository,
	driversRepository database.DriversRepository,
	carsRepository database.CarsRepository) *Domain {
	return &Domain{
		Logger:             logger,
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
