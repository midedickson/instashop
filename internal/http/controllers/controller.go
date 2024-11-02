package controllers

import "github.com/midedickson/instashop/internal/services"

type Controller struct {
	userService    services.IUserService
	productservice services.IProductService
	orderService   services.IOrderService
}

func NewController(userService services.IUserService, productservice services.IProductService, orderServive services.IOrderService) *Controller {
	return &Controller{userService: userService, productservice: productservice, orderService: orderServive}
}
