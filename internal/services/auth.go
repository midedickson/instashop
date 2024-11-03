package services

import (
	"time"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/database/models"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/midedickson/instashop/token"
	"github.com/midedickson/instashop/utils"
)

// required repository interface

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateNewUser(userParam dto.CreateDBUser) (*models.User, error)
	UpdateUser(userParam dto.UpdateDBUser, userEmail string) (*models.User, error)
}

// UserService is responsible for managing user related operations
type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error) {
	passwdHash, err := utils.HashPassword(createUserPayload.Password)
	if err != nil {
		return nil, err
	}
	var userRole string
	if createUserPayload.Role == constants.ADMIN_ROLE {
		userRole = constants.ADMIN_ROLE
	} else {
		userRole = constants.CUSTOMER_ROLE
	}
	newUser, err := u.userRepository.CreateNewUser(
		dto.CreateDBUser{
			Email:        createUserPayload.Email,
			PasswordHash: passwdHash,
			Role:         userRole,
		},
	)
	if err != nil {
		return nil, err
	}
	return newUser.ToEntity(), nil
}

func (u *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

func (u *UserService) GetUserByID(id uint) (*entity.User, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

// Auth - related methods
func (u *UserService) VerifyUserPasswordWithHash(user *entity.User, password string) bool {
	dbUser, err := u.userRepository.GetUserByID(user.ID)
	if err != nil {
		return false
	}
	return utils.CheckPassword(password, dbUser.PasswordHash)
}

func (u *UserService) GenerateJwtTokenForUser(user *entity.User) (string, error) {
	return token.GenerateHS256Token(&token.TokenGenOptions{
		Payload:    user.ToJwtPayload(),
		ExpiryDate: time.Now().UTC().Add(constants.JWT_DEFAULT_EXPIRATION_TIME),
	})
}

func (u *UserService) ActivateUser(activateUserPayload dto.ActivateUserPayload) (*entity.User, error) {
	user, err := u.userRepository.UpdateUser(dto.UpdateDBUser{
		IsActive: true,
	}, activateUserPayload.Email)

	if err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}
