package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Micmada/go-url-shortener/handlers"
)

// SetupRoutes defines all API endpoints and their handlers.
func SetupRoutes(r *gin.Engine) {
	// Group API routes under "/api"
	api := r.Group("/api")
	{
		// URL shortening routes
		api.POST("/shorten", handlers.ShortenURL) // Shorten a URL
		api.GET("/:shortURL", handlers.ResolveURL) // Resolve a short URL to the original

		// User authentication routes
		api.POST("/register", handlers.RegisterUser) // Register a new user
		api.POST("/login", handlers.LoginUser)       // User login and token generation
	}
}
