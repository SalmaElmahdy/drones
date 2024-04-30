package entity

import (
	"gorm.io/gorm"
)

type Medication struct {
	gorm.Model
	name   string  `gorm:"not null;"`
	weight float64 `gorm:"not null;gte=0"`
	code   string  `gorm:"not null;unique"`
	image  string  `gorm:"not null;unique"`
}
