package models

import "time"

// UserKosakata adalah tabel junction/pivot eksplisit antara User dan Kosakata
// Tabel ini juga menyimpan 'Penguasaan' (mastery) dari user terhadap kosakata tertentu.
type UserKosakata struct {
	UserID     uint `gorm:"primaryKey"`
	KosakataID uint `gorm:"primaryKey"`
	Penguasaan int  `json:"penguasaan" gorm:"default:0"`
	CreatedAt  time.Time
}
