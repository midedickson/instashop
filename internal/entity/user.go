package entity

type User struct {
	ID       uint
	Email    string
	IsActive bool
}

func (u *User) ToJwtPayload() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"email": u.Email,
	}
}
