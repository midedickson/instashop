package services

import (
	"errors"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/database/models"
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
	constants.ORDER_COMPLETED: {},
}

type IProductService interface {
	GetAllProducts() ([]*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
}
type OrderRepository interface {
	GetAllOrdersForUser(userID uint) ([]*models.Order, error)
	GetAllOrders() ([]*models.Order, error)
	CreateOrder(createDBOrder dto.CreateDBOrder) (*models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrderStatus(status string, orderId uint) (*models.Order, error)
	CancelOrder(id uint) error
}

type OrderService struct {
	productService  IProductService
	orderRepository OrderRepository
}

func NewOrderService(productService IProductService, orderRepository OrderRepository) *OrderService {
	return &OrderService{productService: productService, orderRepository: orderRepository}
}

func (o *OrderService) CreateOrder(orderPayload dto.CreateOrderPayload, user *entity.User) (*entity.Order, error) {
	// create new order with products from product service
	var orderItems []*dto.CreateOrderItemPayload
	for _, item := range orderPayload.Items {
		product, err := o.productService.GetProductByID(item.ProductID)
		if err != nil {
			continue
		}

		orderItems = append(orderItems, &dto.CreateOrderItemPayload{
			ProductID: product.ID,
			Quantity:  item.Quantity,
		})
	}

	if len(orderItems) == 0 {
		return nil, errors.New("no valid products in order")
	}

	// create new order with products from product service
	createDBOrder := dto.CreateDBOrder{
		UserID: user.ID,
		Items:  orderItems,
	}

	order, err := o.orderRepository.CreateOrder(createDBOrder)

	if err != nil {
		return nil, err
	}

	return order.ToEntity(), nil
}

func (o *OrderService) GetAllOrdersForUser(userID uint) ([]*entity.Order, error) {

	// retrieve all orders from the database for the given user
	orders, err := o.orderRepository.GetAllOrdersForUser(userID)
	if err != nil {
		return nil, err
	}
	return utils.MapConcurrent(orders, func(order *models.Order) *entity.Order { return order.ToEntity() }), nil
}

func (o *OrderService) GetAllOrders() ([]*entity.Order, error) {

	// retrieve all orders from the database for the given user
	orders, err := o.orderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return utils.MapConcurrent(orders, func(order *models.Order) *entity.Order { return order.ToEntity() }), nil
}

func (o *OrderService) GetOrderByID(id uint) (*entity.Order, error) {
	order, err := o.orderRepository.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	return order.ToEntity(), nil
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

	// save the updated order to the database
	dbOrder, err := o.orderRepository.UpdateOrderStatus(updateStatusPayload.Status, orderID)
	if err != nil {
		return nil, err
	}

	return dbOrder.ToEntity(), nil
}

func (o *OrderService) CancelOrder(orderID, userId uint) error {
	// check if order exists
	order, err := o.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	// check if the order is cancelled by the owner
	if int64(order.OwnerID) != int64(userId) {
		return utils.ErrUnauthorized
	}

	// check if order is in the pending state
	if order.Status != constants.ORDER_PENDING {
		return utils.ErrInvalidOrderTransition
	}

	// save the updated order to the database
	err = o.orderRepository.CancelOrder(orderID)
	if err != nil {
		return err
	}
	return nil
}
