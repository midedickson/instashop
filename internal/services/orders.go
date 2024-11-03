package services

import (
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/midedickson/instashop/utils"
)

var orderStatuses = map[string]string{
	constants.ORDER_PENDING:   constants.ORDER_PENDING,
	constants.ORDER_COMPLETED: constants.ORDER_COMPLETED,
	constants.ORDER_SHIPPED:   constants.ORDER_SHIPPED,
	constants.ORDER_CANCELLED: constants.ORDER_CANCELLED,
}

var validOrderTransitions = map[string]map[string]bool{
	constants.ORDER_PENDING: {
		constants.ORDER_SHIPPED:   true,
		constants.ORDER_CANCELLED: true,
	},
	constants.ORDER_SHIPPED: {
		constants.ORDER_COMPLETED: true,
	},
	constants.ORDER_CANCELLED: {},
}

type IProductService interface {
	GetAllProducts() ([]*entity.Product, error)
}

type OrderService struct {
	productService IProductService
}

func NewOrderService(productService IProductService) *OrderService {
	return &OrderService{productService: productService}
}

func (o *OrderService) CreateOrder(orderPayload dto.CreateOrderPayload, user *entity.User) (*entity.Order, error) {
	// create new order with products from product service
	products, err := o.productService.GetAllProducts()
	if err != nil {
		// handle error
		return nil, err
	}

	order := &entity.Order{
		ID:       uint(1),
		OwnerID:  user.ID,
		Status:   "pending",
		Products: products,
	}

	// save the order to the database
	//...

	return order, nil
}

func (o *OrderService) GetAllOrdersForUser(userID uint) ([]*entity.Order, error) {

	// retrieve all orders from the database for the given user
	orders := []*entity.Order{
		{
			ID:      uint(1),
			OwnerID: userID,
			Status:  "pending",
			Products: []*entity.Product{
				{ID: uint(1), Name: "Product 1"},
				{ID: uint(2), Name: "Product 2"},
			},
		},
	}

	return orders, nil
}

func (o *OrderService) GetAllOrders() ([]*entity.Order, error) {

	// retrieve all orders from the database for the given user
	orders := []*entity.Order{
		{
			ID:      uint(1),
			OwnerID: uint(1),
			Status:  "pending",
			Products: []*entity.Product{
				{ID: uint(1), Name: "Product 1"},
				{ID: uint(2), Name: "Product 2"},
			},
		},
	}

	return orders, nil
}

func (o *OrderService) GetOrderByID(id uint) (*entity.Order, error) {
	order := &entity.Order{
		ID:      uint(1),
		OwnerID: uint(23),
		Status:  "pending",
		Products: []*entity.Product{
			{ID: uint(1), Name: "Product 1"},
			{ID: uint(2), Name: "Product 2"},
		},
	}

	return order, nil
}

func (o *OrderService) UpdateOrderStatus(orderID uint, updateStatusPayload dto.UpdateOrderStatusPayload) (*entity.Order, error) {
	// check if order status exists
	if _, ok := orderStatuses[updateStatusPayload.Status]; !ok {
		return nil, utils.ErrInvalidOrderStatus
	}

	// check if order exists
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	// check if order is already in the requested status
	if order.Status == updateStatusPayload.Status {
		return nil, utils.ErrOrderAlreadyInStatus
	}
	// check if order is in a valid transition state
	if _, ok := validOrderTransitions[order.Status][updateStatusPayload.Status]; !ok {
		return nil, utils.ErrInvalidOrderTransition
	}
	// update the status of the order
	order.Status = updateStatusPayload.Status

	// save the updated order to the database
	//...

	return order, nil
}

func (o *OrderService) CancelOrder(orderID, userId uint) error {
	// check if order exists
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	// check if the order is cancelled by the owner
	if order.OwnerID != userId {
		return utils.ErrUnauthorized
	}

	// check if order is in the pending state
	if order.Status != constants.ORDER_PENDING {
		return utils.ErrInvalidOrderTransition
	}

	// update the status of the order
	order.Status = constants.ORDER_CANCELLED

	// save the updated order to the database
	//...

	return nil
}
