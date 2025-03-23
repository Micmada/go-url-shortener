package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/Micmada/go-url-shortener/database"
	"github.com/Micmada/go-url-shortener/models"
)

// ShortenURL handles URL shortening requests.
// It validates and stores the original URL with a generated short URL.
func ShortenURL(c *gin.Context) {
	var req struct {
		OriginalURL string `json:"original_url"`
	}

	// Parse request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the URL has a valid scheme (http or https)
	if !strings.HasPrefix(req.OriginalURL, "http://") && !strings.HasPrefix(req.OriginalURL, "https://") {
		req.OriginalURL = "http://" + req.OriginalURL
	}

	// Generate a unique short URL
	shortURL := generateShortURL()

	// Store the URL in the database
	url := models.URL{OriginalURL: req.OriginalURL, ShortURL: shortURL}
	database.DB.Create(&url)

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// ResolveURL handles redirection from a short URL to the original URL.
func ResolveURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	var url models.URL

	// Find the original URL in the database
	if err := database.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url.OriginalURL)
}

// generateShortURL creates a random 6-character short URL.
func generateShortURL() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "xyz123" // Fallback in case of failure
	}
	return base64.URLEncoding.EncodeToString(b)[:6]
}
