package models

import "time"

// UserLevel adalah tabel junction/pivot eksplisit antara User dan KosakataLevel
type UserLevel struct {
	UserID          uint `gorm:"primaryKey"`
	KosakataLevelID uint `gorm:"primaryKey"`
	CreatedAt       time.Time
}
