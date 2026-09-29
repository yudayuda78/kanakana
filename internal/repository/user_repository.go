package repository

import (
	"math"
	"kanakana/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	GetLevels(userID uint) ([]models.KosakataLevel, error)
	UpdateLevels(userID uint, levels []models.KosakataLevel) error
	GetKosakataMastery(userID uint) ([]models.UserKosakata, error)
	UpdateKosakataMastery(userID uint, kosakataID uint, isCorrect bool) error
	AddExp(userID uint, exp int) (*models.UserDetail, bool, error)
	GetProfile(userID uint) (*models.UserProfileResponse, error)
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

func (r *userRepository) GetKosakataMastery(userID uint) ([]models.UserKosakata, error) {
	var list []models.UserKosakata
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *userRepository) UpdateKosakataMastery(userID uint, kosakataID uint, isCorrect bool) error {
	var userKosakata models.UserKosakata
	
	// Cek apakah record sudah ada
	err := r.db.Where("user_id = ? AND kosakata_id = ?", userID, kosakataID).First(&userKosakata).Error
	
	if err != nil {
		// Jika belum ada, buat record baru
		// Anggap gorm.ErrRecordNotFound, kita insert record baru
		change := 0
		if isCorrect {
			change = 1
		}
		
		newRecord := models.UserKosakata{
			UserID:     userID,
			KosakataID: kosakataID,
			Penguasaan: change,
		}
		return r.db.Create(&newRecord).Error
	}
	
	// Jika sudah ada, update
	if isCorrect {
		userKosakata.Penguasaan += 1
	} else {
		userKosakata.Penguasaan -= 1
	}
	
	if userKosakata.Penguasaan < 0 {
		userKosakata.Penguasaan = 0
	} else if userKosakata.Penguasaan > 20 {
		userKosakata.Penguasaan = 20
	}
	
	return r.db.Save(&userKosakata).Error
}

func (r *userRepository) AddExp(userID uint, exp int) (*models.UserDetail, bool, error) {
	var detail models.UserDetail
	err := r.db.Where("user_id = ?", userID).First(&detail).Error
	if err != nil {
		// Not found, create it
		detail = models.UserDetail{
			UserID: userID,
			Exp:    exp,
			Level:  1,
		}
		if err := r.db.Create(&detail).Error; err != nil {
			return nil, false, err
		}
	} else {
		detail.Exp += exp
	}
	
	oldLevel := detail.Level
	newLevel := int(math.Sqrt(float64(detail.Exp)/50)) + 1
	detail.Level = newLevel
	
	if err := r.db.Save(&detail).Error; err != nil {
		return nil, false, err
	}
	
	isLevelUp := newLevel > oldLevel
	return &detail, isLevelUp, nil
}

func (r *userRepository) GetProfile(userID uint) (*models.UserProfileResponse, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	
	var detail models.UserDetail
	err := r.db.Where("user_id = ?", userID).First(&detail).Error
	if err != nil {
		// If not found, use defaults
		detail = models.UserDetail{Exp: 0, Level: 1}
	}
	
	return &models.UserProfileResponse{
		Name:  user.Name,
		Email: user.Email,
		Exp:   detail.Exp,
		Level: detail.Level,
	}, nil
}
