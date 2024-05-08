package mock

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
)

type MockedMedicationRepository struct{}

func NewMockedMedicationRepository() repository.IMedicationRepository {
	return nil
}

func (MockedMedicationRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return medication, nil
}

func (MockedMedicationRepository) GetByCode(ctx context.Context, code string) (entity.Medication, error) {
	return entity.Medication{}, nil
}
