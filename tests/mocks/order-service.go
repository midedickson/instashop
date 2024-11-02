package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
)

// MockOrderService is a mock implementation of IOrderService
type MockOrderService struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: orderPayload, user
func (m *MockOrderService) CreateOrder(orderPayload dto.CreateOrderPayload, user *entity.User) (*entity.Order, error) {
	args := m.Called(orderPayload, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

// GetAllOrdersForUser provides a mock function with given fields: userID
func (m *MockOrderService) GetAllOrdersForUser(userID uint) ([]*entity.Order, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func (m *MockOrderService) GetAllOrders() ([]*entity.Order, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Order), args.Error(1)
}

// GetOrderByID provides a mock function with given fields: id
func (m *MockOrderService) GetOrderByID(id uint) (*entity.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

// UpdateOrderStatus provides a mock function with given fields: orderID, updateStatusPayload
func (m *MockOrderService) UpdateOrderStatus(orderID uint, updateStatusPayload dto.UpdateOrderStatusPayload) (*entity.Order, error) {
	args := m.Called(orderID, updateStatusPayload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

// CancelOrder provides a mock function with given fields: orderID
func (m *MockOrderService) CancelOrder(orderID, userID uint) error {
	args := m.Called(orderID)
	return args.Error(0)
}
