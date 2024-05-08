package mock

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
)

type MockedDroneRepository struct{}

func (MockedDroneRepository) FindByID(ctx context.Context, id uint) (entity.Drone, error) {
	return entity.Drone{}, nil
}

func (MockedDroneRepository) LoadMedications(ctx context.Context, drone entity.Drone, medication []entity.Medication) error {
	return nil
}

func NewMockedDroneRepository() repository.IDroneRepository {
	return nil
}

func (MockedDroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return drone, nil
}
