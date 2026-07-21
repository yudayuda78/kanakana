package repository

import (
	"kanakana/internal/models"

	"gorm.io/gorm"
)

// KosakataRepository mendefinisikan kontrak operasi database untuk entitas Kosakata.
// Menggunakan interface agar mudah di-mock saat unit testing.
type KosakataRepository interface {
	FindAll() ([]models.Kosakata, error)
	FindByID(id uint) (*models.Kosakata, error)
	FindByKanji(kanji string) ([]models.Kosakata, error)
	FindByReading(reading string) ([]models.Kosakata, error)
	FindByRomaji(romaji string) ([]models.Kosakata, error)
	FindByLevelJLPT(level string) ([]models.Kosakata, error)
	FindAllLevels() ([]models.KosakataLevel, error)
	Create(k *models.Kosakata) error
	Update(k *models.Kosakata) error
	Delete(id uint) error
}

type kosakataRepository struct {
	db *gorm.DB
}

// NewKosakataRepository membuat instance baru KosakataRepository.
func NewKosakataRepository(db *gorm.DB) KosakataRepository {
	return &kosakataRepository{db: db}
}

func (r *kosakataRepository) FindAll() ([]models.Kosakata, error) {
	var list []models.Kosakata
	result := r.db.Preload("Levels").Find(&list)
	return list, result.Error
}

func (r *kosakataRepository) FindByID(id uint) (*models.Kosakata, error) {
	var k models.Kosakata
	result := r.db.Preload("Levels").First(&k, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &k, nil
}

func (r *kosakataRepository) FindByKanji(kanji string) ([]models.Kosakata, error) {
	var list []models.Kosakata
	result := r.db.Preload("Levels").Where("kanji = ?", kanji).Find(&list)
	return list, result.Error
}

func (r *kosakataRepository) FindByReading(reading string) ([]models.Kosakata, error) {
	var list []models.Kosakata
	result := r.db.Preload("Levels").Where("reading = ?", reading).Find(&list)
	return list, result.Error
}

func (r *kosakataRepository) FindByRomaji(romaji string) ([]models.Kosakata, error) {
	var list []models.Kosakata
	result := r.db.Preload("Levels").Where("romaji = ?", romaji).Find(&list)
	return list, result.Error
}

func (r *kosakataRepository) FindByLevelJLPT(level string) ([]models.Kosakata, error) {
	var list []models.Kosakata
	result := r.db.Preload("Levels").
		Joins("JOIN kosakata_level_relations klr ON klr.kosakata_id = kosakatas.id").
		Joins("JOIN kosakata_levels kl ON kl.id = klr.kosakata_level_id").
		Where("kl.name = ?", level).Find(&list)
	return list, result.Error
}

func (r *kosakataRepository) FindAllLevels() ([]models.KosakataLevel, error) {
	var list []models.KosakataLevel
	result := r.db.Find(&list)
	return list, result.Error
}

func (r *kosakataRepository) Create(k *models.Kosakata) error {
	return r.db.Create(k).Error
}

func (r *kosakataRepository) Update(k *models.Kosakata) error {
	if err := r.db.Model(k).Association("Levels").Replace(k.Levels); err != nil {
		return err
	}
	return r.db.Save(k).Error
}

func (r *kosakataRepository) Delete(id uint) error {
	return r.db.Delete(&models.Kosakata{}, id).Error
}
