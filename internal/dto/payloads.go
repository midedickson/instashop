package dto

type UserAuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (p *UserAuthPayload) Validate() bool {
	if p.Email == "" || p.Password == "" {
		return false
	}
	return true
}

type ActivateUserPayload struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (p *ActivateUserPayload) Validate() bool {
	if p.Email == "" || p.Code == "" {
		return false
	}
	return true
}
