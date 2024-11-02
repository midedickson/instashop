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

	// Create a subrouter for authenticated routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middlewares.AuthMiddleware)
}
