package domain

import (
	"context"

	"github.com/ivas1ly/waybill-app/models"
)

func (d *Domain) GetUser(ctx context.Context, id string) (*models.User, error) {
	_, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	return d.UsersRepository.GetUserByID(id)
}
