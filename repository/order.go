package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	Create(ctx context.Context, order []entity.Order) ([]entity.Order, error)
}

type OrderRepository struct {
	client *gorm.DB
}

func NewOrderRepository(client *gorm.DB) IOrderRepository {
	return &OrderRepository{
		client: client,
	}
}

func (oDB OrderRepository) Create(ctx context.Context, order []entity.Order) ([]entity.Order, error) {
	var orders []entity.Order
	err := oDB.client.WithContext(ctx).Create(&orders).Error
	return orders, err
}
