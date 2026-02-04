package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bm-lamine/go-web-essentials.git/internal/config"
	"github.com/bm-lamine/go-web-essentials.git/internal/db"
	"github.com/bm-lamine/go-web-essentials.git/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	database := db.NewClient(cfg.DATABASE_URL)
	db.Migrate(database)

	redis := db.NewCache(cfg.REDIS_URL, 0)
	log.Println("Connected to Redis!")

	// Example usage: set a key
	if err := redis.Set(context.Background(), "example:key", "hello redis", 10*time.Minute).Err(); err != nil {
		log.Fatalf("failed to set key: %v", err)
	}

	log.Printf("starting server on :%s", cfg.PORT)
	if err := http.ListenAndServe(":"+cfg.PORT, server.New()); err != nil {
		log.Fatal(err)
	}
}
