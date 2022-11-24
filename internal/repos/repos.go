package repos

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

type UserCreateParams struct {
	Username   string
	Email      string
	FirstName  *string
	LastName   *string
	ExternalID string
	Scopes     []string
	RoleRank   int16
}

type UserUpdateParams struct {
	FirstName *string
	LastName  *string
	Rank      *int16
	Scopes    *[]string
	ID        string
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	UserByID(ctx context.Context, d db.DBTX, id string) (*db.User, error)
	UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	UserByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error)
	UserByExternalID(ctx context.Context, d db.DBTX, extID string) (*db.User, error)
	UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, params UserCreateParams) (*db.User, error)
	Update(ctx context.Context, d db.DBTX, params UserUpdateParams) (*db.User, error)
	// CreateAPIKey requires an existing user.
	CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error)
}
