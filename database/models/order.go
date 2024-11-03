package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OwnerID  uint       `gorm:"owner_id" json:"owner_id"`
	Owner    User       `gorm:"owner" json:"owner"`
	Status   string     `gorm:"status" json:"status"`
	Products []*Product `gorm:"many2many:order_products;"`
}
