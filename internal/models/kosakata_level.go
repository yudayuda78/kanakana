package models

import "gorm.io/gorm"

// KosakataLevel merepresentasikan level dari kosakata bahasa Jepang (misal N5, N4).
type KosakataLevel struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
