package services

import (
	"time"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/midedickson/instashop/token"
	"github.com/midedickson/instashop/utils"
)

type IUserService interface {
	CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	VerifyUserPasswordWithHash(password string) bool
	GenerateJwtTokenForUser(user *entity.User) (string, error)
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error) {
	// check if user exists first
	if _, err := u.GetUserByEmail(createUserPayload.Email); err == nil {
		return nil, utils.ErrUserAlreadyExists
	}

	return &entity.User{ID: uint(34), Email: createUserPayload.Email}, nil
}

func (u *UserService) GetUserByEmail(email string) (*entity.User, error) {
	return &entity.User{ID: uint(34), Email: email}, nil
}

func (u *UserService) GetUserByID(id uint) (*entity.User, error) {
	return &entity.User{ID: id, Email: "dicksonmide@gmil.com"}, nil
}

// Auth - related methods
func (u *UserService) VerifyUserPasswordWithHash(password string) bool {
	// implement password verification logic here
	return true
}

func (u *UserService) GenerateJwtTokenForUser(user *entity.User) (string, error) {

	return token.GenerateHS256Token(&token.TokenGenOptions{
		Payload:    user.ToJwtPayload(),
		ExpiryDate: time.Now().Add(constants.JWT_DEFAULT_EXPIRATION_TIME),
	})
}
