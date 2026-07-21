package services

import (
	"errors"
	"kanakana/internal/models"
	"kanakana/internal/repository"

	"gorm.io/gorm"
)

// KosakataService mendefinisikan kontrak logika bisnis untuk entitas Kosakata.
type KosakataService interface {
	GetAll() ([]models.Kosakata, error)
	GetByID(id uint) (*models.Kosakata, error)
	GetByKanji(kanji string) ([]models.Kosakata, error)
	GetByReading(reading string) ([]models.Kosakata, error)
	GetByRomaji(romaji string) ([]models.Kosakata, error)
	GetByLevelJLPT(level string) ([]models.Kosakata, error)
	GetAllLevels() ([]models.KosakataLevel, error)
	Create(k *models.Kosakata) error
	Update(id uint, k *models.Kosakata) (*models.Kosakata, error)
	Delete(id uint) error
}

type kosakataService struct {
	repo repository.KosakataRepository
}

// NewKosakataService membuat instance baru KosakataService.
func NewKosakataService(repo repository.KosakataRepository) KosakataService {
	return &kosakataService{repo: repo}
}

func (s *kosakataService) GetAll() ([]models.Kosakata, error) {
	return s.repo.FindAll()
}

func (s *kosakataService) GetByID(id uint) (*models.Kosakata, error) {
	k, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return k, nil
}

func (s *kosakataService) GetByKanji(kanji string) ([]models.Kosakata, error) {
	return s.repo.FindByKanji(kanji)
}

func (s *kosakataService) GetByReading(reading string) ([]models.Kosakata, error) {
	return s.repo.FindByReading(reading)
}

func (s *kosakataService) GetByRomaji(romaji string) ([]models.Kosakata, error) {
	return s.repo.FindByRomaji(romaji)
}

func (s *kosakataService) GetByLevelJLPT(level string) ([]models.Kosakata, error) {
	return s.repo.FindByLevelJLPT(level)
}

func (s *kosakataService) GetAllLevels() ([]models.KosakataLevel, error) {
	return s.repo.FindAllLevels()
}

func (s *kosakataService) Create(k *models.Kosakata) error {
	if err := validate(k); err != nil {
		return err
	}
	return s.repo.Create(k)
}

func (s *kosakataService) Update(id uint, k *models.Kosakata) (*models.Kosakata, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	// Terapkan perubahan ke data yang ada
	existing.Kanji = k.Kanji
	existing.Reading = k.Reading
	existing.Romaji = k.Romaji
	existing.Arti = k.Arti
	existing.Levels = k.Levels

	if err := validate(existing); err != nil {
		return nil, err
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *kosakataService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	return s.repo.Delete(id)
}

// validate memvalidasi field wajib pada Kosakata.
func validate(k *models.Kosakata) error {
	if k.Kanji == "" {
		return errors.New("kanji tidak boleh kosong")
	}
	if k.Reading == "" {
		return errors.New("reading tidak boleh kosong")
	}
	if k.Romaji == "" {
		return errors.New("romaji tidak boleh kosong")
	}
	if k.Arti == "" {
		return errors.New("arti tidak boleh kosong")
	}
	if len(k.Levels) == 0 {
		return errors.New("kosakata harus memiliki setidaknya satu level")
	}
	return nil
}

// ErrNotFound adalah sentinel error untuk data yang tidak ditemukan.
var ErrNotFound = errors.New("kosakata tidak ditemukan")
