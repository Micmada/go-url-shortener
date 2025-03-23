package models

// URL represents a shortened URL entry in the database.
type URL struct {
	ID          uint   `gorm:"primaryKey"`   // Unique ID for the URL record
	OriginalURL string `gorm:"not null"`     // The original full URL
	ShortURL    string `gorm:"uniqueIndex"` // The generated short URL, must be unique
}
