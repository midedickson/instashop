package repository

import (
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/database/models"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/utils"
	"gorm.io/gorm"
)

type ProductProvider interface {
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(updateProductPayload dto.CreateUpdateDBProduct, id uint) (*models.Product, error)
}

type OrderRepository struct {
	db                *gorm.DB
	productRepository ProductProvider
}

func NewOrderRepository(db *gorm.DB, productRepository ProductProvider) *OrderRepository {
	return &OrderRepository{db: db, productRepository: productRepository}
}

func (r *OrderRepository) CreateOrder(createDBOrder dto.CreateDBOrder) (*models.Order, error) {
	order := &models.Order{
		OwnerID: createDBOrder.UserID,
		Status:  constants.ORDER_PENDING,
	}

	for _, item := range createDBOrder.Items {
		product, err := r.productRepository.GetProductByID(item.ProductID)
		if err != nil || item.Quantity > product.Quantity || product.Quantity == 0 {
			continue
		}
		orderItem := &models.OrderItem{
			Product:  product,
			Quantity: item.Quantity,
			Order:    order,
		}
		order.Items = append(order.Items, orderItem)
		// reduce the available product quantity
		r.productRepository.UpdateProduct(dto.CreateUpdateDBProduct{
			Name:     product.Name,
			Price:    int64(product.Price),
			Quantity: product.Quantity - item.Quantity,
		}, product.ID)
	}

	if len(order.Items) == 0 {
		return nil, utils.ErrNoAvailableProducts
	}

	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil

}

func (r *OrderRepository) GetAllOrders() ([]*models.Order, error) {
	var orders []*models.Order
	if err := r.db.Preload("Items").Preload("Items.Product").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetAllOrdersForUser(userID uint) ([]*models.Order, error) {
	var orders []*models.Order
	if err := r.db.Preload("Items").Preload("Items.Product").Where("owner_id =?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("Items").Preload("Items.Product").Where("id =?", id).First(&order).Error; err != nil {
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
