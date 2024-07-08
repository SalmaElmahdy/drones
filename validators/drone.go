package validators

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SalmaElmahdy/drones/repository/entity"
	"github.com/asaskevich/govalidator"
)

func ValidateDroneRequest(droneRequest entity.DroneRequest) error {
	govalidator.TagMap["drone_model"] = govalidator.Validator(func(str string) bool {
		return str == string(entity.Light) || str == string(entity.Middle) || str == string(entity.Cruiser) || str == string(entity.Heavy)
	})
	govalidator.TagMap["state"] = govalidator.Validator(func(str string) bool {
		return str == string(entity.IDLE) || str == string(entity.LOADING) || str == string(entity.LOADED) || str == string(entity.DELIVERING) || str == string(entity.DELIVERED) || str == string(entity.RETURNING)
	})

	if _, err := govalidator.ValidateStruct(droneRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(govalidator.Errors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return fmt.Errorf("[Error]: %s", strings.Join(validationErrors, ", "))
	}
	return nil
}

func ValidateUpdateDroneStateRequest(updatDroneStateRequest entity.UpdateDroneStateRequest) error {
	govalidator.TagMap["state"] = govalidator.Validator(func(str string) bool {
		return str == string(entity.IDLE) || str == string(entity.LOADING) || str == string(entity.LOADED) || str == string(entity.DELIVERING) || str == string(entity.DELIVERED) || str == string(entity.RETURNING)
	})

	if _, err := govalidator.ValidateStruct(updatDroneStateRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(govalidator.Errors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return fmt.Errorf("[Error]: %s", strings.Join(validationErrors, ", "))
	}
	return nil

}

func ValidateLoadDroneState(drone entity.Drone) error {
	if drone.State != entity.IDLE {
		return errors.New("invalid drone state")
	}

	return nil
}

func ValidateLoadDroneBatteryCapacity(drone entity.Drone) error {
	if drone.BatteryCapacity < 25 {
		return errors.New("cannot load medications: battery level is below 25%")
	}
	return nil
}
