package domain

import (
	"context"
	"math"
	"time"

	"github.com/ivas1ly/waybill-app/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (d *Domain) GetAllWaybill(ctx context.Context, limit, offset *int) ([]*models.Waybill, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.WaybillsRepository.GetAllWaybills(limit, offset)
}

func (d *Domain) NewWaybill(ctx context.Context, input models.NewWaybill) (*models.Waybill, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if !user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	car, err := d.CarsRepository.GetCarByID(input.CarID)
	if err != nil {
		d.Logger.Error("Can't find car by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	waybill := &models.Waybill{
		UserID:             user.ID,
		DriverID:           input.DriverID,
		CarID:              input.CarID,
		MileageStart:       car.Mileage,
		FuelRemainingStart: input.FuelRemaining,
		IsActive:           true,
	}
	if input.DateStart != nil {
		t, err := time.Parse(input.DateStart.String(), time.RFC3339)
		if err != nil {
			d.Logger.Error("Can't parse time from request.")
			return nil, gqlerror.Errorf("Internal server error.")

		}
		waybill.DateStart = t
	} else {
		waybill.DateStart = time.Now()
	}

	return d.WaybillsRepository.CreateWaybill(waybill)
}

func (d *Domain) UpdateWaybill(ctx context.Context, id string, input models.UpdateWaybill) (*models.Waybill, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role.IsValid() {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	waybill, err := d.WaybillsRepository.GetWaybillByID(id)
	if err != nil {
		d.Logger.Error("Error while getting waybill by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if user.Role.String() == models.RoleDriver.String() && user.ID != waybill.UserID {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}
	car, err := d.CarsRepository.GetCarByID(waybill.CarID)
	if err != nil {
		d.Logger.Error("Error while getting car by ID.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.FuelFill >= 0 && input.FuelFill < 120 {
		*waybill.FuelFill = math.Round(input.FuelFill*100) / 100
	} else {
		d.Logger.Error("Fuel fill less than zero.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.MileageEnd >= 0 && input.MileageEnd >= waybill.MileageStart {
		*waybill.MileageEnd = input.MileageEnd
	} else {
		d.Logger.Error("Mileage end less than zero or less than Mileage start.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if input.DateEnd != nil && input.DateEnd.After(waybill.DateEnd) {
		t, err := time.Parse(input.DateEnd.String(), time.RFC3339)
		if err != nil {
			d.Logger.Error("Can't parse time from request.")
			return nil, gqlerror.Errorf("Internal server error.")

		}
		waybill.DateEnd = t
	} else {
		waybill.DateEnd = time.Now()
	}

	// Пробег за день
	dayMileage := input.MileageEnd - waybill.MileageStart
	// Расход по норме
	consumptionRate := car.FuelConsumption / 100 * dayMileage
	// Остаток топлива при заезде
	*waybill.FuelRemainingEnd = math.Floor(waybill.FuelRemainingStart+input.FuelFill-consumptionRate) / 100

	// Обновление информации о машине
	car.Mileage = *waybill.MileageEnd
	car.FuelRemaining = *waybill.FuelRemainingEnd
	if _, err := d.CarsRepository.UpdateCar(car.ID, car); err != nil {
		return nil, err
	}

	return d.WaybillsRepository.UpdateWaybill(id, waybill)
}

func (d *Domain) DeleteWaybill(ctx context.Context, id string) (string, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return "", err
	}
	if user.Role.IsValid() &&
		(user.Role.String() != models.RoleAdmin.String() || user.Role.String() != models.RoleMechanic.String()) {
		d.Logger.Error("Forbidden request.")
		return "", gqlerror.Errorf("Forbidden.")
	}

	return d.WaybillsRepository.DeleteWaybill(id)
}

func (d *Domain) GetAllWaybillsByUserID(ctx context.Context, id string, limit, offset *int) ([]*models.Waybill, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role.IsValid() && user.Role.String() == models.RoleDriver.String() && user.ID != id {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.WaybillsRepository.GetAllWaybillsByUserID(id, limit, offset)
}

func (d *Domain) GetWaybill(ctx context.Context, id string) (*models.Waybill, error) {
	user, err := d.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if user.Role.IsValid() && user.Role.String() == models.RoleDriver.String() && user.ID != id {
		d.Logger.Error("Forbidden request.")
		return nil, gqlerror.Errorf("Forbidden.")
	}

	return d.WaybillsRepository.GetWaybillByID(id)
}
