package services

import (
	"kanakana/internal/models"
	"kanakana/internal/repository"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserLevels(userID uint) ([]models.KosakataLevel, error)
	UpdateUserLevels(userID uint, levelIDs []uint) error
	GetMasteryProgress(userID uint) ([]models.UserKosakata, error)
	UpdateMastery(userID uint, kosakataID uint, isCorrect bool) error
	AddExp(userID uint, exp int) (*models.UserDetail, bool, error)
	GetProfile(userID uint) (*models.UserProfileResponse, error)
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

func (s *userService) GetMasteryProgress(userID uint) ([]models.UserKosakata, error) {
	return s.userRepo.GetKosakataMastery(userID)
}

func (s *userService) UpdateMastery(userID uint, kosakataID uint, isCorrect bool) error {
	return s.userRepo.UpdateKosakataMastery(userID, kosakataID, isCorrect)
}

func (s *userService) AddExp(userID uint, exp int) (*models.UserDetail, bool, error) {
	return s.userRepo.AddExp(userID, exp)
}

func (s *userService) GetProfile(userID uint) (*models.UserProfileResponse, error) {
	return s.userRepo.GetProfile(userID)
}
