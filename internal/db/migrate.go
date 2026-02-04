package db

import (
	"log"

	"github.com/bm-lamine/go-web-essentials.git/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
