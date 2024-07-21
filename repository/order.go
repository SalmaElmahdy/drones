package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	Create(ctx context.Context, order []entity.Order) ([]entity.Order, error)
	GetMedicationFromOrder(ctx context.Context, droneID, medicationID uint, state string) (*entity.Order, error)
}

type OrderRepository struct {
	client *gorm.DB
}

func NewOrderRepository(client *gorm.DB) IOrderRepository {
	return &OrderRepository{
		client: client,
	}
}

func (oDB OrderRepository) Create(ctx context.Context, orders []entity.Order) ([]entity.Order, error) {
	err := oDB.client.WithContext(ctx).Create(&orders).Error
	return orders, err
}

func (oDB OrderRepository) GetMedicationFromOrder(ctx context.Context, droneID, medicationID uint, state string) (*entity.Order, error) {
	order := &entity.Order{}
	err := oDB.client.WithContext(ctx).Where("drone_id = ? AND medication_id = ? AND state = ?", droneID, medicationID, state).Last(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
