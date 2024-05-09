package entity

import (
	"gorm.io/gorm"
)

type Drone struct {
	gorm.Model
	SerialNumber    string         `gorm:"not null;size:100;unique"`
	DroneModel      DroneModelEnum `gorm:"not null"`
	WeightLimit     float64        `gorm:"not null;check:weight_limit <= 500"`
	BatteryCapacity uint           `gorm:"not null;check:battery_capacity <= 100"`
	State           DroneStateEnum `gorm:"not null;default:IDLE"`
	Medications     []Medication   `gorm:"many2many:drone_medications;"`
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
	DroneID     uint                `json:"drone_id"`
	Medications []MedicationRequest `json:"medications"`
}
