package main

import (
	"log"
	"net/http"

	"github.com/bm-lamine/go-web-essentials.git/internal/config"
	"github.com/bm-lamine/go-web-essentials.git/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	log.Printf("starting server on :%s", cfg.PORT)

	if err := http.ListenAndServe(":"+cfg.PORT, server.New()); err != nil {
		log.Fatal(err)
	}
}
