package repository

import (
	"context"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/gorm"
)

type IMedicationRepository interface {
	FirstOrCreate(ctx context.Context, medication entity.Medication) (entity.Medication, error)
	GetByCode(ctx context.Context, code string) (entity.Medication, error)
}
type MedicationRepository struct {
	client *gorm.DB
}

func NewMedicationRepository(client *gorm.DB) IMedicationRepository {
	return &MedicationRepository{
		client: client,
	}
}

func (mDB *MedicationRepository) FirstOrCreate(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	err := mDB.client.WithContext(ctx).Where("code = ?", medication.Code).FirstOrCreate(&medication).Error
	return medication, err
}

func (mDB *MedicationRepository) GetByCode(ctx context.Context, code string) (entity.Medication, error) {
	medication := entity.Medication{}
	err := mDB.client.Where("code = ?", code).First(&medication).Error
	return medication, err
}
