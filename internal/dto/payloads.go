package dto

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
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

func (p *CreateProductPayload) Validate() bool {
	if p.Name == "" || p.Price <= 0 || p.Quantity <= 0 {
		return false
	}
	return true
}

type UpdateProductPayload struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

func (p *UpdateProductPayload) Validate() bool {
	return p.Name != ""
}

type CreateOrderItemPayload struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateOrderPayload struct {
	Items []CreateOrderItemPayload `json:"products"`
}

func (p *CreateOrderPayload) Validate() bool {
	return len(p.Items) > 0
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
