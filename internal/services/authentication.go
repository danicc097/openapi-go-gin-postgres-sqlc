package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
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

// GetUserFromToken returns a user from a token.
func (a *Authentication) GetUserFromToken(ctx context.Context, token string) (*db.User, error) {
	return &db.User{}, nil
}

// GetUserFromApiKey returns a user from an api key.
func (a *Authentication) GetUserFromApiKey(ctx context.Context, apiKey string) (*db.User, error) {
	return a.usvc.UserByAPIKey(ctx, a.pool, apiKey)
}

// CreateAccessTokenForUser creates a new token for a user.
func (a *Authentication) CreateAccessTokenForUser(ctx context.Context, user db.User) {
}

// GetClaimFromToken creates a new token for a user.
func (a *Authentication) GetClaimFromToken(ctx context.Context, token string, claim string) any {
	return nil
}
