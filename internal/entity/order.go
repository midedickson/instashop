package entity

import "github.com/midedickson/instashop/constants"

type OrderItem struct {
	ID       uint     `json:"id"`
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
}
type Order struct {
	ID      uint            `json:"id"`
	OwnerID uint            `json:"owner_id"`
	Status  string          `json:"status"`
	Items   []*OrderItem    `json:"order_items"`
	Total   constants.Money `json:"total"`
}
