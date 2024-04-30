package validators

import (
	"fmt"
	"strings"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/asaskevich/govalidator"
)

func ValidateMedicationRequest(medicationRequest entity.MedicationRequest) error {
	if _, err := govalidator.ValidateStruct(medicationRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(govalidator.Errors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return fmt.Errorf("[Error]: %s", strings.Join(validationErrors, ", "))
	}
	return nil
}
