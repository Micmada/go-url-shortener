package handlers

import (
	"net/http"

	"github.com/Micmada/go-url-shortener/database"
	"github.com/Micmada/go-url-shortener/models"
	"github.com/Micmada/go-url-shortener/utils"
	"github.com/gin-gonic/gin"
)

// RegisterUser handles user registration.
// It hashes the password before storing it in the database.
func RegisterUser(c *gin.Context) {
	var user models.User

	// Bind JSON request to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password before storing it
	user.Password = utils.HashPassword(user.Password)

	// Save user to the database
	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user authentication.
// It verifies user credentials and returns a JWT token if successful.
func LoginUser(c *gin.Context) {
	var user models.User
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve user from database by username
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if password matches the hashed password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := utils.GenerateJWT(user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
