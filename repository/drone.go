package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IDroneRepository interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	FindByID(ctx context.Context, id uint) (entity.Drone, error)
	Update(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	GetLoadedMedications(ctx context.Context, id uint) ([]entity.Medication, error)
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

func (dDB *DroneRepository) FindByID(ctx context.Context, id uint) (entity.Drone, error) {
	var drone entity.Drone
	result := dDB.client.First(&drone, id)
	if result.Error != nil {
		return entity.Drone{}, result.Error
	}
	return drone, nil
}

func (dDB *DroneRepository) Update(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	result := dDB.client.Save(&drone)
	if result.Error != nil {
		return entity.Drone{}, result.Error
	}
	return drone, nil
}

func (dDB *DroneRepository) GetLoadedMedications(ctx context.Context, id uint) ([]entity.Medication, error) {
	var drone entity.Drone
	if err := dDB.client.Preload("Medications").First(&drone, id).Error; err != nil {
		return []entity.Medication{}, err
	}
	return drone.Medications, nil
}
