package services

import (
	"kanakana/internal/models"
	"kanakana/internal/repository"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserLevels(userID uint) ([]models.KosakataLevel, error)
	UpdateUserLevels(userID uint, levelIDs []uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUserLevels(userID uint) ([]models.KosakataLevel, error) {
	return s.userRepo.GetLevels(userID)
}

func (s *userService) UpdateUserLevels(userID uint, levelIDs []uint) error {
	var selected []models.KosakataLevel
	for _, id := range levelIDs {
		selected = append(selected, models.KosakataLevel{
			Model: gorm.Model{ID: id},
		})
	}
	return s.userRepo.UpdateLevels(userID, selected)
}
