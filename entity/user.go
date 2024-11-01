package entity

type User struct {
	ID      uint
	Email   string
	IsValid bool
}

func NewUserEntity(id uint, email string) *User {
	return &User{ID: id, Email: email, IsValid: false}
}
