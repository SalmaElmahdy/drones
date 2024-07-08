package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNumber  string
	DroneID      uint
	Drone        Drone
	MedicationID uint
	Medication   Medication
	Quantity     int
	State        OrderStateEnum `gorm:"not null;default:PROCESSING"`
}

type OrderStateEnum string

const (
	PROCESSING OrderStateEnum = "PROCESSING"
	PROCESSED  OrderStateEnum = "PROCESSED"
)
