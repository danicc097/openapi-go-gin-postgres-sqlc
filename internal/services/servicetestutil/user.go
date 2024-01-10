package servicetestutil

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

type CreateUserParams struct {
	DeletedAt  *time.Time
	Role       models.Role
	Scopes     models.Scopes
	WithToken  bool // if true, an access token is created and returned
	WithAPIKey bool // if true, an api key is created and returned
}

type CreateUserFixture struct {
	User   *db.User
	APIKey *db.UserAPIKey
	Token  string
}

// CreateUser creates a new random user with the given configuration.
func (ff *FixtureFactory) CreateUser(ctx context.Context, params CreateUserParams) (*CreateUserFixture, error) {
	randomRepoCreateParams := postgresqltestutil.RandomUserCreateParams(ff.t)
	ucp := services.UserRegisterParams{
		Username:   randomRepoCreateParams.Username,
		Email:      randomRepoCreateParams.Email,
		FirstName:  randomRepoCreateParams.FirstName,
		LastName:   randomRepoCreateParams.LastName,
		ExternalID: randomRepoCreateParams.ExternalID,
		Scopes:     params.Scopes,
		Role:       params.Role,
	}

	// don't use repos for tests
	user, err := ff.svc.User.Register(ctx, ff.d, ucp)
	if err != nil {
		return nil, fmt.Errorf("svc.User.Register: %w", err)
	}

	if params.DeletedAt != nil {
		user, err = ff.svc.User.Delete(ctx, ff.d, user.UserID)
		if err != nil {
			return nil, fmt.Errorf("svc.User.Delete: %w", err)
		}
	}

	var accessToken string
	var apiKey *db.UserAPIKey

	if params.WithAPIKey {
		apiKey, err = ff.svc.Authentication.CreateAPIKeyForUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("svc.Authentication.CreateAPIKeyForUser: %w", err)
		}
	}
	if params.WithToken {
		accessToken, err = ff.svc.Authentication.CreateAccessTokenForUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("svc.Authentication.CreateAccessTokenForUser: %w", err)
		}
	}

	user, err = ff.svc.User.ByID(ctx, ff.d, user.UserID)
	if err != nil {
		return nil, fmt.Errorf("svc.User.ByEmail: %w", err)
	}

	return &CreateUserFixture{
		User:   user,
		APIKey: apiKey,
		Token:  accessToken,
	}, nil
}
