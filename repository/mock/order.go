package mock

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type MockedOrderRepository struct{}

func NewMockedOrderRepository(db *gorm.DB) repository.IOrderRepository {
	return repository.NewOrderRepository(db)
}

func (MockedOrderRepository) Create(ctx context.Context, order []entity.Order) ([]entity.Order, error) {
	return order, nil
}
