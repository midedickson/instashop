package mocks

import (
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockProductService is a mock implementation of IProductService
type MockProductService struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: createProductPayload
func (m *MockProductService) CreateProduct(createProductPayload dto.CreateProductPayload) (*entity.Product, error) {
	args := m.Called(createProductPayload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), args.Error(1)
}

// GetAllProducts provides a mock function with given fields:
func (m *MockProductService) GetAllProducts() ([]*entity.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Product), args.Error(1)
}

// GetProductByID provides a mock function with given fields: id
func (m *MockProductService) GetProductByID(id uint) (*entity.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), args.Error(1)
}

// UpdateProduct provides a mock function with given fields: id, updateProductPayload
func (m *MockProductService) UpdateProduct(id uint, updateProductPayload dto.UpdateProductPayload) (*entity.Product, error) {
	args := m.Called(id, updateProductPayload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), args.Error(1)
}

// DeleteProduct provides a mock function with given fields: id
func (m *MockProductService) DeleteProduct(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
