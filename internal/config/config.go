package config

import (
	"log"
	"os"
)

type Config struct {
	PORT         string
	DATABASE_URL string
	REDIS_URL    string
}

func load() *Config {
	cfg := &Config{
		PORT:         getEnv("PORT", "3000"),
		DATABASE_URL: getEnv("DATABASE_URL", ""),
		REDIS_URL:    getEnv("REDIS_URL", "redis://localhost:6379"),
	}

	if cfg.DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
