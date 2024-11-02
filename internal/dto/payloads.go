package dto

import "github.com/shopspring/decimal"

type UserAuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (p *UserAuthPayload) Validate() bool {
	if p.Email == "" || p.Password == "" {
		return false
	}
	return true
}

type ActivateUserPayload struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (p *ActivateUserPayload) Validate() bool {
	if p.Email == "" || p.Code == "" {
		return false
	}
	return true
}

type CreateProductPayload struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (p *CreateProductPayload) Validate() bool {
	if p.Name == "" || p.Price <= 0 || p.Quantity <= 0 {
		return false
	}
	return true
}

func (p *CreateProductPayload) DecimalPrice() decimal.Decimal {
	return decimal.NewFromFloat(p.Price)
}

type UpdateProductPayload struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (p *UpdateProductPayload) Validate() bool {
	return p.Name != ""
}

type CreateOrderPayload struct {
	ProductIds []uint `json:"products"`
}

func (p *CreateOrderPayload) Validate() bool {
	return len(p.ProductIds) > 0
}

type CancelOrderPayload struct {
	OrderID uint `json:"order_id"`
}

func (p *CancelOrderPayload) Validate() bool {
	return true
}

type UpdateOrderStatusPayload struct {
	Status string `json:"status"`
}

func (p *UpdateOrderStatusPayload) Validate() bool {
	return p.Status != ""
}
