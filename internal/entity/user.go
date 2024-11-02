package entity

type User struct {
	ID       uint
	Email    string
	IsActive bool
	Role     string
}

func (u *User) ToJwtPayload() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"email": u.Email,
		"role":  u.Role,
	}
}
