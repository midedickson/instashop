package entity

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

func (u *User) ToJwtPayload() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"email": u.Email,
		"role":  u.Role,
	}
}
