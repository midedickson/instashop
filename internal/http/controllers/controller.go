package controllers

type Controller struct {
	userService    IUserService
	productservice IProductService
	orderService   IOrderService
}

func NewController(userService IUserService, productservice IProductService, orderServive IOrderService) *Controller {
	return &Controller{userService: userService, productservice: productservice, orderService: orderServive}
}
