package entity

import (
	"gorm.io/gorm"
)

type Medication struct {
	gorm.Model
	Name   string  `gorm:"not null;"`
	Weight float64 `gorm:"not null;gte=0"`
	Code   string  `gorm:"not null;unique"`
	Image  string  `gorm:"not null;"`
}

type MedicationRequest struct {
	Name   string  `json:"name" valid:"required~name is required,matches(^[A-Za-z0-9_-]+$)~allowed only letters|numbers|-|_)"`
	Weight float64 `json:"weight" valid:"required~weight is required"`
	Code   string  `json:"code" valid:"required~code is required,matches(^[A-Z0-9_-]+$)~allowed only upper case letters|underscore|numbers)"`
	Image  string  `json:"image" valid:"required~image is required"`
}
