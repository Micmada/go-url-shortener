package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/Micmada/go-url-shortener/database"
	"github.com/Micmada/go-url-shortener/routes"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Connect to the database
	database.ConnectDB()

	// Setup API routes
	routes.SetupRoutes(r)

	// Root endpoint for basic API info
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Go URL Shortener API!"})
	})

	// Start the server on port 8080
	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}