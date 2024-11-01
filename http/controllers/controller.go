package controllers

import "github.com/midedickson/instashop/usecases"

type Controller struct {
	userUseCase usecases.UserUseCase
}

func NewController(userUseCase usecases.UserUseCase) *Controller {
	return &Controller{userUseCase: userUseCase}
}
