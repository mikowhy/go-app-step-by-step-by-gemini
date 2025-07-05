package database

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-api-test/models"
)

var DB *gorm.DB

func init() {
	dsn := "test.db" // SQLite database file name
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	// Migrate the schema
	DB.AutoMigrate(&models.Item{})

	// Add some initial data if the table is empty
	var count int64
	DB.Model(&models.Item{}).Count(&count)

	if count == 0 {
		log.Println("Adding initial items...")

		items := []models.Item{
			{Name: "Laptop", Description: "A powerful portable computer"},
			{Name: "Mouse", Description: "Wireless optical mouse"},
			{Name: "Keyboard", Description: "Mechanical gaming keyboard"},
			{Name: "Monitor", Description: "27-inch 4K IPS monitor"},
			{Name: "Webcam", Description: "Full HD webcam with microphone"},
		}

		for i := range items {
			items[i].ID = uuid.New().String()
			DB.Create(&items[i])
		}

		log.Println("Initial items added.")
	}

	log.Println("Database migration complete.")
}
