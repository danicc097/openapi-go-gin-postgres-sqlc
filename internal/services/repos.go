// grouped for generation caching purposes
package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// UserRepo defines the datastore/repository handling persisting User records.
type UserRepo interface {
	Upsert(ctx context.Context, d db.DBTX, user *db.User) error
	UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, user *db.User) error
}
