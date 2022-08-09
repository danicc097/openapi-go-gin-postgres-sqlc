// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package internal

import (
	"context"
	"strconv"

	redis "github.com/go-redis/redis/v8"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// NewRedis instantiates the Redis client using configuration defined in environment variables.
func NewRedis(conf *envvar.Configuration) (*redis.Client, error) {
	host, err := conf.Get("REDIS_HOST")
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "conf.Get REDIS_HOST")
	}

	db, err := conf.Get("REDIS_DB")
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "conf.Get REDIS_DB")
	}

	dbi, _ := strconv.Atoi(db)

	rdb := redis.NewClient(&redis.Options{
		Addr: host,
		DB:   dbi,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "rdb.Ping")
	}

	return rdb, nil
}
