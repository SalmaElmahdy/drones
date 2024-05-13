package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/SalmaElmahdy/drones/validators"
)

type IDroneUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
	GetAll(ctx context.Context) ([]byte, error)
	GetLoadedMedications(ctx context.Context, serialNumber string) ([]byte, error)
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

func (d DroneUseCase) GetAll(ctx context.Context) ([]byte, error) {
	drones, err := d.droneRepository.GetAll(ctx)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(drones)
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

func (d DroneUseCase) GetLoadedMedications(ctx context.Context, serialNumber string) ([]byte, error) {
	loadedMedications, err := d.droneRepository.GetLoadedMedications(ctx, serialNumber)
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

	drone, err := d.droneRepository.FindBySerialNumber(ctx, loadMedicationRequest.SerialNumber)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	if drone.State != entity.IDLE {
		err := errors.New("invalid drone state")
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	drone.State = entity.LOADING
	drone, err = d.droneRepository.Update(ctx, drone)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	currentMedicationWeight := 0.0
	for _, reqMedication := range loadMedicationRequest.Medications {

		reqMedicationObj := entity.Medication{
			Name:   reqMedication.Name,
			Weight: reqMedication.Weight,
			Code:   reqMedication.Code,
			Image:  reqMedication.Image,
		}

		medication, err := d.medicationRepository.FirstOrCreate(ctx, reqMedicationObj)
		if err != nil {
			fmt.Printf("[Error]: %v", err.Error())
			return []byte{}, err
		}

		currentMedicationWeight += medication.Weight
		if currentMedicationWeight <= drone.WeightLimit {
			// droneMedication := entity.DroneMedications{
			// 	DroneID:      drone.ID,
			// 	MedicationID: medication.ID,
			// 	OrderNumber:  1,
			// }
			// err := d.droneRepository.AppendMedication(ctx, droneMedication)
			// if err != nil {
			// 	fmt.Printf("[Error]: %v", err.Error())
			// 	return []byte{}, err
			// }
		} else {
			err := errors.New("medications exceed drone's weight limit")
			fmt.Printf("[Error]: %v", err.Error())
			return []byte{}, err
		}

	}

	drone.State = entity.LOADED
	drone, err = d.droneRepository.Update(ctx, drone)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	drone.State = entity.DELIVERED
	drone, err = d.droneRepository.Update(ctx, drone)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(drone)
}
