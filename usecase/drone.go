package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/SalmaElmahdy/drones/validators"
	"gorm.io/gorm"
)

type IDroneUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
	GetLoadedMedications(ctx context.Context, id uint) ([]byte, error)
	LoadMedications(ctx context.Context, request []byte) ([]byte, error)
}

type DroneUseCase struct {
	droneRepository      repository.IDroneRepository
	medicationRepository repository.IMedicationRepository
}

func NewDroneUseCase(droneRepository repository.IDroneRepository, medicationRepository repository.IMedicationRepository) IDroneUseCase {
	return DroneUseCase{
		droneRepository:      droneRepository,
		medicationRepository: medicationRepository,
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

func (d DroneUseCase) GetLoadedMedications(ctx context.Context, id uint) ([]byte, error) {
	loadedMedications, err := d.droneRepository.GetLoadedMedications(ctx, id)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(loadedMedications)
}

func (d DroneUseCase) LoadMedications(ctx context.Context, request []byte) ([]byte, error) {
	loadMedicationRequest := entity.LoadMedicationsRequest{}
	err := json.Unmarshal(request, &loadMedicationRequest)

	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	drone, err := d.droneRepository.FindByID(ctx, loadMedicationRequest.DroneID)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	existingMedications, err := d.droneRepository.GetLoadedMedications(ctx, loadMedicationRequest.DroneID)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	for _, reqMedication := range loadMedicationRequest.Medications {
		medication, err := d.medicationRepository.GetByCode(ctx, reqMedication.Code)
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Printf("[Error]: %v", err.Error())
			return []byte{}, err
		}

		if err == gorm.ErrRecordNotFound {
			medication, err = d.medicationRepository.Create(ctx, reqMedication)
			if err != nil {
				fmt.Printf("[Error]: %v", err.Error())
				return []byte{}, err
			}
		}

		alreadyLoaded := false
		for _, loadedMedication := range existingMedications {
			if loadedMedication.Code == reqMedication.Code {
				alreadyLoaded = true
				break
			}
		}

		if !alreadyLoaded {
			drone.Medications = append(drone.Medications, medication)
		}
	}

	updatedDrone, err := d.droneRepository.Update(ctx, drone)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(updatedDrone)
}
