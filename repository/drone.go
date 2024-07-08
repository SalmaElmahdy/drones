package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IDroneRepository interface {
	GetAll(ctx context.Context) ([]entity.Drone, error)
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	Update(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	FindBySerialNumber(ctx context.Context, serialNumber string) (entity.Drone, error)
	GetLoadedMedications(ctx context.Context, serialNumber string) ([]entity.Medication, error)
	WithTransaction(ctx context.Context, fn func() error) error
}

type DroneRepository struct {
	client *gorm.DB
}

func NewDroneRepository(client *gorm.DB) IDroneRepository {
	return &DroneRepository{
		client: client,
	}
}

func (dDB *DroneRepository) GetAll(ctx context.Context) ([]entity.Drone, error) {
	var drones []entity.Drone
	err := dDB.client.WithContext(ctx).Find(&drones).Error
	return drones, err
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
	return drone, nil
}

func (dDB *DroneRepository) GetLoadedMedications(ctx context.Context, serialNumber string) ([]entity.Medication, error) {
	var drone entity.Drone
	if err := dDB.client.Preload("Medications").Where("serial_number = ?", serialNumber).First(&drone).Error; err != nil {
		return []entity.Medication{}, err
	}
	return []entity.Medication{}, nil
}

func (dDB *DroneRepository) WithTransaction(ctx context.Context, fn func() error) error {
	tx := dDB.client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err := recover(); err != nil {
			tx.Rollback()
			return
		}
	}()

	err := fn()
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
