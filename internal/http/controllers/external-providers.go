package controllers

import (
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
)

type IUserService interface {
	CreateUser(createUserPayload dto.UserAuthPayload) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	VerifyUserPasswordWithHash(password string) bool
	GenerateJwtTokenForUser(user *entity.User) (string, error)
	ActivateUser(activateUserPayload dto.ActivateUserPayload) (*entity.User, error)
}

type IOrderService interface {
	CreateOrder(orderPayload dto.CreateOrderPayload, user *entity.User) (*entity.Order, error)
	GetAllOrdersForUser(userID uint) ([]*entity.Order, error)
	GetOrderByID(id uint) (*entity.Order, error)
	UpdateOrderStatus(orderID uint, updateStatusPayload dto.UpdateOrderStatusPayload) (*entity.Order, error)
	CancelOrder(orderID, userId uint) error
	GetAllOrders() ([]*entity.Order, error)
}

type IProductService interface {
	CreateProduct(createProductPayload dto.CreateProductPayload) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	UpdateProduct(id uint, updateProductPayload dto.UpdateProductPayload) (*entity.Product, error)
	DeleteProduct(id uint) error
}
