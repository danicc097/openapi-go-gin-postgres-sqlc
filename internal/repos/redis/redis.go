// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package redis

import (
	"context"

	redis "github.com/go-redis/redis/v8"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
)

// New instantiates the Redis client using configuration defined in environment variables.
func New() (*redis.Client, error) {
	cfg := internal.Config()

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
		DB:   cfg.Redis.DB,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "rdb.Ping")
	}

	return rdb, nil
}
