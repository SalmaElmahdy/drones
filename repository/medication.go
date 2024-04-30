package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IMedicationRepository interface {
	Create(ctx context.Context, medication entity.Medication) (entity.Medication, error)
}
type MedicationRepository struct {
	client *gorm.DB
}

func NewMedicationRepository(client *gorm.DB) IMedicationRepository {
	return &MedicationRepository{
		client: client,
	}
}

func (mDB *MedicationRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	err := mDB.client.WithContext(ctx).Create(&medication).Error
	return medication, err
}
