package repository

import (
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/database/models"
	"github.com/midedickson/instashop/internal/dto"
	"gorm.io/gorm"
)

type ProductProvider interface {
	GetProductByID(id uint) (*models.Product, error)
}

type OrderRepository struct {
	db                *gorm.DB
	productRepository ProductProvider
}

func NewOrderRepository(db *gorm.DB, productRepository ProductProvider) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(createDBOrder dto.CreateDBOrder) (*models.Order, error) {
	order := &models.Order{
		OwnerID: createDBOrder.UserID,
	}

	for _, item := range createDBOrder.Items {
		product, err := r.productRepository.GetProductByID(item.ProductID)
		if err != nil {
			continue
		}
		orderItem := &models.OrderItem{
			Product:  product,
			Quantity: item.Quantity,
		}
		order.Items = append(order.Items, orderItem)
	}

	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil

}

func (r *OrderRepository) GetAllOrders() ([]*models.Order, error) {
	var orders []*models.Order
	if err := r.db.Preload("items").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetAllOrdersForUser(userID uint) ([]*models.Order, error) {
	var orders []*models.Order
	if err := r.db.Preload("items").Where("owner_id =?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("items").Where("id =?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) CancelOrder(id uint) error {
	var order *models.Order

	if err := r.db.Where("id =?", id).First(&order).Error; err != nil {
		return err
	}

	order.Status = constants.ORDER_CANCELLED
	return r.db.Save(order).Error
}

func (r *OrderRepository) UpdateOrderStatus(status string, orderId uint) (*models.Order, error) {
	var order *models.Order

	if err := r.db.Where("id =?", orderId).First(&order).Error; err != nil {
		return nil, err
	}

	order.Status = status
	return order, r.db.Save(order).Error
}
