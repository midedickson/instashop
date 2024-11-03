package entity

import "github.com/midedickson/instashop/constants"

type Product struct {
	ID       uint
	Name     string
	Price    constants.Money
	Quantity int
}
