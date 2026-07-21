package models

import "gorm.io/gorm"

// Kosakata merepresentasikan satu entri kosakata bahasa Jepang.
// gorm.Model menyertakan field ID, CreatedAt, UpdatedAt, dan DeletedAt (soft delete).
type Kosakata struct {
	gorm.Model
	Kanji     string `json:"kanji"      gorm:"not null"`
	Reading   string `json:"reading"    gorm:"not null"` // Hiragana atau Katakana
	Romaji    string `json:"romaji"     gorm:"not null"`
	Arti      string `json:"arti"       gorm:"not null"`
	Levels    []KosakataLevel `json:"levels"     gorm:"many2many:kosakata_level_relations;"`
}