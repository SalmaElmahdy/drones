package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/SalmaElmahdy/drones/validators"
)

type IMedicationUseCase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
}

type MedicationUseCase struct {
	medicationRepository repository.IMedicationRepository
}

func NewMedicationUseCase(medicationRepository repository.IMedicationRepository) IMedicationUseCase {
	return &MedicationUseCase{
		medicationRepository: medicationRepository,
	}
}

func (m MedicationUseCase) Create(ctx context.Context, request []byte) ([]byte, error) {
	medicationRequest := entity.MedicationRequest{}
	err := json.Unmarshal(request, &medicationRequest)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	err = validators.ValidateMedicationRequest(medicationRequest)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	medication := entity.Medication{
		Name:   medicationRequest.Name,
		Weight: medicationRequest.Weight,
		Code:   medicationRequest.Code,
		Image:  medicationRequest.Image,
	}

	createdMedication, err := m.medicationRepository.Create(ctx, medication)
	if err != nil {
		fmt.Printf("[Error]: %v", err.Error())
		return []byte{}, err
	}
	return json.Marshal(createdMedication)

}
