package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Connect(timeout time.Duration, dbURL string) *redis.Client {
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	opts, err := redis.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)
}
