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

	defer func() {
		sqlDB, err := db.DB()
		assert.NoError(t, err, "Error getting underlying SQL database")
		err = sqlDB.Close()
		assert.NoError(t, err, "Error closing test database")
	}()

	repo := NewDroneRepository(db)

	type args struct {
		ctx   context.Context
		Drone entity.Drone
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Drone
		wantErr string
	}{
		{
			name: "[Test] Create method should create drone",
			args: args{
				ctx:   context.Background(),
				Drone: entity.Drone{SerialNumber: "123", DroneModel: entity.DroneModelEnum("Light"), WeightLimit: 7.5, BatteryCapacity: 10, State: entity.DroneStateEnum("IDLE")},
			},
			want: entity.Drone{SerialNumber: "123", DroneModel: entity.DroneModelEnum("Light"), WeightLimit: 7.5, BatteryCapacity: 10, State: entity.DroneStateEnum("IDLE")},
		},
		{
			name: "[Test] Create method should not create drone with same serial number",
			args: args{
				ctx:   context.Background(),
				Drone: entity.Drone{SerialNumber: "123", DroneModel: entity.DroneModelEnum("Light"), WeightLimit: 7.5, BatteryCapacity: 10, State: entity.DroneStateEnum("IDLE")},
			},
			wantErr: "UNIQUE constraint failed: drones.serial_number",
		},
		{
			name: "[Test] Create method should not create drone with weight more than 500",
			args: args{
				ctx:   context.Background(),
				Drone: entity.Drone{SerialNumber: "111", DroneModel: entity.DroneModelEnum("Light"), WeightLimit: 600, BatteryCapacity: 10, State: entity.DroneStateEnum("IDLE")},
			},
			wantErr: "CHECK constraint failed: chk_drones_weight_limit",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := repo.Create(test.args.ctx, test.args.Drone)
			if err != nil {
				assert.Equal(t, test.wantErr, err.Error())
				return
			}
			got := entity.Drone{}
			err = db.Find(&got).Error
			if err != nil {
				t.Errorf("[Error] Cannot retrieve Drone: %v", err)

			}
			assert.Equal(t, test.want.SerialNumber, got.SerialNumber)
			assert.Equal(t, test.want.DroneModel, got.DroneModel)
			assert.Equal(t, test.want.WeightLimit, got.WeightLimit)
			assert.Equal(t, test.want.BatteryCapacity, got.BatteryCapacity)
			assert.Equal(t, test.want.DroneModel, got.DroneModel)

		})
	}

}
