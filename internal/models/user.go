package models

import "gorm.io/gorm"

// User merepresentasikan entitas pengguna dalam sistem
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"` // json:"-" agar password tidak ikut ter-serialize
	Role     string          `json:"role" gorm:"type:varchar(20);default:'user'"`
	Levels   []KosakataLevel `json:"levels" gorm:"many2many:user_levels;"`
}
