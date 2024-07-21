package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/SalmaElmahdy/drones/validators"
	"github.com/google/uuid"
)

type IDroneUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
	GetAll(ctx context.Context) ([]byte, error)
	UpdateDroneState(ctx context.Context, request []byte) ([]byte, error)
	GetLoadedMedications(ctx context.Context, serialNumber string) ([]byte, error)
	LoadMedications(ctx context.Context, request []byte) ([]byte, error)
}

type DroneUseCase struct {
	droneRepository      repository.IDroneRepository
	medicationRepository repository.IMedicationRepository
	orderRepository      repository.IOrderRepository
}

func NewDroneUseCase(droneRepository repository.IDroneRepository, medicationRepository repository.IMedicationRepository, orderRepository repository.IOrderRepository) IDroneUseCase {
	return DroneUseCase{
		droneRepository:      droneRepository,
		medicationRepository: medicationRepository,
		orderRepository:      orderRepository,
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
	var createdOrders []entity.Order

	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	err = d.droneRepository.WithTransaction(ctx, func() error {
		drone, err := d.droneRepository.FindBySerialNumber(ctx, loadMedicationRequest.SerialNumber)
		if err != nil {
			return err
		}

		if err := validators.ValidateLoadDroneState(drone); err != nil {
			return err
		}

		if err := validators.ValidateLoadDroneBatteryCapacity(drone); err != nil {
			return err
		}

		transitionResult := drone.Transition(entity.LOADING)
		if !transitionResult.Successful {
			return errors.New(transitionResult.Message)
		}

		if createdOrders, err = d.createOrder(ctx, drone, loadMedicationRequest.Medications); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	return json.Marshal(createdOrders)
}

func (d DroneUseCase) UpdateDroneState(ctx context.Context, request []byte) ([]byte, error) {
	updateDroneStateRequest := entity.UpdateDroneStateRequest{}
	err := json.Unmarshal(request, &updateDroneStateRequest)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	err = validators.ValidateUpdateDroneStateRequest(updateDroneStateRequest)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	existDrone, err := d.droneRepository.FindBySerialNumber(ctx, updateDroneStateRequest.SerialNumber)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}

	transitionResult := existDrone.Transition(updateDroneStateRequest.State)
	if !transitionResult.Successful {
		return []byte{}, fmt.Errorf("[Error]: %v", transitionResult.Message)
	}

	existDrone, err = d.droneRepository.Update(ctx, existDrone)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(existDrone)
}

func (d DroneUseCase) createOrder(ctx context.Context, drone entity.Drone, medications []entity.MedicationRequest) ([]entity.Order, error) {
	var orderData []entity.Order
	currentMedicationWeight := 0.0
	uuid := uuid.New()

	for _, reqMedication := range medications {

		reqMedicationObj := entity.Medication{
			Name:   reqMedication.Name,
			Weight: reqMedication.Weight,
			Code:   reqMedication.Code,
			Image:  reqMedication.Image,
		}

		medication, err := d.medicationRepository.FirstOrCreate(ctx, reqMedicationObj)
		if err != nil {
			return nil, err
		}

		//TODO:: need to check if medication already exist in that order
		currentMedicationWeight += medication.Weight
		if currentMedicationWeight <= drone.WeightLimit {
			orderObj := entity.Order{
				OrderNumber:  uuid,
				DroneID:      drone.ID,
				Drone:        drone,
				MedicationID: medication.ID,
				Medication:   medication,
				Quantity:     1,
			}
			orderData = append(orderData, orderObj)
		} else {
			err := errors.New("medications exceed drone's weight limit")
			return nil, err
		}

	}
	createdOrder, err := d.orderRepository.Create(ctx, orderData)
	if err != nil {
		return nil, err
	}
	transitionResult := drone.Transition(entity.LOADED)
	if !transitionResult.Successful {
		return nil, errors.New(transitionResult.Message)
	}

	transitionResult = drone.Transition(entity.DELIVERING)
	if !transitionResult.Successful {
		return nil, errors.New(transitionResult.Message)
	}
	return createdOrder, nil
}
