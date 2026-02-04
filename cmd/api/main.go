package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.load()

	log.Printf("starting server on :%s", cfg.AppPort)

}
