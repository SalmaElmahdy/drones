package repository

import (
	"github.com/SalmaElmahdy/drones/repository/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Drone{}, &entity.Medication{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
