package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/http/controllers"
	"github.com/midedickson/instashop/internal/http/middlewares"
)

func ConnectRoutes(r *mux.Router, controller *controllers.Controller) {
	r.HandleFunc("/", controller.Hello).Methods("GET")
	// authentication routes
	r.Handle("/auth/signup",
		middlewares.Chain(
			http.HandlerFunc(controller.CreateUser),
			middlewares.ValidatePayloadMiddleware(&dto.UserAuthPayload{}, constants.SignupPayloadCtxKey{}),
		)).Methods("POST")
	r.Handle("/auth/login",
		middlewares.Chain(
			http.HandlerFunc(controller.Login),
			middlewares.ValidatePayloadMiddleware(&dto.UserAuthPayload{}, constants.LoginPayloadCtxKey{}),
		)).Methods("POST")
	r.Handle("/auth/verify",
		middlewares.Chain(
			http.HandlerFunc(controller.ActivateUser),
			middlewares.ValidatePayloadMiddleware(&dto.ActivateUserPayload{}, constants.ActivateUserPayloadCtxKey{}),
		)).Methods("POST")

	// subrouter for authenticated routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middlewares.AuthMiddleware)
	// user details route
	protected.HandleFunc("/auth/me", controller.Me).Methods("GET")

	// product management routes
	protected.HandleFunc("/products", controller.GetAllProducts).Methods("GET")
	protected.HandleFunc("/products/{id}", controller.GetProduct).Methods("GET")

	// admin-only management routes
	protected.Handle("/products",
		middlewares.Chain(
			http.HandlerFunc(controller.CreateProduct),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
			middlewares.ValidatePayloadMiddleware(&dto.CreateProductPayload{}, constants.CreateProductPayloadCtxKey{}),
		)).Methods("POST")

	protected.Handle("/products/{id}",
		middlewares.Chain(
			http.HandlerFunc(controller.UpdateProduct),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
			middlewares.ValidatePayloadMiddleware(&dto.UpdateProductPayload{}, constants.UpdateProductPayloadCtxKey{}),
		)).Methods("PUT")

	protected.Handle("/products/{id}",
		middlewares.Chain(
			http.HandlerFunc(controller.DeleteProduct),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
		)).Methods("DELETE")

	// order management routes
	protected.Handle("/orders",
		middlewares.Chain(
			http.HandlerFunc(controller.CreateOrder),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
			middlewares.ValidatePayloadMiddleware(&dto.CreateOrderPayload{}, constants.CreateOrderPayloadCtxKey{}),
		)).Methods("POST")

	protected.HandleFunc("/orders", controller.GetAllOrdersForUser).Methods("GET")

	protected.Handle("/orders/all",
		middlewares.Chain(
			http.HandlerFunc(controller.GetAllOrders),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
		)).Methods("GET")

	protected.Handle("/orders/{id}",
		middlewares.Chain(
			http.HandlerFunc(controller.UpdateOrderStatus),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
			middlewares.ValidatePayloadMiddleware(&dto.UpdateOrderStatusPayload{}, constants.UpdateOrderStatusPayloadCtxKey{}),
		)).Methods("PATCH")

	protected.HandleFunc("/orders/{id}", controller.GetOrderById).Methods("GET")

	protected.HandleFunc("/orders/{id}/cancel", controller.CancelOrder).Methods("PATCH")
}
