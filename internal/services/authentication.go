package services

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Authentication struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
	usvc   *User
}

func NewAuthentication(pool *pgxpool.Pool, logger *zap.Logger) *Authentication {
	return &Authentication{
		pool:   pool,
		logger: logger,
		usvc:   NewUser(pool, logger),
	}
}

func (a *Authentication) GetUserFromToken(ctx context.Context) {
}

func (a *Authentication) GetUserFromApiKey(ctx context.Context) {
}
