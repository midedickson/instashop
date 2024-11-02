package entity

import "github.com/shopspring/decimal"

type Product struct {
	ID       uint
	Name     string
	Price    decimal.Decimal
	Quantity int
}
