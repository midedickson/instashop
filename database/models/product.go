package models

import (
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string          `gorm:"name" json:"name"`
	Price    constants.Money `gorm:"price" json:"price"`
	Quantity int             `gorm:"quantity" json:"quantity"`
}

func (p *Product) ToEntity() *entity.Product {
	return &entity.Product{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
}
