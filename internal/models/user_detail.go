package models

import "gorm.io/gorm"

// UserDetail menyimpan detail tambahan dari User, seperti Experience (EXP).
type UserDetail struct {
	gorm.Model
	UserID uint `gorm:"uniqueIndex;not null"`
	Exp    int  `json:"exp" gorm:"default:0"`
	Level  int  `json:"level" gorm:"default:1"`
}

type UserProfileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Exp   int    `json:"exp"`
	Level int    `json:"level"`
}
