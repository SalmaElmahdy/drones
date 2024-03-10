package repository

import (
	"context"
	"testing"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateDrone(t *testing.T) {
	db, err := setupTestDatabase()
	assert.NoError(t, err, "Error setting up test database")

	// Defer a function to close the database after the test
	defer func() {
		sqlDB, err := db.DB()
		assert.NoError(t, err, "Error getting underlying SQL database")
		err = sqlDB.Close()
		assert.NoError(t, err, "Error closing test database")
	}()

	repo := NewDroneRepository(db)

	testDrone := entity.Drone{
		SerialNumber:    "789",
		DroneModel:      entity.DroneModelEnum("Light"),
		WeightLimit:     7,
		BatteryCapacity: 50,
		State:           entity.DroneStateEnum("IDLE"),
	}

	createdDrone, err := repo.Create(context.Background(), testDrone)

	// Assertions for the first drone creation
	assert.NoError(t, err, "Expected no error while creating drone")
	assert.Equal(t, testDrone.SerialNumber, createdDrone.SerialNumber)
	assert.Equal(t, testDrone.DroneModel, createdDrone.DroneModel)
	assert.Equal(t, testDrone.WeightLimit, createdDrone.WeightLimit)
	assert.Equal(t, testDrone.BatteryCapacity, createdDrone.BatteryCapacity)
	assert.Equal(t, testDrone.State, createdDrone.State)

	// Attempt to create a second drone with the same serial number
	_, err = repo.Create(context.Background(), testDrone)
	assert.Error(t, err, "Expected an error for duplicate serial number")
	assert.Contains(t, err.Error(), "UNIQUE constraint failed: drones.serial_number")

	// Attempt to create a new drone with weight more than 500
	testDrone.SerialNumber = "123"
	testDrone.WeightLimit = 1000
	_, err = repo.Create(context.Background(), testDrone)
	assert.Error(t, err, "Expected an error for wight limit more than 500")
	assert.Contains(t, err.Error(), "CHECK constraint failed: chk_drones_weight_limit")

}
