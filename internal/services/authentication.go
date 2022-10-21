package services

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Authentication struct {
	logger *zap.Logger
	pool   *pgxpool.Pool
}

func NewAuthentication(logger *zap.Logger, pool *pgxpool.Pool) *Authentication {
	return &Authentication{
		logger: logger,
	}
}
