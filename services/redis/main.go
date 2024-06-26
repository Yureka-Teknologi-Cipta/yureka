package yureka_redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Connect(timeout time.Duration, dbURL string) *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	opts, err := redis.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opts)

	if ping := client.Ping(ctx); ping.Err() != nil {
		panic(ping.Err())
	}

	return client
}
