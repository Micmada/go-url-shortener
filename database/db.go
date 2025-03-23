package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Micmada/go-url-shortener/models"
)

// DB is the global database connection instance.
var DB *gorm.DB

// ConnectDB initializes and connects to the SQLite database.
// It also performs automatic migrations for the URL and User models.
func ConnectDB() {
	// Open a connection to the SQLite database
	db, err := gorm.Open(sqlite.Open("shortener.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate ensures the database schema is up to date with the models
	if err := db.AutoMigrate(&models.URL{}, &models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Assign the database instance to the global variable
	DB = db
}
