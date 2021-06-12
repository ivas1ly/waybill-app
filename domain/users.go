package domain

import (
	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/ivas1ly/waybill-app/models"
)

func (d *Domain) GetUser(ctx context.Context, id string) (*models.User, error) {
	_, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	return d.UsersRepository.GetUserByID(id)
}

func (d *Domain) NewUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role.IsValid() && user.Role.String() != models.RoleAdmin.String() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.CreateUser(ctx, input)
}

func (d *Domain) UpdateUser(ctx context.Context, id string, input models.UpdateUser) (*models.User, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	if (user.Role.String() == models.RoleMechanic.String() ||
		user.Role.String() == models.RoleDriver.String()) && user.ID != id {
		return nil, gqlerror.Errorf("Forbidden.")
	}

	u, err := d.UsersRepository.GetUserByID(id)
	if err != nil {
		d.Logger.Error("Error while getting user by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.Password != nil {
		if err := u.HashPassword(*input.Password); err != nil {
			d.Logger.Error("Error while hashing password.")
			return nil, gqlerror.Errorf("Internal server error.")
		}
	}

	if input.Email != nil {
		u.Email = *input.Email
	}

	return d.UsersRepository.UpdateUser(id, u)
}

func (d *Domain) EditUser(ctx context.Context, id string, input models.EditUser) (*models.User, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	if user.Role.String() != models.RoleAdmin.String() {
		return nil, gqlerror.Errorf("Forbidden.")
	}

	u, err := d.UsersRepository.GetUserByID(id)
	if err != nil {
		d.Logger.Error("Error while getting user by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.Role.IsValid() && !(user.Role.String() == input.Role.String()) {
		u.Role = *input.Role
	}

	if input.Email != nil {
		u.Email = *input.Email
	}

	return d.UsersRepository.UpdateUser(id, u)
}

func (d *Domain) DeleteUser(ctx context.Context, id string) (string, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return "", err
	}

	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return "", gqlerror.Errorf("Forbidden.")
	}

	if user.Role.String() != models.RoleAdmin.String() {
		return "", gqlerror.Errorf("Forbidden.")
	}

	return d.UsersRepository.DeleteUser(id)
}
