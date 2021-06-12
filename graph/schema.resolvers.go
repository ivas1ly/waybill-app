package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ivas1ly/waybill-app/graph/generated"
	"github.com/ivas1ly/waybill-app/models"
)

func (r *mutationResolver) Login(ctx context.Context, input models.Login) (*models.AuthResponse, error) {
	return r.Domain.LoginUser(ctx, input)
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (*models.AuthResponse, error) {
	return r.Domain.Refresh(ctx)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	return r.Domain.NewUser(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UpdateUser) (*models.User, error) {
	return r.Domain.UpdateUser(ctx, id, input)
}

func (r *mutationResolver) EditUser(ctx context.Context, id string, input models.EditUser) (*models.User, error) {
	return r.Domain.EditUser(ctx, id, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	return r.Domain.DeleteUser(ctx, id)
}

func (r *mutationResolver) CreateDriver(ctx context.Context, input models.NewDriver) (*models.Driver, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriver(ctx context.Context, id string, input models.UpdateDriver) (*models.Driver, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriver(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCar(ctx context.Context, input models.NewCar) (*models.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCar(ctx context.Context, id string, input models.UpdateCar) (*models.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCar(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWaybill(ctx context.Context, input models.NewWaybill) (*models.Waybill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWaybill(ctx context.Context, id string, input models.UpdateWaybill) (*models.Waybill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditWaybill(ctx context.Context, id string, input models.EditWaybill) (*models.Waybill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteWaybill(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context, limit *int, offset *int) ([]*models.User, error) {
	return r.Domain.UsersRepository.GetUsers(limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.Domain.GetUser(ctx, id)
}

func (r *queryResolver) AllDrivers(ctx context.Context, limit *int, offset *int) ([]*models.Driver, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Driver(ctx context.Context, id string) (*models.Driver, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllCars(ctx context.Context, limit *int, offset *int) ([]*models.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Car(ctx context.Context, id string) (*models.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllWaybills(ctx context.Context, limit *int, offset *int) ([]*models.Waybill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllWaybillsByUserID(ctx context.Context, id string, limit *int, offset *int) ([]*models.Waybill, error) {
	return r.Domain.WaybillsRepository.GetWaybillsByUserID(id, limit, offset)
}

func (r *queryResolver) Waybill(ctx context.Context, id string) (*models.Waybill, error) {
	return r.Domain.WaybillsRepository.GetWaybillByID(id)
}

func (r *queryResolver) CreateReportTable(ctx context.Context, filter models.TableFilter) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
