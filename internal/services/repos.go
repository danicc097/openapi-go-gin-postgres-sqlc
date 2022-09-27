// grouped for generation caching purposes
package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
)

// User defines the datastore/repository handling persisting User records.
// TODO just crud (for impl see if xo for repo and sqlc for services can be used alongside easily
// or need to have some postgen)
type UserRepo interface {
	Create(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error)
}
