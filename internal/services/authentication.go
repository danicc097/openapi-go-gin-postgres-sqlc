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

func NewAuthentication(logger *zap.Logger, usvc *User, pool *pgxpool.Pool) *Authentication {
	return &Authentication{
		logger: logger,
		usvc:   usvc,
		pool:   pool,
	}
}

// GetUserFromAccessToken returns a user from a token.
func (a *Authentication) GetUserFromAccessToken(ctx context.Context, token string) (*db.User, error) {
	return &db.User{}, nil
}

// GetUserFromAPIKey returns a user from an api key.
func (a *Authentication) GetUserFromAPIKey(ctx context.Context, apiKey string) (*db.User, error) {
	return a.usvc.UserByAPIKey(ctx, a.pool, apiKey)
}

// CreateAccessTokenForUser creates a new token for a user.
func (a *Authentication) CreateAccessTokenForUser(ctx context.Context, user *db.User) string {
	return ""
}

// CreateAccessTokenForUser creates a new token for a user.
func (a *Authentication) CreateAPIKeyForUser(ctx context.Context, user *db.User) string {
	return ""
}

// GetClaimFromToken creates a new token for a user.
func (a *Authentication) GetClaimFromToken(ctx context.Context, token string, claim string) any {
	return nil
}
