package models

import "gorm.io/gorm"

// User represents a user account in the database.
type User struct {
	gorm.Model          // Embeds fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Username string `gorm:"unique;not null"` // Unique username, cannot be null
	Password string `gorm:"not null"`        // Hashed password, cannot be null
}
