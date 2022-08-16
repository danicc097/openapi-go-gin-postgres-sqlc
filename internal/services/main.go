package services

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// TODO only use services, not repos. No need to turn redis, etc. into repo for now.
// TODO service interfaces

type Default struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
type Docs struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
type Fake struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
type Pet struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
type Store struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
type User struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
