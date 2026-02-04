package db

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewCache(addr string, dbNum int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,  // e.g. "localhost:6379" or "redis:6379" in Docker network
		DB:   dbNum, // usually 0
	})

	// Test connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	// Optional: configure connection pool
	rdb.Options().PoolSize = 10
	rdb.Options().MinIdleConns = 2
	rdb.Options().ConnMaxIdleTime = 5 * time.Minute

	return rdb
}
