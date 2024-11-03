package controllers

import (
	"net/http"
	"strconv"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/utils"
)

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Create order logic
	createOrderPayload := r.Context().Value(constants.CreateOrderPayloadCtxKey{}).(*dto.CreateOrderPayload)
	userId := r.Context().Value(constants.UserIdCtxKey{}).(uint)

	user, err := c.userService.GetUserByID(userId)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	order, err := c.orderService.CreateOrder(*createOrderPayload, user)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	utils.Dispatch201(w, "Order created successfully", order)
}

func (c *Controller) GetAllOrdersForUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(constants.UserIdCtxKey{}).(uint)
	orders, err := c.orderService.GetAllOrdersForUser(userId)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Orders retrieved successfully", orders)
}

func (c *Controller) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := c.orderService.GetAllOrders()
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Orders retrieved successfully", orders)
}

func (c *Controller) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderIdStr, err := utils.GetPathParam(r, "id")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload, order Id must be an integer", err)
		return
	}
	order, err := c.orderService.GetOrderByID(uint(orderId))
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Order retrieved successfully", order)
}

func (c *Controller) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderIdStr, err := utils.GetPathParam(r, "id")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload, order Id must be an integer", err)
		return
	}
	updateOrderStatusPayload := r.Context().Value(constants.UpdateOrderStatusPayloadCtxKey{}).(*dto.UpdateOrderStatusPayload)

	order, err := c.orderService.UpdateOrderStatus(uint(orderId), *updateOrderStatusPayload)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "Order status updated successfully", order)
}

func (c *Controller) CancelOrder(w http.ResponseWriter, r *http.Request) {
	orderIdStr, err := utils.GetPathParam(r, "id")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload, order Id must be an integer", err)
		return
	}
	userId := r.Context().Value(constants.UserIdCtxKey{}).(uint)

	err = c.orderService.CancelOrder(uint(orderId), userId)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch204(w)
}
