package seeds

import (
	"fmt"
	"math/rand"

	"github.com/SalmaElmahdy/drones/repository/entity"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type DroneSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewDroneSeeder(cfg gorm_seeder.SeederConfiguration) DroneSeeder {
	return DroneSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

var drone_model []string = []string{"Light", "Middle", "Cruiser", "Heavy"}
var state []string = []string{"IDLE", "LOADING", "DELIVERED", "RETURNING"}

func (s *DroneSeeder) Seed(db *gorm.DB) error {
	var drones []entity.Drone
	for i := 0; i < s.Configuration.Rows; i++ {
		indexStr := fmt.Sprint(i)
		drone := entity.Drone{
			SerialNumber:    indexStr,
			DroneModel:      entity.DroneModelEnum(randomPickStr(drone_model)),
			WeightLimit:     float64(rand.Intn(500)),
			BatteryCapacity: uint(rand.Intn(100)),
			State:           entity.DroneStateEnum(randomPickStr(state)),
		}
		drones = append(drones, drone)
	}
	return db.CreateInBatches(drones, s.Configuration.Rows).Error
}

func (s *DroneSeeder) Clear(db *gorm.DB) error {
	return s.SeederAbstract.Delete(db, "drones")
}

func randomPickStr(arr []string) string {

	return arr[rand.Intn(len(arr))]
}
