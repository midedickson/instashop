package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string          `gorm:"name" json:"name"`
	Price    decimal.Decimal `gorm:"price" json:"price"`
	Quantity int             `gorm:"quantity" json:"quantity"`
}
