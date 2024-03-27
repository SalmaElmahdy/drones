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
		{
			name: "[Test] create drone with wrong drone_model should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Dark",
					"weight_limit":10.66,
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: drone_model: Dark does not validate as drone_model",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone without drone_model should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"weight_limit":10.66,
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: drone_model is required",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone without serial_number should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"drone_model":"Light",
					"weight_limit":10.66,
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: serial_number is required",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone with wrong input type serial_number should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"aaaa",
					"drone_model":"Light",
					"weight_limit":10.66,
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: serial_number accepts only numbers",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone with wrong weight_limit should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"weight_limit":1000,
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: weight_limit: 1000 does not validate as range(0|500)",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone without weight_limit should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"battery_capacity":70,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: weight_limit is required",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone with wrong battery_capacity should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"weight_limit":30,
					"battery_capacity":1000,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: battery_capacity: 1000 does not validate as range(0|100)",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone without battery_capacity should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"weight_limit":30,
					"state":"IDLE"
				  }`),
			},
			wantedErr: "[Error]: battery_capacity is required",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
		{
			name: "[Test] create drone with wrong state should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"weight_limit":30,
					"battery_capacity":10,
					"state":"ID"
				  }`),
			},
			wantedErr: "[Error]: state: ID does not validate as state",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},

		{
			name: "[Test] create drone without state should return Error",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1005",
					"drone_model":"Light",
					"weight_limit":30,
					"battery_capacity":10
				  }`),
			},
			wantedErr: "[Error]: state is required",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository()},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				droneUseCase := NewDroneUseCase(test.mocks.mockedDroneRepository)
				got, err := droneUseCase.Create(test.args.ctx, test.args.request)
				fmt.Println("gooooooot", string(got))
				if err != nil {
					assert.EqualError(t, err, test.wantedErr)
					return
				}
				droneRes := entity.Drone{}
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
