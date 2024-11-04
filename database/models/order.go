package models

import (
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/midedickson/instashop/utils"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	Product   *Product `gorm:"product" json:"product"`
	ProductID uint     `gorm:"product_id" json:"product_id"`
	Quantity  int      `gorm:"quantity" json:"quantity"`
	Order     *Order   `gorm:"foreignKey:OrderID" json:"order"`
	OrderID   uint     `gorm:"column:order_id" json:"order_id"`
}

type Order struct {
	gorm.Model
	OwnerID uint         `gorm:"owner_id" json:"owner_id"`
	Owner   User         `gorm:"owner" json:"owner"`
	Status  string       `gorm:"status" json:"status"`
	Items   []*OrderItem `gorm:"foreignKey:OrderID"`
	Total   constants.Money
}

// CalculateTotal calculates the total price for the order
func (o *Order) CalculateTotal() {
	total := 0
	for _, item := range o.Items {
		total += int(item.Product.Price) * item.Quantity
	}

	o.Total = constants.Money(total)
}

// BeforeSave is a GORM hook that ensures the total is calculated before saving
func (o *Order) BeforeSave(tx *gorm.DB) error {
	o.CalculateTotal()
	return nil
}

func (o *Order) ToEntity() *entity.Order {
	orderItems := utils.MapConcurrent(o.Items, func(orderItem *OrderItem) *entity.OrderItem {
		return &entity.OrderItem{
			Product:  orderItem.Product.ToEntity(),
			Quantity: orderItem.Quantity,
		}
	})
	return &entity.Order{
		ID:      o.ID,
		OwnerID: o.OwnerID,
		Status:  o.Status,
		Items:   orderItems,
		Total:   o.Total,
	}
}
