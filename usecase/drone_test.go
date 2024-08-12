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
	db, err := repository.SetupTestDatabase()
	assert.NoError(t, err, "Error setting up test database")
	type mocks struct {
		mockedDroneRepository      repository.IDroneRepository
		mockedMedicationRepository repository.IMedicationRepository
		mockedOrderRepository      repository.IOrderRepository
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
			mocks: mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
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
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				droneUseCase := NewDroneUseCase(test.mocks.mockedDroneRepository, test.mocks.mockedMedicationRepository, test.mocks.mockedOrderRepository)
				got, err := droneUseCase.Create(test.args.ctx, test.args.request)
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

func TestUpdateDroneState(t *testing.T) {
	db, err := repository.SetupTestDatabase()
	assert.NoError(t, err, "Error setting up test database")

	// create needed data
	drones := []entity.Drone{
		{SerialNumber: "1001", DroneModel: entity.Light, WeightLimit: 11, BatteryCapacity: 50, State: entity.IDLE},
		{SerialNumber: "1002", DroneModel: entity.Light, WeightLimit: 11, BatteryCapacity: 50, State: entity.IDLE},
	}
	err = db.Create(&drones).Error
	if err != nil {
		t.Errorf("[Error] Cannot create Drones: %v", err)
	}
	type mocks struct {
		mockedDroneRepository      repository.IDroneRepository
		mockedMedicationRepository repository.IMedicationRepository
		mockedOrderRepository      repository.IOrderRepository
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
			name: "[Test] update drone state method should return updated Drone",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1001",
					"state":"LOADING"
				  }`),
			},
			want: entity.Drone{
				SerialNumber:    "1001",
				DroneModel:      entity.Light,
				WeightLimit:     11,
				BatteryCapacity: 50,
				State:           entity.LOADING},

			mocks: mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
		{
			name: "[Test] update drone state with undefined new state should fail",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1001",
					"state":"TEST"
				  }`),
			},
			wantedErr: "[Error]: state: TEST does not validate as state",

			mocks: mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
		{
			name: "[Test] update not found drone state should fail",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1",
					"state":"LOADING"
				  }`),
			},
			wantedErr: "record not found",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
		{
			name: "[Test] update drone state with wrong defined state should fail",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number":"1002",
					"state":"RETURNING"
				  }`),
			},
			wantedErr: "[Error]: Transition from IDLE to RETURNING is not allowed",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				droneUseCase := NewDroneUseCase(test.mocks.mockedDroneRepository, test.mocks.mockedMedicationRepository, test.mocks.mockedOrderRepository)
				got, err := droneUseCase.UpdateDroneState(test.args.ctx, test.args.request)
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

func TestLoadMedications(t *testing.T) {
	db, err := repository.SetupTestDatabase()
	assert.NoError(t, err, "Error setting up test database")

	// create needed data after setting up the test database
	drones := []entity.Drone{
		{SerialNumber: "1001", DroneModel: entity.Light, WeightLimit: 5, BatteryCapacity: 50, State: entity.IDLE},
	}

	err = db.Create(&drones).Error
	if err != nil {
		t.Errorf("[Error] Cannot create Drones: %v", err)
	}

	type mocks struct {
		mockedDroneRepository      repository.IDroneRepository
		mockedMedicationRepository repository.IMedicationRepository
		mockedOrderRepository      repository.IOrderRepository
	}

	type args struct {
		ctx     context.Context
		request []byte
	}
	tests := []struct {
		name      string
		args      args
		want      []entity.Order
		wantedErr string
		mocks     mocks
	}{

		{
			name: "[Test] load drone with medications should add medications to drone and update drone state",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number": "1001",
					"medications": [
					  {
						"name": "_",
						"weight": 1,
						"code": "WWWWWWW",
						"image": "imagetest"
					  },
					  {
						"name": "_",
						"weight": 1,
						"code": "WWWWWWW",
						"image": "imagetest"
					  },
					  {
						"name": "_",
						"weight": 1,
						"code": "CCC",
						"image": "imagetest"
					  }
					]
				  }`),
			},
			want: []entity.Order{
				{
					Drone: entity.Drone{
						SerialNumber: "1001",
						State:        entity.LOADING,
					},
					Medication: entity.Medication{
						Code: "WWWWWWW",
					},
					Quantity: 2,
					State:    entity.PROCESSING,
				},
				{
					Drone: entity.Drone{
						SerialNumber: "1001",
						State:        entity.LOADING,
					},
					Medication: entity.Medication{
						Code: "CCC",
					},
					Quantity: 1,
					State:    entity.PROCESSING,
				},
			},

			mocks: mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
		{
			name: "[Test] load drone with medications should raise drone invalid",
			args: args{
				ctx: context.Background(),
				request: []byte(`{
					"serial_number": "1001",
					"medications": [
					  {
						"name": "_",
						"weight": 1,
						"code": "WWWWWWW",
						"image": "imagetest"
					  },
					  {
						"name": "_",
						"weight": 1,
						"code": "WWWWWWW",
						"image": "imagetest"
					  },
					  {
						"name": "_",
						"weight": 1,
						"code": "CCC",
						"image": "imagetest"
					  }
					]
				  }`),
			},
			wantedErr: "invalid drone state",
			mocks:     mocks{mockedDroneRepository: mock.NewMockedDroneRepository(db), mockedMedicationRepository: mock.NewMockedMedicationRepository(db), mockedOrderRepository: mock.NewMockedOrderRepository(db)},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				droneUseCase := NewDroneUseCase(test.mocks.mockedDroneRepository, test.mocks.mockedMedicationRepository, test.mocks.mockedOrderRepository)
				got, err := droneUseCase.LoadMedications(test.args.ctx, test.args.request)
				if err != nil {
					assert.EqualError(t, err, test.wantedErr)
					return
				}
				res := []entity.Order{}
				err = json.Unmarshal(got, &res)
				if err != nil {
					fmt.Printf("[Error]: %v", err.Error())
					return
				}
				for key, _ := range test.want {
					assert.Equal(t, test.want[key].Quantity, res[key].Quantity)
					assert.Equal(t, test.want[key].State, res[key].State)
					assert.Equal(t, test.want[key].Drone.SerialNumber, res[key].Drone.SerialNumber)
					assert.Equal(t, test.want[key].Drone.State, res[key].Drone.State)
					assert.Equal(t, test.want[key].Medication.Code, res[key].Medication.Code)
				}
			})
	}
}
