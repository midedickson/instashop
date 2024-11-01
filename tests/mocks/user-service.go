package mocks

import (
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/stretchr/testify/mock"
)

// mocking the user service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error) {
	args := m.Called(createUserPayload)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id uint) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) VerifyUserPasswordWithHash(password string) bool {
	args := m.Called(password)
	return args.Bool(0)
}

func (m *MockUserService) GenerateJwtTokenForUser(user *entity.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
