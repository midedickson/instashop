package usecases

import (
	"github.com/midedickson/instashop/dto"
	"github.com/midedickson/instashop/entity"
	"github.com/midedickson/instashop/utils"
)

type UserUseCase interface {
	CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error)
	GetUser(email string) (*entity.User, error)
}

type UserUseCaseService struct{}

func NewUserUseCaseService() *UserUseCaseService {
	return &UserUseCaseService{}
}

func (u *UserUseCaseService) CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error) {
	// check if user exists first
	if _, err := u.GetUser(createUserPayload.Email); err == nil {
		return nil, utils.ErrUserAlreadyExists
	}

	return &entity.User{ID: uint(34), Email: createUserPayload.Email}, nil
}

func (u *UserUseCaseService) GetUser(email string) (*entity.User, error) {
	return &entity.User{ID: uint(34), Email: email}, nil
}
