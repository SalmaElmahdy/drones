package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IDroneRepository interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
}

type DroneRepository struct {
	client *gorm.DB
}

func NewDroneRepository(client *gorm.DB) IDroneRepository {
	return &DroneRepository{
		client: client,
	}
}

func (dDB *DroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	err := dDB.client.WithContext(ctx).Create(&drone).Error
	return drone, err
}
