package dto

type CreateDBUser struct {
	Email        string `json:"email"`
	Role         string `json:"role"`
	PasswordHash string `json:"password_hash"`
}

type UpdateDBUser struct {
	IsActive bool `json:"is_active"`
}
