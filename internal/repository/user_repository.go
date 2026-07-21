package repository

import (
	"kanakana/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	GetLevels(userID uint) ([]models.KosakataLevel, error)
	UpdateLevels(userID uint, levels []models.KosakataLevel) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetLevels(userID uint) ([]models.KosakataLevel, error) {
	var user models.User
	if err := r.db.Preload("Levels").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Levels, nil
}

func (r *userRepository) UpdateLevels(userID uint, levels []models.KosakataLevel) error {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}
	return r.db.Model(&user).Association("Levels").Replace(levels)
}
