package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/SalmaElmahdy/drones/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDrone(t *testing.T) {
	type mocks struct {
		mockedDroneRepository repository.IDroneRepository
	}
	type args struct {
		ctx     context.Context
		request []byte
	}
	tests := []struct {
		name      string
		args      args
		want      entity.Drone
		wantedErr string
		mocks     mocks
	}{
		{
			name: "[Test] create method should return created drone",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"weight_limit":10.66,
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			want: entity.Drone{
				SerialNumber:    "1005",
				DroneModel:      entity.DroneModelEnum("Light"),
				WeightLimit:     10.66,
				BatteryCapacity: 70,
				State:           entity.DroneStateEnum("IDLE"),
			},
			mocks: mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				droneUseCase := NewDroneUseCase(test.mocks.mockedDroneRepository)
				got, err := droneUseCase.Create(test.args.ctx, test.args.request)
				fmt.Println(string(got))
				if err != nil {
					assert.EqualError(t, err, test.wantedErr)
					return
				}
				droneRes := entity.DroneRequest{}
				err = json.Unmarshal(got, &droneRes)
				if err != nil {
					fmt.Printf("[Error]: %v", err.Error())
					return
				}
				assert.Equal(t, test.want.SerialNumber, droneRes.SerialNumber)
				assert.Equal(t, test.want.DroneModel, droneRes.DroneModel)
				assert.Equal(t, test.want.WeightLimit, droneRes.WeightLimit)
				assert.Equal(t, test.want.BatteryCapacity, droneRes.BatteryCapacity)
				assert.Equal(t, test.want.State, droneRes.State)
			})
	}
}
