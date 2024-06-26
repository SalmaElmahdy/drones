package mock

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
)

type MockedOrderRepository struct{}

func NewMockedOrderRepository() repository.IOrderRepository {
	return nil
}

func (MockedOrderRepository) Create(ctx context.Context, order []entity.Order) ([]entity.Order, error) {
	return order, nil
}
