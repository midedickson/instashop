package controllers

import "github.com/midedickson/instashop/internal/services"

type Controller struct {
	userService services.IUserService
}

func NewController(userService services.IUserService) *Controller {
	return &Controller{userService: userService}
}
