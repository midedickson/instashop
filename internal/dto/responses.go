package dto

import "github.com/midedickson/instashop/internal/entity"

type LoginResponse struct {
	UserDetails *entity.User `json:"user_details"`
	AccessToken string       `json:"access_token"`
}
