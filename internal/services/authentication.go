package services

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Authentication struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}

func NewAuthentication(logger *zap.Logger) *Authentication {
	return &Authentication{
		Logger: logger,
	}
}
