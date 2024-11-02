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
			middlewares.ValidatePayloadMiddleware(&dto.UserAuthPayload{}, constants.SignupPayloadCtxKey{}),
		)).Methods("POST")
	r.Handle("/auth/login",
		middlewares.Chain(
			http.HandlerFunc(controller.Login),
			middlewares.ValidatePayloadMiddleware(&dto.UserAuthPayload{}, constants.SignupPayloadCtxKey{}),
		)).Methods("POST")
	r.Handle("/auth/login/admin",
		middlewares.Chain(
			http.HandlerFunc(controller.AdminLogin),
			middlewares.ValidatePayloadMiddleware(&dto.UserAuthPayload{}, constants.SignupPayloadCtxKey{}),
		)).Methods("POST")
	r.Handle("/auth/verify",
		middlewares.Chain(
			http.HandlerFunc(controller.ActivateUser),
			middlewares.ValidatePayloadMiddleware(&dto.ActivateUserPayload{}, constants.ActivateUserPayloadCtxKey{}),
		)).Methods("POST")

	// subrouter for authenticated routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middlewares.AuthMiddleware)

	// product management routes
	protected.Handle("/products",
		middlewares.Chain(
			http.HandlerFunc(controller.GetAllProducts),
		)).Methods("GET")

	// admin-only management routes
	protected.Handle("/products",
		middlewares.Chain(
			http.HandlerFunc(controller.CreateProduct),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
		)).Methods("POST")

	protected.Handle("/products/{id}",
		middlewares.Chain(
			http.HandlerFunc(controller.UpdateProduct),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
		)).Methods("PUT")

	protected.Handle("/products/{id}",
		middlewares.Chain(
			http.HandlerFunc(controller.DeleteProduct),
			middlewares.PermissionMiddleware(constants.ADMIN_ROLE),
		)).Methods("GET")
}
