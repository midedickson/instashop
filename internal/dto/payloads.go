package dto

type UserAuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *UserAuthPayload) Validate() bool {
	if p.Email == "" || p.Password == "" {
		return false
	}
	return true
}
