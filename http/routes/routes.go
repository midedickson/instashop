package routes

import (
	"github.com/gorilla/mux"
	"github.com/midedickson/instashop/http/controllers"
	"github.com/midedickson/instashop/http/middlewares"
)

func ConnectRoutes(r *mux.Router, controller *controllers.Controller) {
	r.HandleFunc("/", controller.Hello).Methods("GET")
	r.HandleFunc("/auth/signup", controller.CreateUser).Methods("POST")

	// Create a subrouter for authenticated routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middlewares.AuthMiddleware)

}
