package database

import (
	"log"

	"github.com/midedickson/instashop/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDB() {
	d, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database sucessfully")
	DB = d
}

func AutoMigrate() {
	log.Println("Auto Migrating Models...")
	err := DB.AutoMigrate(&models.User{})
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
