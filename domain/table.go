package domain

import (
	"context"
	"fmt"

	"github.com/ivas1ly/waybill-app/models"
)

func (d *Domain) CreateTable(ctx context.Context, filter models.TableFilter) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
