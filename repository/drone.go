package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IDroneRepository interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	FindBySerialNumber(ctx context.Context, serialNumber string) (entity.Drone, error)
	Update(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	GetLoadedMedications(ctx context.Context, serialNumber string) ([]entity.Medication, error)
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

func (dDB *DroneRepository) FindBySerialNumber(ctx context.Context, serialNumber string) (entity.Drone, error) {
	var drone entity.Drone
	result := dDB.client.Where("serial_number = ?", serialNumber).First(&drone)
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

	if err := dDB.client.Preload("Medications").First(&drone, drone.ID).Error; err != nil {
		return entity.Drone{}, err
	}
	return drone, nil
}

func (dDB *DroneRepository) GetLoadedMedications(ctx context.Context, serialNumber string) ([]entity.Medication, error) {
	var drone entity.Drone
	if err := dDB.client.Preload("Medications").Where("serial_number = ?", serialNumber).First(&drone).Error; err != nil {
		return []entity.Medication{}, err
	}
	return drone.Medications, nil
}
