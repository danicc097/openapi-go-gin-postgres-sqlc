// grouped for generation caching purposes
package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// User defines the datastore/repository handling persisting User records.
// TODO need to figure out how to mix and match sqlc and xo
// so far db interface is the same after some template mods.
type UserRepo interface {
	Upsert(ctx context.Context, user *db.User) error
	UserByEmail(ctx context.Context, email string) (*db.User, error)
	Create(ctx context.Context, user *db.User) error
}
