package models

import (
	"github.com/midedickson/instashop/internal/entity"
	"gorm.io/gorm"
)

// user account details

type User struct {
	gorm.Model
	Email        string `gorm:"email" json:"email"`
	Role         string `gorm:"role" json:"role"`
	IsActive     bool   `gorm:"is_active" json:"is_active"`
	PasswordHash string `gorm:"password_hash" json:"password_hash"`
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Email:    u.Email,
		Role:     u.Role,
		IsActive: u.IsActive,
	}
}
