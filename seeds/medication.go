package seeds

import (
	"fmt"
	"math/rand"

	"github.com/SalmaElmahdy/drones/repository/entity"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type MedicationSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewMedicationSeeder(cfg gorm_seeder.SeederConfiguration) MedicationSeeder {
	return MedicationSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *MedicationSeeder) Seed(db *gorm.DB) error {
	var medications []entity.Medication
	for i := 0; i < s.Configuration.Rows; i++ {
		indexStr := fmt.Sprint(i)
		medication := entity.Medication{
			Name:   indexStr,
			Weight: float64(rand.Intn(500)),
			Code:   indexStr,
			Image:  indexStr,
		}
		medications = append(medications, medication)
	}
	return db.CreateInBatches(medications, s.Configuration.Rows).Error
}

func (s *MedicationSeeder) Clear(db *gorm.DB) error {
	return s.SeederAbstract.Delete(db, "medications")
}
