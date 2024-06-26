package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type Drone struct {
	gorm.Model
	SerialNumber    string         `gorm:"not null;size:100;unique"`
	DroneModel      DroneModelEnum `gorm:"not null"`
	WeightLimit     float64        `gorm:"not null;check:weight_limit <= 500"`
	BatteryCapacity uint           `gorm:"not null;check:battery_capacity <= 100"`
	State           DroneStateEnum `gorm:"not null;default:IDLE"`
}
type DroneModelEnum string

const (
	Light   DroneModelEnum = "Light"
	Middle  DroneModelEnum = "Middle"
	Cruiser DroneModelEnum = "Cruiser"
	Heavy   DroneModelEnum = "Heavy"
)

type DroneStateEnum string

const (
	IDLE       DroneStateEnum = "IDLE"
	LOADING    DroneStateEnum = "LOADING"
	LOADED     DroneStateEnum = "LOADED"
	DELIVERING DroneStateEnum = "DELIVERING"
	DELIVERED  DroneStateEnum = "DELIVERED"
	RETURNING  DroneStateEnum = "RETURNING"
)

type DroneRequest struct {
	SerialNumber    string         `json:"serial_number" valid:"required~serial_number is required,int~serial_number accepts only numbers"`
	DroneModel      DroneModelEnum `json:"drone_model" valid:"required~drone_model is required,drone_model"`
	WeightLimit     float64        `json:"weight_limit" valid:"required~weight_limit is required,range(0|500)"`
	BatteryCapacity uint           `json:"battery_capacity" valid:"required~battery_capacity is required,range(0|100)"`
	State           DroneStateEnum `json:"state" valid:"required~state is required,state"`
}

type LoadMedicationsRequest struct {
	SerialNumber string              `json:"serial_number" valid:"required~serial_number is required"`
	Medications  []MedicationRequest `json:"medications" valid:"required~medications are required"`
}

type UpdateDroneStateRequest struct {
	SerialNumber string         `json:"serial_number" valid:"required~serial_number is required,int~serial_number accepts only numbers"`
	State        DroneStateEnum `json:"state" valid:"required~state is required,state"`
}

type TransitionResponse struct {
	Successful bool
	NewState   DroneStateEnum
	Message    string
}

func (d *Drone) Transition(newState DroneStateEnum) TransitionResponse {
	allowedTransitions := map[DroneStateEnum][]DroneStateEnum{
		IDLE:       {LOADING},
		LOADING:    {LOADED},
		LOADED:     {DELIVERING},
		DELIVERING: {DELIVERED},
		DELIVERED:  {RETURNING},
		RETURNING:  {IDLE},
	}

	if _, ok := allowedTransitions[d.State]; !ok {
		return TransitionResponse{
			Successful: false,
			NewState:   d.State,
			Message:    fmt.Sprintf("Invalid transition from %s to %s", d.State, newState),
		}
	}

	for _, allowed := range allowedTransitions[d.State] {
		if allowed == newState {
			d.State = newState
			return TransitionResponse{
				Successful: true,
				NewState:   newState,
				Message:    fmt.Sprintf("Drone state changed from %s to %s", d.State, newState),
			}
		}
	}

	return TransitionResponse{
		Successful: false,
		NewState:   d.State,
		Message:    fmt.Sprintf("Transition from %s to %s is not allowed", d.State, newState),
	}
}
