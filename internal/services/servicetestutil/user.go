package servicetestutil

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

type CreateUserParams struct {
	DeletedAt  *time.Time
	Role       models.Role
	Scopes     models.Scopes
	WithToken  bool // if true, an access token is created and returned
	WithAPIKey bool // if true, an api key is created and returned
}

type CreateUserResult struct {
	User   *db.User
	APIKey *db.UserAPIKey
	Token  string
}

// CreateUser creates a new random user with the given configuration.
func (ff *FixtureFactory) CreateUser(ctx context.Context, params CreateUserParams) (*CreateUserResult, error) {
	ucp := services.UserRegisterParams{
		Username:   testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		Email:      testutil.RandomEmail(),
		FirstName:  pointers.New(testutil.RandomFirstName()),
		LastName:   pointers.New(testutil.RandomLastName()),
		ExternalID: testutil.RandomString(10),
		Scopes:     params.Scopes,
		Role:       params.Role,
	}

	// don't use repos for tests
	user, err := ff.usvc.Register(ctx, ff.pool, ucp)
	if err != nil {
		return nil, fmt.Errorf("usvc.Register: %w", err)
	}

	if params.DeletedAt != nil {
		user, err = ff.usvc.Delete(ctx, ff.pool, user.UserID)
		if err != nil {
			return nil, fmt.Errorf("usvc.Delete: %w", err)
		}
	}

	var accessToken string
	var apiKey *db.UserAPIKey

	if params.WithAPIKey {
		apiKey, err = ff.authnsvc.CreateAPIKeyForUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("authnsvc.CreateAPIKeyForUser: %w", err)
		}
	}
	if params.WithToken {
		accessToken, err = ff.authnsvc.CreateAccessTokenForUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("authnsvc.CreateAPIKeyForUser: %w", err)
		}
	}

	return &CreateUserResult{
		User:   user,
		APIKey: apiKey,
		Token:  accessToken,
	}, nil
}
