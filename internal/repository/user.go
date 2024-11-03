package repository

import (
	"github.com/midedickson/instashop/database/models"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email =?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id =?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) CreateNewUser(userParam dto.CreateDBUser) (*models.User, error) {
	// check if user exists
	if err := r.db.Where("email =?", userParam.Email).First(&models.User{}).Error; err == nil {
		return nil, utils.ErrUserAlreadyExists
	}

	// create new user
	user := &models.User{
		Email:        userParam.Email,
		Role:         userParam.Role,
		IsActive:     false,
		PasswordHash: userParam.PasswordHash,
	}

	err := r.db.Create(user).Error
	return user, err
}

func (r *UserRepository) UpdateUser(userParam dto.UpdateDBUser, email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email =?", email).First(&user).Error; err != nil {
		return nil, err
	}

	user.IsActive = userParam.IsActive

	err := r.db.Save(&user).Error
	return &user, err
}
