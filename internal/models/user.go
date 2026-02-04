package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	UserAgent string `gorm:"not null"`
	IPAddress string

	CreatedAt time.Time
	UpdatedAt time.Time
}
