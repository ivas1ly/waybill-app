package domain

import (
	"context"

	"github.com/ivas1ly/waybill-app/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (d *Domain) GetCar(ctx context.Context, id string) (*models.Car, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.CarsRepository.GetCarByID(id)
}

func (d *Domain) NewCar(ctx context.Context, input models.NewCar) (*models.Car, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	driver := &models.Car{
		Brand:           input.Brand,
		Number:          input.Number,
		Fuel:            input.Fuel,
		Mileage:         input.Mileage,
		FuelConsumption: input.FuelConsumption,
		FuelRemaining:   input.FuelRemaining,
	}

	return d.CarsRepository.CreateCar(driver)
}

func (d *Domain) UpdateCar(ctx context.Context, id string, input models.UpdateCar) (*models.Car, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	car, err := d.CarsRepository.GetCarByID(id)
	if err != nil {
		d.Logger.Error("Error while getting driver by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.Brand != nil {
		car.Brand = *input.Brand
	}
	if input.Number != nil {
		car.Number = *input.Number
	}
	if input.Fuel != nil {
		car.Fuel = *input.Fuel
	}
	if input.Mileage != nil {
		car.Mileage = *input.Mileage
	}
	if input.FuelRemaining != nil {
		car.FuelRemaining = *input.FuelRemaining
	}
	if input.FuelConsumption != nil {
		car.FuelConsumption = *input.FuelConsumption
	}

	return d.CarsRepository.UpdateCar(id, car)
}

func (d *Domain) DeleteCar(ctx context.Context, id string) (string, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return "", err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return "", gqlerror.Errorf("Forbidden.")
	}

	return d.CarsRepository.DeleteCar(id)
}

func (d *Domain) GetAllCars(ctx context.Context, limit, offset *int) ([]*models.Car, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.CarsRepository.GetCars(limit, offset)
}
