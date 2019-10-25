package models

import (
	"time"
)

// Blacklist represents blacklist's IP that aren't allowed to login
type Blacklist struct {
	ID         int `gorm:"primary_key"`
	UserID     int
	IPAddress  string `json:"ip_address" binding:"required"`
	LoginCount int    `gorm:"default:'1'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}