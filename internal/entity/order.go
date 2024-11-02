package entity

type Order struct {
	ID       uint
	OwnerID  uint
	Status   string
	Products []*Product
}
