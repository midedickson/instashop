package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/midedickson/instashop/constants"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func GetDBUrl() string {
	return os.Getenv("DATABASE_URL")
}

func GetPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("Error:", err, "using default port:", constants.PORT)
		return constants.PORT
	}
	return port
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}
