package controllers

import "github.com/midedickson/instashop/internal/services"

type Controller struct {
	userService    services.IUserService
	productservice services.IProductService
}

func NewController(userService services.IUserService, productservice services.IProductService) *Controller {
	return &Controller{userService: userService, productservice: productservice}
}
