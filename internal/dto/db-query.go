package dto

// auth and user
type CreateDBUser struct {
	Email        string `json:"email"`
	Role         string `json:"role"`
	PasswordHash string `json:"password_hash"`
}

type UpdateDBUser struct {
	IsActive bool `json:"is_active"`
}

// product management queries
type CreateUpdateDBProduct struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

// order management queries
type CreateDBOrder struct {
	UserID uint `json:"user_id"`
	Items  []*CreateOrderItemPayload
}
