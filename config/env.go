package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/midedickson/instashop/constants"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("using default port:", constants.PORT)
		return constants.PORT
	}
	return port
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetDBHost() string {
	return getEnvOrDefault("DB_HOST", "localhost")
}

func GetDBPort() string {
	return getEnvOrDefault("DB_PORT", "5432")
}

func GetDBName() string {
	return getEnvOrDefault("DB_NAME", "instashop")
}

func GetDBUser() string {
	return getEnvOrDefault("DB_USER", "postgres")
}

func GetDBPassword() string {
	return getEnvOrDefault("DB_PASSWORD", "password")
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
