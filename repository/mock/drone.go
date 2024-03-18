package mock

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
)

type MockedDroneRepository struct{}

func NewMockedDroneRepository() repository.IDroneRepository {
	return MockedDroneRepository{}
}

func (MockedDroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return drone, nil
}
