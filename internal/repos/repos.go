package repos

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// User defines the datastore/repository handling persisting User records.
type User interface {
	Upsert(ctx context.Context, d db.DBTX, user *db.User) error
	UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, user *db.User) error
}
