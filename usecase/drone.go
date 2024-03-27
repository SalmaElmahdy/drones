package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/SalmaElmahdy/drones/validators"
)

type IDroneUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
}

type DroneUseCase struct {
	droneRepository repository.IDroneRepository
}

func NewDroneUseCase(droneRepository repository.IDroneRepository) IDroneUseCase {
	return DroneUseCase{
		droneRepository: droneRepository,
	}
}
func (d DroneUseCase) Create(ctx context.Context, request []byte) ([]byte, error) {
	droneRequest := entity.DroneRequest{}
	err := json.Unmarshal(request, &droneRequest)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	err = validators.ValidateDroneRequest(droneRequest)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	drone := entity.Drone{
		SerialNumber:    droneRequest.SerialNumber,
		DroneModel:      entity.DroneModelEnum(droneRequest.DroneModel),
		WeightLimit:     droneRequest.WeightLimit,
		BatteryCapacity: droneRequest.BatteryCapacity,
		State:           entity.DroneStateEnum(droneRequest.State),
	}

	createdDrone, err := d.droneRepository.Create(ctx, drone)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(createdDrone)
}
