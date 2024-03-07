package entity

import "gorm.io/gorm"

type Drone struct {
	gorm.Model
	SerialNumber    string         `json:"serial_number" gorm:"size:100"`
	DroneModel      DroneModelEnum `json:"drone_model"`
	WeightLimit     uint           `json:"weight_limit" gorm:"check:weight_limit <= 500"`
	BatteryCapacity uint           `json:"battery_capacity" gorm:"check:battery_capacity <= 100"`
	State           DroneStateEnum `json:"state"`
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
	WeightLimit     uint           `json:"weight_limit"`
	BatteryCapacity uint           `json:"battery_capacity"`
	State           DroneStateEnum `json:"state"`
}
