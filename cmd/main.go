package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/midedickson/instashop/config"
	"github.com/midedickson/instashop/database"
	"github.com/midedickson/instashop/http"
	"github.com/midedickson/instashop/http/controllers"
	"github.com/midedickson/instashop/http/routes"
	"github.com/midedickson/instashop/usecases"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Starting Instashop Application...")

	// start the database connection and auto migrate
	database.ConnectToDB()
	// make sure to close the DB connection on app close
	database.AutoMigrate()

	// create the router
	router := mux.NewRouter()

	// create usecases
	userUseCase := usecases.NewUserUseCaseService()

	// create controller with usecases
	controller := controllers.NewController(userUseCase)

	// connect routes
	routes.ConnectRoutes(router, controller)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// get appplication port
	port := config.GetPort()

	// create application
	app := http.NewApp("0.0.0.0", port, router)

	go func() {
		// create new http app and run in goroutine
		// this allows us continue running any background tasks if any
		log.Fatal(app.Run())
	}()
	fmt.Println("Instashop Application started successfully on port: ", port)

	<-stop
	log.Println("Shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	database.CloseDB()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
