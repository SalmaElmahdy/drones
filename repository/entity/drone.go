package entity

import "gorm.io/gorm"

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
	IDLE      DroneStateEnum = "IDLE"
	LOADING   DroneStateEnum = "LOADING"
	DELIVERED DroneStateEnum = "DELIVERED"
	RETURNING DroneStateEnum = "RETURNING"
)

type DroneRequest struct {
	SerialNumber    string         `json:"serial_number"`
	DroneModel      DroneModelEnum `json:"drone_model"`
	WeightLimit     float64        `json:"weight_limit"`
	BatteryCapacity uint           `json:"battery_capacity"`
	State           DroneStateEnum `json:"state"`
}
