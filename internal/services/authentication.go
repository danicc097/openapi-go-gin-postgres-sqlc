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

func NewAuthentication(logger *zap.Logger, usvc *User) *Authentication {
	return &Authentication{
		logger: logger,
		usvc:   usvc,
	}
}

func (a *Authentication) GetUserFromToken(ctx context.Context) {
}

func (a *Authentication) GetUserFromApiKey(ctx context.Context) {
}
