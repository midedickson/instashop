package database

import (
	"fmt"
	"log"

	"github.com/midedickson/instashop/config"
	"github.com/midedickson/instashop/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetDBHost(),
		config.GetDBUser(),
		config.GetDBPassword(),
		config.GetDBName(),
		config.GetDBPort(),
	)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database sucessfully")
	DB = d
}

func AutoMigrate() {
	log.Println("Auto Migrating Models...")
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migrated DB Successfully")
}

func CloseDB() {
	log.Println("Closing database connection...")
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Database connection closed successfully")
}
