package domain

import (
	"context"

	"github.com/ivas1ly/waybill-app/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (d *Domain) GetDriver(ctx context.Context, id string) (*models.Driver, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.DriversRepository.GetDriverByID(id)
}

func (d *Domain) NewDriver(ctx context.Context, input models.NewDriver) (*models.Driver, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	driver := &models.Driver{
		FirstName:  input.FirstName,
		SecondName: input.SecondName,
		Patronymic: input.Patronymic,
	}

	return d.DriversRepository.CreateDriver(driver)
}

func (d *Domain) UpdateDriver(ctx context.Context, id string, input models.UpdateDriver) (*models.Driver, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	driver, err := d.DriversRepository.GetDriverByID(id)
	if err != nil {
		d.Logger.Error("Error while getting driver by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.FirstName != nil {
		driver.FirstName = *input.FirstName
	}
	if input.SecondName != nil {
		driver.SecondName = *input.SecondName
	}
	if input.Patronymic != nil {
		driver.Patronymic = input.Patronymic
	}
	if input.IsActive != nil {
		driver.IsActive = *input.IsActive
	}

	return d.DriversRepository.Update(id, driver)
}

func (d *Domain) DeleteDriver(ctx context.Context, id string) (string, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return "", err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return "", gqlerror.Errorf("Forbidden.")
	}

	return d.DriversRepository.Delete(id)
}

func (d *Domain) GetAllDrivers(ctx context.Context, limit, offset *int) ([]*models.Driver, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.DriversRepository.GetDrivers(limit, offset)
}
