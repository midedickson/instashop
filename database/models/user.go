package models

import "gorm.io/gorm"

// user account details

type User struct {
	gorm.Model
	Email        string `gorm:"email" json:"email"`
	Role         string `gorm:"role" json:"role"`
	IsActive     string `gorm:"is_active" json:"is_active"`
	PasswordHash string `gorm:"password_hash" json:"password_hash"`
}
