package repos

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

type UserUpdateParams struct {
	FirstName *string
	LastName  *string
	Rank      *int16
	Scopes    *[]string
	ID        string
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, user *db.User) error
	Update(ctx context.Context, d db.DBTX, params UserUpdateParams) (*db.User, error)
	CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error)
}
