package models

// user account details

type User struct {
	ID           int    `json:"account_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
