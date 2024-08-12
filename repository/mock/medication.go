package mock

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type MockedMedicationRepository struct{}

func NewMockedMedicationRepository(db *gorm.DB) repository.IMedicationRepository {
	return repository.NewMedicationRepository(db)
}

func (MockedMedicationRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return medication, nil
}

func (MockedMedicationRepository) GetByCode(ctx context.Context, code string) (entity.Medication, error) {
	return entity.Medication{}, nil
}
