package services

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Fake struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
